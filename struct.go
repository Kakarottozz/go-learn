//结构

package main

import "fmt"
import "encoding/json"
import "os"

type location struct { //类型
	lat  float64
	long float64
}

type partion struct { //变量仍然大写，保证可访问，但编码成json时，用后面的名字
	Lat  float64 `json:"latitude" xml:"latitude"`
	Long float64 `json:"longitude" `
}

func main() {

	var var1 struct { //变量
		Lat  float64 //想要可以被导出，被其他包操作，必须首字母大写
		Long float64
	}
	fmt.Println(var1)

	//初始化方式：字段对应；顺序
	var2 := location{lat: 0.0, long: 0.0}
	var3 := location{0.0, 0.0}

	fmt.Printf("%v\n", var2)
	fmt.Printf("%+v\n", var3)

	// struct赋值给另一个时复制了一份，不再是底层相同的设计了

	//struct json
	var4 := location{lat: 0.3, long: 3.14}
	bytes, err := json.Marshal(var4)
	fmt.Println(err)
	fmt.Println(string(bytes))

}
