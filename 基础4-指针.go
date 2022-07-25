package main

import "fmt"

func pointer1() {
	var a *int
	b := 100
	a = &b
	fmt.Println(a)
	fmt.Println(*a)
}
func pointer2() {
	const (
		Max = 3
	)

	var a = [Max]int{1, 2, 3}
	fmt.Println(&a[1])
	var a_pointer [Max]*int
	for arg := range a {
		fmt.Println(arg)

	}
	for i := 0; i < Max; i++ {
		a_pointer[i] = &a[i]
	}
	fmt.Println(a_pointer)
}
func pointer3() {
	var aa *int
	var aaa **int
	a := 1
	aa = &a
	aaa = &aa
	fmt.Println("a的地址是：", &a, "值为：", a)
	fmt.Println("aa的地址是：", &aa, "值为：", *aa)
	fmt.Println("aaa的地址是：", &aaa, "值为：", **aaa)
	fmt.Println(*aaa)
}
func main() {
	//pointer1()
	//pointer2()   //指针赋值及遍历
	pointer3() // 多级指针
}
