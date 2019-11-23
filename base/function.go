/*
 * Go函数 不支持 嵌套、重载和默认参数
 * 支持以下特性：
 * 无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
 * 左大括号不能另起一行
 *
 * defer 在函数体执行结束后按照调用顺序的 相反顺序 逐个执行
 * 即使函数发生 严重错误 也会执行
 *
 * Go 没有异常机制，但有 panic/recover 模式来处理错误
 * Panic 可以在任何地方引发，但 recover 只有 在defer调用的函数中有效
 */

package base

import "fmt"

// TestFunction 测试函数
func TestFunction() {
	x, y, z := a()
	fmt.Println(x, y, z)

	b(4, 5, 6)

	// 匿名函数
	f := func() {
		fmt.Println("Func no name")
	}
	f()

	// 测试闭包
	fc := closure(10)
	fmt.Println(fc(1))
	fmt.Println(fc(2))

	// 测试defer
	testDefer1()
	testDefer2()
	testDefer3()

	// 测试 恐慌/恢复
	testPanic()
}

// 多个返回值，命名返回值参数
func a() (a, b, c int) {
	a, b, c = 1, 2, 3
	return // return a, b, c
}

// ...不定长度变参
func b(a ...int) {
	fmt.Println(a) // a 为slice
}

// 闭包，定义一个返回值为函数的函数
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x) // 打印x的指针，发现三次相同
		return x + y
	}
}

// 打印 2，1，0，defer 按先进后出的顺序执行
func testDefer1() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // i 为函数调用参数，此时为变量值的拷贝
	}
}

// 打印 3，3，3
func testDefer2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // i 为引用主函数内的一个变量，此时为变量地址的拷贝
		}()
	}
}

// 打印 2，1，0
func testDefer3() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i) // i 在定义defer匿名函数作为参数，此时已获得了变量值的拷贝
		}(i)
	}
}

// panic / recover
func testPanic() {
	fmt.Println("Func begin!")
	panicRecover()
	fmt.Println("Func end!")
}

func panicRecover() {
	// defer 要定义在panic之前，不然无法执行到
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover !!!")
		}
	}()
	panic("Panic")
}
