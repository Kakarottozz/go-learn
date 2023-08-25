package main

import (
	"fmt"
	"math/big"
)

func main() {
	dis := new(big.Int)
	dis.SetString("2400000000000000000000000000000000000", 10)
	fmt.Println(dis)

	//由于go使用utf8变长编码，如果字符出自长短不一的子集，比如中间有一个中文，占两字节，那么len和字符数就不相等
	message := "askj dqiweiq vxc"
	fmt.Println(len(message))

	list := "∂ß∑cc bbaa"
	for i, c := range list { //i是索引，c是list[i]
		fmt.Printf("%v %c\n", i, c)
	}

	for _, c := range list { //_表示不使用
		fmt.Printf("%c\n", c)
	}

	/*强转*/
	age := 41
	fage := float64(age)

	/*go语言中 true/false 不能和 1/0 等价*/

}
