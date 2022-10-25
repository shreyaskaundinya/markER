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

func TestPing(t *testing.T) {
	fmt.Println("TESTING PING")
	requestUrl := "http://localhost:8080/ping"
	res, err := http.Get(requestUrl)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
}



func TestLexer(t *testing.T) {
	fmt.Println("TESTING LEXER")
	requestUrl := "http://localhost:5050/lex"

	f, err := os.Open("./sample_files/sample_diagram.txt")
	
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