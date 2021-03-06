package controllers

import (
	"email/model"
	"email/utils/redisUtil"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func LoginController(email,password string)(model.UserInfoModel,error ){
	userInfo := model.UserInfoModel{}
	err := userInfo.GetUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("数据库查询失败")
		return userInfo,err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("没有这个用户")
		return userInfo,errors.New("没有这个用户")
	}
	userInfo = model.UserInfoModel{}
	err = userInfo.Login(email,password)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("...数据库查询失败")
		return userInfo,err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("请检查邮箱和密码")
		return userInfo,errors.New("请检查邮箱和密码")
	}
	return userInfo,nil
}
func LoginByValidationCode(email, validationCode string)( error,model.UserInfoModel) {
	//func (userInfo *UserInfoModel) GetUserByEmail(email string) error {
	//	return dbutil.LoginDBPool.Table(GetName()).Where("email=?", email).Last(&userInfo).Error
	//}
	//查询该用户是否存在
	userInfo := model.UserInfoModel{}
	err := userInfo.GetUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("数据库查询失败")
		return err,userInfo
	} else if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("没有这个用户")
		return errors.New("没有这个用户"),userInfo
	}
	//查询验证码是否过期
	//func Get(key interface{}) (res string, err error) {
	//username, err := redis.String(RedisConn.Do("GET", key))
	//return username, err
	//}
	res, err := redisUtil.Get(email)
	valiStr := fmt.Sprintf("%s_%s",email,validationCode)
	if res == "" {
		fmt.Printf("%s验证码过期", email)
		return errors.New("验证码或已失效"),userInfo
	}else if res == valiStr{
		return nil,userInfo
	}
	return errors.New("系统错误"),userInfo
}
