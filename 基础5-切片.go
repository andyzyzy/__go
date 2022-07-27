package main

import "fmt"

func printSlice(s []int) {
	// len 是长度，cap是容量
	fmt.Println(len(s), cap(s), s)
}
func addSlice(s []int) { //切片为参数时，也是copy
	s = append(s, 9)
	fmt.Println(s)
}

func addSlice2(s *[]int) {
	*s = append(*s, 9)
	fmt.Println("打印", *s)
}

func main() {
	//make([]T, length, capacity)
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	m := make([]int, 2, 5)
	printSlice(m)       // 2 5 [0 0]
	printSlice(arr[:3]) //3,5,[0,1,2]
	var arr2 []int
	addSlice(arr2)
	fmt.Println(arr2)
	addSlice2(&arr2)
	fmt.Println(arr2)
	m = append(m, 1)
	m = append(m, 3)
	m[0] = 9
	fmt.Printf("长度:%d,容量:%d,值：%v", len(m), cap(m), m)
	fmt.Println()
	slice1 := make([]int, 4, 8)
	slice2 := make([]int, 1, 2)
	slice2[0] = 100
	//copy(slice1,slice2)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(cap(slice1), len(slice1))

}
