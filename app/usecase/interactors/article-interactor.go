package interactor

import (
	"github.com/ROMUALDO-TXT/Go-01/app/domain"
	repository "github.com/ROMUALDO-TXT/Go-01/app/usecase/repositories"
)

type ArticleInteractor struct {
	ArticleRepository repository.ArticleRepository
}

func (ui *ArticleInteractor) Index() (articles domain.Articles, err error) {
	articles, err = ui.ArticleRepository.FindAll()

	return
}

func (ui *ArticleInteractor) Show(id string) (article domain.Article, err error) {
	article, err = ui.ArticleRepository.FindById(id)

	return
}

func (ai *ArticleInteractor) Save(a domain.Article) (article domain.Article, err error) {
	article, err = ai.ArticleRepository.Save(a)

	return
}
