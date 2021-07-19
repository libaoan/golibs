package main

import (
	"fmt"
	"regexp"
)

func main() {

	TestFindString()

}

func TestFindString() {
	//r := regexp.MustCompile("H\\wllo")
	r, _ := regexp.Compile(`H\wllo`)
	fmt.Println(r.FindString("Hello World to Hillo"))

}
