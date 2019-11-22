/*
 * 数组的长度是固定的，初始化时需要「明示」或「暗示」数组的长度
 * 数组的长度是数组类型的组成部分，[2]int 与 [100]int 是不同类型的数组
 * 使用 for ... range 遍历数组
 * 在 Go 语言中，数组是值类型，赋值和传递参数都会发生数组的复制
 * 数组指针是一个指针，它指向了一个数组
 * 指针数组是一个数组，它里面装着指针
 */

package base

import (
	"fmt"
)

var a1 = [5]int{1, 2, 3, 4, 5}
var a2 = new([5]int)
var a3 = [...]int{0: 1, 2: 3}

// TestArray 是一个测试数组的函数
func TestArray() {
	fmt.Println(a1)
	fmt.Printf("%v\n", a2)

	for index, value := range a3 {
		fmt.Printf("下标 = %d, 值 = %d\n", index, value)
	}

	a4 := a1
	fmt.Printf("数组 a1 - 值：%v，指针：%p\n", a1, &a1)
	fmt.Printf("数组 a4 - 值：%v，指针：%p\n", a4, &a4)

	transmitA(a1)

	m, n := 1, 2
	pointerArray := [2]*int{&m, &n}
	fmt.Printf("指针数组%v\n", pointerArray)
	fmt.Printf("数组指针%v\n", &a1)
}

func transmitA(a [5]int) {
	fmt.Printf("传入函数的数组 a - 值：%v，指针：%p\n", a, &a)
}

/* 在同一个package中，可以多个文件中定义init方法
 * 在同一个go文件中，可以重复定义init方法
 * 在同一个package中，不同文件中的init方法的执行按照文件名先后执行各个文件中的init方法
 * 在同一个文件中的多个init方法，按照在代码中编写的顺序依次执行不同的init方法
 */
func init() {
	fmt.Println("base module init function")
}
