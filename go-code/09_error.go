package main

import (
	"errors"
	"fmt"
)
/*练习error*/


//type error interface {
//	Error() string
//}

//*自定义错误*/
type NotNature float64

func (err NotNature)Error() string { //实现了一个Error方法，返回错误提醒字符串。
	return fmt.Sprintf("自然数为大于或等于0的数: %v", float64(err))
}

func Nature(x float64) (float64,error) {  //返回float及error类型的数据
	if x<0 {
		return 0,NotNature(x) //出现不一致，则输出错误
	} else {
		return x,nil
	}
}


func main() {

	err := errors.New("This is an error")
	err1 := fmt.Errorf("This is another error")
	if err != nil {
		fmt.Println(err)
	}


	if err1 != nil {
		fmt.Println(err1) //This is another error
	}

	fmt.Println(Nature(1)) //1 <nil>
	fmt.Println(Nature(-1)) //0 自然数为大于或等于0的数: -1


}
