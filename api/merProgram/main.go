package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/api/merProgram/routers"
)

func main() {
	//router := gin.Default()
	//
	//router.StaticFS("/static", http.Dir("Eleditor"))
	//
	//router.LoadHTMLGlob("templates/*")
	////router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})

	r := gin.Default()

	//设置Options
	//r.Use(middleware.Options)

	//设置 cors 解决跨域问题
	con := cors.DefaultConfig()
	con.AllowAllOrigins = true
	con.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Token"}

	//cors.Default()
	r.Use(cors.New(con))

	//设置 路由

	routers.API(r)

	r.Run(config.Config.API.Port)
}
