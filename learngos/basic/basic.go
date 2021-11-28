package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
  定义变量
 */

var aa = 3
// aa := 3 // 不支持在函数体外定义

var (
	bb = 4
	ss = "aaa"
)

func variable()  {
	var a int
	var s string
	fmt.Printf("%d %q\n", a,s)
}

/**
  赋值初始值
 */
func variableInitialValue() {
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,s, b)
}

/**
  赋值多类型
 */
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

/**
	省略 var
 */
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

/**
  复数
 */
func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	//欧拉公式
	fmt.Println(
		cmplx.Pow(math.E, 1i * math.Pi) + 1 )

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

}

/**
  类型转换 是 强制的（强制类型转换）
 */
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))

	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a * a + b * b)))
	return c
}

/**
   常量的定义
	const 数值可作为各种类型使用
 */
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a * a + b *b))
	fmt.Println(filename, c)
}

/**
  GoLang 中 枚举类型
 */
func enums() {
	/*const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)*/
	const (
		cpp = iota
		java
		python
		_
		golang
	)
	fmt.Println(cpp, java, python,golang)

}

func main() {
	fmt.Println("hello")

	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss)
	euler()
	triangle()
	consts()
	enums()
}
