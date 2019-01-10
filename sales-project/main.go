package main

import (
	"github.com/astaxie/beego"
	"sales-project/models"
	//"sales-project/routers"
)

func main() {
	models.Init()
	beego.Run()
}
