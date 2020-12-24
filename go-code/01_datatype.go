package main

import (
	"fmt"
	"math"
)

func main() {
/*bool类型测试*/
	var b bool = false
	fmt.Println(b)
/*变量*/
	//普通赋值
	var num int = 1
	fmt.Println(num) //1
	//平均赋值
	var num1,num2 = 2,3
	fmt.Println(num1,num2) //2 3
	//
	var (
		num3 int= 3
		num4 int = 4
	)
	fmt.Println(num3,num4) //3 4
/*整数类型*/
	var num8 int8= 027  // 23
	fmt.Println(num8) //按十进制进行输出
	var num16 = 0x39 //57
	fmt.Println(num16) //按十进制进行输出
/*浮点数*/
	var floatnum float32 = 58.3e-2
	fmt.Println(floatnum) //0.583
	num = int(floatnum) //float32转int64
	fmt.Println(num) //0

/*复数类型测试*/
	//var x float32= 10.0
	//var y float32= 13.0
	// cannot use complex(x, y) (type complex64) as type complex128 in assignment
	//64位的复数，需要64位的float，不能是其它类型
	var x float64= 10.0
	var y float64= 13.0
	var name complex128 = complex(x, y)

	fmt.Println(name) //(10+13i)
	var z complex128 = name
	fmt.Println(z) //(10+13i)
	fmt.Println(real(name)) //10
	fmt.Println(imag(name)) //13
	fmt.Println(name == z) //true

/*别名类型byte和rune*/
	//byte == uint8
	//var b1 byte = 256 //constant 256 overflows byte
	var b1 byte = 255 //constant 256 overflows byte
	fmt.Println(b1) //255

	//rune == int32   与Unicode编码对应
	//var r1 rune = math.MaxInt32 + 1 //constant 2147483648 overflows rune
	var r1 rune = math.MaxInt32 //2147483647
	fmt.Println(r1)
	var r2 rune = 'A'
	fmt.Println(r2) //65   输出了A字符代表的Unicode编码号

/*字符串类型*/
	var str1 string = `hello\nabc`  //原生表示法   反引号  在左上角而非 单引号
	var str2 string = "world\nabc"  //解释性表示法  会对转义字符进行解释
	fmt.Println(str1+str2)  //hello\nabcworld
							//abc

/*数组类型*/
	//长度确定
	type nums1 [3]int //只是声明一个[3]int的数组
	//fmt.Println(nums1) //Type nums1 is not an expression
	var nums2 = [3]int{1,2,3}  //起别名     PS:起别名时无需指定类型
	var nums3 = [...]int{1,2,3,4,5}  //起别名
	fmt.Println(nums2)  //[1 2 3]
	fmt.Println(nums3)  //[1 2 3 4 5]
	fmt.Println(nums3[2])  //3
	//修改nums3中的值
	nums3[2] = 9
	fmt.Println(nums3) //[1 2 9 4 5]
	var length = len(nums3)   //PS:起别名时无需指定类型
	fmt.Println(length) //5

/*切片类型*/
	//属于引用类型，长度不确定，底层用数组
	type sliceInt []int //声明的作用？
	s := sliceInt{1, 2, 3} //[1 2 3]   :=是什么操作？
	fmt.Println(s)
	var sliceInt1 = []int{4,5,6,7,8,9}  //[4 5 6 7 8 9]
	fmt.Println(sliceInt1)
	//切片
	var slice1 = sliceInt1[1:4]  //[5 6 7]  前闭后开，类似于python
	fmt.Println(slice1)
	//切片的长度与容量：一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值
	fmt.Println(len(s)) //3
	fmt.Println(len(sliceInt1)) //6
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(s)) //3
	fmt.Println(cap(sliceInt1)) //6
	fmt.Println(cap(slice1)) //5     [5,6,7,8,9]
	//切片值的改变会影响切片原始的值
	slice1[0] = 100
	fmt.Println(sliceInt1)  //[4 100 6 7 8 9]
	//通过切片的第三个值进行限制切片的访问
	var slice2 = sliceInt1[0:2:3]  //3
	fmt.Println(cap(slice2))
	//延展切片
	fmt.Println(slice1[:5])  //原本slice1的值为[5 6 7] 现在输出  [100 6 7 8 9]  意为延展其长度为5，数据参照slice1从哪里切过来的，就从哪里可以观察到   这个5是cap(slice1)
	//在切片中加入第三个索引后， var slice1 = numbers3[1:4:4] 将无法通过slice1访问到number3的值中的第五个元素。
	//使用内建函数append可以不受限制的扩展,并返回一个新的切片值,与
	fmt.Println(slice1) //[100 6 7]
	var slice3 = append(slice1, 40,41,42)
	fmt.Println(sliceInt1) // [4 100 6 7 8 9]  与最初是的切片值的来源将没有关系
	fmt.Println(slice3) //[100 6 7 40 41 42]

	//copy函数会直接对其第一个参数值进行对应位置的修改
	var slice4 = []int{0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println(copy(slice4,slice1)) //3   意为替换了3个
	fmt.Println(slice4)  //[100 6 7 0 0 0 0 0]
	
/*测试注*/
	fmt.Println("========================================")
	fmt.Println(math.MaxFloat32) //3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) //1.7976931348623157e+308
	fmt.Println(math.MaxInt8) //127
	fmt.Println(math.MaxInt16) //32767
	fmt.Println(math.MaxInt32) //2147483647
	fmt.Println(math.MaxInt64) //9223372036854775807
	fmt.Println(math.MaxUint8) //255
	fmt.Println(math.MaxUint16) //65535
	fmt.Println(math.MaxUint32) //4294967295
	var maxuint64 uint64= math.MaxUint64
	fmt.Println(maxuint64)  // 18446744073709551615
	//fmt.Println(math.MaxUint64) //错误： constant 18446744073709551615 overflows int，   https://stackoverflow.com/questions/16474594/how-can-i-print-out-an-constant-uint64-in-go-using-fmt
	//报这个错的原因：常量本身是无类型的，常量的类型依赖于其上下文语义，在fmt.Println()时，常量被用作interface{}，因此编译器无法知道您要使用哪种具体类型。  ？？被用作interface？？
	//对于整型常量，在fmt.Println()时，默认为int. 而常量溢出了int的范围（此处int与int64位宽相同），产生编译错误，通过传递uint64（num），可以通知编译器您希望将该值视为uint64。
	fmt.Println(uint64(math.MaxUint64))  // 18446744073709551615
	//fmt.Println()输出时是以int格式进行输出的 ???

	const e = .6326
	const f = 12.
	fmt.Println(e+f) //12.6326

	const P = 1.3e5
	fmt.Println(P) //130000

	const Avogadro = 6.02214129e23  // 阿伏伽德罗常数
	const Planck   = 6.62606957e-34 // 普朗克常数

	fmt.Println(Avogadro) // 6.02214129e+23
	fmt.Println(Planck)	 //6.62606957e-34

}
