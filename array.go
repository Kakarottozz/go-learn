package main

import "fmt"

func main() {
	var planets [8]string //8个string的数组
	fmt.Println(planets)
	//复合字面值
	planets1 := [5]string{"1", "2", "3", "4", "5"}
	planets2 := [...]string{"1", "2", "3", "4", "2"}
	fmt.Println(planets2)
	for i, p := range planets1 { //默认获得索引和值
		fmt.Println(i, p)
	}

	//数组的赋值和按值传递都会发生拷贝，一般用切片传参
	//不同长度的数组是不同的类型

	var board [8][8]byte
	board[0][0] = 'c'

	for column := range board[1] {
		board[1][column] = 'x'
	}

	//切片，左闭右开,索引不能是负数，不同于python
	slice := "abcde"
	slice1 := slice[0:3] //不能自己赋给自己，赋值给新变量不会改变原数组
	fmt.Println(slice1)
	slice = slice[0:3]

	dwarfs := []string{"a", "b", "c"} //切片的声明
	dwarfs = append(dwarfs, "d")      //切片可以append

	dwarfs1 := dwarfs[0:3]
	dwarfs[0] = "x"

	//capacity相同的数组底层是同一个数组,如果添加元素没有导致扩容，那么切片是数组的一个窗口，因为底层是同一个数组
	//如果添加元素很多，导致数组扩容，那么切片变成历史，因为扩容使数组分开存储，新数组的改变与原数组（切片）无关
	fmt.Println(len(dwarfs), cap(dwarfs))
	fmt.Println(dwarfs)
	fmt.Println(dwarfs1)

	dwarfs = append(dwarfs, "e", "f", "g", "h", "i", "j", "k")
	dwarfs[0] = "y"
	fmt.Println(dwarfs)
	fmt.Println(dwarfs1)

	//三索引，第三个索引代表容量限制，不能超过容量限制

	arr1 := []string{
		"a", "b", "c", "d", "e", "f", "g",
	}
	arr2 := arr1[0:4:4]
	arr3 := arr1[0:6]
	arr3[0] = "x"
	//不管原数组改，还是切片出来的新数组改，只要容量相同，底层是同一个数组，就会影响其他数组
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr4 := append(arr2, "h") //arr2限制了容量，则arr4一定会赋值一个新的底层数组
	arr4[4] = "y"             //更改不会影响原数组
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	//使用make预分配，避免append时再出现扩容，多余的内存分配操作
	dwar := make([]string, 0, 10)      //如果没有第三个参数，表示大小和容量都是0（或指定的数）
	dwar = append(dwar, "a", "b", "c") //数组大小不超过10时不会扩容

}
