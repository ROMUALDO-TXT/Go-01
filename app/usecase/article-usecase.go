package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ROMUALDO-TXT/Go-01/app/domain"
)

var Articles domain.Articles

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	Articles = []domain.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}
