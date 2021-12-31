package block

type BlockHeader struct {
	RawData          headRawData `json:"raw_data"`
	WitnessSignature string      `json:"witness_signature"`
}

type headRawData struct {
	Number         int64  `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int    `json:"version"`
	Timestamp      int    `json:"timestamp"`
}

type Transaction struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature  []string           `json:"signature"`
	TxID       string             `json:"txID"`
	RawData    TransactionRawData `json:"raw_data"`
	RawDataHex string             `json:"raw_data_hex"`
	Visible    bool               `json:"visible"`
}

type TransactionRawData struct {
	Contract []struct {
		Type      string `json:"type"`
		Parameter struct {
			TypeUrl string `json:"type_url"`
			Value   struct {
				Data            string `json:"data"`
				OwnerAddress    string `json:"owner_address"`
				ContractAddress string `json:"contract_address"`
				Amount          int64  `json:"amount"`
				ToAddress       string `json:"to_address"`
			} `json:"value"`
		} `json:"parameter"`
	} `json:"contract"`
	RefBlockBytes string `json:"ref_block_bytes"`
	RefBlockHash  string `json:"ref_block_hash"`
	Expiration    int64  `json:"expiration"`
	FeeLimit      int64  `json:"fee_limit"`
	Timestamp     int    `json:"timestamp"`
}

type TransactionInfo struct {
	Id              string   `json:"id"`
	BlockNumber     int64    `json:"blockNumber"`
	BlockTimeStamp  int64    `json:"blockTimeStamp"`
	ContractResult  []string `json:"contractResult"`
	ContractAddress string   `json:"contract_address"`
	Receipt         struct {
		OriginEnergyUsage int64  `json:"origin_energy_usage"`
		EnergyUsageTotal  int64  `json:"energy_usage_total"`
		NetUsage          int64  `json:"net_usage"`
		Result            string `json:"result"`
	} `json:"receipt"`
	Result     string `json:"result"`
	ResMessage string `json:"resMessage"`
}
