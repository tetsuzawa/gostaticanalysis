package a

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"cloud.google.com/go/spanner"
)

var a int

const b = 2

func Rule_osopen() {
	// 1. call os.Open
	f, _ := os.Open("xxx")
	// 2. call *File.close
	defer f.Close()
}

func Rule_httpclient() {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	// 1. create *http.Client
	client := &http.Client{}
	// 2. call client.DO
	resp, _ := client.Do(req)
	// 3. call response.Body.Close()
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}

func Rule_error() {
	// 1. returns err
	err := func() error { return errors.New("error") }
	// 2. check error
	if err != nil {
		//hoge
	}
}

func Rule_spanner_tx() {
	client := &spanner.Client{}
	// 1. open transaction
	tx := client.ReadOnlyTransaction()
	// 2. close transaction
	tx.Close()
}
