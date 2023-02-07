package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	res := New()
	assert.Equal(t, res, Config{})
}

func TestSetupMasterDB(t *testing.T) {
	_ = SetupMasterDB(Config{
		Postgres: PostgresConfig{
			MaxOpenConnections: 4,
			MaxIdleConnections: 2,
			MaxIdleLifetime:    10000,
			Master: PSQL{
				Host:     "127.0.0.1",
				Port:     5432,
				Schema:   "public",
				DBName:   "aplus_incentive_db",
				User:     "amarthaplus",
				Password: "some-password",
			},
		},
	}, nil, true)
}

func TestSetupSlaveDB(t *testing.T) {
	_ = SetupSlaveDB(Config{
		Postgres: PostgresConfig{
			MaxOpenConnections: 4,
			MaxIdleConnections: 2,
			MaxIdleLifetime:    10000,
			Master: PSQL{
				Host:     "127.0.0.1",
				Port:     5432,
				Schema:   "public",
				DBName:   "aplus_incentive_db",
				User:     "amarthaplus",
				Password: "some-password",
			},
		},
	}, nil, true)
}

func TestSetupRedis(t *testing.T) {
	_ = SetupRedis(Config{
		Redis: Redis{
			Host:     "localhost:6379",
			Password: "",
		},
	})
}
