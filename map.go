package main

import "fmt"

func main() {
	temp := map[string]int{
		"a": 1,
	}
	t1 := temp["a"]
	temp["b"] = 2
	fmt.Println(t1)
	t2 := temp["c"] //不存在不会报错，而是返回零值
	fmt.Println(t2)

	if c, ok := temp["c"]; ok { //c是元素，ok表示是否存在
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
	//map不传副本
	map1 := map[string]string{
		"a": "1",
		"b": "2",
	}
	map2 := map1

	map1["a"] = "2" //map1与map2同一底层数组，改map1，map2也会变
	fmt.Println(map1)
	fmt.Println(map2)

	delete(map1, "a")

	//make初始化map，要么字面值初始化，要么make预分配初始化
	map3 = make(map[float64]int, 8) //8个键值对

	for key, value := range map3 { //不一定按原本顺序

	}

	groups := make(map[float64][]float64) //key是float64，value是float64的切片
	//go没有set数据结构

}
