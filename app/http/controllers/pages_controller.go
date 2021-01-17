package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

// 因为go中没有类这个说法，所以我们就创建一个结构体struct来代替类，struct 上的方法代替类中的方法
// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	data := "此博客是用以记录变编程笔记，如你有反馈或建议，请联系\"+\"<a href=\\\"mailto:wlight@gmail.com\\\">wlight</a>"
	view.Render(w, view.D{"data":data}, "abouts.about")
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到:(</h1><p>如有疑惑，请联系我们。</p>)")
}
