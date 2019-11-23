package base

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("map module init function")
}

// TestMap 测试函数
func TestMap() {
	var m map[int]string
	// make创建一个map
	m = make(map[int]string)
	m[1] = "OK"
	fmt.Println(m)

	// 创建一个map[int]string类型的slice
	s := make([]map[int]string, 5, 10)
	for i := range s {
		// 初始化每一个map，并为属性“1”赋值
		s[i] = make(map[int]string)
		s[i][1] = "GOOD"
	}
	fmt.Println(s)

	// 字面量法创建map
	mp := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	sl := make([]int, len(mp)) // len(mp)获取key-value对数
	ix := 0
	// 遍历map
	for k := range mp {
		sl[ix] = k
		ix++
	}
	// 整形数组排序
	sort.Ints(sl)
	fmt.Println(sl)

	// 将map[int]string变换成map[string]int的形式
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := make(map[string]int)
	for k := range m1 {
		m2[m1[k]] = k
	}
	fmt.Println(m1, m2)
}
