package client

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/net/ghttp"
)

func NewClient(network string) *Client {
	client := Client{}
	switch network {
	case "shasta":
		client.Host = Shasta
		break
	case "nile":
		client.Host = NILE
		break
	default:
		client.Host = MAIN
	}
	return &client
}

func (t *Client) OpenDebug() *Client {
	t.Debug = true
	return t
}

func (t *Client) CloseDebug() *Client {
	t.Debug = false
	return t
}

func (t *Client) Post(uri string, params, result interface{}) (err error) {
	httpClient := ghttp.NewClient()

	httpClient.SetHeaderMap(t.Headers)

	response, err := httpClient.ContentJson().Post(t.Host+uri, params)
	if err != nil {
		return
	}
	defer response.Close()

	if t.Debug {
		response.RawDump()
	}
	err = gvar.New(response.ReadAll()).Scan(&result)
	return
}

func (t *Client) Get(uri string, result interface{}) (err error) {
	httpClient := ghttp.NewClient()

	httpClient.SetHeaderMap(t.Headers)

	response, err := httpClient.ContentJson().Get(t.Host + uri)
	if err != nil {
		return
	}
	defer response.Close()

	if t.Debug {
		response.RawDump()
	}

	err = gvar.New(response.ReadAll()).Scan(&result)
	return
}
