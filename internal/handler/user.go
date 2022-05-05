package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ttu-backend/apiv1"
	"ttu-backend/internal/model"
)

var (
	User = hUser{}
)

type hUser struct{}

func (h *hUser) UserLogin(ctx context.Context, req *apiv1.UserLoginReq) (res *apiv1.UserLoginRes, err error) {
	fmt.Println(req.Password)
	var user model.User
	buf, err1 := ioutil.ReadFile("config/user.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	err2 := json.Unmarshal(buf, &user)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(user.Id)

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
	buf, err1 := ioutil.ReadFile("config/user.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	err2 := json.Unmarshal(buf, &user)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(user.Id)

	out := req.Username == user.Id && req.OldPassword == user.Password
	if out == false {
		return &apiv1.UserChangePasswordRes{
			Code:   "0",
			Status: "密码输入错误",
		}, nil
	} else {
		user.Password = req.NewPassword
		result, _ := json.MarshalIndent(user, "", "    ")
		_ = ioutil.WriteFile("config/user.json", result, 0644)
		return &apiv1.UserChangePasswordRes{
			Code:   "1",
			Status: "修改成功",
		}, nil
	}
}
