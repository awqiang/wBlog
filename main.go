package main

import (
	"github.com/awqiang/wBlog/models"
	"github.com/awqiang/wBlog/routers"
	"github.com/awqiang/wBlog/setting"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	r := routers.InitRouter()
	r.Run(":8080")
}
