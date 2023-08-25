package main

import "fmt"
import "os"
import "errors"

func openx(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close() //在函数结束之前执行

	_, err = fmt.Fprintln(f, "file write") //文件写入
	if err != nil {
		return err
	}

	return err
}

type openerror struct{}

func (oe openerror) Error() string { //实现Error接口，可通过断言判断错误类型是否是自定义类型
	return "openError"
}

func main() {

	if err := openx("file.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//
	var x error
	x = errors.New("error")
	fmt.Println(x)

	//类型断言
	var oe openerror
	var e error
	e = oe
	if err, ok := e.(openerror); ok {
		fmt.Println(err)
		fmt.Println("type of error is openError")
	}

	panic("xxx") //主动引发panic
	//先使用错误值error相关的，其次使用panic，最后使用exit 来处理错误
	//panic会将defer的语句执行，exit不会

}
