package search

import (
	"context"
	gw "github.com/P2PCloud/miner-search-api/api/github.com/P2PCloud/miner-search-api"
	"github.com/P2PCloud/miner-search-api/internal/app/models"
	"github.com/P2PCloud/miner-search-api/internal/app/repository"
	"github.com/P2PCloud/miner-search-api/internal/app/test/mocks"
	"reflect"
	"testing"
)

func Test_appImpl_FindOffer(t *testing.T) {
	r := mocks.NewRepositoryMock(mocks.MCtrl{T: t})
	r.GetOffersByParamsMock.Set(func(p context.Context, p1 *models.VmType) (r models.OfferCollection, r1 error) {
		return nil, nil
	})

	r1 := mocks.NewRepositoryMock(mocks.MCtrl{T: t})
	r1.GetOffersByParamsMock.Set(func(p context.Context, p1 *models.VmType) (r models.OfferCollection, r1 error) {
		return models.OfferCollection{{
			Id:        1,
			VmTypeId:  "2",
			Miner:     "3",
			Available: 4,
			Pps:       5,
		}}, nil
	})
	r1.GetVmTypeByIdMock.Set(func(p string) (r *models.VmType) {
		return &models.VmType{
			Region:    "1",
			CpuSingle: 2,
			CpuMulti:  3,
			Ram:       4,
			DiskSpeed: 5,
			DiskSize:  6,
			NetSpeed:  7,
		}
	})

	type fields struct {
		serviceVersion string
		r              repository.Repository
	}
	type args struct {
		ctx     context.Context
		request *gw.FindOfferRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gw.FindOfferResponse
		wantErr bool
	}{

		{"sucess empty", fields{r: r}, args{request: &gw.FindOfferRequest{}}, &gw.FindOfferResponse{
			Offers: []*gw.Offer{},
		}, false},
		{"sucess one offer", fields{r: r1}, args{request: &gw.FindOfferRequest{}}, &gw.FindOfferResponse{
			Offers: []*gw.Offer{{
				Id:    1,
				Pps:   5,
				Miner: "3",
				VmType: &gw.VmType{
					Region:    "1",
					CpuSingle: 2,
					CpuMulti:  3,
					Ram:       4,
					DiskSpeed: 5,
					DiskSize:  6,
					NetSpeed:  7,
				},
			}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := appImpl{
				serviceVersion: tt.fields.serviceVersion,
				r:              tt.fields.r,
			}
			got, err := a.FindOffer(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOffer() got = %v, want %v", got.Offers, tt.want)
			}
		})
	}
}
