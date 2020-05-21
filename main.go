package main

import (
	"email/routers"
	config "email/utils/conf"
	"email/utils/dbutil"
	etcdIni "email/utils/etcd"
	"email/utils/redisUtil"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	etcdIni.InitEtcd()
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
	routers.StartHttpServer(18080)

	//fmt.Println("running,not deadlock")
	//waitQueue := make(chan int)
	//waitQueue <- 1
	//go func() {
	//	<- waitQueue
	//}()
	//close(waitQueue)
	//server,err := net.Listen("tcp","localhost:8080")
	//if err != nil{
	//	panic(err)
	//}
	//for {
	//	conn,err := server.Accept()
	//	if err!=nil{
	//		panic("server")
	//	}
	//	fmt.Printf("rec %s",conn.RemoteAddr())
	//
	//}


}


