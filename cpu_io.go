package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

// IndexHandler returns a simple message
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "stay here...")
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}
	http.HandleFunc("/", IndexHandler)
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

	//time.Sleep(10 * time.Second)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
