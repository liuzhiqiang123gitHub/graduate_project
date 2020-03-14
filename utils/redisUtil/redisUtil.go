package redisUtil

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var(
	RedisConn redis.Conn
)
func InitRedis(addr string)(err error)  {
	RedisConn, err = redis.Dial("tcp", "148.70.248.33:6379",redis.DialPassword("liuzhi19972123"))
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	return err
}
//插入数据
func Set(key,value interface{},exp uint)  {
	_, err := RedisConn.Do("SET", "mykey", "superWang", "EX", exp)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}
//Get
func Get(key interface{})(res string,err error)  {
	username, err := redis.String(RedisConn.Do("GET", key))
	return username,err
}
//Del
func Delete(key string)error  {
	_, err := RedisConn.Do("DEL", key)
	return err
}
