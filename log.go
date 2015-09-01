package main

import (
	
)

func crash_check(e error) {
	if e != nil {
		panic(e)
	}
}