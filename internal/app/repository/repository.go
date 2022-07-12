package repository

import (
	"context"

	"github.com/P2PCloud/miner-search-api/internal/app/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type Repository interface {
	GetOffersByParams(ctx context.Context, params *models.Offer) (models.OfferCollection, error)
}

type repoImpl struct {
	dataRaw string
	data    models.OfferCollection
}

func NewRepository(data string) (Repository, error) {
	r := &repoImpl{
		dataRaw: data,
		data:    make(models.OfferCollection, 0),
	}

	err := r.parseData([]byte(data))
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *repoImpl) parseData(dataRaw []byte) error {
	if len(dataRaw) == 0 {
		return errors.New("data doesn't set")
	}

	if err := jsoniter.Unmarshal(dataRaw, &r.data); err != nil {
		return errors.Wrap(err, "can't unmarshal data")
	}

	return nil
}
