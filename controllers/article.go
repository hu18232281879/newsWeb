package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"newsWeb/models"
	"math"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowIndex() {
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	articles := new([]models.Article)
	//qs.All(articles)
	count, _ := qs.Count()
	pageSize := 2
	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))
	qs.Limit(pageSize, (pageIndex-1)*pageSize).All(articles)
	this.Data["count"] = count
	this.Data["pageCount"] = pageCount
	this.Data["articles"] = articles
	this.Data["pageIndex"] = pageIndex
	this.TplName = "index.html"
}
func (this *ArticleController) ShowAddArticle() {
	this.TplName = "add.html"
}
func (this *ArticleController) HandleAddArticle() {
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	_, header, err := this.GetFile("uploadname")
	if articleName == "" || content == "" || err != nil {
		fmt.Println("获取数据错误", err)
		this.TplName = "add.html"
		return
	}
	if header.Size > 1000000 {
		this.Data["errmsg"] = "图片过大,请重新上传,限制在1000000字节以内"
		this.TplName = "add.html"
		return
	}
	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式不正确,请重新上传"
		this.TplName = "add.html"
		return
	}
	fileName := time.Now().Format("20060102150405")
	err = this.SaveToFile("uploadname", "static/img/"+fileName+ext)
	if err != nil {
		fmt.Println("this.SaveToFile err:", err)
		this.Data["errmsg"] = "文件存储失败,请重新上传"
		this.TplName = "add.html"
		return
	}
	o := orm.NewOrm()
	article := new(models.Article)
	article.Title = articleName
	article.Content = content
	article.Image = "static/img/" + fileName + ext
	o.Insert(article)
	this.Redirect("/index", 302)
}
func (this *ArticleController) ShowContent() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Redirect("/index", 302)
		return
	}
	o := orm.NewOrm()
	article := new(models.Article)
	article.Id = id
	err = o.Read(article, "Id")
	if err != nil {
		this.Redirect("/index", 302)
		return
	}
	this.Data["article"] = article
	this.TplName = "content.html"
	article.ReadCount += 1
	o.Update(article)
}
