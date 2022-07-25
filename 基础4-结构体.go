package main

import "fmt"

type Book struct {
	name string
	page int
}

func printf1(b1, b2 Book) {
	fmt.Println(b1.name, b1.page)
	fmt.Println(b2.name, b2.page)
}

func f1() {
	var b1, b2 Book
	b1.name = "西游记"
	b1.page = 200
	b2.name = "三国"
	b2.page = 300
	printf1(b1, b2)
}
func printf2(b1, b2 *Book) {
	fmt.Println(b1.name, b1.page)
	fmt.Println(b2.name, b2.page)
}
func f2() {
	var b1, b2 Book
	b1.name = "水浒"
	b1.page = 100
	b2.name = "红楼梦"
	b2.page = 400
	printf2(&b1, &b2)
}
func main() {
	//f1() // 结构体作为参数
	f2() // 结构体地址作为参数
}
