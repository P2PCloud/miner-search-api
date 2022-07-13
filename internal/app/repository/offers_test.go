package repository

import (
	"context"
	"github.com/P2PCloud/miner-search-api/internal/app/models"
	"reflect"
	"testing"
)

func Test_repoImpl_GetOffersByParams(t *testing.T) {
	type fields struct {
		dataRaw string
		data    db
	}
	type args struct {
		ctx    context.Context
		params *models.VmType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.OfferCollection
		wantErr bool
	}{
		{"success empty", fields{}, args{}, models.OfferCollection{}, false},
		{"success one row", fields{data: db{
			VmOffers: models.OfferCollection{{
				Id:        1,
				VmTypeId:  "2",
				Miner:     "3",
				Available: 4,
				Pps:       5,
			}},
			VmTypes: models.VmTypeMap{"2": {
				Region:    "1",
				CpuSingle: 2,
				CpuMulti:  3,
				Ram:       4,
				DiskSpeed: 5,
				DiskSize:  6,
				NetSpeed:  7,
			}}}}, args{params: &models.VmType{}}, models.OfferCollection{{
			Id: 1, VmTypeId: "2", Miner: "3", Available: 4, Pps: 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				dataRaw: tt.fields.dataRaw,
				data:    tt.fields.data,
			}
			got, err := r.GetOffersByParams(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOffersByParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOffersByParams() got = %#v, want %#v", got[0], tt.want)
			}
		})
	}
}
