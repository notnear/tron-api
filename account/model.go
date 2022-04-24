package account

type Account struct {
	Address string              `json:"address"`
	Balance int64               `json:"balance"`
	Trc20   []map[string]string `json:"trc20"`
	Hex     Hex
}

type Accounts struct {
	Data    []Account `json:"data"`
	Success bool      `json:"success"`
	Meta    struct {
		At       int `json:"at"`
		PageSize int `json:"page_size"`
	}
}

type Hex []byte
