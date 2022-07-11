package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"net/http"
	"os"
	"strings"

	gw "github.com/P2PCloud/miner-search-api/api/github.com/P2PCloud/miner-search-api"
	"github.com/P2PCloud/miner-search-api/internal/app/search"
	"github.com/P2PCloud/miner-search-api/internal/config"
	"github.com/P2PCloud/miner-search-api/pkg/middleware"
	"github.com/golang/glog"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var Version = ""

const MaxReceiveMessageSize = 50000000

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		return errors.New("os.Args len error")
	}

	config.InitConfig(argsWithoutProg[0])

	var group errgroup.Group

	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusEntry := logrus.NewEntry(logger)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", config.Cfg.GRPC.URL)
	if err != nil {
		logrusEntry.WithError(err).Fatal("Failed to listen")
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
	}

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.GRPCRecover,
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
		),
	)

	// Attach the Cart service to the server
	app, err := search.NewApp(Version)
	if err != nil {
		logrusEntry.WithError(err).Fatal("cannot init app")
	}
	gw.RegisterSearchServiceServer(s, app)

	// Serve gRPC server
	logrusEntry.Debugf("Serving gRPC on %s", config.Cfg.GRPC.URL)
	group.Go(func() error {
		return s.Serve(lis)
	})

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		config.Cfg.GRPC.URL,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxReceiveMessageSize)),
	)
	if err != nil {
		logrusEntry.WithError(err).Fatal("Failed to dial server")
	}

	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{OrigName: true, EmitDefaults: true},
		),
	)

	// Register Search service
	err = gw.RegisterSearchServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logrusEntry.WithError(err).Fatal("Failed to register gateway")
	}

	gwServer := &http.Server{
		Addr: config.Cfg.HTTP.URL,
		Handler: middleware.HTTPRecover(
			apmhttp.Wrap(removeTrailingSlash(gwMux)),
		),
	}

	// TODO отдача api.swagger.json
	logrusEntry.Debugf("Serving gRPC-Gateway on http://%s", config.Cfg.HTTP.URL)
	group.Go(func() error {
		return gwServer.ListenAndServe()
	})

	return group.Wait()
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
