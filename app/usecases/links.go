package usecases

import (
	"fmt"
	"github.com/I-Limar/link_shortening/app/domain"
	"math/rand"
)

type LinksInteractor struct {
	repo domain.LinksStorer
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Link struct {
	Short  string
	Link   string
	Status string
}

func NewLinksInteractor(repo domain.LinksStorer) *LinksInteractor {
	return &LinksInteractor{
		repo: repo,
	}
}

func (l LinksInteractor) GetLink(req *Link) (*domain.Link, error) {
	request := &domain.Link{
		Short: req.Short,
	}
	res, err := l.repo.GetItem(request)
	fmt.Println(res)
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
