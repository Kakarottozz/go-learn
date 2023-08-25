package main

import "fmt"

func zhankai(worlds ...string) int {
	//worlds是指针，改worlds会影响外部传入的实参
	return len(worlds)
}

func main() {
	len1 := zhankai("1", "2", "3")
	fmt.Println(len1)

	worlds := []string{"1", "2", "3", "4"}
	len2 := zhankai(worlds...) //...展开传入
	fmt.Println(len2)
}
