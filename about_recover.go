package koans

import (
	"fmt"
	"strings"
)

func aboutRecover() {
	func() {
		defer func() {
			if r := recover(); r != nil {
				f := fmt.Sprint(r)
				assert(strings.Split(f, ":")[0] == "runtime error")
			}
		}()
		a := 0
		fmt.Println(1 / a)
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				f := fmt.Sprint(r)
				assert(strings.Split(f, ":")[0] == "runtime error")
			}
		}()
		c := make(chan int)
		<-c
	}()
}
