package data_conversion

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"io"
	"log"
)

func HbaseDb() {
	client := gohbase.NewClient("127.0.0.1:9191")
	var hf []func(hrpc.Call) error
	// 只查询固定的列{cf: [col1, col2]}
	Families := make(map[string][]string)
	temp := []string{"name", "age"}
	Families["f1"] = temp
	hf = append(hf, hrpc.Families(Families))
	// f 设定过滤配置
	var f filter.Filter
	// 条件1：限制返回条数
	f = filter.NewPageFilter(100)
	hf = append(hf, hrpc.Filters(f))
	// 条件2：前缀过滤
	var str string = "003"
	//var data []byte = []byte(str)
	f = filter.NewPrefixFilter([]byte(str))
	hf = append(hf, hrpc.Filters(f))
	getRequest, err := hrpc.NewScanStr(context.Background(), "表名", hf...)
	if err != nil {
		fmt.Println(err.Error())
	}
	scan := client.Scan(getRequest)
	fmt.Println(scan)
	var res []*hrpc.Result
	for {
		getRsp, err := scan.Next()
		if err == io.EOF || getRsp == nil {
			break
		}
		if err != nil {
			log.Print(err)
		}
		res = append(res, getRsp)
	}
	fmt.Println(res)
}
