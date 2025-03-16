package main

import "fmt"

func changePointer(p *int) {
	v := 5
	p = &v
}

func main() {
	v := 5
	p := &v
	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}
