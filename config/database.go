package config

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"

	"github.com/spf13/viper"
)

const (
	databaseHost                  = "DATABASE_HOST"
	databasePort                  = "DATABASE_PORT"
	databaseName                  = "DATABASE_NAME"
	databaseUser                  = "DATABASE_USER"
	databasePassword              = "DATABASE_PASSWORD"
)

func NewDatabaseConn(ctx context.Context) *pgx.Conn {
	conn, err := pgx.ConnectConfig(ctx, databaseConfig())
	if err != nil {
		panic(err)
	}
	return conn
}

func databaseConfig() *pgx.ConnConfig {
	cfg, err := pgx.ParseConfig("sslmode=require")
	if err != nil {
		panic(err)
	}

	cfg.Host = viper.GetString(databaseHost)
	cfg.Port = uint16(viper.GetInt(databasePort))
	cfg.Database = viper.GetString(databaseName)
	cfg.User = viper.GetString(databaseUser)
	cfg.Password = viper.GetString(databasePassword)
	cfg.ConnectTimeout = time.Second * 10
	return cfg
}
