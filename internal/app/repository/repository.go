package repository

import (
	"context"
	"sort"

	"github.com/P2PCloud/miner-search-api/internal/app/models"
	"github.com/P2PCloud/miner-search-api/internal/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type Repository interface {
	GetOffersByParams(ctx context.Context, params *models.VmType) (models.OfferCollection, error)
	GetVmTypeById(id string) *models.VmType
	GetMinerReputationById(id string) float32
	CompareOffer(offer *models.Offer, params *models.VmType) bool
}

type repoImpl struct {
	dataRaw  string
	data     db
	Defaults config.Defaults
}

type db struct {
	VmTypes           models.VmTypeMap             `json:"vmTypes"`
	VmOffers          models.OfferCollection       `json:"vmOffers"`
	MinerPayouts      models.MinerPayoutCollection `json:"minerPayouts"`
	UserReports       models.UserReportCollection  `json:"userReports"`
	BusinessActivity  map[string]int32
	UserReportsPpsSum map[string]int32
	UserReputation    map[string]float32
	VotingPower       map[string]int32
	MinerReputation   map[string]float32
}

func NewRepository(d config.Db) (Repository, error) {
	r := &repoImpl{
		dataRaw:  d.Data,
		Defaults: d.Defaults,
		data: db{
			VmTypes:           models.VmTypeMap{},
			VmOffers:          make(models.OfferCollection, 0),
			MinerPayouts:      make(models.MinerPayoutCollection, 0),
			UserReports:       make(models.UserReportCollection, 0),
			BusinessActivity:  map[string]int32{},
			UserReportsPpsSum: map[string]int32{},
			UserReputation:    map[string]float32{},
			VotingPower:       map[string]int32{},
			MinerReputation:   map[string]float32{},
		},
	}

	err := r.parseData([]byte(d.Data))
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

	r.calcBusinessActivity()
	r.calcUserReportsPpsSum()
	r.calcUserReputation()

	return nil
}

func makeDivided(logs []ur, num int) [][]ur {
	var divided [][]ur

	chunkSize := (len(logs) + num - 1) / num

	for i := 0; i < len(logs); i += chunkSize {
		end := i + chunkSize

		if end > len(logs) {
			end = len(logs)
		}

		divided = append(divided, logs[i:end])
	}

	return divided
}

type ur struct {
	user       string
	reputation float32
}

func (r *repoImpl) calcUserReputation() {
	var sorted []ur
	for _, offer := range r.data.VmOffers {
		if _, ok := r.data.UserReputation[offer.Miner]; !ok {
			urs := r.data.UserReportsPpsSum[offer.Miner]
			if urs < r.Defaults.UserReputation {
				urs = r.Defaults.UserReputation
			}

			ba := r.data.BusinessActivity[offer.Miner]
			rr := float32(ba) / float32(urs)
			sorted = append(sorted, ur{user: offer.Miner, reputation: rr})
			r.data.UserReputation[offer.Miner] = rr
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].reputation > sorted[j].reputation
	})

	divided := makeDivided(sorted, 5)

	for i, dd := range divided {
		for _, d := range dd {
			if i == 0 {
				r.data.VotingPower[d.user] = 4
			} else if i == 1 {
				r.data.VotingPower[d.user] = 2
			}
			if i == 2 || i == 3 {
				r.data.VotingPower[d.user] = 1
			}
			if i == 4 {
				r.data.VotingPower[d.user] = 0
			}
		}
	}
}

func (r *repoImpl) calcUserReportsPpsSum() {
	for _, report := range r.data.UserReports {
		if _, ok := r.data.UserReportsPpsSum[report.User]; !ok {
			r.data.UserReportsPpsSum[report.User] = 0
		}
		r.data.UserReportsPpsSum[report.User] += report.Pps
	}
}

func (r *repoImpl) calcMinerReputation() {
	for k, v := range r.data.VotingPower {
		ba, ok := r.data.BusinessActivity[k]
		if !ok {
			ba = 1
		}

		urs := r.data.UserReportsPpsSum[k]
		d := urs * v
		if d < r.Defaults.MinerReputation {
			d = r.Defaults.MinerReputation
		}

		r.data.MinerReputation[k] = float32(ba) / float32(d)
	}
}

func (r *repoImpl) calcBusinessActivity() {
	for _, payout := range r.data.MinerPayouts {
		if _, ok := r.data.BusinessActivity[payout.Miner]; !ok {
			r.data.BusinessActivity[payout.Miner] = 0
		}
		r.data.BusinessActivity[payout.Miner] += payout.Amount
	}
}
