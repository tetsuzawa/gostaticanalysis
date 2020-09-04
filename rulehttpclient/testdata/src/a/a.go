package a

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1() {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}

func f2() {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
