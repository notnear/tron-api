package client

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/notnear/tron-api/account"
)

const (
	// TronBytePrefix is the hex prefix to address
	TronBytePrefix = byte(0x41)
)

func (t *Client) CreateAddress() (pri string, address string) {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	if len(priv.D.Bytes()) != 32 {
		for {
			priv, err := btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(priv.D.Bytes()) == 32 {
				break
			}
		}
	}
	a := pubkeyToAddress(priv.ToECDSA().PublicKey)
	address = a.String()
	pri = hex.EncodeToString(priv.D.Bytes())
	return
}

func pubkeyToAddress(p ecdsa.PublicKey) account.Hex {
	address := crypto.PubkeyToAddress(p)

	addressTron := make([]byte, 0)
	addressTron = append(addressTron, TronBytePrefix)
	addressTron = append(addressTron, address.Bytes()...)
	return addressTron
}

func (t *Client) GetAccount(addr string) (result *account.Account, err error) {
	err = t.Post("/wallet/getaccount", struct {
		Address string `json:"address"`
		Visible bool   `json:"visible"`
	}{Address: addr, Visible: true}, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) GetAccounts(addr string) (result *account.Accounts, err error) {
	err = t.Get("/v1/accounts/"+addr, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) ValidateAddress(addr string) (Is bool, err error) {
	var result struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}
	err = t.Post("/wallet/validateaddress", struct {
		Address string `json:"address"`
	}{Address: addr}, &result)

	if err != nil {
		return
	}
	Is = result.Result
	return
}
