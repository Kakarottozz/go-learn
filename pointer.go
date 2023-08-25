package main

import "fmt"

type structxx struct {
	a int
	b int
}

func (p *structxx) xfunc() { //如果某些类的方法用了指针，最好所有方法都用指针，不用就都不用
	p.a++ //函数中想要改变成员变量必须绑定成指针
}

func main() {
	/*go语言不允许对指针操作，只能对取地址和解引用 */
	answer := 3
	var address *int //声明一个指针
	address = &answer
	fmt.Printf("%T\n", address) /* *int 类型 */
	/*指针的相等是值相等，也就是存的地址相等 */

	//复合字面值，也就是结构体的字面值，前面可以放&

	timmy := &structxx{1, 2} //对创建的临时变量取地址
	(*timmy).a = 3           //加不加*效果一样
	timmy.a = 4

	//指向数组的指针
	superman := &[3]string{"1", "2", "3"}
	fmt.Println(superman)
	//操作时也不用显式的解引用
	fmt.Println(superman[1:3])

	//虽然函数是与指针绑定，但使用对象调用也可以，会自动转为指针调用
	px := structxx{1, 2}
	px.xfunc()
}
