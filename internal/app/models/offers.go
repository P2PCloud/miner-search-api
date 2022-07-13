package models

type VmType struct {
	Region    string `json:"region"`
	CpuSingle int32  `json:"cpuSingle"`
	CpuMulti  int32  `json:"cpuMulti"`
	Ram       int32  `json:"ram"`
	DiskSpeed int32  `json:"diskSpeed"`
	DiskSize  int32  `json:"diskSize"`
	NetSpeed  int32  `json:"netSpeed"`
}

type VmTypeCollection []*VmType
type VmTypeMap map[string]*VmType

type Offer struct {
	Id        int32  `json:"id"`
	VmTypeId  string `json:"vm_type_id"`
	Miner     string `json:"miner"`
	Available int32  `json:"available"`
	Pps       int32  `json:"pps"`
}

type OfferCollection []*Offer

type MinerPayout struct {
	Amount int32  `json:"amount"`
	User   string `json:"user"`
	Miner  string `json:"miner"`
}

type MinerPayoutCollection []*MinerPayout

type UserReport struct {
	Pps   int32  `json:"pps"`
	User  string `json:"user"`
	Miner string `json:"miner"`
}

type UserReportCollection []*UserReport
