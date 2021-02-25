package main

import (
	"fmt"
	"net/http"
	//"io"
)


type HttpHandler struct  {}

func (hh HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//Create a response binary data
	data := []byte("Hello World!\n")  // slice of bytes
        fmt.Println(req.URL)
	res.Write(data)
}

func main() {


       // handle '/' route to 'http.DefaultServeMux'
       http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	       fmt.Fprint(res, "Hello World!\n")
       })

       //handle '/hello/golang' route to 'http.DefaultServeMux'
       http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
	       fmt.Fprint(res, "Hello Golang!\n")
       })
       // listen and serve using 'http.DefaultServeMux'
       http.ListenAndServe(":9000", nil)
}
