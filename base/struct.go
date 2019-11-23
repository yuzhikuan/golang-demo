/*
 * Go没有class
 * struct名称遵循可见性规则
 */

package base

import "fmt"

// 定义一个普通的struct
type person struct {
	Name string
	Age  int
}

// 定义一个嵌套匿名struct的struct
type people struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

// 定义一个匿名字段的struct
type human struct {
	string
	int
}

// Animal struct
type Animal struct {
	Name string
	Age  int
}

// Cat struct --- 嵌入struct作为匿名字段看起来像继承，但不是继承
type Cat struct {
	Animal
	Climb bool
}

// TestStruct 测试函数
func TestStruct() {
	// 使用字面量法初始化struct
	p := person{
		Name: "joe",
		Age:  13,
	}

	// 使用.操作符赋值
	/* p := person{}
	p.Name = "joe"
	p.Age = 19 */

	fmt.Println(p)

	// 验证struct作为函数参数传递的情况
	transmit1(p)
	transmit2(&p)
	fmt.Println(p)

	pe := people{
		Name: "zkyucn",
		Age:  33,
	}
	pe.Contact.Phone = "18575508763" // 使用.操作符给匿名的struct赋值
	pe.Contact.City = "shenzhen"
	fmt.Println(pe)

	h := human{"laoyu", 22} // 类型顺序必须正确
	fmt.Println(h)

	// 嵌套struct作为匿名字段的struct初始化
	c := Cat{Animal: Animal{Name: "xiaohua", Age: 2}, Climb: true}
	c.Animal.Name = "miaomiao"
	c.Age = 3 // 被嵌套的struct的字段，已经绑定到了外层的struct
	fmt.Println(c)
}

// struct类型作为函数参数，此时参数是 值 的拷贝
func transmit1(per person) {
	per.Age = 15
	fmt.Println(per)
}

// *struct类型作为函数参数，此时参数是 指针
func transmit2(per *person) {
	per.Age = 19
	fmt.Println(*per)
}
