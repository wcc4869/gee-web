package main

import (
	"fmt"
	"gee"

	"net/http"
)
//第三次，封装到gee

func main()  {
	r:=gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}








//  第二次封装实现ServeHTTP
//type Engine struct {
//}
//
//func (engine Engine) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
//	switch r.URL.Path {
//	case "/":
//		fmt.Fprintf(w,"url.path = %q\n",r.URL.Path)
//	case "/hello":
//		for k,v := range r.Header {
//			fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
//		}
//	default:
//		fmt.Fprintf(w,"404 not found : %s\n",r.URL)
//	}
//}
//func main()  {
//	engine := new(Engine)
//	log.Fatal(http.ListenAndServe(":9999",engine))
//}




// 第一次直接调用
//func main() {
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/hello", helloHandleqr)
//	log.Fatal(http.ListenAndServe(":9999",nil))
//}
//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "url.path = %q\n", r.URL.Path)
//}
//
//func helloHandleqr(w http.ResponseWriter, r *http.Request) {
//	for k, v := range r.Header {
//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//	}
//}
