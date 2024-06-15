package repository

import (
	"github.com/azeek21/blog/models"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepositroy(db *gorm.DB) ArticleRepository {
	return ArticleRepo{
		db: db,
	}
}

func (r ArticleRepo) GetArticleById(articleId uint) (*models.Article, error) {
	res := &models.Article{
		Model: gorm.Model{ID: articleId},
	}
	err := r.db.First(res).Error
	return res, err
}

func (r ArticleRepo) CreateArticle(article *models.Article) (*models.Article, error) {
	err := r.db.Create(article).Error
	return article, err
}

func (r ArticleRepo) Update(article *models.Article) (*models.Article, error) {
	err := r.db.Save(article).Error
	return article, err
}
func (r ArticleRepo) Delete(article *models.Article) (bool, error) {
	article.DeletedAt = gorm.DeletedAt{}
	err := r.db.Delete(article).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
