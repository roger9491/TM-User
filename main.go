package main

import (
	// init
	_ "tm-user/init/configinit"
	_ "tm-user/init/loginit"
	_ "tm-user/init/sqlinit"

	"tm-user/global"
	"tm-user/global/database"

	"tm-user/init/sqlinit"

	"os"

	"log"
	"tm-user/init/routerinit"
	"tm-user/router/userrouter"
)

func main() {

	// 設置日誌輸出配置

	f, err := os.OpenFile("mrmaster_log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		// test
		log.Fatal("OpenFile is failed")
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("test test ")

	// 初始化資料庫
	sqlinit.InitMySQL(database.Username, database.Password, database.Host, database.Port, database.DBName)

	// 加載路由
	routerinit.Include(userrouter.UserApi)
	r := routerinit.InitRouters()

	err = r.Run(global.IP + ":" + global.Port) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Println("err ", err.Error())
	}
}
