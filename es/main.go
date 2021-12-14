package main

import (
	"context"
	"hb_es/es"
)

type Test struct {
	Age   int `json:"age"`
	Name  string `json:"name"`
	Id    int `json:"id"`
	Score string `json:"score"`
}

func main() {
	es.NewEs(context.Background())

	// 添加
	//name := "hb_test"
	//data := `{
	//"name": "wali",
	//"id": 24,
	//"age": 30,
	//"score": "980"
	//}`
	//es.Es.AddData(name,data)

	// 跟据id查找
	//var t Test
	//json.Unmarshal(result.Source,&t)
	//fmt.Println(t.Age)
	//result,_ := es.Es.FindId(name,"GVIiuX0BxNDfSOnYbyU0")

	//查找全部
	//result,_ := es.Es.FindAllQuery(name)

	// 条件查找
	//boolQ := elastic.NewBoolQuery()
	//boolQ.Must(elastic.NewMatchQuery("name", "三"))
	//boolQ.Filter(elastic.NewRangeQuery("Score").Gt(2))
	//res, _ :=  es.Es.FindConditionQuery(name,boolQ)




}


