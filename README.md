# GF-go

# 1. go语言介绍

## 1.1 go的由来

> Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。

## 1.2 go的特性

> - 自动垃圾回收
> - 更丰富的内置类型
> - 函数多返回值
> - 错误处理
> - 匿名函数和闭包
> - 类型和接口
> - 并发编程
> - 反射
> - 语言交互性

## 1.3 go代码案例解读

> - 包声明
> - 引入包
> - 函数
> - 变量
> - 语句 & 表达式
> - 注释

```go
package main 

import "fmt"  
func main() { 
   /* Always Hello, World! */
   fmt.Println("Hello, World!")
}
```

- **解析:**

1.  package main定义了包名。**（必须在源文件中非注释的第一行指明这个文件属于哪个包）。**在go应用（一系列go文件及其它文件）的main包中，开始运行，故每个 Go 应用程序都包含一个名为 main 的包。

2. import "fmt"告诉编译器程序运行需要用fmt包。fmt 包实现了格式化 IO（输入/输出）的函数。

3. func main() 是程序开始执行的函数，main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数  **（如果有 init() 函数则会先执行该函数）**。

4. **{}中"{"不可以单独放一行。**

5. 注释： 

   - // 单行注释      

   - /* 块注释（ 多行注释） 不可以嵌套使用，一般用于包的文档描述或注释成块的代码片段。*/   

6. fmt.Println(...) 将字符串输出到控制台，并在最后自动增加换行字符 \n。等价于 fmt.Print("hello, world\n") 。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。

## 1.4 go语言编译与执行

- 编写hello.go（复制上边的文本）

```go
go run hello.go   //编译并运行

go build hello.go //编译生成二进制文件
./hello  // 再进行运行
```

# 2. go环境配置

## 2.1 IDE选择

> IDE需知：vscode与goland，推荐使用vscode，vscode免费，goland收费。
>
> vscode地址：https://code.visualstudio.com/
>
> goland地址：[https://www.jetbrains.com/go/](https://www.jetbrains.com/go/)

## 2.2 go安装包

> 下载地址：
>
> 1. [https://golang.org/](https://golang.org/)    
> 2. [Go下载 - Go语言中文网 - Golang中文社区 (studygolang.com)](https://studygolang.com/dl)

## 2.3 go配置

> - 不管是哪个系统，需要配置go安装环境到path环境变量里面。   **go/bin目录添加到path中**
> - goland只需要选择go对应的sdk配置即可，也就是把go安装的路径添加进去，使用安装的go环境进行编译，不需要安装任何额外插件。
> - vscode需要安装额外插件，在插件市场里面直接搜索go，随后进行安装即可。在安装的时候容易出现下载失败问题，此时需要更换为国内源，如下设置：

```go
windows设置：
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

```go
MacOS or Linux
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

![image-20201214203851365](https://cdn.jsdelivr.net/gh/lizhangjie316/img/2020/20201214203941.png)

---

# 3. 数据类型、关键字、标识符

## 3.1 基本数据类型

> 1. bool类型
> 2. 数字型
> 3. 字符串类型
> 4. 复数类型

> 代码：go-code/01_datetype.go

```go
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

```

### 3.1.1 常量和变量

- 普通赋值

```go
// var 变量名称 变量类型 = 值
var num int = 1
```

- 平行赋值

```go
var num1,num2 int = 1, 2
```

- 多行赋值

```go
var (
    num1 int = 1
    num2 int = 2
)
```

### 3.1.2 整数类型



### 3.1.3 浮点数类型



### 3.1.4 复数类型



### 3.1.5 别名类型



### 3.1.6 字符串类型





### 3.1.7 数组类型



### 3.1.8 切片类型



### 3.1.9 字典类型



### 3.1.10 通道类型



### 3.1.11 注意

```go
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
```

## 3.2 派生类型

> - 指针类型（Pointer）
> - 数组类型
> - 结构化类型(struct)
> - Channel 类型
> - 函数类型
> - 切片类型
> - 接口类型（interface）
> - Map 类型

## 3.3 关键字

### 3.3.1  25个关键字或保留字

break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var

### 3.3.2  36 个预定义标识符

append bool byte cap close complex complex64 complex128 uint16
copy false float32 float64 imag int int8 int16 uint32
int32 int64 iota len make new nil panic uint64
print println real recover string true uint uint8 uintptr

### 3.3.3 额外知识点

- 程序一般由关键字、常量、变量、运算符、类型和函数组成。
- 程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
- 程序中可能会使用到这些标点符号：.、,、;、: 和 …  

## 3.4 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~ Z和a~ z)数字(0~9)、下划线“_”组成的序列，但是第一个字符必须是字母或下划线而不能是数字。







---

# 参考

## 手册



## 网页

1. [Go 语言教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/go/go-tutorial.html)  **by 菜鸟教程**
2. [Go语言入门教程（原C语言中文网） (biancheng.net)](http://c.biancheng.net/golang/)  **by C语言中文网**
3. [Go语言 30 分钟入门教程 ](https://studygolang.com/articles/13958)  **by go语言中文网**





##  博客



## 书籍



