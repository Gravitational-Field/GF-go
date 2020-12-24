package main

import (
	"fmt"
	"strconv"
)

func main() {

/*修改字符串的内容*/
	//s := "hello"
	//b := []byte(s)
	//b[0] = 'g'
	//s = string(b)
	//fmt.Println(s) //gello
	//fmt.Println(len(s))  //5

/*含中文的字符串修改*/
	//s := "hello中国你好"
	//fmt.Println(s)  //hello中国你好
	//fmt.Println(utf8.RuneCountInString(s)) //9
	//
	//r := []rune(s)
	//fmt.Println(r)
	//for i:=0; i<len(r); i++ {
	//	fmt.Printf("%c",r[i])
	//}

/*strings包*/
	//var str string = "This is an example of a string"
	//fmt.Println(strings.HasPrefix(str,"Th")) //true
	//fmt.Println(strings.HasSuffix(str,"ng")) //true
	//fmt.Println(strings.Contains(str,"an")) //true


/*strconv包*/

	i,err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i)  //-42
	fmt.Println(err)  //<nil>
	fmt.Println(s)  //-42



}
