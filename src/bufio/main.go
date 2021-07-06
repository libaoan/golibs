package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// readSlice()
	//readBytes()
	//readString()
	readLine()
}

func readSlice() {
	// base read
	reader := bufio.NewReader(strings.NewReader("this is a strings\n and with return.\n"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line is '%s'\n", line)
	line, _ = reader.ReadSlice('\n')
	fmt.Printf("the other line is '%s'\n", line)

	// outof buffer size
	reader = bufio.NewReaderSize(strings.NewReader("this is a strings\n and with return.\n"), 16)
	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)

}

func readBytes() {
	// base read
	reader := bufio.NewReader(strings.NewReader("this is a strings\n and with return.\n"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line is '%s'\n", line)
	line, _ = reader.ReadBytes('\n')
	fmt.Printf("the other line is '%s'\n", line)

	// outof buffer size
	reader = bufio.NewReaderSize(strings.NewReader("this is a strings\n and with return.\n"), 16)
	line, err := reader.ReadBytes('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
	line, err = reader.ReadBytes('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
}

func readString() {
	reader := bufio.NewReader(strings.NewReader("this is a strings\n and with return.\n"))
	line, _ := reader.ReadString('\n')
	fmt.Printf("the line is '%#v'\n", line)
	line, _ = reader.ReadString('\n')
	fmt.Printf("the other line is '%#v'\n", line)
}

func readLine() {
	// baseio
	reader := bufio.NewReader(strings.NewReader("this is a strings\n and with return.\n"))
	line, flag, _ := reader.ReadLine()
	for flag {
		l, f, _ := reader.ReadLine()
		flag = f
		line = append(line, l...)
	}
	fmt.Printf("the line is '%s'\n", line)
	line, flag, _ = reader.ReadLine()
	for flag {
		l, f, _ := reader.ReadLine()
		flag = f
		line = append(line, l...)
	}
	fmt.Printf("the other line is '%s'\n", line)

	// ouf of buffer size
	reader = bufio.NewReaderSize(strings.NewReader("this is a strings out of buffer\n and with return.\n"), 16)
	line, flag, err := reader.ReadLine()
	for flag || err != nil {
		l, f, e := reader.ReadLine()
		flag = f
		err = e
		line = append(line, l...)
	}
	fmt.Printf("the line is '%s'\n", line)
	line, flag, _ = reader.ReadLine()
	for flag {
		l, f, _ := reader.ReadLine()
		flag = f
		line = append(line, l...)
	}
	fmt.Printf("the other line is '%s'\n", line)
}
