package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	ListAll("/root")
}


func ListAll(path string){
	fileInfos, err := ioutil.ReadDir(path)
        if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileInfos{
		if info.IsDir(){
			fmt.Println(info.Name())
                        ListAll(path+"/"+info.Name())
		}else{
			fmt.Println(info.Name())
		}
	}
	
}
