package main

import "fmt"

type coordinate struct {
	Lat  float64
	Long float64
}

func (c coordinate) xfunc() float64 { //前面的括号表示与coordinate关联，相当于this指针

}
func NewCoordinate() coordinate { //约定俗成，作为coordinate的构造函数
	return coordinate{0.4, 0.99}
}
func NewCoordinate(lat float64, long float64) coordinate { // 作为coordinate的有参构造函数
	return coordinate{lat, long}
}

func main() {
	error.New() //包名.函数名 不会因为New名字容易重复而造成函数调用的重复

}
