package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/reversersed/taskservice/pkg/postgres"
)

type ServerConfig struct {
	Environment string `env:"SERVICE_ENVIRONMENT" env-description:"Service environment" env-default:"debug"`
	Url         string `env:"SERVICE_HOST_URL" env-required:"true" env-description:"Server listening address"`
	Port        int    `env:"SERVICE_HOST_PORT" env-required:"true" env-description:"Server listening port"`
}
type Config struct {
	Database *postgres.DatabaseConfig
	Server   *ServerConfig
}

var cfg *Config
var once sync.Once

func Load(envPath string) (*Config, error) {
	var e error
	once.Do(func() {
		server := new(ServerConfig)
		database := new(postgres.DatabaseConfig)

		if err := cleanenv.ReadConfig(envPath, server); err != nil {
			desc, _ := cleanenv.GetDescription(cfg, nil)

			e = fmt.Errorf("%v: %s", err, desc)
			return
		}
		if err := cleanenv.ReadConfig(envPath, database); err != nil {
			desc, _ := cleanenv.GetDescription(cfg, nil)

			e = fmt.Errorf("%v: %s", err, desc)
			return
		}

		cfg = &Config{
			Server:   server,
			Database: database,
		}
	})
	if e != nil {
		return nil, e
	}
	return cfg, nil
}
