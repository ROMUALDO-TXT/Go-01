package controller

import (
	"encoding/json"
	"net/http"

	i "github.com/ROMUALDO-TXT/Go-01/app/interface"
	repository "github.com/ROMUALDO-TXT/Go-01/app/interface/repositories"
	interactor "github.com/ROMUALDO-TXT/Go-01/app/usecase/interactors"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	ArticleInteractor interactor.ArticleInteractor
	Logger            i.Logger
}

func NewArticleController(sqlHandler i.SQLHandler, logger i.Logger) *ArticleController {
	return &ArticleController{
		ArticleInteractor: interactor.ArticleInteractor{
			ArticleRepository: &repository.ArticleRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

func (ac *ArticleController) Index(w http.ResponseWriter, r *http.Request) {
	ac.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	Articles, err := ac.ArticleInteractor.Index()
	if err != nil {
		ac.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}

func (ac *ArticleController) Show(w http.ResponseWriter, r *http.Request) {
	ac.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)

	Articles, err := ac.ArticleInteractor.Show(vars["id"])
	if err != nil {
		ac.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}
