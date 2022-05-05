package geteways

import (
	"github.com/I-Limar/link_shortening/app/usecases"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (g *GateWeb) RedirectTo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	request := usecases.Link{
		Short: vars["short"],
	}

	resp, err := g.links.GetLink(&request)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Страницы по данной ссылке не существует"))
		return
	}
	http.Redirect(w, r, resp.Link, http.StatusSeeOther)
}
