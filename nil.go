package main

import "fmt"
import "sort"

type person struct {
	a int
}

func (p *person) xfunc() {
	//即使指针是nil，但仍然是person类型，还是会进到函数中，所以panic的控制最好放在函数中
	if p == nil {
		return
	}
	p.a++
}

// 对函数变量为nil的情况作处理
func xfunc(s string, less func(i int, j int) bool) {
	if less == nil {
		less = func(i int, j int) bool { return i < j }
	}
	sort.Slice(s, less)
}

func main() {
	var nobody *person

	fmt.Println(nobody)
	nobody.xfunc() //先进入函数才报错

	//函数变量的默认值也是nil

	var funcvar func(a int, b string) string
	fmt.Println(funcvar == nil)

	//对nil的map访问不会报错，会返回false
	var soup map[string]string

	if value, ok := soup["fuck"]; ok {
		fmt.Println(value) //ok为false
	}

	//接口类型的变量只有在类型和值都为nil时才是nil，即是一个空接口且没有赋值
	var p interface{}
	fmt.Printf("%T %v %v\n", p, p, p == nil) //<nil> <nil> true

	var x *int
	p = x
	fmt.Printf("%T %v %v\n", p, p, p == nil) //*int <nil> false

	fmt.Printf("%#v\n", p) //%#v打印接口类型变量的内部表示 (*int)(nil)

}
