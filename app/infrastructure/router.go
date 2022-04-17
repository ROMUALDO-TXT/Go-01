package infrastructure

import (
	"net/http"
	"os"

	"github.com/ROMUALDO-TXT/Go-01/app/interfaces"
	"github.com/ROMUALDO-TXT/Go-01/app/usecase"
	"github.com/gorilla/mux"
)

func HandleRequests(logger interfaces.Logger) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", usecase.HomePage)
	router.HandleFunc("/articles", usecase.ReturnAllArticles)
	router.HandleFunc("/article/{id}", usecase.ReturnSingleArticle)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router); err != nil {
		logger.LogError("%s", err)
	}

}
