package block

type Block struct {
	BlockID      string        `json:"blockID"`
	BlockHeader  BlockHeader   `json:"block_header"`
	Transactions []Transaction `json:"transactions"`
}

type Blocks struct {
	Block []Block `json:"block"`
}

type SmartContractRet struct {
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
	Transaction Transaction `json:"transaction"`
}

type ConstantContractRet struct {
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
	EnergyUsed     int         `json:"energy_used"`
	ConstantResult []string    `json:"constant_result"`
	Transaction    Transaction `json:"transaction"`
}
