package main

import (
	"bytes"
	"fmt"
)

func main() {
	//TestContains()
	//TestCount()
	TestRunes()
}

func TestContains() {
	fmt.Println(bytes.Contains([]byte("你 好 Golang"), []byte("Go")))
}

func TestCount() {
	fmt.Println(bytes.Count([]byte("你 好 Golang"), []byte("")), len("你 好 Golang"))
}

func TestRunes() {
	s := "你好Goalng！"
	rs := bytes.Runes([]byte(s))

	for i, v := range s {
		fmt.Println(i, string(v))
	}
	fmt.Println("---")
	for i, v := range rs {
		fmt.Println(i, string(v))
	}
}
