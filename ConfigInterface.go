package webconfig

// [ConfigInterface]

type ConfigInterface interface {
	Name() string
	setLoaded(bool)
	IsLoaded() bool
}
