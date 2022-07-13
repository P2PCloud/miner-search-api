package repository

import (
	"context"
	"github.com/P2PCloud/miner-search-api/internal/app/models"
	"sort"
)

const offerCollectionLimit = 100

func (r *repoImpl) GetOffersByParams(ctx context.Context, params *models.VmType) (models.OfferCollection, error) {
	out := make(models.OfferCollection, 0)

	for _, offer := range r.data.VmOffers {
		if r.CompareOffer(offer, params) {
			out = append(out, offer)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return r.GetMinerReputationById(out[i].Miner) < r.GetMinerReputationById(out[j].Miner)
	})

	if len(out) > offerCollectionLimit {
		out = out[:offerCollectionLimit]
	}

	return out, nil
}

func (r *repoImpl) GetVmTypeById(id string) *models.VmType {
	return r.data.VmTypes[id]
}
func (r *repoImpl) GetMinerReputationById(id string) float32 {
	return r.data.MinerReputation[id]
}

func (r *repoImpl) CompareOffer(offer *models.Offer, params *models.VmType) bool {
	curVmType := r.GetVmTypeById(offer.VmTypeId)

	if params.Region != "" && curVmType.Region != params.Region {
		return false
	}
	if params.CpuSingle != 0 && curVmType.CpuSingle < params.CpuSingle {
		return false
	}
	if params.CpuMulti != 0 && curVmType.CpuMulti < params.CpuMulti {
		return false
	}
	if params.Ram != 0 && curVmType.Ram < params.Ram {
		return false
	}
	if params.DiskSpeed != 0 && curVmType.DiskSpeed < params.DiskSpeed {
		return false
	}
	if params.DiskSize != 0 && curVmType.DiskSize < params.DiskSize {
		return false
	}
	if params.NetSpeed != 0 && curVmType.NetSpeed < params.NetSpeed {
		return false
	}

	return true
}
