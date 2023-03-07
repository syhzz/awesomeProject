package controllers

import (
	"BeegoDemo/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	tag := this.GetString("tag")
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var articleList []models.Article
	// 查询的总数量
	var num int
	if len(tag) > 0 {
		//按照指定的标签搜索
		articleList, _ = models.QueryArticlesWithTag(tag, page)
		num = len(articleList)
	} else {
		articleList, _ = models.FindArticleWithPage(page)
		num = models.QueryArticleRowNum()
	}
	this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page, num)
	this.Data["HasFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.LoginUser)
	this.Data["content"] = models.MakeHomeBlocks(articleList, this.IsLogin)
	this.TplName = "home.html"
}
