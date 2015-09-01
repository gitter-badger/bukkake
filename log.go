package main

import (
	"fmt"
)

func crash_check(e error) {
	if e != nil {
		panic(e)
	}
}

func def_check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}