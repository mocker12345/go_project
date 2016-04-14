package main

import (
	"fmt"
)

func main() {
	v := [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() {
			fmt.Println("closure_defer i=", i)
		}()
		v[i] = func() {
			fmt.Println("closure i=", i)
		}
	}
	for _, f := range v {
		f()
	}

}
