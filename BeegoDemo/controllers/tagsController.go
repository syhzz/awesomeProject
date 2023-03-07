package controllers

import (
	"BeegoDemo/models"
	"fmt"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	tagsMap := models.HandleTagsListData(tags)
	fmt.Println(tagsMap)
	this.Data["Tags"] = tagsMap
	this.TplName = "tags.html"
}
