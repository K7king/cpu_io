package main

import (
	"io/ioutil"
	"log"
	"runtime"
	"time"
)

func main() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	for i := 0; i < n; i++ {
		go func() {
			for {
				err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(1)
			}
		}()
	}
}
