package env

import "github.com/google/wire"

// ProvideConfig exports the EnvConfig for Wire.
func ProvideConfig() *Config {
	return EnvConfig
}

// Add this to your existing code
var EnvProviderSet = wire.NewSet(ProvideConfig)
