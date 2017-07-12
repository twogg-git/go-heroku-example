// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import (
	"fmt"
	"net/http"
	"os"
  	"time"
	"math/rand"
)

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	
  	fmt.Fprintln(res, "Hello, datapool testing")
  
  	now := time.Now()
  	fmt.Fprintln(res, now)

    dataPool := [5]int{1, 2, 3, 4, 5}
  	fmt.Fprintln(res, len(dataPool))
  
  	var dataPool2 [150]int
  	for i:= 0; i < len(dataPool2); i++ {
      	dataPool2[i] = rand.Intn(100) 
  	}
  
  	fmt.Fprintln(res, "Datapool size", len(dataPool2))
  	for data := range dataPool2 {
      fmt.Fprintln(res, data)
  	} 
  
 	fmt.Fprintln(res, "Yes it's alive!!!")

}
