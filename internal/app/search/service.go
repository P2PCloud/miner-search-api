package search

import (
	"context"
	"github.com/P2PCloud/miner-search-api/internal/app/repository"

	gw "github.com/P2PCloud/miner-search-api/api/github.com/P2PCloud/miner-search-api"
)

type App interface {
	FindOffer(ctx context.Context, request *gw.FindOfferRequest) (*gw.FindOfferResponse, error)
}

type appImpl struct {
	serviceVersion string

	r repository.Repository
}

func NewApp(version string, r repository.Repository) (App, error) {
	return &appImpl{
		serviceVersion: version,

		r: r,
	}, nil
}
