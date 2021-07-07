package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//readSlice()
	//readBytes()
	//readString()
	//readLine()
	//peek()
	scanner()
	//readwrite()
}

func readSlice() {
	// base read
	reader := bufio.NewReader(strings.NewReader("this is a strings\n and with return.\n"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line is '%s'\n", line[:len(line)+5])
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the other line is '%s'\n", line)
	fmt.Printf("the other line is '%s'\n", n)

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
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the other line is '%s'\n", line)
	fmt.Println(string(n))

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

func peek() {
	// TODO 需要再研究一下
	Peek := func(reader *bufio.Reader) {
		line, _ := reader.Peek(14)
		fmt.Printf("%s\n", line)
		time.Sleep(1)
		fmt.Printf("%s\n", line)
	}

	reader := bufio.NewReaderSize(strings.NewReader("this is a strings out of buffer\t and with return.\t"), 14)
	go reader.ReadBytes('\t')
	go Peek(reader)
	time.Sleep(1e8)
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text(), "Ctrl-D to exit...") // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	file.WriteString("this is a line.\n and this is another line\t.")
	// 将文件 offset 设置到文件开头
	file.Seek(0, os.SEEK_SET)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readwrite() {
	r := bufio.NewReader(strings.NewReader("this is a line.\n and this is another line\n"))
	w := bufio.NewWriter(os.Stdout)
	rw := bufio.NewReadWriter(r, w)
	line, _, _ := rw.ReadLine()
	n, _ := rw.Write(line)
	rw.Flush()
	fmt.Println(n)
}
