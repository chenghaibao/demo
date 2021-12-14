package es

import (
	"context"
	"github.com/olivere/elastic/v7"
)

// https://blog.csdn.net/p1049990866/article/details/117254708
// https://blog.csdn.net/tflasd1157/article/details/81981915
// 说明表在最下方  这两个链接实参考的
type Elastic struct {
	Client *elastic.Client
	Ctx    context.Context
}

var Es *Elastic

func NewEs(ctx context.Context) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	Es = &Elastic{
		Client: client,
		Ctx:    ctx,
	}
}

func (this *Elastic) CreateIndex(index string) error {
	_, err := this.Client.CreateIndex(index).Do(this.Ctx)
	return err
}

func (this *Elastic) AddData(index string, data string) error {
	_, err := this.Client.Index().Index(index).BodyJson(data).Do(this.Ctx)
	//_, err := this.Client.Index().Index(index).Id().BodyJson(data).Do(context.Background())
	return err
}

// 查询
func (this *Elastic) FindId(index string, id string) (*elastic.GetResult, error) {
	data, err := this.Client.Get().Index(index).Id(id).Do(this.Ctx)
	//_, err := this.Client.Index().Index(index).Id().BodyJson(data).Do(context.Background())
	return data, err
}

// 查询
func (this *Elastic) FindQuery(termQuery elastic.Query) (*elastic.SearchResult, error) {
	data, err := this.Client.Search().Query(termQuery).Do(this.Ctx)
	//this.Client.Search().Index(name).Query(termQuery).SortBy(sorter).FetchSourceContext(elastic.NewFetchSourceContext(true).Include("id", "name")).Sort("age", true).From(0).Size(10).Pretty(true).Do(context.Background())
	return data, err
}

// 查询
func (this *Elastic) FindConditionQuery(index string, Query elastic.Query) (*elastic.SearchResult, error) {
	data, err := this.Client.Search(index).Query(Query).Do(this.Ctx)
	return data, err
}

// 取所有
func (this *Elastic) FindAllQuery(index string) (*elastic.SearchResult, error) {
	data, err := this.Client.Search(index).Do(this.Ctx)
	//this.Client.Search().Index(name).Query(termQuery).Sort("age", true).From(0).Size(10).Pretty(true).Do(context.Background())
	return data, err
}

//简单分页
func (this *Elastic) ListQuery(index string, size int, page int) (*elastic.SearchResult, error) {
	res, err := this.Client.Search(index).Size(size).From((page - 1) * size).Do(this.Ctx)
	return res, err
}

/**
  // term 查询
  // termQuery := elastic.NewTermQuery("name", "wali")
  // elastic.NewQueryStringQuery("last_name:Smith")

  // 条件查询range match模糊
  //boolQ := elastic.NewBoolQuery()
  //boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
  //boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))


  // 取数据方法两种 自己 看着拿
  func printEmployee(res *elastic.SearchResult, err error) {
  	// Extend 1
  	//if err != nil {
  	//	print(err.Error())
  	//	return
  	//}
  	//var typ Test
  	//for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
  	//	t := item.(Test)
  	//	fmt.Printf("%#v\n", t)
  	//}
  	// Extend 2

  	//if *searchResult.Hits.TotalHits > 0 {
  	//	fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
  	//
  	//	// Iterate through results
  	//	for _, hit := range searchResult.Hits.Hits {
  	//		// hit.Index contains the name of the index
  	//
  	//		// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
  	//		var t Tweet
  	//		err := json.Unmarshal(*hit.Source, &t)
  	//		if err != nil {
  	//			// Deserialization failed
  	//		}
  	//
  	//		// Work with tweet
  	//		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
  	//	}
  	//}
  }
*/
