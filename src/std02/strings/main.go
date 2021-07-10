package main

import (
	"fmt"
	"strings"
)

func main() {
	//TestCompare()
	//TestEqualFold()
	//TestContains()
	//TestCount()
	//TestFields()
	TestSplitX()

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

func TestCount() {
	s1 := "Golang 基础库 调试"
	s2 := " 调试"
	fmt.Println("Count: ", s1, s2, strings.Count(s1, s2))

	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(len("谷歌中国"))
	// todo: if substr是空，则返回Unicode个数+1
	fmt.Println(strings.Count("谷歌中国", ""))

}

func TestFields() {
	s := "中共|中央=人民广播* 电视台"
	fs := strings.FieldsFunc(s, func(c rune) bool {
		s := string(c)
		if s == "|" || s == "=" || s == " " || s == "*" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(s, fs, strings.Fields(s))
}

func TestSplitX() {
	s := "中共|中央|人民|广播电视台"
	fmt.Println(s, strings.Split(s, "|"))
	fmt.Println(s, strings.SplitN(s, "|", 2))
	fmt.Println(s, strings.SplitAfter(s, "|"))
	fmt.Println(s, strings.SplitAfterN(s, "|", 2))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}
