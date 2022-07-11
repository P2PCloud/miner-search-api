package search

import (
	"context"

	gw "github.com/P2PCloud/miner-search-api/api/github.com/P2PCloud/miner-search-api"
)

type App interface {
	FindOffer(ctx context.Context, request *gw.FindOfferRequest) (*gw.FindOfferResponse, error)
}

type appImpl struct {
	serviceVersion string
}

func (a appImpl) FindOffer(ctx context.Context, request *gw.FindOfferRequest) (*gw.FindOfferResponse, error) {
	return &gw.FindOfferResponse{
		Offers: []*gw.Offer{{Region: "moscow"}},
	}, nil
}

func NewApp(version string) (App, error) {
	return &appImpl{
		serviceVersion: version,
	}, nil
}
