package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//hello world
//import "fmt"
//
//func main() {
//	fmt.Println("hello world")
//}

//函数定义声明
//func funcName(params) result{
//	body
//}
/*
关键字 func
函数名字 funcName
函数的参数 params，用来定义形参和变量名和类型，可以有一个参数，多个或者没有
返回的函数值 result 用于定义返回值的类型型，如果没有返回值，省略即可，也可以有多个返回值
函数体 body 在这里写函数逻辑
*/

func sum(a int, b int) int {
	return a + b
}

// 多指返回
func dec(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能为负数")
	}
	return a - b, nil
}

// 命名返回参数
func f(a int, b int) (sum int, err error) {
	if a <= 0 || b <= 0 {
		sum = 0
		err = errors.New("a,b取值错误")
		return
	}
	sum = a + b
	err = nil
	return
}

// 可变参数
func sum1(params ...int) int {
	temp := 0
	for _, i := range params {
		temp += i
	}
	return temp
}

//函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；
//函数名称首字母大写代表共有函数，不同的包也可以调用
//任何一个函数都会从属于一个包
//tips:Go没有public private这样的修饰符来修饰函数公有还是私有，而是通过函数名称大小写来代表，更简洁

//匿名函数和闭包：顾名思义，匿名函数就是没有名字的函数，这是他和正常函数的主要区别，下面示例中,sum2嗲表一个匿名函数，这个sum2时一个函数类型的变量，并不是函数的名字
/**
函数类型的变量和函数名在语言中的作用不同：

函数类型的变量是一个特殊的类型，它可以存储一个函数的引用。这个类型的变量通常被称为函数指针或函数类型的指针。函数指针可以赋值给其他函数指针，也可以用作函数的参数或返回值类型。

函数名则是表示一个函数的标识符。你可以使用函数名来调用函数，并且可以传递函数名作为参数传递给其他函数。

在很多编程语言中，函数名和函数可以看作是等价的，因此可以将函数名赋值给一个函数类型的变量。在这种情况下，函数名实际上是一个指向该函数的指针。

在其他一些语言中，函数名和函数是不同的，在这些语言中不能将函数名直接赋值给函数类型的变量，必须使用特殊的语法来表示函数的引用。
*/
func show() {
	sum2 := func(a int, b int) int {
		return a + b
	}
	fmt.Println(sum2(1, 2))
}

// 方法：方法和函数是两个概念，但是非常相似，不同在于方法必须有一个接收者，这个接收者是一个类型，这样方法就和类型绑定在一起，成为这个类型的方法
type Age uint

func (age Age) string() {
	fmt.Println(age)

}
func (age *Age) Modify() {
	*age = Age(30)
}

