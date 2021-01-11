package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 0. 初始化变量
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordComfirm: r.PostFormValue("password_comfirm"),
	}

	// 2.表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 4.1 有错误发生，大隐书局
		view.RenderSimple(w, view.D{
			"Errors":errs,
			"User":_user,
		}, "auth.register")
	} else {
		//4.2 验证通过--入库，并跳转到首页
		_user.Create()
		if _user.ID > 0 {
			http.Redirect(w,r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}

	// 3. 表单不通过--重新显示表单
}
