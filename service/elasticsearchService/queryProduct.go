package elasticsearchService

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

//
//var (
//	//请求主体
//	requestBodySourceStr = `
//	{
//		"query":{%v},
//		"sort":[%v]
//	}`
//	//距离排序
//	sortGeoDistance = `{
//      "_geo_distance": {
//        "location": {
//          "lat": %v,
//          "lon": %v
//        },
//        "order":         "%v",
//        "unit":          "%v",
//        "distance_type": "%v"
//      }
//    }`
////	"term": { "categoryIds": "77" },
////
////		"filter" : {
////"term": { "categoryIds": "77" },
////"geo_distance" : {
////"distance" : "10km",
////"distance_type": "plane",
////"location" : {
////"lat": 31.2287250000,
////"lon": 121.4751860000
////}
////}
////}
//{ "match": { "description":   "ttt" }}
//	//查询所有
//	queryMatchAll = `{
//            "must" : {
//                "match_all" : {}
//            }`
//
//	queryMatch = ` "must": []`
//
//	//查询过滤 -- 距离(经纬度)过滤
//	queryFilterGeoDistance = `{
//                "geo_distance" : {
//                	"_name":"distance",
//                    "distance" : "%v",
//                    "distance_type": "%v",
//                    "location" : {
//                       "lat": %v,
//                		"lon": %v
//                    }
//                }`
//
//	queryBodyString = `
//{
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
//}
//`
//)
//
//排序
//"sort": [
//{ "currentPrice": { "order": "asc" }},
//{ "_score": { "order": "desc" }},
//{
//"_geo_distance": {
//"location": {
//"lat": 31.2287250000,
//"lon": 121.4751860000
//},
//"order":         "desc",
//"unit":          "km",
//"distance_type": "plane"
//}
//},
//{ "sales": { "order": "desc" }}
//]

type (
	RequestBodySource struct {
		Query RequestBodySourceQuery   `json:"query"`
		Sort  []map[string]interface{} `json:"sort,omitempty"`
	}

	RequestBodySourceQuery struct {
		Bool QueryBool `json:"bool"`
	}

	QueryBool struct {
		Must   map[string]interface{} `json:"must"`
		Should map[string]interface{} `json:"should"`
		Filter *QueryBoolFilter       `json:"filter,omitempty"`
	}

	QueryBoolMustName struct {
		Name string `json:"name"`
	}

	QueryBoolMustDescription struct {
		Description string `json:"description"`
	}

	QueryBoolMustCategoryId struct {
		CategoryId uint64 `json:"categoryIds"`
	}

	QueryBoolFilter struct {
		GeoDistance *GeoDistanceFilter `json:"geo_distance,omitempty"`
	}

	GeoDistanceFilter struct {
		Location     GeoDistanceLocation `json:"location"`
		Distance     string              `json:"distance"`      //距离 范围 10km
		DistanceType string              `json:"distance_type"` //plane
	}

	RequestSort struct {
		Order string `json:"order"` //asc , desc
	}

	RequestGeoDistanceSort struct {
		GeoDistance GeoDistanceSort `json:"geo_distance"`
	}

	GeoDistanceSort struct {
		Location     GeoDistanceLocation `json:"location"`
		Unit         string              `json:"unit"`          //km
		Order        string              `json:"order"`         //asc , desc
		DistanceType string              `json:"distance_type"` //plane
	}

	GeoDistanceLocation struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
)

type SearchProductModel struct {
	From           int     `json:"from"`         //分页参数 显示应该跳过的初始结果数量，默认是 0
	Size           int     `json:"size"`         //分页参数 显示应该返回的结果数量，默认是 10
	Distance       int     `json:"distance"`     //筛选距离 单位 km
	CategoryId     uint64  `  json:"categoryId"` //商品分类Id
	QueryText      string  `  json:"text"`       //搜索名称 泛搜索 包括 商品名称
	Lat            float64 ` json:"lat"`         //维度
	Lon            float64 ` son:"lon"`          //经度
	SearchSortType int     `  json:"sortType"`   //排序类型 1、默认 3、销量 正序 5、销量 倒叙  7、价格 正序 9、价格 倒叙 11、距离 最近
}

