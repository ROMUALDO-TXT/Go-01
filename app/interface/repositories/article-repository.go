package repository

import (
	"github.com/ROMUALDO-TXT/Go-01/app/domain"
	interfaces "github.com/ROMUALDO-TXT/Go-01/app/interface"
)

type ArticleRepository struct {
	SQLHandler interfaces.SQLHandler
}

var articles = []domain.Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func (pr *ArticleRepository) FindAll() (articles domain.Articles, err error) {
	return articles, nil
}

func (pr *ArticleRepository) FindById(id string) (article domain.Article, err error) {

	for _, article := range articles {
		if article.Id == id {
			return article, nil
		}

	}

	return
}

// func (pr *ArticleRepository) CreateArticle(w http.ResponseWriter, r *http.Request) {

// 	reqBody, err := ioutil.ReadAll(r.Body)

// 	if err

// }
