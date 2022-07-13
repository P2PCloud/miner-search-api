package search

import (
	"context"

	gw "github.com/P2PCloud/miner-search-api/api/github.com/P2PCloud/miner-search-api"
	"github.com/P2PCloud/miner-search-api/internal/app/models"
)

func (a appImpl) FindOffer(ctx context.Context, request *gw.FindOfferRequest) (*gw.FindOfferResponse, error) {
	offers, err := a.r.GetOffersByParams(ctx,
		&models.VmType{
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
		vmType := a.r.GetVmTypeById(offer.VmTypeId)

		gwOffers[i] = &gw.Offer{
			VmType: &gw.VmType{
				Region:    vmType.Region,
				CpuSingle: vmType.CpuSingle,
				CpuMulti:  vmType.CpuMulti,
				Ram:       vmType.Ram,
				DiskSpeed: vmType.DiskSpeed,
				DiskSize:  vmType.DiskSize,
				NetSpeed:  vmType.NetSpeed,
			},
			Miner: offer.Miner,
			Pps:   offer.Pps,
		}
	}
	return &gw.FindOfferResponse{
		Offers: gwOffers,
	}, nil
}
