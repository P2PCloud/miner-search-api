package models

type Offer struct {
	Region      string `json:"region"`
	CpuSingle   int32  `json:"cpuSingle"`
	CpuMulti    int32  `json:"cpuMulti"`
	Ram         int32  `json:"ram"`
	DiskSpeed   int32  `json:"diskSpeed"`
	DiskSize    int32  `json:"diskSize"`
	NetSpeed    int32  `json:"netSpeed"`
	MinerRating int32  `json:"minerRating"`
}

type OfferCollection []*Offer
