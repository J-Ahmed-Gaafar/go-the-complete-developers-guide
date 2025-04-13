package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil // Wrong Implementation that compiles correctly
	fmt.Println(string(bs))
	fmt.Println("just wrote", len(bs), "bytes")
	return len(bs), nil // Correct Implementation
}
