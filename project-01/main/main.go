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

	fmt.Println("------------------------------------------------------------------------")
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

	fmt.Println("------------------------------------------------------------------------")
	sum, total := add(1, 2, 3)
	// sum, _ := add(1, 2, 3)
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

	fmt.Println("------------------------------------------------------------------------")
	div2, err := customError(10, 0)
	if err != nil {
		fmt.Println("自定义异常", err)
		// 可以通过panic终端程序
		// panic(err)
	} else {
		fmt.Println("除法调用", div2)
	}

	fmt.Println("------------------------------------------------------------------------")

	var scores [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("常规定义", scores)
	scores2 := [...]int{1, 2}
	fmt.Println("类型推断+长度自动计算", scores2)
	scores3 := [...]int{1: 4, 0: 5, 6: 6}
	fmt.Println("类型推断+长度自动计算+k,v定义", scores3)

	var sum3 int
	for i := 0; i < len(scores); i++ {
		sum3 += scores[i]
	}
	fmt.Println("数组", scores, ",平均数是：", sum3/len(scores))
	fmt.Printf("数组的指针地址是%p,类型是%T,指针类型是%T;数组第一个空间的指针地址是%p,类型是%T,指针类型是%T\n", &scores3, scores3, &scores3, &scores3[0], scores3[0], &scores3[0])
	var ptr1 *int = &scores[0]
	var ptr2 *[5]int = &scores
	// 通过指针操作值（）
	(*ptr2)[0] = 10
	fmt.Println("从数组指针地址取值是", *ptr1, *ptr2)
	// fmt.Println("请录入一个成绩")
	// fmt.Scanln(&scores[1])
	// fmt.Println("数组", scores, ",平均数是：", sum3/len(scores))
	// key value 遍历
	for k, v := range scores2 {
		// 这里对于数组k是下标0开始
		fmt.Println(k, " ", v)
	}

	// 区别：多维数组后面维度长度要指定
	scores4 := [...][9]int{{0: 1, 6: 2}, {1: 5, 8: 9}}
	fmt.Println("二维数据初始化", scores4)
	fmt.Println("------------------------------------------------------------------------")

	var slice []int
	// 这里的start:end都可以忽略不写，如[2:][:5][:]
	slice = scores[2:5]
	// 切片容量自动扩容。容量<1024时，以2倍递增。容量超过1024时，容量变为原来的1.25倍。如果一次插入后长度大于原容量2倍时会以新长度为基准
	// 切片扩容机制是开辟新的连续空间并拷贝原值，再插入
	fmt.Printf("切片类型：%T\n", slice)
	fmt.Println("原数据", scores, "切片  1:3  包前不包后", slice, "长度为", len(slice), "容量为", cap(slice))
	// make(切片类型，长度，容量)
	var slice2 []int = make([]int, 8, 20)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	slice2[3] = 4
	fmt.Printf("切片类型：%T\n", slice2)
	fmt.Println("切片", slice2, "长度为", len(slice2), "容量为", cap(slice2))
	// 切片类型这里[]没有...计算长度
	slice3 := []int{1, 2}
	slice3 = append(slice3, 1, 2)
	slice3 = append(slice3, slice2...)
	fmt.Printf("切片类型：%T\n", slice3)
	fmt.Println("切片", slice3, "长度为", len(slice3), "容量为", cap(slice3))

	slice4 := []int{1, 2}
	// 原值是[1,2],方法中将0下标改成10并新增2个元素[3，4]
	// 方法内切片 [10 2 3 4] 长度为 4 容量为 4
	transferSlice(slice4)
	// 最终slice4切片为[10 2] 长度为 2 容量为 2
	// 第一个值改了但是超出部分并没有在slice4,因为append方法是开辟新的连续空间并拷贝原值，再插入
	fmt.Println("切片传递", slice4, "长度为", len(slice4), "容量为", cap(slice4))

	// copy(target,source),多出的长度不会复制
	copy(slice2, slice3)
	fmt.Println("切片", slice2, "长度为", len(slice2), "容量为", cap(slice2))
	fmt.Println("------------------------------------------------------------------------")

	var map1 map[int]int
	map1 = make(map[int]int, 10)
	map1[1] = 1
	map1[2] = 2
	map1[1] = 3
	fmt.Println("映射", map1)
	// 由于动态扩容，创建时可以不指定长度
	map2 := make(map[int]string)
	map2[1] = "1"
	fmt.Println("映射", map2, len(map2))
	map3 := map[int]string{1: "1", 2: "2"}
	fmt.Println("映射", map3, len(map3))
	// 删除
	delete(map1, 1)
	fmt.Println("映射 删除", map1, len(map1))
	// 清空，1重新初始化，2遍历删除
	map2 = make(map[int]string)
	fmt.Println("映射 清空", map2, len(map2))
	// 查找
	num6, hasNum6 := map1[6]
	fmt.Println("映射", map1, "查找6", hasNum6, num6)

	map5 := make(map[int]string)
	map5[1] = "1"
	// 映射传递后，方法内所有操作都会影响原映射
	transferMap(map5)
	fmt.Println("映射传递", map5, len(map5))

	// 加深难度,map的value可以是任何值
	map4 := map[string]map[int]string{"张": {1: "1", 2: "2"}, "王": {3: "3"}}
	fmt.Println("映射", map4, len(map4), len(map4["王"]))

	fmt.Println("------------------------------------------------------------------------")
	// 演示一下管道的使用
	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	//2. 看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	//3. 向管道写入数据
	intChan <- 10
	intChan <- 211
	intChan <- 50
	// //如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 98 //注意点, 当我们给管写入数据时，不能超过其容量
	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3
	//5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3
	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	//num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom jack"
	dog := Dog{"金毛巡回犬", "金色", 12}
	allChan <- dog
	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan
	newDog := <-allChan //从管道中取出的Dog是什么?
	fmt.Printf("newDog=%T , newDog=%v\n", newDog, newDog)
	//对于数据使用前，使用类型断言
	newDog2, isDog := newDog.(Dog)
	if isDog {
		fmt.Printf("newDog2.Name=%v", newDog2.Type)
	} else {
		fmt.Printf("不是一个Dog类型")
	}

	fmt.Println("------------------------------------------------------------------------")
	// 直接定义有默认值
	var husky Dog
	husky.Type = "哈士奇"
	husky.Age = 45
	husky.Color = "红色"
	fmt.Println("哈士奇", husky)
	// 定义并初始化
	goldenRetriever := Dog{"金毛巡回犬", "金色", 12}
	fmt.Println("金毛巡回犬", goldenRetriever)
	// new的方式返回指针
	var borderCollie *Dog = new(Dog)
	(*borderCollie).Type = "边牧"
	(*borderCollie).Age = 6
	// go编译器底层对指针取值进行了推断，可以不写* ！！
	borderCollie.Color = "银色"
	fmt.Println("边牧", *borderCollie)
	// &方式返回指针
	var komaki *Dog = &Dog{"古牧", "灰色", 15}
	komaki.Age = 16
	var komaki2 *Dog2 = &Dog2{}
	// Dog和Dog2一模一样所以可以强转
	*komaki2 = (Dog2)(*komaki)
	fmt.Println("古牧", *komaki, *komaki2)

	fmt.Println("------------------------------------------------------------------------")
	var dog2 = Dog{"哈士奇", "黑白", 12}
	res := dog2.bit(3) // 哈士奇汪汪汪
	// 传递指针调用 这里可以简写，编译器可以自动处理，会自动加上&和*
	// (&dog2).bit2(3)
	dog2.bit2(3)
	fmt.Println("bit方法调用", res, dog2)

	fmt.Println("------------------------------------------------------------------------")
	num7 := 1
	transferPtr(&num7)
	fmt.Println("指针传递", num7)

	fmt.Println("执行完成")
}

type Dog2 struct {
	// 变量名字大写外界可以访问属性
	Type  string
	Color string
	Age   int
}

type Dog struct {
	// 变量名字大写外界可以访问属性
	Type  string
	Color string
	Age   int
}

func transferPtr(a *int) {
	*a = 10
}

func (dog Dog) bit(bitTime int) (res string) {
	res += dog.Type
	for i := 0; i < bitTime; i++ {
		res += "汪"
	}
	return
}

func (dog *Dog) bit2(bitTime int) (res string) {
	dog.Age = 13
	dog.Color = "金色"
	return
}

func transferMap(map1 map[int]string) {
	map1[1] = "2"
	map1[2] = "2"
}

func transferSlice(slice []int) {
	slice[0] = 10
	slice = append(slice, 3)
	slice = append(slice, 4)
	fmt.Println("方法内切片", slice, "长度为", len(slice), "容量为", cap(slice))
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

func customError(a int, b int) (div int, err error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	} else {
		div = a / b
		return div, nil
	}
}
