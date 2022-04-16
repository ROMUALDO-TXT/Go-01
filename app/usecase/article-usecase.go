package usecase

import (
	"app/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

var Articles domain.Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	Articles = []domain.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	json.NewEncoder(w).Encode(Articles)
}
