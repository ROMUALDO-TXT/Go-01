package infrastructure

import (
	"net/http"
	"os"

	interfaces "github.com/ROMUALDO-TXT/Go-01/app/interface"
	controller "github.com/ROMUALDO-TXT/Go-01/app/interface/controllers"
	"github.com/gorilla/mux"
)

func HandleRequests(sqlHandler interfaces.SQLHandler, logger interfaces.Logger) {
	router := mux.NewRouter().StrictSlash(true)
	homepageController := controller.NewHomepageController(sqlHandler, logger)
	articleController := controller.NewArticleController(sqlHandler, logger)

	router.HandleFunc("/", homepageController.Index)
	router.HandleFunc("/articles", articleController.Index)
	router.HandleFunc("/article", articleController.Save).Methods("POST")
	router.HandleFunc("/article/{id}", articleController.Show)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router); err != nil {
		logger.LogError("%s", err)
	}

}
