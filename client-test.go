package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doRequest() {
	defer wg.Done()

	url := "http://localhost:7771/ping"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	time.Sleep(time.Millisecond * 200)
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))
}

func doWork(parallel, loop int) {
	x := 0

	for x < loop {

		wg.Add(parallel)

		for i := 1; i <= parallel; i++ {
			go doRequest()
			x += 1
			if x >= loop {
				wg.Add((parallel - i) * -1)
				wg.Wait()
				break
			}
		}

		wg.Wait()

	}
}

func main() {
	start := time.Now()
	doWork(1, 1) // os argumentos de quantidade de operações e rotinas paralelas seão informadas aqui
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}
