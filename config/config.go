package config

type Config struct {
	Port                         int    `env:"PORT"`
	PostgresqlUser               string `env:"PG_USER"`
	PostgresqlPassword           string `env:"PG_PASSWORD"`
	PostgresqlDriver             string `env:"PG_DRIVER"`
	PostgresqlDBName             string `env:"PG_NAME"`
	PostgresqlHost               string `env:"PG_HOST"`
	PostgresqlPort               int    `env:"PG_PORT"`
	PostgresqlSchema             string `env:"PG_SCHEMA"`
	PostgresqlMaxOpenConnections int    `env:"PG_MAX_OPEN_CONNS"`
	PostgresqlMaxIdleConnections int    `env:"PG_MAX_IDLE_CONNS"`
	PostgresqlConnectionMaxLife  int    `env:"PG_CONN_MAX_LIFE"`
}
