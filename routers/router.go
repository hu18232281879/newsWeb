package routers

import (
	"newsWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/index",&controllers.ArticleController{},"get:ShowIndex")
    beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")
    beego.Router("/content",&controllers.ArticleController{},"get:ShowContent")
    beego.Router("/update",&controllers.ArticleController{},"get:ShowEditContent;post:EditContent")
    beego.Router("/delete",&controllers.ArticleController{},"get:DeleteContent")
    beego.Router("/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")

    }
