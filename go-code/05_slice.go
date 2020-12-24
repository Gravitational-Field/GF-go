package main

import "fmt"

func main() {
	//var slice []int
	//a := []int{1, 2, 3}
	////b := a[1:3]
	//c := make([]int, 3)
	//fmt.Printf("len:%d,cap:%d\n",len(slice),cap(slice))
	//copy(c, a)
	//fmt.Println(a)
	//fmt.Println(c)


	a := []int{1,2,3}
/*声明b切片时，其长度比a切片长,进行复制*/
	//b := make([]int,5)  //[1 2 3 0 0]
/*声明b切片时，其长度比a切片短,进行复制*/
	//b := make([]int,1) //[1]
/*声明b切片时，其长度被定义为0*/
	b := make([]int,0)  //[]
	copy(b,a)

	fmt.Println(a)
	fmt.Println(b)

}
