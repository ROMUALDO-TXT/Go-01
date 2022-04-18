package repository

import (
	"github.com/ROMUALDO-TXT/Go-01/app/domain"
)

type ArticleRepository interface {
	FindAll() (domain.Articles, error)
	FindById(id string) (domain.Article, error)
	//Save(domain.Article) (int64, error)
	//Update
	//Delete
}
