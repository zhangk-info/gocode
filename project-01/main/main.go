package main

import (
	"errors"
	"fmt"
	"gocode/project-01/utils"
	"unsafe"
)

// 全局变量 必需用var声明
// var version = "1.0"
// var time = "2023/05/10"
var (
	version = "1.0"
	time    = "2023/05/10"
)

func main() {
	// go build test.go + test.exe  = go run test.go
	fmt.Println("hello world")

	// 变量
	var age int = 1
	fmt.Println("age is ", age)
	// 不初始化定义类型会有默认
	var a, b int
	fmt.Println("var a,b int：a is", a, "and b is", b)
	// 类型推断
	// var name1,sex2 = "golang","male"
	name, sex := "golang", "male"
	// Print 原样输出 且不换行
	// Printf 格式输出 且不换行 https://studygolang.com/pkgdoc fmt
	// Println  值+空格原样输出（不会格式化） 且换行
	// %v — 值的默认格式
	// %s — 字符串
	// %d — 10进制数值
	// %T — type(值)
	fmt.Println("Println：name is", name, ",sex is", sex)
	fmt.Printf("Printf\\n：name is %s and type is %T and size is %d\n", name, name, unsafe.Sizeof(name))
	// e-2表示乘以10(-2次幂)
	fmt.Println("科学计数法：-3.14==-314E-2", -314e-2, -3.14 == -314e-2)

	fmt.Println("age对应的存储空间的地址为：", &age)
	var ptr *int = &age
	fmt.Println("ptr这个指针指向的数值为：", *ptr)

	sum, _ := add(1, 2, 3)
	sum, total := add(1, 2, 3)
	fmt.Println("add函数调用:", sum, total)

	num := 10
	fmt.Println("内存地址是：", &num)
	change(&num)
	fmt.Println("传递内存地址作为入参（指针），改变指针指向的内存的值", num)

	var t = test
	fmt.Printf("test2的类型是%T,test的类型是%T\n", t, test)
	fmt.Println("通过变量调用")
	t(1)
	fmt.Println("将函数作为入参传递")
	test2(t)

	type myInt int
	var num1 myInt = 1
	fmt.Printf("myInt的类型%T\n", num1)
	fmt.Printf("强转后的类型%T\n", int(num1))

	fmt.Println("将函数作为自定义类型")
	test3(t)

	sum2, total2 := add2(1, 2, 3)
	fmt.Println("函数返回值命名add2函数调用:", sum2, total2)

	fmt.Println("调用其他包的函数", utils.Add(1, 2))
	div1 := div(10, 0)
	fmt.Println("除法调用", div1)

	err := customError(10, 0)
	if err != nil {
		fmt.Println("自定义异常", err)
		// 可以通过panic终端程序
		// panic(err)
	}

	fmt.Println("执行完成")
}

func add(a int, b int, total int) (int, int) {
	total += a + b
	return a + b, total
}

// 使用名称直接返回
func add2(a int, b int, beforeTotal int) (sum int, total int) {
	sum = a + b
	total = beforeTotal + a + b
	return
}

// 参数类型为指针
func change(ptr *int) {
	// 对地址对应的变量进行赋值 *a返回指针变量a的值
	*ptr = 30
}

type funcTest func(int)

func test(num int) {
	fmt.Println("test", num)
}

func test2(a func(int)) {
	a(2)
}

func test3(a funcTest) {
	a(3)
}

func div(a int, b int) (div int) {
	defer func() {
		//调用recover内置函数，可以捕获错误
		err := recover()
		// 如果没有捕获错误，返回值为零值：nil
		if err != nil {
			fmt.Println("div捕获到错误", err)
		}
	}() // 加上匿名函数的调用
	div = a / b
	return
}

func customError(a int, b int) (err error) {
	if b == 0 {
		return errors.New("除数不能为0")
	} else {
		return nil
	}
}
