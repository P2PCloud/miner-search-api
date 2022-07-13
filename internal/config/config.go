package config

import (
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

var Cfg config

const defaultConfigFilename = "config.yml"

type config struct {
	Log struct {
		Level       string
		SentryDSN   string
		Environment string
	}
	GRPC struct {
		URL string
	}
	HTTP struct {
		URL string
	}
	DB Db
}

type Defaults struct {
	UserReputation  int32
	MinerReputation int32
}

type Db struct {
	Defaults Defaults
	Data     string
}

func InitConfig(configFilepath string) {
	decimal.DivisionPrecision = 2

	if configFilepath == "" {
		configFilepath = defaultConfigFilename
	}

	viper.SetConfigFile(configFilepath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		logrus.New().WithError(err).Fatal("cannot load config")
	}
}
