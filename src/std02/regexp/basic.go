package main

import (
	"fmt"
	"regexp"
)

func main() {

	//TestFindString()
	TestFindAllString()
	TestMatchString()
	//TestFindAllStringIndex()

}

func TestFindString() {
	//r := regexp.MustCompile("H\\wllo")
	r, _ := regexp.Compile(`H\wllo`)
	fmt.Println(r.FindString("Hello World to Hillo"))
}

func TestFindAllString() {
	s := "Hello World!"
	r := regexp.MustCompile(`l..`)
	fmt.Println(r.FindAllString(s, -1))

	r = regexp.MustCompile(`H[uioe]llo`)
	// Will print 'Hullo'.
	fmt.Println(r.FindAllString("Hello Regular Expression. Hullo again.", -1))
}

func TestMatchString() {
	//r := regexp.MustCompile(`^C:\\\\`)
	r := regexp.MustCompile("^C:\\\\")
	fmt.Println(r.MatchString(`C:\\Windows`))
	fmt.Println(r.MatchString(`~C:\\Windows`))

	r = regexp.MustCompile(`H[^uio]llo`)
	fmt.Printf("%v ", r.MatchString("Hillo")) // false
	fmt.Printf("%v ", r.MatchString("Hallo")) // true
	fmt.Printf("%v ", r.MatchString("H9llo")) // true

	fmt.Println()
	r = regexp.MustCompile(`Jim|Tim`)
	fmt.Printf("%v", r.MatchString("Dickie, Tom and Tim")) // true
	fmt.Printf("%v", r.MatchString("Jimmy, John and Jim")) // true

	t := regexp.MustCompile(`Santa Clara|Santa Barbara`)
	s := "Clara was from Santa Barbara and Barbara was from Santa Clara"
	//                   -------------                      -----------
	fmt.Printf("%v", t.FindAllStringIndex(s, -1))
	// [[15 28] [50 61]]

	u := regexp.MustCompile(`Santa (Clara|Barbara)`) // Equivalent
	v := "Clara was from Santa Barbara and Barbara was from Santa Clara"
	//                   -------------                      -----------
	fmt.Printf("%v", u.FindAllStringIndex(v, -1))
	// [[15 28] [50 61]]

}

func TestFindAllStringIndex() {
	s := "How much wood would a woodchuck chuck in Hollywood?"
	r := regexp.MustCompile(`\bwood`)
	fmt.Println(r.FindAllStringIndex(s, -1))
}
