package Greetings

import "fmt"

// Hello 返回对指定人员的问候。
func Hello(name string) string {
	// 返回将姓名嵌入消息中的问候。
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
