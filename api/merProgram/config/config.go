package config

import (
	"fmt"
	"github.com/ququgou-shop/modules/resource"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
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
	JWT struct {
		Secret      string
		AdminSecret string
	}
	API struct {
		Name         string `default:"app name"`
		Port         string `default:"8080"`
		RelativePath string `json:"relativePath"`
		Version      string `json:"version"`
	}
	WX struct {
		AppID     string `required:"true"`
		AppSecret string `required:"true"`
	}
	ImgService struct {
		Url      string `json:"url"`
		SavePath string `json:"savepath"`
		QiniuUrl string `json:"qiniuurl"` //七牛云 图片utl
	}

	UploadConfigForImage resource.UploadConfigModel

	Cache struct {
		Enabled bool `json:"enabled"`
		Redis   struct {
			Host     string `json:"host"`
			Password string `json:"password"`
			Database int    `json:"database"`
		}
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
		p := filepath.FromSlash("/src/github.com/ququgou-shop/api/merProgram/conf.yaml")

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
