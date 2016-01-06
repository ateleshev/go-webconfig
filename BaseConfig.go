package webconfig

// [BaseConfig]

type BaseConfig struct {
	ConfigInterface

	loaded bool
}

func (this *BaseConfig) setLoaded(loaded bool) { // {{{
	this.loaded = loaded
} // }}}

func (this *BaseConfig) IsLoaded() bool { // {{{
	return this.loaded
} // }}}
