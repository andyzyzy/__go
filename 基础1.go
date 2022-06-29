package main

import "fmt"

func main() {
	// 单行注释
	/*
		多行注释
	*/
	var a, b, c = 1, "2", 3
	fmt.Println(a, b, c) // pintln 换行打印
	fmt.Println("v1" + "v2")

	var url = "http://ahc.day-care.cn/#/admin/staff/index?age=%d&name=%s"
	var age = 15
	var name = "andy"
	var urlPath = fmt.Sprintf(url, age, name) // 格式化字符串
	fmt.Printf(urlPath)
}
