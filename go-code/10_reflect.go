package main

import (
	"fmt"
	"reflect"
)

func hello()  {
	fmt.Println("hello go")
}

func main() {

	/*反射实践*/
	/*1. 通过反射修改内容*/
	var f float64 = 2.58
	p := reflect.ValueOf(&f)
	v1 := p.Elem()
	t1 := reflect.TypeOf(f)
	fmt.Println("p",p.CanSet()) //p false
	if v1.CanSet() {
		v1.SetFloat(3.67)  //修改了的值为3.67
		fmt.Println("改变了值")
	}

	fmt.Println(f) //3.67
	fmt.Println(t1)  //float64

	/*2.通过反射调用方法*/
	hl := hello
	fh := reflect.ValueOf(hl)
	fh.Call(nil)  //被调用，输出hello go


	fmt.Println("=================================")
	var Num float64 = 3.14
	/*1. 接口类型变量转变为反射类型对象*/
	v := reflect.ValueOf(Num) //3.14
	t := reflect.TypeOf(Num) //float64

	fmt.Println(v)
	fmt.Println(t)

	/*2. 反射类型对象转换回接口类型变量*/
	origin := v.Interface().(float64) //3.14
	fmt.Println(origin)


	defer func() {
		fmt.Println("有错误产生")
		if err:=recover();err!=nil{ // err==nil 代表着正常
			fmt.Println(err)
			recover()
			fmt.Println("被抓住后...")
		}
	}()
	/*3. 反射类型对象必须是可写的，才可以进行修改*/
	fmt.Println(v.CanSet()) //false
	if !v.CanSet() {
		v.SetFloat(6.28) //出现异常，向上传了
		fmt.Println(v)
	}
	fmt.Println("end!!!!!")  //没有被输出来

}
