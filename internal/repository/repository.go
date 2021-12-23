package repository

import "context"

type Repository interface {
	GetLink(ctx context.Context, shortLink string) (string, error)
	GetShortLink(link string) (string, error)
	AddLink(link string) error
}
