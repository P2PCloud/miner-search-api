package search

import (
	"context"
	"github.com/P2PCloud/miner-search-api/internal/app/models"
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

func (a appImpl) FindOffer(ctx context.Context, request *gw.FindOfferRequest) (*gw.FindOfferResponse, error) {
	offers, err := a.r.GetOffersByParams(ctx,
		&models.Offer{
			Region:    request.Region,
			CpuSingle: request.CpuSingle,
			CpuMulti:  request.CpuMulti,
			Ram:       request.Ram,
			DiskSpeed: request.DiskSpeed,
			DiskSize:  request.DiskSize,
			NetSpeed:  request.NetSpeed,
		},
	)
	if err != nil {
		return nil, err
	}

	gwOffers := make([]*gw.Offer, len(offers))
	for i, offer := range offers {
		gwOffers[i] = &gw.Offer{
			Region:    offer.Region,
			CpuSingle: offer.CpuSingle,
			CpuMulti:  offer.CpuMulti,
			Ram:       offer.Ram,
			DiskSpeed: offer.DiskSpeed,
			DiskSize:  offer.DiskSize,
			NetSpeed:  offer.NetSpeed,
		}
	}
	return &gw.FindOfferResponse{
		Offers: gwOffers,
	}, nil
}

func NewApp(version string, r repository.Repository) (App, error) {
	return &appImpl{
		serviceVersion: version,

		r: r,
	}, nil
}
