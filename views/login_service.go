package views

import (
	"email/model"
	httputils2 "email/utils/httputils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `form:"email" json:"email" bind:"required"`
	Password string `form:"password" json:"password" bind:"required"`
}
type GetLoginRsp struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func LoginController(c *gin.Context) {
	req := &LoginReq{}
	rsp := GetLoginRsp{}
	if err :=c.Bind(req);err != nil{
		fmt.Printf("%+v",req)
		err := errors.New("invalid params")
		//clog.Logger.Warning("LoginController failed to %v", err.Error())
		fmt.Printf("LoginController failed to %v", err.Error())
		httputils2.ResponseError(c, rsp, err.Error())
		return
	}
	userInfo :=model.UserInfoModel{}
	err := userInfo.Login(req.Email,req.Password)
	if err != nil {
		fmt.Println("登陆失败")
		httputils2.ResponseError(c, rsp, err.Error())
		return
	}
	httputils2.ResponseOk(c, "rsp", "")
	fmt.Printf("LoginController req=%+v ", req)
	return
}

type GetPasswordReq struct {
	Email    string `form:"graduate_login" json:"graduate_login" bind:"required"`
	Password string `form:"password" json:"password" bind:"required"`
}
type GetPasswordByEmailRsp struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func GetPasswordByEmail(c *gin.Context) {
	req := &LoginReq{}
	rsp := GetPasswordByEmailRsp{}
	if err:=c.Bind(req) ;err!=nil {
		err := errors.New("invalid params")
		fmt.Printf("GetPasswordByEmail failed to %v", err.Error())
		httputils2.ResponseError(c, rsp, err.Error())
		return
	}
	fmt.Printf("LoginController req=%+v ", req)

}
