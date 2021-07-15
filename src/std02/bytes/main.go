package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//TestRunes()
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
	//TestTrim()
	//TestReader()
	TestBufferString()

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

func TestCompare() {
	s1 := "hello world"
	s2 := "heLlo world"

	fmt.Println("Compare", s1, s2, bytes.Compare([]byte(s1), []byte(s2)))

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

	fmt.Println("Compare", s1, s2, bytes.EqualFold([]byte(s1), []byte(s2)))

	switch {
	case strings.ToLower(s1) == strings.ToLower(s2):
		fmt.Println("Compare", s1, s2, true)
	default:
		fmt.Println("Compare", s1, s2, false)
	}

	s1 = "壹"
	s2 = "一"
	fmt.Println("EqualFold With Unicode", s1, s2, bytes.EqualFold([]byte(s1), []byte(s2)))
}

func TestContains() {
	s1 := "Hello World"
	s2 := "Worlds"

	s1 = "你好，Golang"
	s2 = "好吗"

	fmt.Println("Contains", s1, s2, bytes.Contains([]byte(s1), []byte(s2)))
	fmt.Println("ContainAny", s1, s2, bytes.ContainsAny([]byte(s1), s2))

	//fmt.Printf("%s", strconv.QuoteToASCII("好"))
	s3 := rune(0x597d)
	fmt.Println("ContainRune", s1, s2, bytes.ContainsRune([]byte(s1), s3))
}

func TestCount() {
	s1 := "Golang 基础库 调试"
	s2 := " 调试"
	fmt.Println("Count: ", s1, s2, bytes.Count([]byte(s1), []byte(s2)))

	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(len("谷歌中国"))
	// todo: if substr是空，则返回Unicode个数+1
	fmt.Println(bytes.Count([]byte("谷歌中国"), []byte("")))

}

