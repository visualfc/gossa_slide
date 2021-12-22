//!play go run main.go
package main

import "fmt"

type T struct {
	n int
}

func (t *T) String() string {
	return fmt.Sprintf("T:%v", t.n)
}
func main() {
	fmt.Println("Hello World", &T{100})
}
