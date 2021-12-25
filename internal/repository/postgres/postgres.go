package postgres

import (
	"context"
	conf "shortlink/pkg/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbGetLink      = "SELECT link FROM shortlink WHERE shortlink = $1"
	dbGetShortLink = "SELECT shortlink FROM shortlink WHERE link = $1"
	dbAddLink      = "INSERT INTO shortlink(link, shortlink) VALUES ($1, $2)"
)

type Postgres struct {
	db      *pgxpool.Pool
	context context.Context
}

func (p *Postgres) InitPostgres(cfg *conf.Config) error {
	p.context = context.Background()
	poolConfig, err := conf.NewPoolConfig(cfg)
	if err != nil {
		return err
	}
	poolConfig.MaxConns = cfg.Connections

	connect, err := conf.NewConnection(poolConfig)
	if err != nil {
		return err
	}
	//ping
	_, err = connect.Exec(p.context, ";")
	if err != nil {
		return err
	}
	p.db = connect
	return nil
}

func (p *Postgres) GetLink(shortLink string) (string, error) {
	var link string
	err := p.db.QueryRow(p.context, dbGetLink, shortLink).Scan(&link)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (p *Postgres) GetShortLink(link string) (string, error) {
	var shortlink string
	err := p.db.QueryRow(p.context, dbGetShortLink, link).Scan(&shortlink)
	if err != nil {
		return "", err
	}
	return shortlink, nil
}

func (p *Postgres) AddLink(link, shortlink string) error {
	_, err := p.db.Exec(p.context, dbAddLink, link, shortlink)
	return err
}
