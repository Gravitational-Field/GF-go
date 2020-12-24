package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type ListNode struct {
	Val int
	Next *ListNode
}

type Person struct {
	ID string
	int // 匿名字段
}

type A struct {
	X,Y int
}

type B struct {
	A //匿名结构体
	Name string
	X int
}

type C struct {
	A
	B
	X int
}

type D struct {
	arr [5]int
	slice []int
	Map map[string]int

}


func main() {
//	创建结构体实例
	s1 := new(Student)
	s2 := Student{"james",36}
	s3 := &Student{
		Name: "wade",
		Age: 35,
	}

	fmt.Println(*s1) //{ 0}
	fmt.Println(s2) //{james 36}
	fmt.Println(*s3) //{wade 35}

	fmt.Println(s2.Name,s2.Age) //james 36

	p := new(Person)
	p.ID = "001"
	p.int = 100
	fmt.Println(*p) //{001 100}

	fmt.Println(ToJson(s3))

	b := new(B)
	b.X = 10
	b.Y = 20
	b.Name= "BBBBB"
	fmt.Println("*b: ",*b)

	c := new(C)
	c.X = 10
	c.Y = 11
	fmt.Println("*c: ",*c)  //*c:  {{0 11} {{0 0} } 10}

/*复杂结构体*/
	d := &D{
		arr: [5]int{1,2,3},
		slice: []int{4,5,6,7},
		Map: map[string]int{"id":1,"age":18},
	}

	fmt.Println("*d:",*d) // *d: {[1 2 3 0 0] [4 5 6 7] map[age:18 id:1]}

}

func ToJson(s *Student) (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}
