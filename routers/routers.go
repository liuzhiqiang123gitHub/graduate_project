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
		//用户找回密码,暂时使用邮箱
		login.GET("/get_password_by_email", views.GetPasswordByEmail)

	}
	err := router.Run(fmt.Sprintf("%s:%d", "0.0.0.0", port))
	fmt.Println(err)
}
