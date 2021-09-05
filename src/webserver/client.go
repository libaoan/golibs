package webserver

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Client(input []string) ([]bool,error) {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	transData, err := json.Marshal(input)
	if err != nil{
		return nil,err
	}

	client := &http.Client{Transport: transport}
	res,err := client.Post("https://127.0.0.1:8080/strings","text/json",bytes.NewReader(transData))
	if err != nil{
		return nil,err
	}
	defer res.Body.Close()

	if res.StatusCode != 200{
		return nil,errors.New("500")
	}

	body,err := ioutil.ReadAll(res.Body)
	if err != nil{
		return nil,err
	}

	var result []bool
	err = json.Unmarshal(body,&result)
	if err != nil{
		return nil,err
	}
	return result,err
}
