package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/bienbk/first_golang/routers"
)

func main() {
	beego.Run()
}
