package config

type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enabled            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enabled
}

func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
