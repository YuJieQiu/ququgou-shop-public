package elasticsearchService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// product 相关

const (
	ELASTIC_INDEX = "product"
	//ELASTIC_TYPE  = "product"
	REFRESH = "false" //refresh 参数 值 "wait_for" ,"false" default ,"" or "true" 使用默认值
)

//获取index 自定义 映射信息
func getIndexMapping() string {

	//1、这样
	//data := map[string]interface{}{
	//	"settings": map[string]interface{}{
	//		"number_of_shards": 1,
	//	},
	//}

	//或者
	//2、
	//type CreateIndexBody struct {
	//	Settings *CreateIndexSettings `json:"settings,omitempty"`
	//	...
	//}
	//
	//type CreateIndexSettings struct {
	//	NumberOfShards int `json:"number_of_shards,omitempty"`
	//	...
	//}
	//
	//mapping := &CreateIndexBody{
	//	Settings: &CreateIndexSettings{NumberOfShards: 1},
	//}
	//body, err := json.Marshal(mapping)
	//client.CreateIndex(...).Body(string(body)).Do()

	//3、

	//"settings":{
	//	"number_of_shards":1,
	//		"number_of_replicas":0
	//},
	mapping := `{
	"mappings":{
		"properties":{
				"name":{
					"type":"text",
					"analyzer": "ik_max_word"
				},
				"categoryInfo":{
					"type":"text",
					"analyzer": "ik_max_word"
				},
				"location":{"type":"geo_point"},
				"createdTime":{"type":"date","format": "yyyy-MM-dd HH:mm:ss"},
				"updatedTime":{"type":"date", "format": "yyyy-MM-dd HH:mm:ss"}
		}
	}
}`

	return mapping
}

//判断是否存在 Index,如果不存在  创建
func IndexExists(elastic *ElasticClient, ctx context.Context) (error, bool) {
	if ctx == nil {
		ctx = context.Background()
	}
	exists, err := elastic.client.IndexExists(ELASTIC_INDEX).Do(ctx)
	if err != nil {
		// Handle error
		//panic(err)
		return err, false
	}

	if !exists {
		// Create a new index.
		createIndex, err := elastic.client.CreateIndex(ELASTIC_INDEX).Body(getIndexMapping()).Do(ctx) //BodyString(mapping). error
		if err != nil {
			// Handle error
			//panic(err)
			return err, false
		}
		if !createIndex.Acknowledged {
			return errors.New("createIndex Acknowledged false"), false
			// Not acknowledged
		}

		elastic.client.PutMapping()
	}

	return nil, true
}

