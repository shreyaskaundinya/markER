package main

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)


func TestParser(t *testing.T) {
	fmt.Println("TESTING PARSER")
	requestUrl := "http://localhost:5050/parse"

	f, err := os.Open("./sample_files/sample_simple_diagram.txt")
	
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		os.Exit(1)
	}

	b := make([]byte, 2000)
	f.Read(b)
	s := string(b[:])

	j := map[string]string{
		"code": s,
	}

	jsonData, _ := json.Marshal(j)

	request, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}
	
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

}