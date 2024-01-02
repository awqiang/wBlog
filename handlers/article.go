package handlers

import (
	"github.com/awqiang/wBlog/service"
	"github.com/awqiang/wBlog/status"
	"github.com/awqiang/wBlog/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //要校验

	articleService := service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}
	if !exists {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	article, err := articleService.Get()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}
	util.Success(c, article)

}

func GetArticles(c *gin.Context) {
	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
	}

	page := 1
	if arg := c.Query("page"); arg != "" {
		page, _ = strconv.Atoi(arg)
	}
	c.Query("page")

	articleService := service.Article{
		State:    state,
		PageNum:  page,
		PageSize: 10,
	}

	total, err := articleService.Count()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total
	util.Success(c, data)
}

type AddArticleForm struct {
	Title     string `form:"title" valid:"Required;MaxSize(100)"`
	Content   string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

func AddArticle(c *gin.Context) {
	var (
		form AddArticleForm
	)

	articleService := service.Article{
		Title:     form.Title,
		Content:   form.Content,
		State:     form.State,
		CreatedBy: form.CreatedBy,
	}
	if err := articleService.Add(); err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

type EditArticleForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Title      string `form:"title" valid:"Required;MaxSize(100)"`
	Content    string `form:"content" valid:"Required;MaxSize(65535)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //要校验
	var (
		form = EditArticleForm{ID: id}
	)

	articleService := service.Article{
		ID:         form.ID,
		Title:      form.Title,
		Content:    form.Content,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}
	if !exists {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	err = articleService.Edit()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //要校验

	articleService := service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}
	if !exists {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	err = articleService.Delete()
	if err != nil {
		util.Error(c, int(status.ApiCode.FAILED), status.ApiCode.GetMessage(status.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}
