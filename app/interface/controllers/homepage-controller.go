package controller

import (
	"fmt"
	"net/http"

	i "github.com/ROMUALDO-TXT/Go-01/app/interface"
	interactor "github.com/ROMUALDO-TXT/Go-01/app/usecase/interactors"
)

type HomepageController struct {
	HomepageInteractor interactor.HomepageInteractor
	Logger             i.Logger
}

func NewHomepageController(_ i.SQLHandler, logger i.Logger) *HomepageController {
	return &HomepageController{
		HomepageInteractor: interactor.HomepageInteractor{},
		Logger:             logger,
	}
}

func (hp *HomepageController) Index(w http.ResponseWriter, r *http.Request) {
	hp.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	msg := hp.HomepageInteractor.Index()

	fmt.Fprintf(w, msg)
}
