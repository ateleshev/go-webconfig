package webconfig

// [MainConfig]

func NewMainConfig() *MainConfig { // {{{
	return &MainConfig{}
} // }}}

type MainConfig struct {
	BaseConfig

	AppName    string // Name of application
	ServerType string // HTTP|FCGI
	UseThread  bool
	MaxProcs   int // [-1, 0, 1, ..], -1: used runtime.NumCPU(), 0: ignored this parameter
}

func (this *MainConfig) Name() string { // {{{
	return CONFIG_MAIN
} // }}}

func (this *MainConfig) IsServerType(serverType string) bool { // {{{
	return this.ServerType == serverType
} // }}}
