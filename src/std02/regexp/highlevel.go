package main

import (
	"fmt"
	"regexp"
)

func main() {
	TestFindAllStringSubmatch()
}

func TestFindAllStringSubmatch() {

	s := "Nobody expects the Spanish inquisition."
	re, _ := regexp.Compile(`(e)(.)`) // Prepare our regex
	result_slice := re.FindAllStringSubmatch(s, -1)
	fmt.Printf("%v", result_slice)

	s = "Mr. Leonard Spock"
	re1, _ := regexp.Compile(`(Mr)(s)?\. (\w+) (\w+)`)
	fmt.Println(re1.FindAllString(s, -1))
	// 分组匹配
	result := re1.FindStringSubmatch(s)

	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}

	fmt.Println("命名捕获")
	re = regexp.MustCompile("(?P<first_char>.)(?P<middle_part>.*)(?P<last_char>.)")
	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch("Super", -1)

	md := map[string]string{}
	for i, n := range r2[0] {
		fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
		md[n1[i]] = n
	}
	fmt.Printf("The names are  : %v\n", n1)
	fmt.Printf("The matches are: %v\n", r2)
	fmt.Printf("The first character is %s\n", md["first_char"])
	fmt.Printf("The last  character is %s\n", md["last_char"])

	fmt.Println("非捕获分组")
	s = "Mrs. Leonora Spock"
	re2, _ := regexp.Compile(`Mr(?:s)?\. (\w+) (\w+)`)
	result = re2.FindStringSubmatch(s)
	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}

	fmt.Println("标志位：不区分大小写")
	s = "Never say never."
	r, _ := regexp.Compile(`(?i)^n`)   // 是否是以 'N' 或 'n' 开头？
	fmt.Printf("%v", r.MatchString(s)) // true, 不区分大小写

	fmt.Println("标志位：关闭默认的贪婪匹配")

	r, _ = regexp.Compile(`'.*'`)
	res := r.FindString(" 'abc','def','ghi' ")
	fmt.Println("<%v>", res)

	r, _ = regexp.Compile(`'.*?'`)
	res = r.FindString(" 'abc','def','ghi' ")
	fmt.Println("<%v>", res)

	r, _ = regexp.Compile(`(?U)'.*'`)
	res = r.FindString(" 'abc','def','ghi' ")
	fmt.Println("<%v>", res)

	fmt.Println("标志位：匹配\n")
	r, _ = regexp.Compile(`a.`)
	s = "atlanta\narkansas\nalabama\narachnophobia"
	ress := r.FindAllString(s, -1)
	fmt.Printf("<%v>", ress)
	fmt.Println("打开匹配\n")
	r, _ = regexp.Compile(`(?s)a.`)
	s = "atlanta\narkansas\nalabama\narachnophobia"
	ress = r.FindAllString(s, -1)
	fmt.Printf("<%v>", ress)

	fmt.Println("标志位：^/$ 匹配换行符")

	r, _ = regexp.Compile(`a$`) // without flag
	s = "atlanta\narkansas\nalabama\narachnophobia"
	res2 := r.FindAllStringIndex(s, -1)
	fmt.Printf("<%v>\n", res2)

	fmt.Println("打开标志位m")
	t, _ := regexp.Compile(`(?m)a$`) // with flag
	u := "atlanta\narkansas\nalabama\narachnophobia"
	res3 := t.FindAllStringIndex(u, -1)
	fmt.Printf("<%v>", res3)

}
