package main

import (
	"email/routers"
	"fmt"
	"gitee.com/liuzhiqiang9696/utils.git/dbutil"
	"gitee.com/liuzhiqiang9696/utils.git/redisUtil"
)

func main() {

	err := redisUtil.InitRedis("")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dbutil.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("配置初始化成功")
	routers.StartHttpServer(18080)
	//res,err :=redisUtil.Get("123")
	//if res == ""{
	//	fmt.Println(err)
	//}

}
