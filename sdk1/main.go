package main

import (
	"fmt"
	"github.com/chenghaibao/sdk_demo/logic/test"
)

func main() {
	aa := &test.HbFirst{}
	aa.SetAge(12)
	fmt.Println(aa.GetAge())
}
