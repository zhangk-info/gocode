package main

import (
	"errors"
	"fmt"

	"time"

	"github.com/jinzhu/now"
)

var (
	location, err = time.LoadLocation("Asia/Shanghai")

	myConfig = &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02 15:04:05"},
	}

	Now = time.Now().Format("2006-01-02 15:04:05")
)

func main() {
	fmt.Println("hello world", Now)

	var age int = 1
	fmt.Println("age is", age)
	var a, b int
	fmt.Println("a is", a, "b is", b)
	name, sex := "golang", "male"
	fmt.Println("name is", name, "sex is", sex)
	fmt.Println("科学计数法: -3.14 == -314e-2", -3.14 == -314e-2)
	var ptr *int = &age
	fmt.Println("age的指针地址是", ptr, "指针ptr的值是", *ptr)
	sum, total := add22(1, 2, 3)
	fmt.Println("函数调用", sum, total)
	num := 10
	change2(&num, add22)
	fmt.Println("change()", num)

	fmt.Println("div函数调用", div2(10, 0))
	div22, err := customError2(10, 0)
	fmt.Println("div2函数调用", div22, err)

	fmt.Println("------------------------------------------------------------------------")

	var score2s [5]int = [5]int{1, 2, 3, 4, 5}
	score2s1 := [...]int{1, 2}
	score2s2 := [...]int{0: 1, 1: 2, 8: 0}
	fmt.Println("数组定义的3种方式", score2s, score2s1, score2s2)

	var slice5 = score2s[0:3]
	slice6 := make([]int, 1, 2)
	slice6[0] = 1
	slice7 := []int{1, 2}
	fmt.Println("切片定义的3中方式", slice5, slice6, slice7)

	var map6 map[int]int
	// 映射必须初始化才能使用
	map6 = make(map[int]int, 10)
	map6[1] = 1
	var map7 = make(map[int]int)
	map8 := map[int]int{1: 1, 2: 2}
	fmt.Println("映射定义的3中方式", map6, map7, map8)

}

func add22(a int, b int, beforeTotal int) (sum int, total int) {
	sum = a + b
	total = beforeTotal + sum
	return
}

func change2(a *int, add func(int, int, int) (int, int)) {
	// 对于基础数据类型这里的*省略会编译报错
	*a = 20
	sum, total := add(1, 2, 3)
	fmt.Println("函数作为参数传递并调用", sum, total)
}

// sum sub mul div
func div2(a int, b int) (c int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("div函数捕获到恐慌", err)
		}
	}()
	c = a / b
	return c
}

func customError2(a int, b int) (c int, err error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	} else {
		return a / b, nil
	}
}
