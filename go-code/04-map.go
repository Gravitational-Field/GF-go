package main

import "fmt"

func main() {

	//var m1 map[string]int
	//m2 := make(map[int]interface{},100)
	m3 := map[string]string{
		"name":"james",
		"age":"35",
	}
	//m1["key1"] = 1
	//m2[1] = 1

	m3["key1"] = "v1"
	m3["key2"] = "v2"
	m3["key3"] = "v3"
	m3["key3"] = "v0"

	//fmt.Println(len(m3))
	//
	//fmt.Println(m3["name"])

	//确定键值对存在
	//if value, ok:= m3["name"]; ok{
	//	fmt.Println(value)
	//}


	//遍历
	//for key,value := range m3{
	//	fmt.Println("key: ", key, " value: ", value)
	//}

	//delete(m3,"key3")

	//遍历
	//for key,value := range m3{
	//	fmt.Println("key: ", key, " value: ", value)
	//}

	 m := make(map[string]func(a, b int) int)
	 m["add"] = func(a, b int) int {
		 return a+b
	 }
	 m["multi"] = func(a, b int) int {
	 	return a*b
	 }

	 fmt.Println(m["add"](3,2))
	 fmt.Println(m["multi"](3,2))


}
