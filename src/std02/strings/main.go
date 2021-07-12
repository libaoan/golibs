package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//TestCompare()
	//TestEqualFold()
	//TestContains()
	//TestCount()
	//TestFields()
	//TestSplitX()
	//TestXFix()
	//TestIndexX()
	//TestJoin()
	//TestRepeat()
	//TestMap()
	//TestReplace()
	//TestUpperLower()
	//TestTitle()
	TestTrim()

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

func TestXFix() {
	s := "中国的边境是万里长城吗,zzz"
	fmt.Println(strings.HasPrefix(s, "中"))
	fmt.Println(strings.HasSuffix(s, "zzz"))
	fmt.Println(strings.HasSuffix(s, ""))
}

func TestIndexX() {
	s := "中国的边境是万里长城吗,zzz"
	s1 := "长"
	fmt.Println(s, s1, strings.Index(s, s1))
	fmt.Println(s, "z", strings.IndexByte(s, byte('z')))

	// todo: 中文index问题需要再调试
	s2 := strconv.QuoteToASCII(s1)
	s2 = s2[1 : len(s2)-1]
	s2 = strings.Trim(s2, "\\u")
	fmt.Println(s2)
	u, _ := strconv.ParseInt(s2, 16, 32)
	fmt.Println(s, s1, strings.IndexRune(s, rune(u)))
	fmt.Println(s, s1, strings.IndexFunc(s, func(c rune) bool {
		return c == rune(u)
	}))
}

func TestJoin() {
	s1 := []string{"你好", "Golang", "!"}
	fmt.Println(strings.Join(s1, " "))
}

func TestRepeat() {
	s1 := "Hello"
	fmt.Println(s1 + strings.Repeat(" 好吗 ", 3))
}

func TestMap() {
	s := "你好xs& 我们都8_43在这里等.你。。。"
	s2 := strings.Map(func(c rune) rune {
		if unicode.Is(unicode.Han, c) {
			return c
		} else {
			return -1
		}
	}, s)
	fmt.Println("过滤汉字", s, s2)
	s = "abCdEfgHI"
	s2 = strings.Map(func(c rune) rune {
		switch {
		case c >= 'A' && c <= 'Z':
			return c + 32
		case c >= 'a' && c <= 'z':
			return c
		default:
			return -1
		}
	}, s)
	fmt.Println("大写转小写，过滤非法字符", s, s2)
}

func TestReplace() {
	fmt.Println("oink oink oink", "|", "k", "|", "ky", "|", strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println("oink oink oink", "|", "oink", "|", "moo", "|", strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("oink oink oink", "|", "oink", "|", "moo", "|", strings.Replace("oink oink oink", "oink", "moo", 0))
	fmt.Println("oink oink oink", "|", "oink", "|", "moo", "|", strings.ReplaceAll("oink oink oink", "oink", "moo"))
}

func TestUpperLower() {
	fmt.Println(strings.ToLower("HELLO WORLD"))
	fmt.Println(strings.ToLower("Ā Á Ǎ À"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "壹"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD"))
	fmt.Println(strings.ToLower("Önnek İş"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.ToUpper("ā á ǎ à"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "一"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world"))
	fmt.Println(strings.ToUpper("örnek iş"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}

func TestTitle() {
	fmt.Println(strings.Title("hElLo wOrLd"))
	fmt.Println(strings.ToTitle("hElLo wOrLd"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))
	fmt.Println(strings.Title("āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}

func TestTrim() {
	x := "!!!@@@你好,!@#$ Gophers###$$$"
	fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println(strings.TrimPrefix(x, "!"))
	fmt.Println(strings.TrimSuffix(x, "$"))

	f := func(r rune) bool {
		return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	}
	fmt.Println(strings.TrimFunc(x, f))
	fmt.Println(strings.TrimLeftFunc(x, f))
	fmt.Println(strings.TrimRightFunc(x, f))
}
