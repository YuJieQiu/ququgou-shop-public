package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/adminuser"
	"github.com/ququgou-shop/modules/appConfig"
	"github.com/ququgou-shop/modules/banner"
	"github.com/ququgou-shop/modules/label"
	"github.com/ququgou-shop/modules/merchant"
	"github.com/ququgou-shop/modules/order"
	"github.com/ququgou-shop/modules/payment"
	"github.com/ququgou-shop/modules/product"
	"github.com/ququgou-shop/modules/resource"
	"github.com/ququgou-shop/modules/shopcart"
	"github.com/ququgou-shop/modules/user"

	_ "github.com/go-sql-driver/mysql"
)

var Config struct {
	DB struct {
		Driver   string `default:"mysql"`
		Host     string `default:"127.0.0.1"`
		Port     string `default:"3306"`
		Name     string `default:"lbs"`
		User     string `default:"root"`
		Password string `default:"test123"`
	}
}

func init() {
	var (
		yamlFile []byte
		err      error
	)
	//GOPRODUCTWEB
	//获取运行环境
	runenv := os.Getenv("GOPRODUCTWEB")
	if runenv == "TEST" {
		//获取当前目录下的配置
		yamlFile, err = ioutil.ReadFile("conf.yaml")

	} else {
		p := filepath.FromSlash("/src/github.com/ququgou-shop/cmd/conf.yaml")

		//TODO:test
		gopath := os.Getenv("GOPATH")

		yamlFile, err = ioutil.ReadFile(gopath + p)
	}

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	//configor.Load(&Config, "conf.yml")
	fmt.Printf("config: %#v", Config)
}

var (
	mysqlConn *gorm.DB
	err       error
)

func setupMysqlConn() {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8", Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Name)
	fmt.Println(connectionString)

	mysqlConn, err = gorm.Open(Config.DB.Driver, connectionString)

	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()

	if err != nil {
		panic(err)
	}
	mysqlConn.LogMode(true)
}

func MysqlConn() *gorm.DB {

	if mysqlConn == nil {
		setupMysqlConn()
	}

	return mysqlConn
}

//初始化表
func main() {
	var input string = ""
	fmt.Print("input init table y/n?")
	for input != "exit" {
		var inputReader *bufio.Reader
		inputReader = bufio.NewReader(os.Stdin)
		input, _ = inputReader.ReadString('\n')
		input = strings.Replace(input, " ", "", -1)  //删除空格
		input = strings.Replace(input, "\n", "", -1) // 去除换行符
		if strings.ToLower(input) == "y" {
			initTable()
		}
	}
	fmt.Print("end")
}

func initTable() {

	var db *gorm.DB

	db = MysqlConn()

	tx := db.Begin()

	var moduleAdminUser = adminuser.ModuleAdminUser{DB: tx}
	var moduleAppConfig = appConfig.ModuleAppConfig{DB: tx}
	var moduleBanner = banner.ModuleBanner{DB: tx}
	var moduleMerchant = merchant.ModuleMerchant{DB: tx}
	var moduleOrder = order.ModuleOrder{DB: tx}
	var modulePayment = payment.ModulePayment{DB: tx}
	var moduleProduct = product.ModuleProduct{DB: tx}
	var moduleResource = resource.ModuleResource{DB: tx}
	var moduleShopCart = shopcart.ModuleShopCart{DB: tx}
	var moduleUser = user.ModuleUser{DB: tx}
	var moduleLabel = label.ModuleLabel{DB: tx}
	var err error

	err = moduleAdminUser.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleAppConfig.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleBanner.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleMerchant.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleOrder.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = modulePayment.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleProduct.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleResource.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleShopCart.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleUser.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	err = moduleLabel.CreateTable()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	tx.Commit()

	fmt.Println("success!")
}
