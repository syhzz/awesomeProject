package controllers

import (
	"BeegoDemo/models"
	"fmt"
	"log"
)

type DeleteArticleController struct {
	BaseController
}

func (this *DeleteArticleController) Get() {
	artId, _ := this.GetInt64("id")
	fmt.Println("删除 ID：", artId)
	_, err := models.DeleteArticle(artId)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/", 302)
}
