package config

const (
	modeSandbox = iota + 1
	modeProduction
)

type Config struct {
	Mode     int    `url:"-"`
	Version  string `url:"Version,omitempty"`
	ShopID   string `url:"ShopID"`
	ShopPass string `url:"ShopPass"`
}

func New(id, pass string) *Config {
	c := &Config{
		Mode:     modeSandbox,
		ShopID:   id,
		ShopPass: pass,
	}
	return c
}
