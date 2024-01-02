package service

import "github.com/awqiang/wBlog/models"

type Article struct {
	ID         int
	Title      string
	Content    string
	State      int
	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) Get() (*models.Article, error) {
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles []*models.Article
	)

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}

	return maps
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"title":      a.Title,
		"content":    a.Content,
		"created_by": a.CreatedBy,
		"state":      a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"title":       a.Title,
		"content":     a.Content,
		"state":       a.State,
		"modified_by": a.ModifiedBy,
	})
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}
