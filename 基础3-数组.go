package main

import (
	"fmt"
)

func getAverage(a [5]int, b int) float32 {
	var (
		sum int
		//int     i
	)

	for i := 0; i < b; i++ {
		sum += a[i]
	}
	avge := float32(sum) / float32(b)
	return avge
}
func Func1() {
	fmt.Println("向函数传递数据组")
	b1 := [5]int{9.0, 11, 20, 21, 22}
	fmt.Println(getAverage(b1, 5))
}
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

	z := [2][4]int{{1, 2, 3}, {4, 5, 6, 7}}
	fmt.Println(z[0][2])
	fmt.Println(z[1][3])
	// 函数传递数组为参数
	//void myFunction(param []int)
	Func1()

	i := 7
	var j int = i
	fmt.Println(j)
	fmt.Println(i)
	var t int
	if t {
		fmt.Println(t)
	}
}
