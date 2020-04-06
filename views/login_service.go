package views

import (
	"email/controllers"
	"email/model"
	"email/utils/email"
	"email/utils/httputils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type GetLoginRsp struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func LoginController(c *gin.Context) {
	req := &LoginReq{}
	//rsp := GetLoginRsp{}
	if err := c.Bind(req); err != nil {
		fmt.Printf("%+v", req)
		err := errors.New("invalid params")
		//clog.Logger.Warning("LoginController failed to %v", err.Error())
		fmt.Printf("LoginController failed to %v", err.Error())
		httputils.ResponseError(c, "", err.Error())
		return
	}
	fmt.Printf("LoginController req=%+v ", req)
	rspdata,err := controllers.LoginController(req.Email,req.Password)
	if err !=nil{
		fmt.Println(err)
		fmt.Println(rspdata)
		httputils.ResponseError(c, rspdata, err.Error())
		return
	}
	//fmt.Println(rspdata)
	httputils.ResponseOk(c, rspdata, "")
	return
}

type GetLoginByValidationCodeReq struct {
	Email        string `form:"email" json:"email" binding:"required"`
	ValidateCode string `form:"validate_code" json:"validate_code" binding:"required"`
}
type GetLoginByValidationCodeRsp struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func LoginByValidationCode(c *gin.Context) {
	req := &GetLoginByValidationCodeReq{}
	//rsp := GetLoginByValidationCodeRsp{}
	if err := c.Bind(req); err != nil {
		fmt.Printf("%+v", req)
		err := errors.New("invalid params")
		//clog.Logger.Warning("LoginController failed to %v", err.Error())
		fmt.Printf("LoginControllerByValidation failed to %v", err.Error())
		httputils.ResponseError(c, "", err.Error())
		return
	}
	//验证邮箱
	if !email.EmailValidate(req.Email) {
		err := errors.New("邮箱格式不合法")
		//clog.Logger.Warning("LoginController failed to %v", err.Error())
		fmt.Printf("LoginControllerByValidation failed to %v", err.Error())
		httputils.ResponseError(c, "", err.Error())
		return
	} else if len(req.ValidateCode) != 6 {
		err := errors.New("请输入正确的验证码")
		//clog.Logger.Warning("LoginController failed to %v", err.Error())
		fmt.Printf("LoginControllerByValidation failed to %v", err.Error())
		httputils.ResponseError(c, "", err.Error())
		return
	}
	err,data := controllers.LoginByValidationCode(req.Email, req.ValidateCode)
	if err != nil {
		httputils.ResponseError(c, data, err.Error())
		return
	}
	httputils.ResponseOk(c, data, "")
	return
}

type BackendGetAllUsersReq struct {
	Email string `form:"email" json:"email"`
}
type BackendGetAllUsersRsp struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func BackendGetAllUsers(c *gin.Context) {
	req := &BackendGetAllUsersReq{}
	rsp := BackendGetAllUsersRsp{}
	if err := c.Bind(req); err != nil {
		fmt.Printf("%+v", req)
		err := errors.New("invalid params")
		//clog.Logger.Warning("LoginController failed to %v", err.Error())
		fmt.Printf("BackendGetAllUsers failed to %v", err.Error())
		httputils.ResponseError(c, rsp, err.Error())
		return
	}
	userList := model.UserList{}
	err := userList.GetAllUser()
	if err != nil {
		fmt.Println("查询失败")
		err = errors.New("查询失败")
		httputils.ResponseError(c, rsp, err.Error())
		return
	}

	httputils.ResponseOk(c, userList, "")

}
