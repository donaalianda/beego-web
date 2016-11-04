package main

import (
	_ "beegoweb/beego-web/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.Run()
	fmt.Println("hello")

}
