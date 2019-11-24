/*
 * 反射 Reflect
 * 想要利用反射修改对象状态，前提是interface.data 是 settable，即pointer-interface
 * 通过反射可以 “动态” 调用方法
 */

package base

import (
	"fmt"
	"reflect"
)

// User 定义了一个struct
type User struct {
	ID   int
	Name string
	Age  int
}

// Manager 定义
type Manager struct {
	User
	Title string
}

// Speak User结构体绑定了speak方法
func (u User) Speak(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}

// Info 打印入参的属性和方法
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	// 对反射目标对象进行类型验证，看是否为reflect struct
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Param is't reflect struct")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	// 遍历出struct的字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// 遍历出struct的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("Methods:\n%6s: %v \n", m.Name, m.Type)
	}
}

// Set 通过反射对 struct类型 进行修改
func Set(o interface{}) {
	v := reflect.ValueOf(o)

	// 校验被反射者的类型是否正确，是否可以修改
	if v.Kind() == reflect.Ptr && v.Elem().CanSet() {
		v = v.Elem()
	} else {
		fmt.Println("被反射者类型不正确")
	}

	// 取出要修改的字段，并校验是否存在
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
	}

	// 判断要修改的字段的类型，并调用类型对应的方法去修改其值
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}

// TestReflect 测试函数
func TestReflect() {
	u := User{1, "OK", 12}
	Info(u)

	// 对嵌套的struct通过反射进行取值操作
	m := Manager{User: User{1, "OK", 12}, Title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) // 取匿名字段的属性ID

	// 通过反射对 简单类型 进行修改
	x := 123
	v := reflect.ValueOf(&x) // 此时接受的是一个指针
	v.Elem().SetInt(999)
	fmt.Println(x)

	// 测试通过反射修改struct类型的字段值
	Set(&u)
	fmt.Println(u)

	// 通过反射，调用struct的方法
	val := reflect.ValueOf(u)
	mv := val.MethodByName("Speak")
	args := []reflect.Value{reflect.ValueOf("joe")} // 构建参数slice
	mv.Call(args)                                   // Call需要一个slice作为参数序列
}
