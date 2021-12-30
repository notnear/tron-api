package client

import "tron-api/block"

func (t *Client) GetNowBlock() (result block.Block, err error) {
	err = t.Post("/wallet/getnowblock", nil, &result)
	if err != nil {
		return
	}
	return
}

func (t *Client) GetBlockByNum(num int64) (result block.Block, err error) {
	err = t.Post("/wallet/getblockbynum", struct {
		Num int64 `json:"num"`
	}{Num: num}, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) GetBlockById(hashId string) (result block.Block, err error) {
	err = t.Post("/wallet/getblockbyid", struct {
		Value string `json:"value"`
	}{Value: hashId}, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) GetBlockByLatestNum(num int32) (result block.Blocks, err error) {
	err = t.Post("/wallet/getblockbylatestnum", struct {
		Num int32 `json:"num"`
	}{Num: num}, &result)

	if err != nil {
		return
	}
	return
}

func (t *Client) GetBlockByLimitNext(startNum, endNum int32) (result block.Blocks, err error) {
	err = t.Post("/wallet/getblockbylimitnext", struct {
		StartNum int32 `json:"startNum"`
		EndNum   int32 `json:"endNum"`
	}{StartNum: startNum, EndNum: endNum}, &result)

	if err != nil {
		return
	}
	return
}
