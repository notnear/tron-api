package client

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"tron-api/block"
	"tron-api/common"
)

func (t *Client) GetTransactionById(hashId string) (result *block.Transaction, err error) {
	err = t.Post("/wallet/gettransactionbyid", struct {
		Value string `json:"value"`
	}{Value: hashId}, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) CreateTransaction(owner, to string, amount int64) (result *block.Transaction, err error) {
	err = t.Post("/wallet/createtransaction", struct {
		Owner   string `json:"owner_address"`
		To      string `json:"to_address"`
		Amount  int64  `json:"amount"`
		Visible bool   `json:"visible"`
	}{Owner: owner, To: to, Amount: amount, Visible: true}, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) BroadcastTransaction(tran *block.Transaction) (txId string, err error) {
	if len(tran.Signature) < 0 {
		err = errors.New("transaction need sign")
		return
	}
	var result struct {
		Code    string `json:"code"`
		TxId    string `json:"txid"`
		Message string `json:"message"`
	}
	err = t.Post("/wallet/broadcasttransaction", tran, &result)

	if err != nil {
		return
	}
	if result.Code != "" {
		err = errors.New(result.Message)
	} else {
		txId = result.TxId
	}
	return
}

func (t *Client) SignTransaction(transaction *block.Transaction, privateKey string) (*block.Transaction, error) {
	privateBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	if privateBytes == nil {
		return nil, errors.New("privateKey is error")
	}
	priv := crypto.ToECDSAUnsafe(privateBytes)
	defer zeroKey(priv)

	s, er := common.Hex2Bytes(transaction.RawDataHex)
	if er != nil {
		return nil, er
	}

	h256h := sha256.New()
	h256h.Write(s)
	hash := h256h.Sum(nil)
	signature, err := crypto.Sign(hash, priv)
	if err != nil {
		return nil, err
	}
	var signatures []string
	signatures = append(signatures, common.Bytes2Hex(signature))

	transaction.Signature = signatures

	return transaction, nil
}

func zeroKey(k *ecdsa.PrivateKey) {
	b := k.D.Bits()
	for i := range b {
		b[i] = 0
	}
}
