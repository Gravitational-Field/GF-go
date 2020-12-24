package main

import (
	"fmt"
)

func main() {

	//var arr1 = [5]int{}
	//var arr2 = []int{1,2,3,4,5}
	//var arr3 = [5]int{3:10}
	//fmt.Println(arr1)
	//fmt.Println(arr2)
	//fmt.Println(arr3)
	//
	///*使用循环为arr1赋初值*/
	//for i := 0; i < len(arr1); i++ {
	//	arr1[i] = i*10
	//}
	//
	//fmt.Println(arr1)
	//
	///*遍历数组*/
	//for index,value := range arr1{
	//	fmt.Printf("index: %d, value: %d\n", index, value)
	//}

	var a [5]int
	fmt.Println(a)
	test(a)
	fmt.Println(a)

}

func test(a [5]int) {
	a[1] = 2
	fmt.Println(a)
}
