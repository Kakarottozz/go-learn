package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 一般最后一个返回值是错误，没有错误则返回nil
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