//
func SearchProductList(ec *ElasticClient, m *SearchProductModel) (error, []ElasticProductModel) {
	ctx := context.Background()

	var query RequestBodySource
	var qBody RequestBodySourceQuery
	var qBool QueryBool
	var qFilter *QueryBoolFilter
	var qqSortArr []map[string]interface{}
	qSort := make(map[string]interface{})
	qMust := make(map[string]interface{})
	qShould := make(map[string]interface{})

	//查询条件
	if m.CategoryId > 0 {
		qMust["term"] = QueryBoolMustCategoryId{CategoryId: m.CategoryId}
	}
	qMust["match"] = QueryBoolMustName{Name: m.QueryText}

	//增加相关性
	qShould["match"] = QueryBoolMustDescription{Description: m.QueryText}

	m.Distance = 0
	//过滤字段
	if m.Lat != 0 && m.Lon != 0 && m.Distance > 0 {
		//距离过滤
		qFilter = &QueryBoolFilter{}
		qFilter.GeoDistance = &GeoDistanceFilter{
			DistanceType: "plane",
			Location: GeoDistanceLocation{
				Lat: m.Lat,
				Lon: m.Lon,
			},
			Distance: strconv.Itoa(m.Distance) + "km",
		}
	}

	//排序
	if m.SearchSortType == 1 {
		if m.Lon != 0 && m.Lat != 0 {
			qSort["_geo_distance"] = GeoDistanceSort{
				Order:        "desc",
				Unit:         "km",
				DistanceType: "plane",
				Location: GeoDistanceLocation{
					Lat: m.Lat,
					Lon: m.Lon,
				},
			}
			qqSortArr = append(qqSortArr, qSort)
			qSort = make(map[string]interface{})
		}

		qSort["sales"] = RequestSort{Order: "desc"}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})

		qSort["currentPrice"] = RequestSort{Order: "desc"}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})
	} else if m.SearchSortType == 3 { //销量正序 sales
		qSort["sales"] = RequestSort{Order: "desc"}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})
	} else if m.SearchSortType == 5 { //销量倒序
		qSort["sales"] = RequestSort{Order: "asc"}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})
	} else if m.SearchSortType == 7 { //价格正序
		qSort["currentPrice"] = RequestSort{Order: "desc"}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})
	} else if m.SearchSortType == 9 { //价格倒序
		qSort["currentPrice"] = RequestSort{Order: "asc"}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})
	} else if m.SearchSortType == 11 { //距离 最近
		qSort["_geo_distance"] = GeoDistanceSort{
			Order:        "desc",
			Unit:         "km",
			DistanceType: "plane",
			Location: GeoDistanceLocation{
				Lat: m.Lat,
				Lon: m.Lon,
			},
		}
		qqSortArr = append(qqSortArr, qSort)
		qSort = make(map[string]interface{})
	}

	qBool.Must = qMust
	qBool.Should = qShould

	if qFilter != nil {
		qBool.Filter = qFilter
	}

	qBody.Bool = qBool

	query.Query = qBody
	query.Sort = qqSortArr

	querystrg, _ := json.Marshal(query)

	queryStr := string(querystrg)

	fmt.Println(queryStr)

	searchResult, err := ec.client.Search().
		Index(ELASTIC_INDEX).
		Source(queryStr).
		From(m.From).Size(m.Size).Pretty(true).
		Do(ctx)

	if err != nil {
		return err, nil
	}

	var data ElasticProductModel
	var products []ElasticProductModel

	//自己序列号 对象返回
	//if searchResult.Hits.TotalHits.Value > 0 {
	//	for _, i := range searchResult.Hits.Hits {
	//
	//		err := json.Unmarshal(i.Source, &data)
	//		products = append(products, data)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		}
	//	}
	//}

	//使用库里的方法序列号 返回
	for _, item := range searchResult.Each(reflect.TypeOf(data)) {

		if t, ok := item.(ElasticProductModel); ok {
			products = append(products, t)
		}
	}

	if len(products) < 1 {
		// TODO:根据商品名没有查询到

		//查询相关商品
		//1、分类
		//2、描述
		//3、商品标签

	}

	return nil, products
}
