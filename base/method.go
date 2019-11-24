package base

import "fmt"

// A struct
type A struct {
	Name string
}

// B struct
type B struct {
	Name string
}

// TZ 定义了一个 int 的类型别名 TZ
type TZ int

// TestMethod 测试函数
func TestMethod() {
	a := A{}
	a.Print()
	fmt.Println(a)

	b := B{}
	b.Print()
	fmt.Println(b)

	var tz TZ
	tz.Print()       // Method Value
	(*TZ).Print(&tz) // Method Expression

	var t TZ
	t.Increase(100)
	fmt.Println(t)
}

// Print 将*A传递给reciver
func (a *A) Print() {
	a.Name = "AA" // 此处可以修改A struct的字段
	fmt.Println("A")
}

// Print 将B以值得方式传递给reciver
func (b B) Print() {
	b.Name = "BB" // 此处不会修改B struct的原始字段的值
	fmt.Println("B")
}

// Print 将类型别名指针传递给reciver
func (a *TZ) Print() {
	fmt.Println("TZ")
}

// Increase 测试 类型 和 类型别名 的相互转换
func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}
