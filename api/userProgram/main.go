package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/ququgou-shop/api/userProgram/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/api/userProgram/routers"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

var (
	g errgroup.Group
)

type configuration struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func (c *configuration) getConf() *configuration {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// @title USER WEB API
// @version 1.0
// @description This is a swag test http server.

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7001
// @BasePath /api/v1
// @query.collection.format multi
// @schemes http https
// @x-example-key {"key": "value"}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//设置Options
	//r.Use(middleware.Options)

	//设置 cors 解决跨域问题
	con := cors.DefaultConfig()
	con.AllowAllOrigins = true
	con.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Token"}
	//cors.Default()
	r.Use(cors.New(con))

	//设置 路由
	fmt.Println()
	routers.API(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(config.Config.API.Port)
}
