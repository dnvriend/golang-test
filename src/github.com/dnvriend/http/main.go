package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://www.google.nl")
	handleError(err)
	//fmt.Println(resp)
	bs, err2 := ioutil.ReadAll(resp.Body)
	handleError(err2)
	//bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

type textFileReader struct {}

func (textFileReader) Read(bs []byte) (int, error) {
	return "Information from a text file"
}