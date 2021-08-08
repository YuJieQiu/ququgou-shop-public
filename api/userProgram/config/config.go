package config

import (
	"fmt"
	"github.com/ququgou-shop/modules/resource"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//conf.yaml 配置文件信息
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
	WXUser struct {
		AppID     string `required:"true"`
		AppSecret string `required:"true"`
	}
	ImgService struct {
		Url      string `json:"url"`
		SavePath string `json:"savepath"`
		QiniuUrl string `json:"qiniuurl"` //七牛云 图片utl
	}

	UploadConfigForImage resource.UploadConfigModel

	TemplatePoster struct {
		InputUrl   string `json:"input_url"`                                          //url
		Format     string `json:"format" default:"jpg"`                               //输出图片文件格式
		BinaryPath string `json:"binary_path" default:"/usr/local/bin/wkhtmltoimage"` //wkhtmltoimage 路径

	}

	ProjectPath struct {
		Path string `json:"path"` //程序路径
	}

	Cache struct {
		Enabled bool `json:"enabled"`
		Redis   struct {
			Host     string `json:"host"`
			Password string `json:"password"`
			Database int    `json:"database"`
		}
	}

	Elastic struct {
		Host    string `json:"host"`    //TODO 这里可以是数组 集群
		Enabled bool   `json:"enabled"` //是否启用
	}

	Wxmpapi struct { //微信公众平台接口
		Url             string `json:"url"`
		TemplateSendUrl string `json:"templateSendUrl"`
		Enabled         bool   `json:"enabled"` //是否启用
	}
}

func init() {
	var (
		yamlFile []byte
		err      error
	)
	//TODO:待优化
	//这里的一个问题是，如果本地调试的话 配置文件不能根据绝对路径读取到
	//GOPRODUCTWEB
	//获取运行环境
	runenv := os.Getenv("GOPRODUCTWEB")
	if runenv == "TEST" {
		//获取当前目录下的配置
		yamlFile, err = ioutil.ReadFile("conf.yaml")

	} else {
		p := filepath.FromSlash("/src/github.com/ququgou-shop/api/userProgram/conf.yaml")

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
