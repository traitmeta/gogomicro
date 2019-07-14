package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	us "github.com/songxuexian/gogomicro/user-srv/proto/user"
	"net/http"
	"time"
)

var serviceClient us.UserService

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("sxx.micro.book.srv.user", client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Logf("参数解析错误")
		http.Error(w, "参数解析错误", 400)
		return
	}
	resp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	if resp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = resp.Success

		// 干掉密码返回
		resp.User.Pwd = ""
		response["data"] = resp.User

	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
