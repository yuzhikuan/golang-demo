/*
 * Interface
 * 接口是一个或多个方法签名的集合
 * 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口
 * 接口只有方法声明，没有实现，没有数据字段
 * 接口可以匿名嵌入其他接口，或嵌入到结构中
 * 空接口可以作为任何类型数据的容器
 */

package base

import "fmt"

// USB 定义了一个接口
type USB interface {
	Connect() bool
}

// Connecter 定义了一个结构
type Connecter struct {
	Name string
}

// Connect 给Connecter添加一个method
func (pc Connecter) Connect() bool {
	fmt.Println("Connected:", pc.Name)
	return true
}

// TestInterface 测试函数
func TestInterface() {
	var usb USB
	usb = Connecter{Name: "PhoneConnecter"} // Connecter结构实现了USB接口
	usb.Connect()

	Disconnect1(usb)
	Disconnect2(usb)
}

// Disconnect1 入参为USB类型，使用if判断类型
func Disconnect1(usb USB) {
	if pc, ok := usb.(Connecter); ok { // 类型判断 usb.(Connecter)
		fmt.Println("Disconnected1:", pc.Name)
		return
	}
	fmt.Println("Unknown decive.")
}

// Disconnect2 入参为 空接口interface{} 类型，使用switch判断类型
func Disconnect2(usb interface{}) {
	switch v := usb.(type) {
	case Connecter:
		fmt.Println("Disconnected2:", v.Name)
	default:
		fmt.Println("Unknown decive.")
	}
}
