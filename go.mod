module github.com/ququgou-shop

go 1.13

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.36.0
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go v0.0.0-20190204201341-e444a5086c43
	golang.org/x/build => github.com/golang/build v0.0.0-20190307215223-c78805dbabc8
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190228161510-8dd112bcdc25
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190306152737-a1d7652674e8
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net => github.com/golang/net v0.0.0-20190301231341-16b79f2e4e95
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190306144031-151b6387e3f2
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190308023053-584f3b12f43e
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190308142131-b40df0fb21c3
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.1.0
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.0
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bradfitz/gomemcache v0.0.0-20190329173943-551aad21a668
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/kr/text v0.2.0 // indirect
	github.com/olivere/elastic/v7 v7.0.12
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.4.3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/silenceper/wechat v2.0.1+incompatible
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/tools v0.0.0-20200507152607-625332f3c5da // indirect
	gopkg.in/yaml.v2 v2.2.8
	qiniupkg.com/x v7.0.8+incompatible // indirect
)
