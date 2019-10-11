package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
	"newsWeb/models"
	"math"
	"mime/multipart"
	"path"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowIndex() {
	pageIndex, err := this.GetInt("pageIndex")
	typeName:=this.GetString("select")
	if err != nil {
		pageIndex = 1
	}
	errmsg := this.GetString("errmsg")
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	articles := new([]models.Article)
	//qs.All(articles)
	var count int64
	if typeName==""{
		count, _ = qs.RelatedSel("ArticleType").Count()
	}else {
		count, _ = qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count()
	}

	pageSize := 2
	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))
	if typeName==""{
		qs.Limit(pageSize, (pageIndex-1)*pageSize).RelatedSel("ArticleType").All(articles)
	}else {
		qs.Limit(pageSize, (pageIndex-1)*pageSize).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(articles)
	}
	this.Data["count"] = count
	this.Data["pageCount"] = pageCount
	this.Data["articles"] = articles
	this.Data["pageIndex"] = pageIndex
	articleTypes:=new([]models.ArticleType)
	o.QueryTable("ArticleType").All(articleTypes)
	this.Data["articleTypes"]=articleTypes
	this.Data["errmsg"] = errmsg
	this.Data["typeName"]=typeName
	this.TplName = "index.html"
}
func (this *ArticleController) ShowAddArticle() {
	o := orm.NewOrm()
	articleTypes := new([]models.ArticleType)
	o.QueryTable("ArticleType").All(articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "add.html"
}
func (this *ArticleController) HandleAddArticle() {
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	typeName := this.GetString("select")
	_, header, err := this.GetFile("uploadname")
	if articleName == "" || content == "" || err != nil {
		fmt.Println("获取数据错误", err)
		this.TplName = "add.html"
		return
	}
	ext := FileCheck(header, this)
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
	articleType := new(models.ArticleType)
	articleType.TypeName = typeName
	err = o.Read(articleType, "TypeName")
	if err != nil {
		this.Redirect("/index", 302)
		return
	}
	article.ArticleType = articleType
	article.Image = "static/img/" + fileName + ext
	o.Insert(article)
	this.Redirect("/index", 302)
}
func (this *ArticleController) ShowContent() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Redirect(`/index?errmsg="获取文章失败"`, 302)
		return
	}
	o := orm.NewOrm()
	article := new(models.Article)
	article.Id = id
	err = o.Read(article, "Id")
	if err != nil {
		this.Redirect(`/index?errmsg="获取文章失败"`, 302)
		return
	}
	this.Data["article"] = article
	this.TplName = "content.html"
	article.ReadCount += 1
	o.Update(article)
}
func (this *ArticleController) ShowEditContent() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Redirect(`/index?errmsg="Id获取错误"`, 302)
		return
	}
	o := orm.NewOrm()
	article := new(models.Article)
	article.Id = id
	err = o.Read(article)
	if err != nil {
		this.Redirect(`/index?errmsg="文章未查询到!"`, 302)
		return
	}
	this.Data["article"] = article
	this.TplName = "update.html"
}
func (this *ArticleController) EditContent() {
	id, _ := this.GetInt("id")
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	_, header, err := this.GetFile("uploadname")
	if articleName == "" || content == "" || err != nil {
		this.Data["errmsg"] = "文章未查询到1"
		this.TplName = "update.html"
		return
	}
	ext := FileCheck(header, this)
	o := orm.NewOrm()
	article := new(models.Article)
	article.Id = id
	err = o.Read(article)
	if err != nil {
		this.Data["errmsg"] = "文章未查询到2"
		this.TplName = "update.html"
		return
	}
	article.Title = articleName
	article.Content = content
	fileName := time.Now().Format("20060102150405")
	err = this.SaveToFile("uploadname", "static/img/"+fileName+ext)
	if err != nil {
		fmt.Println("this.SaveToFile err:", err)
		this.Data["errmsg"] = "文件存储失败,请重新上传"
		this.TplName = "update.html"
		return
	}
	article.Image = "static/img/" + fileName + ext
	o.Update(article)
	this.Redirect("/index", 302)
}
func (this *ArticleController) DeleteContent() {
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	article := new(models.Article)
	article.Id = id
	err := o.Read(article)
	if err != nil {
		this.Redirect("/index?errmsg=删除文章失败1", 302)
		return
	}
	_, err = o.Delete(article)
	if err != nil {
		this.Redirect("/index?errmsg=删除文章失败2", 302)
		return
	}
	this.Redirect("/index", 302)
}
func (this *ArticleController) ShowAddType() {
	articleTypes := new([]models.ArticleType)
	o := orm.NewOrm()
	qs := o.QueryTable("ArticleType")
	qs.All(articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "addType.html"
}
func (this *ArticleController) HandleAddType() {
	typeName := this.GetString("typeName")
	if typeName == "" {
		this.Redirect("/addType", 302)
		return
	}
	o := orm.NewOrm()
	articleType := new(models.ArticleType)
	articleType.TypeName = typeName
	_, err := o.Insert(articleType)
	if err != nil {
		fmt.Println("o.Insert err:", err)
		this.Redirect("/addType", 302)
		return
	}
	this.Redirect("/addType", 302)
}

//工具类
func FileCheck(header *multipart.FileHeader, this *ArticleController) string {
	if header.Size > 1000000 {
		this.Data["errmsg"] = "图片过大,请重新上传,限制在1000000字节以内"
		this.TplName = "add.html"
		return ""
	}
	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式不正确,请重新上传"
		this.TplName = "add.html"
		return ""
	}
	return ext
}
