package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 单行注释
	/*
		多行注释
	*/
	var a, b, c = 1, "222", 3
	fmt.Println(a, b, c) // pintln 换行打印
	fmt.Println("v1" + "v2")

	var url = "http://ahc.day-care.cn/#/admin/staff/index?age=%d&name=%s"
	var age = 15
	var name = "andy"
	var urlPath = fmt.Sprintf(url, age, name) // 格式化字符串
	fmt.Println(urlPath)

	n1 := nm() // 调用函数
	fmt.Println(n1)

	const c1, c2, c3, c4 = 200, 111, 222, 333 // 定义常量
	fmt.Println(c1, c2, c3, c4)

	const length = 100
	const width = 20
	fmt.Sprintf("面积：%d", length*width) // 常量计算
	fmt.Println(fmt.Sprintf("面积：%d", length*width))
	const (
		unknown = 0
		female  = 1
		male    = 2
	) // 常量枚举

	fmt.Println(fmt.Sprintf("男：%d,女：%d,位置：%d", unknown, female, male))

	const (
		a1 = "abc"
		b1 = len(a1)
		cc = unsafe.Sizeof(b) //16字节，128位
	)
	fmt.Println(a1, b1, cc)

	//iota 用法
	const (
		i1 = 2 << iota // 2
		i2 = iota      // 1
		i3 = 2 << iota // 8
		i4             // 16
		i5             // 32
	)
	fmt.Println(i1, i2, i3, i4, i5)
	// switch 条件语句
	var g = 21
	var n = "andy"

	switch g {
	case 10:
		n = "gd"
	case 20, 30:
		n = "jack"
	case 40:
		n = "sb"
	default:
		n = "123"
	}
	switch {
	case n == "jack":
		fmt.Println(n)
	case n == "gd":
		fmt.Println(n)
	default:
		fmt.Println("123")
	}

	var number = 3
	switch {
	case number > 0:
		fmt.Println(1)
	case number > 2:
		fmt.Println(2)
	default:
		fmt.Println(3)

	}

	var iz int16
	fmt.Println(iz)
	iz = 9
	ccc := 0
	fmt.Println(ccc)

}

func nm() int16 {
	//a := 5
	return 5
}
