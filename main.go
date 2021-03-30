package main

import (
	"fengzigneg/godocker/common"
	"fengzigneg/godocker/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = routes.CollectRoute(r)
	panic(r.Run(":1111")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

