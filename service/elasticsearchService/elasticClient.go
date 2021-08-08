package elasticsearchService

import (
	"github.com/olivere/elastic/v7"
)

type ElasticClient struct {
	client *elastic.Client
}

//创建新的客户端连接
func NewElasticClient(conf *ElasticConfigModel) *ElasticClient {
	//如果 host 地址 不是 本地 127.0.0.0:9200 这种 使用 NewClient 创建是失败的 ，它会嗅探寻找节点
	//这个时候 使用 SimpleClient 创建
	//SimpleClient只是禁用了一些功能，例如通过我上面链接的“嗅探”过程自动查找添加到集群中的新节点。您可以通过Wiki中所述的NewClient中的选项禁用嗅探。
	//不确定您的设置有什么问题。群集输出看起来不错，嗅探应该选择http_address中的设置
	//参考 地址 https://github.com/olivere/elastic/issues/312

	client, err := elastic.NewSimpleClient(elastic.SetURL(conf.Host))
	//client, err := elastic.NewClient(elastic.SetURL(conf.Host))

	//elastic.SetSniff(false)
	if err != nil {
		// Handle error
		panic(err)
	}

	//TODO: Elasticsearch 版本判断(Ping the Elasticsearch server to get e.g. the version number)
	//ctx := context.Background()
	//_, _, err = client.Ping(conf.Host).Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}

	return &ElasticClient{client: client}
}

//判断连接是否成功
//func PingElasticsearch() {
//	ctx := context.Background()
//	info, code, err := client.Ping(conf.Host).Do(ctx)
//	if err != nil {
//		// Handle error
//		panic(err)
//	}
//}
