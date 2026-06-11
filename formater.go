package logxgo

type Option func(*Config)

func WithLevel(level string) Option {
	return func(config *Config) {
		config.Level = level
	}
}

func WithModule(module string) Option {
	return func(config *Config) {
		config.Module = module
	}
}

func WithJSON() Option {
	return func(config *Config) {
		config.JSON = true
	}
}