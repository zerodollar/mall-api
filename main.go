package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	_ "github.com/zerodollar/mall-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
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

func getLocalIP() (error, string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err, ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return nil, ipnet.IP.String()
			}
		}
	}
	return fmt.Errorf("NO IP"), ""
}
func getTrustClient() []string {
	trustIP := []string{"http*", "http://127.0.0.1*", "http://172.17.42.1*", "http://172.31.25.56*", "http://54.245.112.93*"}
	if err, ip := getLocalIP(); err == nil {
		trustIP = append(trustIP, "http://"+ip+"*")
	}
	//TODO 从DB中获取信任的客户端

	trustIP = append(trustIP, os.Getenv("TRUST_URL"))
	fmt.Println(trustIP)
	return trustIP
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http*"}, //getTrustClient(), //  []string{"http://localhost*", "http://127*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}
