package webconfig

// DatabaseConfig

func NewDatabaseConfig() *DatabaseConfig { // {{{
	return &DatabaseConfig{}
} // }}}

type DatabaseConfig struct {
	BaseConfig

	Driver        string
	DSN           string
	MaxIdleConns  int
	MaxOpenConns  int
	SingularTable bool
	LogMode       bool
}

func (this *DatabaseConfig) Name() string {
	return CONFIG_DATABASE
}
