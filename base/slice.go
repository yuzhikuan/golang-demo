package base

import "fmt"

// TestSlice 测试函数
func TestSlice() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)
	// 从数组创建slice
	s1 := a[5:10]
	fmt.Println(s1)

	// 使用make创建slice
	s2 := make([]int, 3, 10)
	fmt.Println(len(s2), cap(s2))

	a2 := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s3 := a2[2:5] // [c, d, e]
	// reslice
	s4 := s3[1:3] // [d, e]
	// 超出索引的部分指向原数组
	s5 := s3[3:5] // [f, g]
	fmt.Println(string(s3), string(s4), string(s5), len(s3), cap(s3))

	s6 := make([]int, 3, 6)
	fmt.Printf("%p\n", s6)
	// append 超出容量会重新创建
	s6 = append(s6, 1, 2, 3, 4)
	fmt.Printf("%v %p\n", s6, s6)

	a3 := [5]int{1, 2, 3, 4, 5}
	s7 := a3[2:5]
	s8 := a3[1:4]
	fmt.Println(s7, s8)
	// 两个slice指向同一个数组，修改其中一个slice，会影响两一个slice（共同部分）
	s7[0] = 8
	fmt.Println(s7, s8)

	// append超出容量会新建一个slice，不在指向原数组，这是再修改一个slice，另一个不收影响
	s8 = append(s8, 1, 1, 1, 1, 1)
	s7[0] = 9
	fmt.Println(s7, s8)

	// copy会覆盖相同索引的值，超出目标slice的部分会丢弃
	s9 := []int{1, 2, 3, 4, 5}
	s10 := []int{7, 8, 9, 0}
	copy(s9, s10)
	fmt.Println(s9, s10)
}
