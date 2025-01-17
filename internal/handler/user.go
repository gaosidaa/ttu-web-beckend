package handler

import (
	"context"
	"fmt"
	g "github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"ttu-backend/apiv1"
	"ttu-backend/internal/model"
)

var (
	User = hUser{}
)

type hUser struct{}

func (h *hUser) UserLogin(ctx context.Context, req *apiv1.UserLoginReq) (res *apiv1.UserLoginRes, err error) {
	fmt.Println(req.Username)
	fmt.Println(req.Password)
	var user model.User
	user.Id = g.Cfg("user_config").MustGet(ctx, "id").String()
	user.Password = g.Cfg("user_config").MustGet(ctx, "password").String()

	fmt.Println(user.Id)
	fmt.Println(user.Password)
	out := req.Username == user.Id && req.Password == user.Password
	if out == false {
		return &apiv1.UserLoginRes{
			Code:   "0",
			Status: "账号或密码错误",
		}, nil
	} else {
		return &apiv1.UserLoginRes{
			Code:   "1",
			Status: "登录成功",
		}, nil
	}
}

func (h *hUser) UserChangePassword(ctx context.Context, req *apiv1.UserChangePasswordReq) (res *apiv1.UserChangePasswordRes, err error) {
	var user model.User
	user.Id = g.Cfg("user_config").MustGet(ctx, "id").String()
	user.Password = g.Cfg("user_config").MustGet(ctx, "password").String()

	out := req.Username == user.Id && req.OldPassword == user.Password
	if out == false {
		return &apiv1.UserChangePasswordRes{
			Code:   "0",
			Status: "密码输入错误",
		}, nil
	} else {
		user.Password = req.NewPassword
		byte1, _ := yaml.Marshal(user)
		_ = ioutil.WriteFile(
			"config/user_config", byte1, 0644)
		return &apiv1.UserChangePasswordRes{
			Code:   "1",
			Status: "修改成功",
		}, nil
	}
}
