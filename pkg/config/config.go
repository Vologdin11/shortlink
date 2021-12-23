package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DbName      string
	Timeout     int
	Connections int32
}

func NewPoolConfig(cfg *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	return poolConfig, nil
}

func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	connect, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	return connect, nil
}

func NewConfig(host, port, username, password, dbName string, timeout int, connections int32) *Config {
	return &Config{
		Host:        host,
		Port:        port,
		Username:    username,
		Password:    password,
		DbName:      dbName,
		Timeout:     timeout,
		Connections: connections,
	}
}
