package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
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

func NewConfig() (*Config, error) {
	if err := InitConfig(); err != nil {
		return nil, err
	}
	return &Config{
			Host:        viper.GetString("host"),
			Port:        viper.GetString("port"),
			Username:    viper.GetString("username"),
			Password:    viper.GetString("password"),
			DbName:      viper.GetString("dbName"),
			Timeout:     viper.GetInt("timeout"),
			Connections: viper.GetInt32("connections"),
		},
		nil
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("configs/dbconfig.yml")
	return viper.ReadInConfig()
}
