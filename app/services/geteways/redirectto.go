package geteways

import (
	"database/sql"
	"errors"
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
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Для данной страницы не создавалась короткая ссылка", http.StatusNotFound)
		} else {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

		logrus.Error(err)
		return
	}

	http.Redirect(w, r, resp.Link, http.StatusSeeOther)
}
