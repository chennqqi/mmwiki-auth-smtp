package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"strings"
)

/*
{
     // 认证接口信息，必须返回，成功为空，失败非空
    "message": "",
    // 认证成功用户数据，必须返回, 成功时非空
    "data": {
        "mobile": "111111111111", // 用户手机号，string，可以为空
        "phone": "010-9929921", // 用户电话，string，可以为空

        "email": "root@mmWiki.com", // 用户邮箱，string，不能为空
        "given_name": "王哈哈", // 用户姓名，string，不能为空

        "department": "广告事业部.技术部.系统开发组", // 用户所属部门，string，可以为空
        "position": "高级JAVA开发工程师", // 职位，string，可以为空
        "location": "B座32层E区1002", // 工位位置，string，可以为空
        "im": "QQ：12211", // 即时聊天工具，string，可以为空
    }
}
*/
type AuthResp struct {
	Data    *User  `json:"data"`
	Message string `json:"message"`
}

func main() {
	var (
		server   string
		listen   string
		path     string
	//	userDict string
		port     int
	)
	flag.StringVar(&path, "path", "/smtplogin", "set path")
	flag.StringVar(&server, "server", "", "set server")
	flag.StringVar(&listen, "listen", ":8080", "set http listen")
	// flag.StringVar(&userDict, "user", "user.yaml", "set users config file")
	flag.IntVar(&port, "port", 25, "set smtp port, 465 for ssl smtp")
	flag.Parse()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("username")
		pass := r.FormValue("password")

		d := NewDialer(server, port, user, pass)
		err := d.DialAndAuth()

		var resp AuthResp
		if err == nil {
			resp.Message = ""
			resp.Data = &User{
				Email:     user,
				GivenName: strings.Split(user, "@")[0],
			}
		} else {
			resp.Message = err.Error()
		}

		txt, _ := json.Marshal(resp)
		w.Write(txt)
	})
	http.ListenAndServe(listen, nil)
}
