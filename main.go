package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/zerodollar/mall-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 使用flag，需要命令行传入，但是docker无法通过命令行传入
// 1.1 通过 －e传入环境变量, 镜像中需要sh，
// 1.2 通过 －e传入环境变量, 用os.Getenv获取，命令行无法获取，不能用flag
// 2. 通过 docker run 后参数，可以支持flag
// 3. 通过 卷管理挂接配置文件，从配置文件读取

func init() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("ARGV", i, os.Args[i])
	}
	dburl := flag.String("db", "root:rootpw@tcp(127.0.0.1:3306)/mall", "mysql url")
	flag.Parse()
	fmt.Println(*dburl)
	orm.RegisterDataBase("default", "mysql", *dburl)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
