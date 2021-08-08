# ququgou-shop

### 基本信息介绍
- 使用Go开发的一套电商平台系统，主要功能模块有 **用户模块**、**管理员模块**、**商品模块**、**订单模块****、**支付模块**、**静态资源模块**、**缓存模块**、**搜索模块**、**微信接口相关模块**、**商户模块**等

### 项目的优点
- 

- 项目中前端项目如下:  
    - [x] [用户端小程序端项目地址](https://github.com/YuJieQiu/ququgou-xiaochengxu)
    - [x] [商户端小程序端项目地址](https://github.com/YuJieQiu/ququgou-merchant)
    - [ ] 后台管理项目

### 安装使用
- 运行 `go run main.go`

### 项目结构
- **[api](/api/README.md)**
> 对外提供的API接口，包括用户端和后台管理

- **[library](/library/README.md)**
> 项目通用库

- **[modules](/modules/README.md)**
> 项目模块，每个模块包含数据库 model，和一些简单的GURD

- **[service](/service/README.md)**
> 服务，主要业务逻辑

- **cmd**
- **vendor**
> 项目所有的依赖库