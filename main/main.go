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

/*
结构体：
结构体是一种聚合类型，里面可以包含任意类型的值，这些值就是我们定义的结构体的成员，也常称为字段。在Go语言中定义一个结构体需要使用type+struct关键字组合
在下面例子中，自定义结构体person，有两个字段
*/
type person struct {
	name string
	age  uint
}

// 结构体声明使用
var p person

//or 简短声明法

/*
接口
在Go语言中，接口(interface)是一组方法签名(方法名称、参数和返回值类型的集合)定义的一种类型。接口类型是一个抽象的类型，它并不关心这些方法实现的细节，只关心方法的定义。因此，接口类型允许我们定义一个对象可以实现的多种行为，而不必关心对象的具体实现。

接口定义关键字为 `interface`，定义样例如下：

```go
type 接口名 interface {
    方法名1 (参数列表) 返回类型列表 // 方法签名1
    方法名2 (参数列表) 返回类型列表 // 方法签名2
    ...
}
```

其中，接口名是一个标识符，接口中可以包含一个或多个方法签名。对于方法签名，只需要定义方法的名称、参数列表以及方法的返回值类型。方法的返回类型可以是单个类型或者是由多个类型组成的一个元组。

注意，接口中只包含方法的声明，不包括方法的实现代码。要实现一个接口类型，需要在一个类型上定义这些方法。

示例：

```go
// 定义接口
type Animal interface {
    Speak() string
}

// 定义一个实现动物接口的结构体
type Dog struct {
}

// 实现动物接口中的 Speak 方法
func (d Dog) Speak() string {
    return "woof"
}
```

在上面的示例中，定义了一个名为 `Animal` 的接口，它定义了一个名为 `Speak` 的方法。然后定义了一个 `Dog` 结构体类型，在 `Dog` 结构体类型上实现 `Animal` 接口的 `Speak` 方法。这个 `Speak` 方法传回了一个字符串 "woof"，表示狗类的叫声。这样，类型 `Dog` 就是一个实现了接口 `Animal` 的类型。
*/

// 举例
type Stringer interface {
	String() string
}

//接口的是闲着必须是一个具体的类型
//func (p person) String() string{
//	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
//}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// 以值类型接收者实现接口时，不管类型本身，还是该类型的指针类型，都实现了该接口
// 以指针类型接收者实现接口的时候，只有对应的指针上实现了该接口
func (p *person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

//工厂函数
//在Go语言中，一个工厂函数可以用来创建新的对象，而不是直接使用结构体字面量来创建。
//
//工厂函数是一个普通函数，不同于普通函数的是，它返回一个结构体变量或接口的实例。通常，工厂函数有以下的命名约定，使其更加易于理解：
//
//- `New` + 结构体名称 或
//- `New` + 接口名称
//
//下面是一个简单的例子，演示如何使用工厂函数创建和初始化一个结构体的实例：
//
//```go
//package main
//
//import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

// 工厂函数
func NewPerson(name string, age int, address string) *Person {
	return &Person{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

//
//```
//
//上面的代码定义了一个 `Person` 结构体和一个命名为 `NewPerson` 的工厂函数，用于创建新的 `Person` 结构体实例。`NewPerson`函数接受三个参数，返回一个指向新分配的 `Person` 结构体的指针，该结构体含有传递的三个参数。最后在 `main` 函数中，调用 `NewPerson` 函数创建一个指向 `Person` 实例的指针，并使用打印语句将结构体中的字段打印到标准输出。

// 初始化结构体使用{}初始化字段值
// 工厂函数，返回一个error接口，其实实现时*errorString
func New(text string) error {
	return &errorString{text}
}

// 结构体，内部一个字段s，存储错误信息
type errorString struct {
	s string
}

// 实现error接口
func (e *errorString) Error() string {
	return e.s
}

/*
Go中没有继承的概念，所以结构，接口之间也没有父子关系，Go语言提倡组合，利用组合和接口达到代码复用的目的
*/
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

// 结构体组合
type Animal struct {
	name string
}

func (a *Animal) Eat() {
	fmt.Printf("%s is eating...", a.name)
}

type Dog struct {
	*Animal
}

func (d *Dog) Bark() {
	fmt.Printf("%s is barking...", d.name)
}

type Cat struct {
	*Animal
}

func (c *Cat) Meow() {
	fmt.Printf("%s is meowing...", c.name)
}

//在 Go 语言中，类型断言（Type Assertion）用于将一个接口类型的变量转换为另一个具体类型的变量。
//
//类型断言语法如下：
//
//```go
//x.(T)
//```
//
//其中，`x` 为接口类型变量，`T` 为断言的具体类型。该语法会将 `x` 转换为类型 `T`，如果断言成功，则返回一个具体类型 `T` 的值以及一个 `true` 值；否则返回一个该类型的零值以及一个 `false` 值。
//
//例如，我们可以将一个 `interface{}` 类型的变量 `x` 转换为一个 `int` 类型的变量：
//
//```go
//func foo(x interface{}) {
//    i, ok := x.(int)
//    if ok {
//        fmt.Println("x is an int:", i)
//    } else {
//        fmt.Println("x is not an int")
//    }
//}
//
//func main() {
//    foo("hello")
//    foo(123)
//}
//```
//
//在上面的代码中，`foo` 函数接受一个 `interface{}` 类型的参数 `x`，然后通过断言将其转换为 `int` 类型。如果转换成功，则输出 `x is an int` 并打印出具体的值 `i`；否则输出 `x is not an int`。
//
//在 `main` 函数中，我们分别调用 `foo` 函数，并传递了一个字符串和一个整数参数。第一次调用中，由于字符串无法转换为整数类型，所以输出 `x is not an int`；第二次调用中，整数参数成功转换为了整数类型并输出。
//
//需要注意的是，在使用类型断言时，如果断言的类型不是接口类型的底层类型或其实现类，则会导致运行时错误。因此，我们在使用类型断言时需要谨慎处理可能出现的错误情况。

func main() {
	per := NewPerson("Alice", 20, "Beijing")
	poo := person{"Makiyo", 30}
	printString(&poo)
	fmt.Println(poo.name, poo.age, per.Name)
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
