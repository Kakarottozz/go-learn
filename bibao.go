package main

import "fmt"

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0

}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset //返回的匿名函数的实现中调用了形参s
	}
}

// 闭包就是由于匿名函数封闭并包围作用域中的变量而得名。	封闭了s
func main() {
	sensor := calibrate(realSensor, 5) //sensor作为返回值，已经出了calibrate的作用域
	fmt.Println(sensor())              //但sensor依然能够调用形参s

	var k float64 = 293.0
	f1 := func() float64 {
		return k //与var k是同一个变量
	}
	fmt.Println(f1()) //293
	k++
	fmt.Println(f1()) //294
}
