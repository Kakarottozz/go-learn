package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //chan是关键字

	for i := 0; i < 5; i++ {
		go thread(i, c)
	}
	//通过channel通信
	//发送操作会阻塞等待接受
	//接受操作会阻塞等待接受到信号
	for i := 0; i < 5; i++ {
		goid := <-c
		fmt.Println(goid, "	has finished")
	}
}

func thread(id int, c chan int) {
	time.Sleep(time.Second * 2)
	fmt.Println(id, "	work")
	c <- id
}
