package main

import "fmt"
import "time"

// 接口声明
type talker interface {
	talk() string
}

var t talker

// 接口实现
type martain struct{}

// 接口实现
func (m martain) talk() string {
	return "talk talk"
}

// 接口声明
type laser int

// 接口实现
func (l laser) talk() string {
	return "laser laser"
}

// 接口的后实现,mytimer是一个接口，因为声明了interface，就成了所有 有YearDay()和Hour()的类的接口
type mytimer interface {
	YearDay() int
	Hour() int
}

// time.Time可以直接换成上层的接口mytimer
func xfunc(t mytimer) {
	fmt.Println(t.YearDay()) //通过接口调用函数
	fmt.Println(t.Hour())
	fmt.Println(t)
}

func main() {
	t = martain{}
	fmt.Println(t.talk())

	t = laser(3)

	fmt.Println(t.talk())

}

/*标准库实现了许多含有单个方法的接口，比如Stringer，自定义的类实现string函数就是继承Stringer，
fmt通过调用Stringer就可以调用到自定义的string*/
