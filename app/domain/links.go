package domain

type LinksStorer interface {
	GetItem(req *Link) (resp *Link, err error)
	SetItem(req *Link) error
}

type Link struct {
	Short string
	Link  string
}
