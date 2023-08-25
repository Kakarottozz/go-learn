package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil { //如果defer语句执行了recover，则panic不再退出程序而是继续执行
			fmt.Println(e)
		}
	}()
	var zero int
	x := 43 / zero
	fmt.Println(x)
	// panic("xxx")
}
