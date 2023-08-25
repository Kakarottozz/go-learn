package main

import "fmt"

type base struct {
	a int
	b int
}

type son struct {
	b base
	c int
}

func xfunc(b *base) {
	b.a++
	b.b++
}

func main() {

	px := son{c: 1}

	xfunc(&px.b) //即可获得内部结构体的地址，称为内部指针

}

/*隐式的指针：某些内置集合类型
map在被赋值或作为参数传入时，不会被复制，map就是一种隐式的指针
func xfunc(m *map[string]string) *map多余，直接写map就行，map本身就是隐式的指针，
map的key和value也都可以是指针


slice内部只有三个成员变量：数组指针，指向原数组，slice的容量，slice的长度
指向slice的指针本质上是修改以上三个属性 
*/
