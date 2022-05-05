package usecases

import (
	"crypto/md5"
	"fmt"
	"github.com/I-Limar/link_shortening/app/domain"
)

type LinksInteractor struct {
	repo domain.LinksStorer
	host string
	port int
}

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

func (l LinksInteractor) Shorting(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func (l LinksInteractor) DefaultLink() *Link {
	return &Link{
		Host: l.host,
		Port: l.port,
	}
}
