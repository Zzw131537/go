/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-24 18:45:03
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indeHander)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indeHander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PAth = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
