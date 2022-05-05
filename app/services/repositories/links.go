package repositories

import (
	"github.com/I-Limar/link_shortening/app/db"
	"github.com/I-Limar/link_shortening/app/domain"
)

type LinksRepo struct {
	executor *db.Executor
}

func NewLinksRepo(service *db.DBService) *LinksRepo {
	return &LinksRepo{
		executor: db.NewExecutor(service, linkQuery),
	}
}

func (r *LinksRepo) GetItem(req *domain.Link) (*domain.Link, error) {
	var link string
	querySource := r.makeQuerySource()
	err := r.executor.GetRow(&link, "link", querySource, req.Short)
	resp := domain.Link{
		Link: link,
	}
	return &resp, err
}

func (r *LinksRepo) SetItem(req *domain.Link) error {
	querySource := r.makeQuerySource()
	err := r.executor.SetRow("setlink", querySource, req.Link, req.Short)

	return err
}

func (r *LinksRepo) makeQuerySource() map[string]interface{} {
	querySource := map[string]interface{}{
		"Tables": db.Tables,
	}

	return querySource
}

var linkQuery = `
{{- define "link" -}}
SELECT link FROM {{.Tables.Links}} WHERE short = $1
{{end}}
{{- define "setlink" -}}
INSERT into {{.Tables.Links}} (link, short) values ($1, $2)
{{end}}
`
