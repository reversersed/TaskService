package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/reversersed/taskservice/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	file, _ := os.OpenFile(fmt.Sprintf("%s/.env", dir), os.O_CREATE|os.O_TRUNC, os.ModeAppend)

	file.WriteString(`
SERVICE_ENVIRONMENT = debug
SERVICE_HOST_URL = localhost
SERVICE_HOST_PORT = 1001

POSTGRES_HOST = db
POSTGRES_PORT = 1000
POSTGRES_PASSWORD = dbpass
POSTGRES_USER = root
POSTGRES_DB = base`)

	cfg, e := Load(fmt.Sprintf("%s/.env", dir))

	file.Close()

	assert.NoError(t, e)
	if assert.NotNil(t, cfg) {
		excepted := &Config{
			Database: &postgres.DatabaseConfig{
				Host:          "db",
				Port:          1000,
				Password:      "dbpass",
				User:          "root",
				Database:      "base",
				MigrationPath: "/migrations",
			},
			Server: &ServerConfig{
				Environment: "debug",
				Url:         "localhost",
				Port:        1001,
			},
		}

		assert.EqualValues(t, excepted, cfg)
	}
}
