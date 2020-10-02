package main

import (
	"fmt"
	"net/http"
)

type Employee struct {
	Name   string
	Number int
}

func main() {
	resp, err := http.Get("http://localhost:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	for true {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}
