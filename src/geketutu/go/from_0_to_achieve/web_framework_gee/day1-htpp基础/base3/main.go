/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-24 19:00:51
 */
package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Head[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
