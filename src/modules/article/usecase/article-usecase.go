package usecase

import (
	"fmt"
	"json"
	"net/http"

	"./modules/article/model"
)

var Articles []model.Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	Articles = []model.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	json.NewEncoder(w).Encode(Articles)
}
