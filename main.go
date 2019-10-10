package main

import (
	_ "newsWeb/routers"
	"github.com/astaxie/beego"
	_ "newsWeb/models"
)

func main() {
	beego.AddFuncMap("prePage", ShowPrepage)
	beego.AddFuncMap("nextPage", ShowNextPage)
	beego.Run()
}

func ShowPrepage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}
func ShowNextPage(pageIndex, pageCount int) int {
	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1

}
