package main

import (
	"fmt"
)

func main() {
	// 获取问候消息并打印。
	message := Greetings.Hello("Gladys")
	fmt.Println(message)
}
