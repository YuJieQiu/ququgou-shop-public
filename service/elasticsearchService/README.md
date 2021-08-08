# Elastic Search 模块
    通过Elastic 保存商品信息。并且利用其特性进行 分词 搜索商品信息等。
## go  Elastic Search 7.0 安装
    export  GOSUMDB=off
    
    export  GOSUMDB=sum.golang.google.cn
    
    export GOPROXY=https://goproxy.cn
    
    export GO111MODULE=on
    
    go get github.com/olivere/elastic/v7

#####  安装
     基于java 需要安装JDK,且对内存有要求，最低8G内存起步
     官网下载 https://www.elastic.co/cn/products/elasticsearch
     教程 https://www.elastic.co/guide/cn/elasticsearch/guide/cn/running-elasticsearch.html#sense
     Linux 安装 java  https://linuxize.com/post/install-java-on-centos-7/#prerequisites
     Jdk 官网地址 https://www.oracle.com/technetwork/java/javase/downloads/jdk12-downloads-5295953.html
     只要运行时 jre 就可以了： https://www.oracle.com/technetwork/java/javase/downloads/jre8-downloads-2133155.html

### 使用
    golang 使用 开源库(github.com/olivere/elastic/v7) 进行操作Elastic
    
### 方法 
######(暂时只是下面几个方法)    
    - 创建
    - 更新
    - 删除
    - Get 根据Idnex ID 获取, 
    - 精准字段搜索 termQuery
    - 字段全文搜索 matchQuery
    - 所有字段全文搜索 stringQuer   y