package account

type Account struct {
	Address string              `json:"address"`
	Balance int64               `json:"balance"`
	Trc20   []map[string]string `json:"trc20"`
	Hex     Hex
}

type Hex []byte
