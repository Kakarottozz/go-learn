package main

import "fmt"

type talker interface {
	talk() string
}

type martain struct{}

func (m martain) talk() string {
	return "martain"

}

type laser struct{}

func (l *laser) talk() string {
	return "laser"
}

func shout(t talker) {
	t.talk()
}

func main() {
	shout(martain{})
	shout(&martain{})
	// shout(laser{}) 报错，相当于是*laser满足了接口而不是laser
	shout(&laser{})
}
