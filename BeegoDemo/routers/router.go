package routers

import (
	"BeegoDemo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	beego.Router("article/update", &controllers.UpdateArticleController{})
	beego.Router("article/delete", &controllers.DeleteArticleController{})
	//标签
	beego.Router("/tags", &controllers.TagsController{})
	beego.Router("/album", &controllers.AlbumController{})
	beego.Router("/upload", &controllers.UploadFileController{})
	beego.Router("/aboutme", &controllers.AboutMeController{})

}
