package greeting

import "fmt"

// HelloWorld should have a comment documenting it.
func HelloWorld() string {
	return "Hello, World!"
}

func greeting() {
	fmt.Println(HelloWorld())
}
