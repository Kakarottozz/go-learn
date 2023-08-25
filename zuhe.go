package main

import "fmt"

//go没有继承，但组合可以实现继承的效果

// 方法的转发
type base struct{}
type son struct {
	b base
}

func (b base) report() {
	fmt.Println("base.report")
}

func (s son) report() {
	s.b.report() //包装了其他类的同名方法，转发
}

// 使用 struct嵌入 实现方法的转发
type base1 struct{}
type son1 struct {
	base1 //只声明类型，称为struct嵌入，son1可以直接调用base1的方法和属性
	//如果有base2，base2也有report方法，那么son直接调用report会产生歧义，解决方法是
	//写一个接受son的report方法，内部封装base1或base2的report，
	//因为son本身的方法优先级高于嵌入的方法，就没有歧义了
}

func (b base1) report() {
	fmt.Println("base.report")
}

func main() {
	var s son1
	s.report()       //son1可以直接调用base1的方法
	s.base1.report() //也可以直接用类型名，go会为其生成一个临时变量来调用report方法

}
