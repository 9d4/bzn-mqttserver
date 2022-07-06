package config

import (
	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

type Config struct {
	// List of addresses that the server will listen to
	Addresses []string
	v         *viper.Viper
}

// create new Config
func New() *Config {
	c := &Config{
		Addresses: []string{":1883"},
		v:         viper.New(),
	}

	c.v.SetConfigType("yaml")

	return c
}

// find and load config from given paths
func LoadConfig(paths ...string) error {
	Cfg = New()

	if err := Cfg.LoadConfig(paths...); err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadConfig(paths ...string) error {
	for _, p := range paths {
		c.v.AddConfigPath(p)
	}

	if err := c.v.ReadInConfig(); err != nil {
		return err
	}

	if err := c.v.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
