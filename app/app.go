package app

import (
	"fmt"
	"github.com/I-Limar/link_shortening/app/db"
	"github.com/I-Limar/link_shortening/app/services/geteways"
	"github.com/I-Limar/link_shortening/app/services/repositories"
	"github.com/I-Limar/link_shortening/app/usecases"
	"github.com/gorilla/mux"
	"net/http"
)

type Config struct {
	DB       db.DBConfig
	HTTPPort int `envconfig:"http_port" required:"true"`
}

func NewApp(cfg *Config) (http.Handler, error) {

	dbService, err := db.NewDBService(cfg.DB)

	gateWeb := annexBuild(&dbService)

	if err != nil {
		return nil, fmt.Errorf("create connection to DB: %v", err)
	}

	serverMux := mux.NewRouter()
	serverMux.HandleFunc("/", gateWeb.IndexPage)
	serverMux.HandleFunc("/to/{short}", gateWeb.RedirectTo)

	return serverMux, nil
}

func annexBuild(dbService *db.DBService) *geteways.GateWeb {
	// create repositories
	linkRepo := repositories.NewLinksRepo(dbService)

	// create interactor
	linksInteract := usecases.NewLinksInteractor(linkRepo)

	getWeb := geteways.NewGateWeb(linksInteract)

	return getWeb
}
