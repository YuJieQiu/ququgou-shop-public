package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/api/userProgram/config"
)

var (
	dbConfig  = config.Config.DB
	mysqlConn *gorm.DB
	err       error
)

func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	fmt.Println(connectionString)

	mysqlConn, err = gorm.Open(config.Config.DB.Driver, connectionString)

	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()

	if err != nil {
		panic(err)
	}
	mysqlConn.LogMode(true)
}

//TODO:连接 资源 分配 待检查
func MysqlConn() *gorm.DB {

	if mysqlConn == nil {
		setupMysqlConn()
	}

	return mysqlConn
}
