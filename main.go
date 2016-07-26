package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("%v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println("%v\n", string(b))
}
