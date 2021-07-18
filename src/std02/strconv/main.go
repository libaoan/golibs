package main

import (
	"fmt"
	"strconv"
)

func main() {
	//TestParseInt()
	//TestFormatInt()
	//TestParseFormatBool()
	TestParseFormatFloat()
}

func TestParseInt() {
	n, err := strconv.ParseInt("-128", 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	u, err := strconv.ParseUint("-128", 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)

	nint, err := strconv.Atoi("-128")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nint)
}

func TestFormatInt() {

	fmt.Println(strconv.FormatInt(-128, 10))
	fmt.Println(strconv.FormatUint(128, 10))
	fmt.Println(strconv.Itoa(-128))

	bs := []byte{}
	bs = strconv.AppendInt(bs, -128, 10)
	bs = strconv.AppendUint(bs, 128, 10)
	fmt.Println(string(bs))

}

func TestParseFormatBool() {
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.Quote(strconv.FormatBool(true)))
	fmt.Println(string(strconv.AppendBool([]byte{}, false)))
}

func TestParseFormatFloat() {
	fmt.Println(strconv.ParseFloat("89.123", 64))
	fmt.Println(strconv.FormatFloat(89.123, 'e', 3, 32))
	fmt.Println(strconv.FormatFloat(89.123, 'f', 3, 32))
	fmt.Println(strconv.FormatFloat(89.123, 'g', 3, 32))

	s := strconv.FormatFloat(1234.5678, 'g', 6, 64)
	fmt.Println(strconv.ParseFloat(s, 64))
}
