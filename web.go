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
  
  	fmt.Fprintln(res, "Adding data into the pool")
  	const poolSize int = 100
  	const randLimit int = 250
  	var dataPool2 [poolSize]int
  
  	dataPool2[0] = 150
  	dataPool2[1] = 152
  	/*for i:= 0; i < poolSize; i++ {
      	dataPool2[i] = int(rand.Intn(randLimit))
      	msg := string(i) + " : " + string(rand.Intn(randLimit)) + " > " + string(dataPool2[i]) 
      	fmt.Fprintln(res, msg)
  	}*/
  
  	msg := "Testing "+poolSize + " randLimit "+randLimit
  	fmt.Fprintln(res, msg)
  
  	fmt.Fprintln(res, dataPool2[0])
  	fmt.Fprintln(res, dataPool2[1])
  
  	fmt.Fprintln(res, "Datapool size", len(dataPool2))
  	for data := range dataPool2 {
      fmt.Fprintln(res, data)
  	} 
  
 	fmt.Fprintln(res, "Yes it's alive!!!")

}
