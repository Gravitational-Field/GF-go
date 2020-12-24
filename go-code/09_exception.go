package main

import "fmt"

/*
真正的异常处理：panic和recover
 */
func main() {
	//fmt.Println("hello go!")
	//panic(errors.New("i am a error"))  //出现异常，后边不会被执行
	//fmt.Println("hello go again") //不会被执行

	//f()
	defer func() {
		fmt.Println("我是defer里面第一个打印函数")
		if err:=recover();err!=nil{ // err==nil 代表着正常
			fmt.Println(err)
			recover()
		}
		fmt.Println("我是defer里面第二个打印函数")
	}()
	f()
	fmt.Println("end！")


	//1
	//我是defer里面第一个打印函数
	//我是panic
	//我是defer里面第二个打印函数
}

func f() {
	fmt.Println("1")
	panic("我是panic")
	fmt.Println("2")
}
