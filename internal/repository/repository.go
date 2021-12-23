package repository

type Repository interface {
	GetLink(shortLink string) (string, error)
	GetShortLink(link string) (string, error)
	AddLink(link string) error
}
