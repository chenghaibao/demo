package main

func main() {

	// 自动生成文档思路
	// ast + generate http://c.biancheng.net/view/4442.html
	//t1 := reflect.TypeOf((*controller.TestInterface)(nil)).Elem()
	//for i := 0; i < t1.NumMethod(); i++ {
	//	f := t1.Method(i)
	//	println("---------")
	//	println(f.Name)
	//	println(f.Type.String(), f.Type.Name())
	//	//todo golang不能想php和java的jvm 一样进行string创建对象(类),需要通过字符串映射的模式来实现功能
	//	println(f.Index)
	//	println(f.Func.String())
	//	println("---------")
	//}

	// todo 反射,容器，Iop 切面，
	// todo 不入侵业务    专注于业务不关注于框架
	// todo 微服务 先拆成服务 ，在进行优化
}
