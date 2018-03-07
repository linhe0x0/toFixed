package main

import (
	"fmt"

	. "github.com/sqrthree/toFixed"
)

func main() {
	fmt.Println(ToFixed(1.2345678, 2))
	fmt.Println(ToFixed(1.2345678, 5))
	fmt.Println(ToFixed(1.2345678, 9))
	fmt.Println(ToFixed(1.2345678, 14))
	fmt.Println(ToFixed(92234, 14))
}
