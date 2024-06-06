package subpackagename

import "fmt"

// 多层级引用

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome subpackagename!", name)
	return message
}
