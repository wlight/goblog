package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！<h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "此博客是用以记录变编程笔记，如你有反馈或建议，请联系"+"<a href=\"mailto:wlight@gmail.com\">wlight</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到:(</h1>"+"<p>如有问题，请联系我们。</p>")
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
