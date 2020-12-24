package main

import "fmt"

func main() {
	//
	//var name1 = 100
	//name2:=200
	//fmt.Println(name1)
	//fmt.Println(name2)
	//name := identifier  // 出现在 := 左侧的变量不能是已经被声明过的，否则会导致编译错误

	//var identifier int = 100
	//var identifier1,identifier2 int = 200,300
	//fmt.Println(identifier+identifier1+identifier2) //600

	// 声明一个变量并初始化
	//var a = "RUNOOB"
	//fmt.Println(a)  //RUNOOB

	// 没有初始化就为零值
	//var b int
	//fmt.Println(b)
	//
	//// bool 零值为 false
	//var c bool
	//fmt.Println(c)
	//
	//var d string
	//fmt.Println(d)
	//
	//var a *int
	//fmt.Println(a)
	//
	//var intVal int
	//intVal,intVal1 := 1,2
	//fmt.Println(intVal+intVal1) //3

	const(
		he = 100
		ha = 200
		xi  //此时xi与ha等值
	)
	fmt.Println(he,ha,xi)
	//fmt.Println(a,b,c)  //5 1 2


}
