package repository

import (
	"github.com/ROMUALDO-TXT/Go-01/app/domain"
)

type ArticleRepository interface {
	FindAll() (domain.Articles, error)
	FindById(id string) (domain.Article, error)
	Save(a domain.Article) (domain.Article, error)
	//Update
	//Delete
}
