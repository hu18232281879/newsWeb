package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}
type Article struct {
	Id        int       `orm:"pk;auto"`
	Title     string    `orm:"size(50);unique"`
	Content   string    `orm:"size(500)"`
	Time      time.Time `orm:"type(datetime);auto_now_add"`
	ReadCount int       `orm:"default(0)"`
	Image     string    `orm:"null"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/newsWeb")
	orm.RegisterModel(new(User),new(Article))
	orm.RunSyncdb("default", false, true)
}
