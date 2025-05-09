package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

const Url = "https://cataas.com/cat"
const maxLoad = 10

func getCat() []byte {
	resp, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < maxLoad; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			file := getCat()
			err := os.WriteFile(fmt.Sprintf("cat%d.jpg", i), file, 0666)
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
}
