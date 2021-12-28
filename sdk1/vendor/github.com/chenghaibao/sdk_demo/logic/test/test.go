package test

import "fmt"

type HbInterface interface {
	echoName() string
}

type HbFirst struct {
	HbInterface
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (this *HbFirst) SetName(name string) {
	this.Name = name
}

func (this *HbFirst) GetName() string {
	return this.Name
}

func (this *HbFirst) SetAge(Age int) {
	this.Age = Age
}

func (this *HbFirst) GetAge() int {
	return this.Age
}

func (this *HbFirst) echoName() string {
	fmt.Println(this.Age)
	return this.Name
}
