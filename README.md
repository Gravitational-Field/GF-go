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

## 3.1 数据类型

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





## 3.2 关键字

### 3.2.1  25个关键字或保留字

break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var

### 3.2.2  36 个预定义标识符

append bool byte cap close complex complex64 complex128 uint16
copy false float32 float64 imag int int8 int16 uint32
int32 int64 iota len make new nil panic uint64
print println real recover string true uint uint8 uintptr

### 3.2.3 额外知识点

- 程序一般由关键字、常量、变量、运算符、类型和函数组成。
- 程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
- 程序中可能会使用到这些标点符号：.、,、;、: 和 …  

## 3.3 标识符

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



