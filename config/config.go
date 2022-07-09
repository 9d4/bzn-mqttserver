package config

import (
	"github.com/spf13/viper"
)

var (
	Cfg    *Config
	EnvCfg *EnvConfig
)

func init() {
	Cfg = New()
	EnvCfg = NewEnv()
}

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

type EnvConfig struct {
	// Data source name for GORM. see https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	DSN string
	v   *viper.Viper
}

// create new EnvConfig
func NewEnv() *EnvConfig {
	ec := &EnvConfig{
		v: viper.New(),
	}

	ec.v.SetConfigType("env")
	ec.v.SetConfigName(".env")
	ec.v.AddConfigPath(".")
	ec.v.AutomaticEnv()
	ec.v.ReadInConfig()

	return ec
}

func LoadEnv() error { return EnvCfg.LoadConfig() }

func (ec *EnvConfig) LoadConfig() error {
	if err := ec.v.ReadInConfig(); err != nil {
		return err
	}

	if err := ec.v.Unmarshal(ec); err != nil {
		return err
	}

	return nil
}
