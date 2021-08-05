package main

import (
	"fmt"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

func main() {
	//testUnicode()
	//TestUft8()
	TestUtf16()
}

func testUnicode() {
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))
	single = '\ufe35'
	fmt.Println(unicode.IsControl(single))

	digit := '1'
	fmt.Println(unicode.IsDigit(digit))
	fmt.Println(unicode.IsNumber(digit))

	letter := 'Ⅷ'
	fmt.Println(unicode.IsDigit(letter))
	fmt.Println(unicode.IsNumber(letter))

	han := '你'
	fmt.Println(unicode.IsDigit(han))
	fmt.Println(unicode.Is(unicode.Han, han))
	fmt.Println(unicode.In(han, unicode.Gujarati, unicode.White_Space))
}

func TestUft8() {
	word := []byte("界")

	fmt.Println(utf8.Valid(word[:2]))
	fmt.Println(utf8.ValidRune('界'))
	fmt.Println(utf8.ValidString("世界"))

	fmt.Println(utf8.RuneLen('界'))

	fmt.Println(utf8.RuneCount(word))
	fmt.Println(utf8.RuneCountInString("世界"))

	p := make([]byte, 3)
	utf8.EncodeRune(p, '好')
	fmt.Println(p)
	fmt.Println(utf8.DecodeRune(p))
	fmt.Println(utf8.DecodeRuneInString("你好"))
	fmt.Println(utf8.DecodeLastRune([]byte("你好")))
	fmt.Println(utf8.DecodeLastRuneInString("你好"))

	fmt.Println(utf8.FullRune(word[:2]))
	fmt.Println(utf8.FullRuneInString("你好ss"))

	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))
}

func TestUtf16() {
	words := []rune{'𝓐', '𝓑'}

	u16 := utf16.Encode(words)
	fmt.Println(u16)
	fmt.Println(utf16.Decode(u16))

	r1, r2 := utf16.EncodeRune('𝓐')
	fmt.Println(r1, r2)
	fmt.Println(utf16.DecodeRune(r1, r2))
	fmt.Println(utf16.IsSurrogate(r1))
	fmt.Println(utf16.IsSurrogate(r2))
	fmt.Println(utf16.IsSurrogate(1234))
}
