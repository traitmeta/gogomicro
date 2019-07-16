package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	auth "github.com/songxuexian/gogomicro/auth/proto/auth"
	us "github.com/songxuexian/gogomicro/user-srv/proto/user"
	"net/http"
	"time"
)

var (
	serviceClient us.UserService
	authClient    auth.AuthService
)

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("sxx.micro.book.srv.user", client.DefaultClient)
	authClient = auth.NewAuthService("sxx.micro.book.srv.auth", client.DefaultClient)
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
		log.Logf("[Login] 密码校验完成，生成token...")

		// 生成token
		resp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   uint64(resp.User.Id),
			UserName: resp.User.Name,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Logf("[Login] token %s", resp2.Token)
		response["token"] = resp2.Token

		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: resp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

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

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
