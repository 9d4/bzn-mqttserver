package config

import (
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	Cfg  *Config
	Env *EnvConfig
)

type Config struct {
	// List of addresses that the server will listen to
	Addresses []string
}

type EnvConfig struct {
	DSN string
}

func LoadEnv(path string) {

}

func LoadConfig(paths ...string) error {
	Cfg = &Config{}
	c := viper.New()

	for _, p := range paths {
		c.AddConfigPath(p)
	}

	c.SetConfigType("yaml")
	
	c.SetDefault("addresses", []string{":1883"})

	if err := c.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			jww.FATAL.Fatal(err)
		}
	}

	if err := c.Unmarshal(Cfg); err != nil {
		jww.ERROR.Fatal(err)
	}

	return nil
}
