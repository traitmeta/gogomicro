package config

type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enable"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func (m defaultMysqlConfig) GetEnabled() bool {
	panic("implement me")
}

func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

func (m defaultMysqlConfig) GetEnable() bool {
	return m.Enable
}

func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
