package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	ListAll(".")
	path := "./src/std01/ioutil/ioutil.go"
	content, err := ReadAll(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}

	// Discard
	content, err = os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	n, err := ioutil.Discard.Write(content)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The length of content of file", path, n)
	}

	//TempDir and TempFile
	newpath, err := ioutil.TempDir("", "prefix")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Construct new path", newpath)
	}

	f, err := ioutil.TempFile("", "prefix")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	n, err = f.Write(content)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Write content in file %s with length %d", f.Name(), n)
	}
}

func ListAll(path string) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			fmt.Println(path + "/" + info.Name())
			ListAll(path + "/" + info.Name())
		} else {
			fmt.Println(path + "/" + info.Name())
		}
	}
}

func ReadAll(path string) ([]byte, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
