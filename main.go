package main

import (
	"email/routers"
	config "email/utils/conf"
	"email/utils/dbutil"
	"email/utils/redisUtil"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := redisUtil.InitRedis(config.Conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dbutil.InitDb(config.Conf)
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
