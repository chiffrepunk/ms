package main

import (
	"fmt"
	"github.com/dchest/uniuri"
)

func split(pw string, size int) []string {
	ss := make([]string, 0, len(pw)/size+1)
	for len(pw) > 0 {
		if len(pw) < size {
			size = len(pw)
		}
		ss, pw = append(ss, pw[:size]), pw[size:]

	}
	return ss
}

func main() {
	password := uniuri.NewLen(20)
    	fmt.Print(split(password, 4))	
}

