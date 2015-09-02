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

func rgb(color int) float32 {
	return float32(color) / 255
}
