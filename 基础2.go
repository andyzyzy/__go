package main

import (
	"fmt"
	"math"
	"time"
)

func Chan(ch chan int, stopCh chan bool) {
	for i := 1; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func Select() {
	ch := make(chan int)
	stopCh := make(chan bool)
	go Chan(ch, stopCh) // 并行
	for {
		select {
		case i := <-ch:
			fmt.Println("接受", i)
		case j := <-ch:
			fmt.Println("接受2", j)
		case _ = <-stopCh:
			goto end // 结束
		}
	}
end:
}

func For() {
	var s = []string{"测试", "开发", "产品"}
	for i, v := range s {
		fmt.Println(i, v)
	}

	sum := 1
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	i3 := 1
	for i3 < 10 {
		fmt.Println(i3)
		i3 += 1
	}
}
func compare(a int, b int) int {
	/*
		if 布尔表达式 {
		   // 在布尔表达式为 true 时执行
		}
	*/
	var result int
	if a < b {
		result = b
	} else {
		result = a
	}
	return result
}

func Exchage(s []string) [2]string {
	var restult [2]string
	//fmt.Println(s[0])
	//fmt.Println(restult)

	restult[0] = s[1]
	restult[1] = s[0]
	fmt.Println(restult[1])
	return restult

}
func Func() {

	a := 100
	b := 200
	v := compare(a, b)
	fmt.Println(v)
}

func Func2() {
	var a = []string{"谷歌", "百度"}
	//fmt.Println(a)
	v := Exchage(a)
	fmt.Println(v)
}

func excharge2(x, y string) (string, string) {
	return y, x
}
func Func3() {
	s1, s2 := excharge2("5", "6")
	fmt.Println(s1, s2)
}
func swap(a, b *int) {
	var temp int
	fmt.Println(*a)
	temp = *a
	*a = *b
	*b = temp
}

func Func4() {
	a := 100
	b := 200
	fmt.Println(&a)
	fmt.Println(&b)
	swap(&a, &b)
	fmt.Println("a等于:", a)
	fmt.Println(fmt.Sprintf("b等于：%d", b))
	fmt.Println(&a)
	fmt.Println(&b)
}

func Func5() {
	getSquareRoot := func(b float64) float64 {
		return math.Sqrt(b)
	}
	fmt.Println(getSquareRoot(9)) //函数作为实参，灵活调用
}

type cb func(int) int // type别名
func Func6() {
	Call(1, CallBack)
	Call(3, func(i int) int {
		fmt.Println("回调函数：", i)
		return i
	})
}
func Call(a int, b cb) {
	b(a)
}
func CallBack(x int) int {
	fmt.Println("回调函数:", x)
	return x
}

// 闭包
func getSequence() func() int {
	s := 0
	return func() int {
		s += 1
		return s
	}

}
func Func7() {
	gs := getSequence()
	fmt.Println(gs())
	fmt.Println(gs())
	fmt.Println(gs())
}

type Circle struct {
	radius float64
}

func (c Circle) getAreas() float64 {
	return c.radius * c.radius
}

func (c *Circle) charges(radius float64) {
	c.radius = radius
}
func Func8() {
	var c Circle
	c.radius = 100.00
	fmt.Println(c.getAreas())
	c.charges(20)
	fmt.Println(c.getAreas())
}

func main() {
	//Select() // select 信道接受消息，阻塞接收
	//For()  // for 循环
	//Func()  // 函数
	//Func2()
	//Func3()
	//Func4() //值交换
	//Func5() //函数作为实参，灵活调用
	//Func6() // 函数作为参数传递，实现回调。--重要
	//Func7() // 闭包
	// Func8() // 结构体
}
