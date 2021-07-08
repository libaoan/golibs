package main

import (
	"fmt"
	"strings"
)

func main() {
	//TestCompare()
	//TestEqualFold()
	TestContains()
}

func TestCompare() {
	s1 := "hello world"
	s2 := "heLlo world"

	fmt.Println("Compare", s1, s2, strings.Compare(s1, s2))

	switch {
	case s1 > s2:
		fmt.Println("Compare", s1, s2, 1)
	case s1 < s2:
		fmt.Println("Compare", s1, s2, -1)
	default:
		fmt.Println("Compare", s1, s2, 0)
	}
}

func TestEqualFold() {
	s1 := "hello world"
	s2 := "heLlo world"

	fmt.Println("Compare", s1, s2, strings.EqualFold(s1, s2))

	switch {
	case strings.ToLower(s1) == strings.ToLower(s2):
		fmt.Println("Compare", s1, s2, true)
	default:
		fmt.Println("Compare", s1, s2, false)
	}

	s1 = "壹"
	s2 = "一"
	fmt.Println("EqualFold With Unicode", s1, s2, strings.EqualFold(s1, s2))
}

func TestContains() {
	s1 := "Hello World"
	s2 := "Worlds"

	s1 = "你好，Golang"
	s2 = "好吗"

	fmt.Println("Contains", s1, s2, strings.Contains(s1, s2))
	fmt.Println("ContainAny", s1, s2, strings.ContainsAny(s1, s2))

	//fmt.Printf("%s", strconv.QuoteToASCII("好"))
	s3 := rune(0x597d)
	fmt.Println("ContainRune", s1, s2, strings.ContainsRune(s1, s3))
}
