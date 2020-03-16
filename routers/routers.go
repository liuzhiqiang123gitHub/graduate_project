package routers

import (
	"email/utils/httputils"
	"email/views"
	"fmt"
	"github.com/gin-gonic/gin"
)

const ServerAdmin = "http://localhost:4420"

func StartHttpServer(port int) {
	router := gin.New()
	router.Use(httputils.ReqData2Form())
	login := router.Group("/v2")
	{
		//用户登录
		login.POST("/login", views.LoginController)
		//验证码登录
		login.POST("/login_with_validation_code", views.LoginByValidationCode)
		//后台获取注册人数
		login.POST("/backend_get_all_user", views.BackendGetAllUsers)
	}
	err := router.Run(fmt.Sprintf("%s:%d", "0.0.0.0", port))
	fmt.Println(err)
}
