package account

import (
	"fmt"
	"github.com/notnear/tron-api/common"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

func (a *Account) ToTrx() (trx float64) {
	balance := decimal.NewFromInt(a.Balance)
	if balance.Cmp(decimal.NewFromInt(0)) > 0 {
		trx, _ = balance.Div(decimal.NewFromFloat(math.Pow10(6))).Float64()
	}
	return
}

func ParseTRC20NumericProperty(data string) (*big.Int, error) {
	if common.Has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) == 64 {
		var n big.Int
		_, ok := n.SetString(data, 16)
		if ok {
			return &n, nil
		}
	}
	return nil, fmt.Errorf("Cannot parse %s", data)
}

func Base58ToHex(s string) (Hex, error) {
	addr, err := common.DecodeCheck(s)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (h Hex) Hex() string {
	return common.ToHex(h[:])
}

func (h Hex) String() string {
	if h[0] == 0 {
		return new(big.Int).SetBytes(h.Bytes()).String()
	}
	return common.EncodeCheck(h.Bytes())
}

func (h Hex) Bytes() []byte {
	return h[:]
}
