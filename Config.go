package webconfig

import (
//	"fmt"
//	"io/ioutil"
//	"path"
)

// [Config]

func LoadConfig(appConfig ConfigInterface, dir string) (*Config, *ConfigErrors) { // {{{
	config := NewConfig(appConfig)
	return config, config.Load(dir)
} // }}}

func NewConfig(appConfig ConfigInterface) *Config { // {{{
	return &Config{
		App:      appConfig,
		Main:     NewMainConfig(),
		Server:   NewServerConfig(),
		Database: NewDatabaseConfig(),
	}
} // }}}

type Config struct {
	BaseConfig

	App      ConfigInterface
	Main     *MainConfig
	Server   *ServerConfig
	Database *DatabaseConfig
}

func (this *Config) HasApp() bool { // {{{
	return this.App != nil && this.App.IsLoaded()
} // }}}

func (this *Config) HasMain() bool { // {{{
	return this.Main != nil && this.Main.IsLoaded()
} // }}}

func (this *Config) HasServer() bool { // {{{
	return this.Server != nil && this.Server.IsLoaded()
} // }}}

func (this *Config) HasDatabase() bool { // {{{
	return this.Database != nil && this.Database.IsLoaded()
} // }}}

func (this *Config) Load(dir string) *ConfigErrors { // {{{
	var err error
	errs := make(ConfigErrors, 0)

	if err = Load(dir, this.App); err != nil {
		errs[this.App.Name()] = err
	}

	if err = Load(dir, this.Main); err != nil {
		errs[this.Main.Name()] = err
	}

	if err = Load(dir, this.Server); err != nil {
		errs[this.Server.Name()] = err
	}

	if err = Load(dir, this.Database); err != nil {
		errs[this.Database.Name()] = err
	}

	this.setLoaded(true)

	return &errs
} // }}}
