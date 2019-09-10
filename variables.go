package main

import (
	"math/cmplx"
	"fmt"
	"math"
)

// var aa = 3
// var ss= "kkk" //在函数外面定义变量必须用var
// var bb = true //作用域在包内部

var (
	aa = 3
	ss = "kkk" //在函数外面定义变量必须用var
	bb = true //作用域在包内部
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a,  s)
}
func variableInitValue(){
	var a,b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction(){
	var a, b, c, s   = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}
func variableShorter(){
	a, b, c, s := 3, 4, true, "abc" //省略var，用冒号:定义，只能在函数内使用
	b = 5
	fmt.Println(a, b, c, s)
}
// 内建变量类型  加u是无符号整数，不加u是有符号整数
func euler() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))

	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b = 3, 4  // const 数值可作为各种类型使用
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举类型
func enums() {
	const (
		cpp = iota  // 自增值
		_
		python
		golang
		javascript
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
// 变量类型写在变量名之后
// 编译器可推测变量类型
// 没有char，只有rune
// 原生支持复数类型
func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()

	triangle()

	consts()

	enums()
}

//类型转换是强制的