package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8888", nil)) //nil 代表使用标准库中的实例处理。第二个参数，则是我们基于net/http标准库实现Web框架的入口。
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Url.Path=%q\n", req.URL.Path) //%q 单引号围绕的字符字面值，由Go语法安全地转义
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
