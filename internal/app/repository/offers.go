package repository

import (
	"context"

	"github.com/P2PCloud/miner-search-api/internal/app/models"
)

func (r *repoImpl) GetOffersByParams(ctx context.Context, params *models.Offer) (models.OfferCollection, error) {
	out := make(models.OfferCollection, 0)

	for _, offer := range r.data {
		if params.Region != "" && offer.Region != params.Region {
			continue
		}

		out = append(out, offer)
	}
	return out, nil
}
