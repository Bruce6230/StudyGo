package main

import (
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

// 定义变量 var name type = expression
func main() {
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

}
