package elasticsearchService

//
type ElasticSearchQueryType int

const (
	EnumMatchQuery ElasticSearchQueryType = 1 //	匹配查询 该match查询是用于执行全文搜索（包括模糊匹配选项）的标准查询。

	EnumTermQuery = 3 //精准查询

	EnumStringQuery = 5 //
)