func TestFields() {
	s := "中共|中央=人民广播* 电视台"
	fs := bytes.FieldsFunc([]byte(s), func(c rune) bool {
		s := string(c)
		if s == "|" || s == "=" || s == " " || s == "*" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(s)
	for _, c := range fs {
		fmt.Printf("%s >>", c)
	}
	fmt.Println()
	for _, c := range bytes.Fields([]byte(s)) {
		fmt.Printf("%s >>", c)
	}
}

func TestSplitX() {
	s := "中共|中央|人民|广播电视台"
	fmt.Println(s, bytes.Split([]byte(s), []byte("|")))
	fmt.Println(s, bytes.SplitN([]byte(s), []byte("|"), 2))
	fmt.Println(s, bytes.SplitAfter([]byte(s), []byte("|")))
	fmt.Println(s, bytes.SplitAfterN([]byte(s), []byte("|"), 2))

	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}

func TestXFix() {
	s := "中国的边境是万里长城吗,zzz"
	fmt.Println(bytes.HasPrefix([]byte(s), []byte("中")))
	fmt.Println(bytes.HasSuffix([]byte(s), []byte("zzz")))
	fmt.Println(bytes.HasSuffix([]byte(s), []byte("")))
}

func TestIndexX() {
	s := "中国的边境是万里长城吗,zzz"
	s1 := "长"
	fmt.Println(s, s1, bytes.Index([]byte(s), []byte(s1)))
	fmt.Println(s, "z", bytes.IndexByte([]byte(s), byte('z')))

	index := bytes.IndexByte([]byte(s), []byte(s1)[0])
	fmt.Println(s, index)
	index = bytes.IndexRune([]byte(s), bytes.Runes([]byte(s1))[0])
	fmt.Println(s, index)

}

func TestJoin() {
	s1 := [][]byte{[]byte("你好"), []byte("Golang"), []byte("!")}
	fmt.Printf("%s\n", bytes.Join(s1, []byte(" ")))
}

func TestRepeat() {
	s1 := "Hello"
	fmt.Printf("%s%s\n", s1, bytes.Repeat([]byte(" 好吗 "), 3))
}

func TestMap() {
	s := []byte("你好xs& 我们都8_43在这里等.你。。。")
	s2 := bytes.Map(func(c rune) rune {
		if unicode.Is(unicode.Han, c) {
			return c
		} else {
			return -1
		}
	}, s)
	fmt.Println("过滤汉字", string(s), string(s2))
	s = []byte("abCdEfgHI")
	s2 = bytes.Map(func(c rune) rune {
		switch {
		case c >= 'A' && c <= 'Z':
			return c + 32
		case c >= 'a' && c <= 'z':
			return c
		default:
			return -1
		}
	}, s)
	fmt.Println("大写转小写，过滤非法字符", string(s), string(s2))
}

func TestReplace() {
	fmt.Println("oink oink oink", "|", "k", "|", "ky", "|", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Println("oink oink oink", "|", "oink", "|", "moo", "|", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
	fmt.Println("oink oink oink", "|", "oink", "|", "moo", "|", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), 0))
	fmt.Println("oink oink oink", "|", "oink", "|", "moo", "|", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))
}

func TestUpperLower() {
	fmt.Println(string(bytes.ToLower([]byte("HELLO WORLD"))))
	fmt.Println(string(bytes.ToLower([]byte("Ā Á Ǎ À"))))
	fmt.Println(string((bytes.ToLowerSpecial(unicode.TurkishCase, []byte("壹")))))
	fmt.Println(string(bytes.ToLowerSpecial(unicode.TurkishCase, []byte("HELLO WORLD"))))
	fmt.Println(string(bytes.ToLower([]byte("Önnek İş"))))
	fmt.Println(string(bytes.ToLowerSpecial(unicode.TurkishCase, []byte("Önnek İş"))))

	fmt.Println(string(bytes.ToUpper([]byte("hello world"))))
	fmt.Println(string(bytes.ToUpper([]byte("ā á ǎ à"))))
	fmt.Println(string(bytes.ToUpperSpecial(unicode.TurkishCase, []byte("一"))))
	fmt.Println(string(bytes.ToUpperSpecial(unicode.TurkishCase, []byte("hello world"))))
	fmt.Println(string(bytes.ToUpper([]byte("örnek iş"))))
	fmt.Println(string(bytes.ToUpperSpecial(unicode.TurkishCase, []byte("örnek iş"))))
}

func TestTitle() {
	fmt.Printf("%s\n", bytes.Title([]byte("hElLo wOrLd")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("hElLo wOrLd")))
	fmt.Printf("%s\n", bytes.ToTitleSpecial(unicode.TurkishCase, []byte("hElLo wOrLd")))
	fmt.Printf("%s\n", bytes.Title([]byte("āáǎà ōóǒò êēéěè")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("āáǎà ōóǒò êēéěè")))
}

func TestTrim() {
	x := "!!!@@@你好,!@#$ Gophers###$$$"
	fmt.Printf("%s\n", bytes.Trim([]byte(x), "@#$!%^&*()_+=-"))
	fmt.Printf("%s\n", bytes.TrimLeft([]byte(x), "@#$!%^&*()_+=-"))
	fmt.Printf("%s\n", bytes.TrimRight([]byte(x), "@#$!%^&*()_+=-"))
	fmt.Printf("%s\n", bytes.TrimSpace([]byte(" \t\n Hello, Gophers \n\t\r\n")))
	fmt.Printf("%s\n", bytes.TrimPrefix([]byte(x), []byte("!")))
	fmt.Printf("%s\n", bytes.TrimSuffix([]byte(x), []byte("$")))

	f := func(r rune) bool {
		return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	}
	fmt.Printf("%s\n", bytes.TrimFunc([]byte(x), f))
	fmt.Printf("%s\n", bytes.TrimLeftFunc([]byte(x), f))
	fmt.Printf("%s\n", bytes.TrimRightFunc([]byte(x), f))
}

func TestReader() {
	r := bytes.NewReader([]byte("我们都用Golang语言！"))
	n, err := r.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n %d bytes write to %v", n, os.Stdout)
}

func TestBufferString() {
	bs := bytes.NewBufferString("我们都喜欢Golang!")
	bs.WriteTo(os.Stdout)

	fmt.Println()

	bs = bytes.NewBuffer([]byte("我们都喜欢Golang!"))
	bs.WriteTo(os.Stdout)
}
