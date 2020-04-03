package main

import (
	"email/routers"
	"fmt"
	"gitee.com/liuzhiqiang9696/utils.git/dbutil"
	"gitee.com/liuzhiqiang9696/utils.git/redisUtil"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := redisUtil.InitRedis("148.70.248.33:6379")
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
