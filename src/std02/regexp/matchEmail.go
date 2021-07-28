package main

import (
	"fmt"
	"regexp"
)

func main() {

	s := "zai libaoan@huawei.com libaoan2@huawei.com libaoan@zte.com libaoan@_.com"
	getEmail(s)
}

func getEmail(text string) {
	re, _ := regexp.Compile(`(\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3})`)

	emails := re.FindAllStringSubmatch(text, -1)
	for _, email := range emails {
		fmt.Println(email[1])
	}
	emailss := re.FindAllString(text, -1)
	for _, email := range emailss {
		fmt.Println(email)
	}

}
