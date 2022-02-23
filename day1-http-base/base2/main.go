package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {}

// 通过查看net/http的源码可以发现，Handler是一个接口，需要实现方法 ServeHTTP ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch  req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q\n",req.URL.Path)
	case "/hello":
		for k, v := range req.Header{
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 not found: %s\n", req.URL)
	}
}

func main()  {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9992", engine))
}

