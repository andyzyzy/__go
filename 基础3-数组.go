package main

import "fmt"

func main() {
	//var variable_name [SIZE] variable_type
	var a [10]int // 声明
	a[1] = 100
	fmt.Println(a[2])

	b := [...]int{1, 2, 3} //可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
	fmt.Println(b[2])

	c := [...]float32{4: 0.1, 2: 1.1} //  将索引为 2 和 4 的元素初始化
	fmt.Println(c[2])

	var d [100]int

	for i := 0; i < 100; i++ {
		d[i] = i
	}

	for i := 99; i >= 0; i-- {
		fmt.Printf("d[%d]=%d", i, d[i])
		fmt.Println()
	}

	// 多维数组
	var e [][]int
	f := []int{1, 2, 3}
	g := []int{4, 5, 6}
	e = append(e, f)
	e = append(e, g)
	fmt.Println(e[0][2])
	fmt.Println(e[1])

	for index := range e {
		fmt.Println(index)
	}

	// 函数传递数组为参数
	//void myFunction(param []int)
	//{
	//.
	//.
	//.
	//}

}
