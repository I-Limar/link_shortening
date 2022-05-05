package geteways

import "github.com/I-Limar/link_shortening/app/usecases"

type GateWeb struct {
	links *usecases.LinksInteractor
}

func NewGateWeb(links *usecases.LinksInteractor) *GateWeb {
	return &GateWeb{
		links: links,
	}
}
