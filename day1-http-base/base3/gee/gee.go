package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// 只要实现了 ServeHTTP 方法，相当于实现了 Handler接口，那么所有的http请求都会转发到实现了 ServeHTTP 方法的实例 此处为engine
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// New is the constructor of gee.Engine
func New() *Engine  {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine * Engine) addRoute(method string, pattern string, handler HandlerFunc){
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc){
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc){
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr, engine)
}

