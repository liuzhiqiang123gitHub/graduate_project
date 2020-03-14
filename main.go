package main

import "email/routers"

func main() {
	//err := redisUtil.InitRedis("")
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//res,err := redisUtil.Get("level")
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	routers.StartHttpServer(18080)


}