/*
tips:在调用方法的时候，传递的接收者本质上都是副本，只不过一个是这个值副本，一个是指向这个值指针的副本。
示例中调用指针接收者方法的时候，使用的是一个值类型的变量，并不是一个指针类型，其实这里使用指针变量调用也是可以的，如下面的代码所示：
(&age).Modify
Go语言编译器自动：
如果使用一个值类型变量调用指针接受者的方法，Go语言编译器会自动帮我们取指针调用，以满足指针接收者的要求
如果一个在指针类型变量调用值类型结束这方法，Go语言会自动帮我们转义，大大提高开发者效率，但是要注意bug
*/
// 定义变量 var name type = expression
func main() {
	aaa := Age(10)
	fmt.Println(aaa)
	aaa.Modify()
	aaa.string()
	age := Age(25)
	age.string()
	show()
	fmt.Println(sum1(1, 2, 3, 4))
	fmt.Println(f(10, 2))
	fmt.Println(dec(10, 2))
	var i int = 10
	fmt.Println(i)
	//也可以不用指定变量类型
	var (
		j = 0
		k = 1
	)
	fmt.Println(j + k)
	var (
		float32 = 3.2
		flag    = false
		s       = "hello"
	)
	fmt.Println(float32, flag, s)
	/**
	0为变量默认
	**/
	//变量简短声明 结构：变量名:=表达式
	t := 10
	fl := false
	fmt.Println(t, fl)
	//指针
	var pi = &i
	*pi = 15
	what := pi
	fmt.Println(*pi + *what)
	//常量
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)
	//Go是强类型语言

	//int转string
	i2s := strconv.Itoa(i)
	//string转int
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)
	//判断s的前缀是否是h
	fmt.Println(strings.HasPrefix(s, "h"))
	//再s中查找字符串l
	fmt.Println(strings.Index(s, "l"))
	//将s全部转为大写
	fmt.Println(strings.ToUpper(s))
	/**
	len(s string) int：返回字符串s的长度。

	s[i]：获取字符串s中的第i个字符。

	s[i:j]：获取字符串s中从第i个字符到第j-1个字符组成的子串。

	s + t：将两个字符串s和t拼接起来。

	strings.Contains(s, substr string) bool：判断字符串s是否包含子串substr，返回bool类型。

	strings.Index(s, sep string) int：在字符串s中查找sep第一次出现的位置，返回位置索引，如果不存在返回-1。

	strings.Count(s, sep string) int：计算字符串s中sep出现的次数。

	strings.Replace(s, old, new string, n int) string：将字符串s中前n个old子串替换为new子串，返回新字符串。

	strings.Split(s, sep string) []string：将字符串s按照sep分隔成多个子串组成的切片。

	strings.ToLower(s string) string：将字符串s中的字母全部转换为小写。

	strings.ToUpper(s string) string：将字符串s中的字母全部转换为大写。

	strings.TrimSpace(s string) string：去除字符串s前后的空格。
	**/
	//r:=strings.Index("飞雪无情","飞")
	//fmt.Println(r)
	//条件
	//if i > 0 {
	//	fmt.Println("+" + strconv.Itoa(i))
	//} else {
	//	fmt.Println("-" + strconv.Itoa(i))
	//}
	////选择结构
	//switch i {
	//case 1:
	//	fallthrough
	//case 15:
	//	fmt.Println("yes")
	//default:
	//	fmt.Println("没有匹配")
	//}
	////fallthrough是一个控制流关键字，可用于在 switch 语句中，继续执行下一个分支的代码块。
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//for i >= 0 {
	//	sum += i
	//	i--
	//}
	//fmt.Println(sum, i)
	//for {
	//	sum += i
	//	i++
	//	if i > 100 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	//数组，集合类型
	//数组声明：array:=[5]string{"a","b","c","d","e"}
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	for j := 0; j < 5; j++ {
		fmt.Printf("数组索引：%d,对应值：%s\n", j, array[j])
	}
	//for range形式 第一个是索引，第二个是数组的值
	//arr := []string{"hello", "world", "Go"}
	//for index, value := range arr {
	//	fmt.Printf("Index: %d, Value: %s\n", index, value)
	//}
	//arr := []string{"hello", "world", "Go"}
	//for _, value := range arr {
	//	fmt.Println(value)
	//}

	//slice切片 左闭右开
	//arr := []string{"hello", "world", "Go"}
	slice := array[0:3]
	sl := array[:]
	fmt.Println(sl)
	fmt.Println(slice)
	for j := 0; j < 5; j++ {
		fmt.Printf("数组索引：%d,对应值：%s\n", j, array[j])
	}
	sl[1] = "pro"
	fmt.Println(sl)
	//切片是一个拥有三个字段的数据结构，分别是指向数组的指针data,长度Len,容量cap

	//使用 make 函数创建切片时，需要指定切片的长度和容量：
	l := make([]int, 5, 10) // 创建长度为 5，容量为 10 的切片
	for index, value := range l {
		fmt.Println(index, value)
	}
	//通过append对一个切片追加元素，返回新切片
	tt := append(l, 10)
	fmt.Println(tt)
	//向切片中添加单个元素，可以通过 append 函数添加单个元素到切片的末尾：
	//f := []int{1, 2, 3}
	//f = append(f, 4)
	//fmt.Println(f) // 输出 [1 2 3 4]
	//向切片中添加多个元素,可以通过 append 函数同时添加多个元素到切片的末尾：
	//f := []int{1, 2, 3}
	//f = append(f, 4, 5, 6)
	//fmt.Println(f) // 输出 [1 2 3 4 5 6]
	//将一个切片添加到另一个切片的末尾,可以通过 append 函数将一个切片添加到另一个切片的末尾：
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println(s1) // 输出 [1 2 3 4 5 6]
	//Map声明初始化
	nameAgeMap := make(map[string]int)
	nameAgeMap["A"] = 10
	//声明一个空map：
	//var m map[string]int
	//使用make函数初始化一个map：
	//m := make(map[string]int)
	//使用键值对初始化一个Map
	//m := map[string]int{
	//	"apple":  5,
	//	"orange": 2,
	//	"banana": 3,
	//}
	//	需要注意的是，对于一个空的map，不能直接添加键值对。如果需要添加键值对，必须先使用make函数或键值对初始化map，然后才能进行添加。例如：
	//	m := make(map[string]int)
	//	m["apple"] = 5
	//	m["orange"] = 2
	//	m["banana"] = 3
	//map遍历
	m := map[string]int{
		"apple":  5,
		"orange": 2,
		"banana": 3,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	str := "hello world"
	bs := []byte(str)
	fmt.Println(bs)
	fmt.Println(bs[0])
	fmt.Println(len(str))
	for m, n := range str {
		fmt.Println(m, n)
	}
	//函数声明
	//func funcName(params) result{
	//	body
	//}
	/*
		func 关键字
		funcName 函数名字
		params 函数的参数，用来定义形参的变量名和类型，可以有一个参数也可以有多个，也可以没有
		result 返回的函数值，用于定义返回值的类型，如果没有返回值 省略即可，也可以有多个返回值
		body函数体
	*/
	//func sum(a int,b int) int{
	//	return a+b;
	//}

}
