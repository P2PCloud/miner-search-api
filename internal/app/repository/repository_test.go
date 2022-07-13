package repository

import (
	"reflect"
	"testing"

	"github.com/P2PCloud/miner-search-api/internal/app/models"
)

func Test_repoImpl_calcUserReputation(t *testing.T) {
	type fields struct {
		data db
	}
	tests := []struct {
		name    string
		fields  fields
		checkDb db
	}{
		{"check", fields{db{
			VmOffers: models.OfferCollection{
				{Miner: "0"},
				{Miner: "1"},
				{Miner: "2"},
				{Miner: "3"},
				{Miner: "4"},
				{Miner: "5"},
				{Miner: "6"},
				{Miner: "7"},
				{Miner: "8"},
				{Miner: "9"},
			},
			UserReputation: map[string]float32{},
			UserReportsPpsSum: map[string]int32{
				"0": 10,
				"1": 20,
				"2": 30,
				"3": 40,
				"4": 50,
				"5": 60,
				"6": 70,
				"7": 80,
				"8": 90,
				"9": 100,
			},
			BusinessActivity: map[string]int32{
				"0": 100,
				"1": 100,
				"2": 100,
				"3": 100,
				"4": 100,
				"5": 100,
				"6": 100,
				"7": 100,
				"8": 100,
				"9": 100,
			},
			VotingPower: map[string]int32{},
		}}, db{
			UserReputation: map[string]float32{
				"0": 10,
				"1": 5,
				"2": 3.3333333,
				"3": 2.5,
				"4": 2,
				"5": 1.6666666,
				"6": 1.4285715,
				"7": 1.25,
				"8": 1.1111112,
				"9": 1,
			},
			VotingPower: map[string]int32{
				"0": 4,
				"1": 4,
				"2": 2,
				"3": 2,
				"4": 1,
				"5": 1,
				"6": 1,
				"7": 1,
				"8": 0,
				"9": 0,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				data: tt.fields.data,
			}
			r.calcUserReputation()

			if !reflect.DeepEqual(r.data.UserReputation, tt.checkDb.UserReputation) {
				t.Fatal("UserReputation not correct")
			}
			if !reflect.DeepEqual(r.data.VotingPower, tt.checkDb.VotingPower) {
				t.Fatal("VotingPower not correct")
			}
		})
	}
}
