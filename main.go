package main

import (
	"gee"
	"net/http"
)


func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

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
