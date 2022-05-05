package geteways

import (
	"github.com/I-Limar/link_shortening/app/services/validators"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func (g *GateWeb) IndexPage(w http.ResponseWriter, r *http.Request) {
	link := g.links.DefaultLink()
	templ, err := template.ParseFiles("app/templates/index.html")

	if err != nil {
		logrus.Error(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		if !validators.IsValidUrl(r.FormValue("s")) {
			link.Status = "Неправильный формат ссылки"
		} else {
			//TODO тут нужна проверка на существование ссылки, если уже есть вернуть ее
			link.Link = r.FormValue("s")
			link.Short = g.links.Shorting(link.Link)
			err := g.links.SetLink(link)

			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}
			link.Status = "Сокращение было выполнено успешно."
		}
	}
	err = templ.Execute(w, link)
	if err != nil {
		logrus.Error(err)
	}
}
