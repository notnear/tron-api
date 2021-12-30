package client

import (
	"fmt"
	"math/big"
	"tron-api/account"
	"tron-api/block"
	"tron-api/common"
)

const (
	trc20TransferMethodSignature = "0xa9059cbb"
	trc20TransferEventSignature  = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	trc20NameSignature           = "0x06fdde03"
	trc20SymbolSignature         = "0x95d89b41"
	trc20DecimalsSignature       = "0x313ce567"
	trc20BalanceOf               = "0x70a08231"
)

type SmartParams struct {
	OwnerAddress     string `json:"owner_address"`
	ContractAddress  string `json:"contract_address"`
	FunctionSelector string `json:"function_selector"`
	Parameter        string `json:"parameter"`
	FeeLimit         int32  `json:"fee_limit"`
	CallValue        int32  `json:"call_value"`
	Visible          bool   `json:"visible"`
}

func (t *Client) TriggerSmartContract(params SmartParams) (result *block.ConstantContractRet, err error) {
	err = t.Post("/wallet/triggersmartcontract", params, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) TriggerConstantContract(params SmartParams) (result *block.ConstantContractRet, err error) {
	err = t.Post("/wallet/triggerconstantcontract", params, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) ContractBalanceOf(owner, contract string) (*big.Int, error) {
	params := SmartParams{
		OwnerAddress:     owner,
		ContractAddress:  contract,
		FunctionSelector: "balanceOf(address)",
		Visible:          true,
	}
	addrB, _ := account.Base58ToHex(owner)

	parameter := "0x70a08231000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-2:] + addrB.Hex()[2:]
	params.Parameter = parameter
	result, err := t.TriggerConstantContract(params)
	if err != nil {
		return nil, err
	}
	r, err := account.ParseTRC20NumericProperty(result.ConstantResult[0])
	if err != nil {
		return nil, fmt.Errorf("contract address %s: %v", contract, err)
	}

	return r, nil
}

func (t *Client) Trc20send(from, to, contract string, amount *big.Int, feeLimit int32) (result *block.ConstantContractRet, err error) {
	addrB, err := account.Base58ToHex(to)
	if err != nil {
		return nil, err
	}
	ab := common.LeftPadBytes(amount.Bytes(), 32)
	req := "0xa9059cbb000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-4:] + addrB.Hex()[4:]
	req += common.Bytes2Hex(ab)

	params := SmartParams{
		OwnerAddress:     from,
		ContractAddress:  contract,
		FunctionSelector: "transfer(address,uint256)",
		Parameter:        req,
		Visible:          true,
		FeeLimit:         feeLimit,
	}

	result, err = t.TriggerSmartContract(params)
	return
}
