package elasticClientBase

import (
	"fmt"

	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/service/elasticsearchService"
)

var ElasticClient *elasticsearchService.ElasticClient

func init() {
	fmt.Printf("elasticClient init")

	opts := &elasticsearchService.ElasticConfigModel{
		Host: config.Config.Elastic.Host,
	}

	ElasticClient = elasticsearchService.NewElasticClient(opts)
}
