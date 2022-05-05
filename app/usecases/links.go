package usecases

import (
	"github.com/I-Limar/link_shortening/app/domain"
	"math/rand"
)

type LinksInteractor struct {
	repo domain.LinksStorer
	host string
	port int
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Link struct {
	Short  string
	Link   string
	Status string
	Host   string
	Port   int
}

func NewLinksInteractor(repo domain.LinksStorer, host string, port int) *LinksInteractor {
	return &LinksInteractor{
		repo: repo,
		host: host,
		port: port,
	}
}

func (l LinksInteractor) GetLink(req *Link) (*domain.Link, error) {
	request := &domain.Link{
		Short: req.Short,
	}
	res, err := l.repo.GetItem(request)
	response := domain.Link{Link: res.Link}

	return &response, err
}

func (l LinksInteractor) SetLink(req *Link) error {
	request := &domain.Link{
		Short: req.Short,
		Link:  req.Link,
	}
	err := l.repo.SetItem(request)
	return err
}

func (l LinksInteractor) Shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (l LinksInteractor) DefaultLink() *Link {
	return &Link{
		Host: l.host,
		Port: l.port,
	}
}
