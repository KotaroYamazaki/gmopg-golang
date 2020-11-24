package client

import (
	"fmt"
	"gmopg-card/config"

	"github.com/evalphobia/go-gmo-pg/request"
)

const (
	endpointSandbox    = "https://pt01.mul-pay.jp"
	endpointProduction = "https://pt01.mul-pay.jp"
)

type Client struct {
	Config *config.Config `url:"-"`
	Option request.Option `url:"-"`
}

type parameter struct {
	Common Client      `url:",squash"`
	Extra  interface{} `url:",squash"`
}

func New(conf *config.Config) Client {
	return Client{
		Config: conf,
	}
}

func (c Client) Call(path string, param interface{}, result interface{}) error {
	p := parameter{
		Common: c,
		Extra:  param,
	}

	if c.Config.IsProduction() {
		return request.CallPOST(fmt.Sprintf("%s%s", endpointProduction, path), p, c.Option, result)
	}
	return request.CallPOST(fmt.Sprintf("%s%s", endpointSandbox, path), p, c.Option, result)

}
