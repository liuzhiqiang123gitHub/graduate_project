package controllers

import (
	"email/model"
	"email/utils/redisUtil"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func LoginByValidationCode(email, validationCode string) error {
	//func (userInfo *UserInfoModel) GetUserByEmail(email string) error {
	//	return dbutil.LoginDBPool.Table(GetName()).Where("email=?", email).Last(&userInfo).Error
	//}
	//查询该用户是否存在
	userInfo := model.UserInfoModel{}
	err := userInfo.GetUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("数据库查询失败")
		return err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("没有这个用户")
		return errors.New("没有这个用户")
	}
	//查询验证码是否过期
	//func Get(key interface{}) (res string, err error) {
	//username, err := redis.String(RedisConn.Do("GET", key))
	//return username, err
	//}
	res, err := redisUtil.Get(validationCode)
	if res == "" {
		fmt.Printf("%s验证码过期", email)
		return errors.New("验证码过期")
	}
	return nil
}
