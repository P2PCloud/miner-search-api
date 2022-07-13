package repository

import (
	"context"
	"sort"

	"github.com/P2PCloud/miner-search-api/internal/app/models"
)

const offerCollectionLimit = 100

func (r *repoImpl) GetOffersByParams(ctx context.Context, params *models.Offer) (models.OfferCollection, error) {
	out := make(models.OfferCollection, 0)

	for _, offer := range r.data {
		if compareOffer(offer, params) {
			out = append(out, offer)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].MinerRating < out[j].MinerRating
	})

	if len(out) > offerCollectionLimit {
		out = out[:offerCollectionLimit]
	}

	return out, nil
}

func compareOffer(offer, params *models.Offer) bool {
	if params.Region != "" && offer.Region != params.Region {
		return false
	}
	if params.CpuSingle != 0 && offer.CpuSingle != params.CpuSingle {
		return false
	}
	if params.CpuMulti != 0 && offer.CpuMulti != params.CpuMulti {
		return false
	}
	if params.Ram != 0 && offer.Ram != params.Ram {
		return false
	}
	if params.DiskSpeed != 0 && offer.DiskSpeed != params.DiskSpeed {
		return false
	}
	if params.DiskSize != 0 && offer.DiskSize != params.DiskSize {
		return false
	}
	if params.NetSpeed != 0 && offer.NetSpeed != params.NetSpeed {
		return false
	}

	return true
}
