package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ROMUALDO-TXT/Go-01/app/domain"
	"github.com/gorilla/mux"
)

var articles = []domain.Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticles")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range articles {
		if article.Id == key {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
		}
	}

}
