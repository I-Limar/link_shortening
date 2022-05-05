package geteways

import (
	"github.com/I-Limar/link_shortening/app/services/validators"
	"github.com/I-Limar/link_shortening/app/usecases"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func (g *GateWeb) IndexPage(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("app/templates/index.html")
	var link usecases.Link

	if r.Method == "POST" {
		if !validators.IsValidUrl(r.FormValue("s")) {
			link.Status = "Ссылка имеет неправильный формат!"
			link.Link = ""
		} else {
			link.Link = r.FormValue("s")
			link.Short = g.links.Shorting()
			err := g.links.SetLink(&link)

			if err != nil {
				logrus.Error(err)
				link.Status = "Произошла внутренняя ошибка"
				link.Link = ""
				link.Short = ""
				templ.Execute(w, link)
				return
			}
			link.Status = "Сокращение было выполнено успешно"
		}
	}
	templ.Execute(w, link)
}