//新增
func AddProduct(ec *ElasticClient, data *ElasticProductModel) (error, *elastic.IndexResponse) {

	ctx := context.Background()

	if err, ok := IndexExists(ec, ctx); err != nil || !ok {
		return err, nil
	}

	put, err := ec.client.Index().
		Index(ELASTIC_INDEX).
		//Type(ELASTIC_TYPE). 新版本移除type
		Id(data.Guid).
		BodyJson(&data).
		Refresh(REFRESH).
		Do(ctx)

	if err != nil {
		// Handle error
		//panic(err)
		return err, nil
	}
	return nil, put
	//fmt.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//更新 (如果不存在该数据 就创建)
func UpdateProduct(ec *ElasticClient, data *ElasticProductModel) (error, *elastic.UpdateResponse) {
	ctx := context.Background()

	if err, ok := IndexExists(ec, ctx); err != nil || !ok {
		return err, nil
	}

	// Update a tweet by the update API of Elasticsearch.
	update, err := ec.client.Update().
		Index(ELASTIC_INDEX).
		//Type(ELASTIC_TYPE). 新版本移除type
		Id(data.Guid).
		Upsert(&data).
		Do(ctx)

	if err != nil {
		// Handle error
		//panic(err)
		return err, nil
	}

	return nil, update
	//fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)
}

//删除
func DeleteProduct(ec *ElasticClient, data *ElasticProductModel) (error, *elastic.DeleteResponse) {
	ctx := context.Background()

	if err, ok := IndexExists(ec, ctx); err != nil || !ok {
		return err, nil
	}

	// Update a tweet by the update API of Elasticsearch.
	del, err := ec.client.Delete().
		Index(ELASTIC_INDEX).
		//Type(ELASTIC_TYPE). 新版本移除type
		Id(data.Guid).Do(ctx)

	if err != nil {
		// Handle error
		return err, nil
		//panic(err)
	}

	return nil, del
	//fmt.Printf("New version of tweet %q is now %d\n", del.Id)
}

//
func GetProduct(ec *ElasticClient, guid string) (error, *ElasticProductModel) {
	var data ElasticProductModel

	ctx := context.Background()

	// Get tweet with specified ID
	getRes, err := ec.client.Get().
		Index(ELASTIC_INDEX).
		//Type(ELASTIC_TYPE). 新版本移除type
		Id(guid).
		Do(ctx)

	if err != nil {
		// Handle error
		//panic(err)
		return err, nil
	}

	if getRes.Found {
		fmt.Println(string(getRes.Source))
		//fmt.Printf("Got document %s in version %d from index %s, type %s\n", getRes.Id, getRes.Version, getRes.Index, getRes.Type)
		if err := json.Unmarshal(getRes.Source, &data); err != nil {
			//panic(err)
			return err, nil
		}
	}

	return nil, &data
}

type ProductQueryModel struct {
	QueryType   int    //查询类型
	QueryFields string //查询字段 可以 单个 或者 多个  [ "title", "body" ]
	QueryText   string //查询文字
	SortField   string //排序字段 (可以多字段排序等 ,暂时没有写)
	Ascending   bool   //升序
	From        int    //分页参数 显示应该跳过的初始结果数量，默认是 0
	Size        int    //分页参数 显示应该返回的结果数量，默认是 10
}

// var querystrg = `
// {
//    "query": {
//        "bool" : {
//            "must" : {
//                "match_all" : {}
//            },
//            "filter" : {
//                "geo_distance" : {
//                	"_name":"distance",
//                    "distance" : "10km",
//                    "distance_type": "plane",
//                    "location" : {
//                       "lat": 31.2287250000,
//                		"lon": 121.4751860000
//                    }
//                }
//            }
//        }
//    },
//    "sort": [
//    {
//      "_geo_distance": {
//        "location": {
//          "lat": 31.2287250000,
//          "lon": 121.4751860000
//        },
//        "order":         "desc",
//        "unit":          "km",
//        "distance_type": "plane"
//      }
//    }
//  ]
// }
// `

//查询 product
//多字段查询
//相关性排序
//查询是否存在
func QueryProduct(ec *ElasticClient, m *ProductQueryModel) (error, *[]ElasticProductModel) {
	ctx := context.Background()

	//查询方式
	q := getQueryTypeObject(m)
	//termQuery := elastic.NewTermQuery("user", "olivere") 精准查询
	//q := elastic.NewQueryStringQuery(m.Name) //全局文字查询
	//q := elastic.NewMatchQuery(m.QueryName, m.QueryText) //匹配查询 该match查询是用于执行全文搜索（包括模糊匹配选项）的标准查询。

	s := ec.client.Search().
		Index(ELASTIC_INDEX).
		Query(q)

	//res, err := ec.client.Search().
	//	Index(ELASTIC_INDEX).
	//	Source(querystrg).
	//	From(m.From).Size(m.Size).Pretty(true).
	//	Do(ctx)

	if m.SortField != "" {
		s = s.Sort(m.SortField, m.Ascending)
	}

	searchResult, err := s.From(m.From).Size(m.Size).Pretty(true).Do(ctx)

	//searchResult, err := ec.client.Search().
	//	Index(ELASTIC_INDEX).
	//	Query(q).           //
	//	Sort("user", true). // sort by "user" f
	//	// ield, ascending user.keyword
	//	From(m.From).Size(m.Size). // take documents 0-9
	//	Pretty(true).              // pretty print request and response JSON
	//	Do(ctx)                    // execute

	if err != nil {
		// Handle error
		//panic(err)
		return err, nil
	}

	var data ElasticProductModel
	var products []ElasticProductModel

	for _, item := range searchResult.Each(reflect.TypeOf(data)) {
		if t, ok := item.(ElasticProductModel); ok {
			products = append(products, t)
		}
	}

	return nil, &products
}

//获取查询对象
func getQueryTypeObject(m *ProductQueryModel) elastic.Query {

	switch ElasticSearchQueryType(m.QueryType) {
	case EnumMatchQuery:
		return getMatchQuery(m)
	case EnumTermQuery:
		return getTermQuery(m)
	case EnumStringQuery:
		return getStringQuery(m)
	default:
		return getMatchQuery(m) //默认
	}

}

func getMatchQuery(m *ProductQueryModel) *elastic.MatchQuery {
	q := elastic.NewMatchQuery("name", m.QueryText).Operator("OR").QueryName("skuInfo.attributeInfo") //Analyzer("ik_max_word")

	//elastic.NewMultiMatchQuery() 多字段查询

	return q
}

func getTermQuery(m *ProductQueryModel) *elastic.TermQuery {
	q := elastic.NewTermQuery(m.QueryFields, m.QueryText)

	return q
}

func getStringQuery(m *ProductQueryModel) *elastic.QueryStringQuery {
	q := elastic.NewQueryStringQuery(m.QueryText).Field("skuInfo.attributeInfo").Field("name")

	return q
}
