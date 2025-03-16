package main

import (
	"fmt"
	"math"
)

func foo() (k float64) {
	i := 36.0
	defer func(){k = math.Sqrt(k)}()
	defer func(){k -= 11}()
	defer func(){k += 10}()
	return i
}

func main(){
	fmt.Println(foo())
}
