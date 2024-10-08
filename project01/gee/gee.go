package gee
import (
	"fmt"
	"net/http"
)

const ListenAddrPort = ":8089"


type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}
func (engine *Engine) Run(addr string) (err error) {
	fmt.Println(":op",engine.router)
	return http.ListenAndServe(addr, engine)
}

// func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	   switch req.URL.Path {
// 	   case "/":
// 		   fmt.Fprintf(w, "URL = %q\n", req.URL.Path)
// 	   case "/hello":
// 		   for k, v := range req.Header {
// 			   fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 		   }
// 	   default:
// 		   fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
// 	   }
// }

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}