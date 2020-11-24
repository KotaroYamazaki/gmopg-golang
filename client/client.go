package client

const (
	endpointSandbox    = ""
	endpointProduction = ""
)

type Client struct {
	Config *config.Config `url:"-"`
}
