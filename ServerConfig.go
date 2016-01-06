package webconfig

import (
	"fmt"
)

// [ServerConfig]

func NewServerConfig() *ServerConfig { // {{{
	return &ServerConfig{}
} // }}}

type ServerConfig struct {
	BaseConfig

	Host      string
	Port      int
	PoolSize  int // Size of poll
	QueueSize int // Size of Queue
	LogMode   bool
}

func (this *ServerConfig) Name() string { // {{{
	return CONFIG_SERVER
} // }}}

func (this *ServerConfig) Addr() string { // {{{
	return fmt.Sprintf("%s:%d", this.Host, this.Port)
} // }}}
