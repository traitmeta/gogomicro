package config

type ConsulConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

type defaultConsulConfig struct {
	Enabled bool   `json:"enabled"`
	Port    int    `json:"port"`
	Host    string `json:"host"`
}

func (c defaultConsulConfig) GetEnabled() bool {
	return c.Enabled
}

func (c defaultConsulConfig) GetPort() int {
	return c.Port
}

func (c defaultConsulConfig) GetHost() string {
	return c.Host
}
