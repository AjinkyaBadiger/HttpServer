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
       // create a new handler
       handler := HttpHandler{}

       // listen and serve
       http.ListenAndServe(":9000", handler)
}
