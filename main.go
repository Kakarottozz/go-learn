package main

import "fmt"
import "math/rand"

var num = 10 //作用域为package，不可用短声明
func main() {
	/*fmt使用*/
	fmt.Println("hello world!")

	/*短声明，分支判断*/
	if x := rand.Intn(3) + 1; x == 2 { //这种不带var的就是短声明
		fmt.Println("x == 2")
	}

	/*指定变量类型*/
	// var fl = 3.14 //默认是float64
	// var pi64 float64 = math.Pi
	// var pi32 float32 = math.Pi

	// var price float64 //初始化称为零值，0.0

	//格式化输出
	var third = 1.0 / 3
	fmt.Printf("%.3f", third)

	var var1 int = 2
	var var2 uint = 4
	var3 := 5

	/*
		int8 	uint8
		int16 	uint16
		int32 	uint32
		int64 	uint64

		int		uint	针对目标设备优化，旧设备32位，新设备64位

	*/

	fmt.Println("%T is type of %[1]v\n")

}

type cel float64 //typedef
func (c cel) fun(x float64) cel { //与cel类关联，cel类的对象可通过.运算符调用
	return cel(x - 1)
}

//一等函数,函数是头等的，可以
/*
1.将函数赋给变量				函数对象
2.将函数作为参数传递给函数		函数指针
3.将函数作为返回值类型
*/

// 1.将函数赋给变量
func faker() cel { return cel(3.14) }

// 2.将函数作为参数传递给函数
func foker(x int, sensor func() cel) {
	fmt.Println("666")
}

// 3.将函数作为返回值类型
type FakerType func() cel //typedef一个函数
func foker1(x int, sensor FakerType) {
	fmt.Println("666")
}
func foker2() FakerType {
	var f FakerType
	return f
}

func main1() {
	f1 := faker
	f1()
	var f2 func() cel //声明一个函数对象，这个函数对象没有入参，返回值为cel
	f2()

	foker(3, faker) //参数
}

func main3() {
	//匿名函数,定义在函数中
	f := func() {
		fmt.Println("444")
	}
	f()

	// 或者定义并直接执行
	func() {
		fmt.Println("444")
	}()
}
