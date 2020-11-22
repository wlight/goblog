package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hometHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World!<h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录变编程笔记，如你有反馈或建议，请联系\"+\"<a href=\\\"mailto:wlight@gmail.com\\\">wlight</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到:(</h1>"+"<p>如有问题，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建新的文章")
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1：设置标头
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		// 2:继续处理请求
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hometHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	// 文章详情
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")

	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")

	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 html
	router.Use(forceHTMLMiddleware)

	// 通过命名路由获取URL示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL:", homeURL)

	articleURL, _ := router.Get("articles.show").URL("id", "32")
	fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", router)
}
