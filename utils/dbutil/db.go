package dbutil

import (
	"database/sql"
	config "email/utils/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	LoginDBPool *gorm.DB
)

//初始化go-sql-driver/mysql 连接池
func InitDbPool(config *config.MysqlConfig) (*sql.DB, error) {

	dbPool, err := sql.Open("mysql", config.MysqlConn)
	if nil != err {
		return nil, err
	}
	dbPool.SetMaxOpenConns(config.MysqlConnectPoolSize)
	dbPool.SetMaxIdleConns(config.MysqlConnectPoolSize / 2)

	err = dbPool.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("init db pool OK")
	return dbPool, nil
}

//初始化gorm 连接池
func InitGormDbPool(config *config.MysqlConfig, setLog bool) (err error) {

	LoginDBPool, err = gorm.Open("mysql", config.MysqlConn)
	if err != nil {
		fmt.Println("init db err : ", config, err)
		return err
	}

	LoginDBPool.DB().SetMaxOpenConns(config.MysqlConnectPoolSize)
	LoginDBPool.DB().SetMaxIdleConns(config.MysqlConnectPoolSize / 2)
	if setLog {
		LoginDBPool.LogMode(true)
		//db.SetLogger(clog.Logger)
	}
	LoginDBPool.SingularTable(true)

	err = LoginDBPool.DB().Ping()
	if err != nil {
		return err
	}
	//	fmt.Println("init db pool OK")

	return nil
}
func InitDb(dbConf config.Configure) error {
	mysqlConf := &config.MysqlConfig{
		MysqlConn:            dbConf.MysqlSetting["db"].MysqlConn,
		MysqlConnectPoolSize: dbConf.MysqlSetting["db"].MysqlConnectPoolSize,
	}
	err := InitGormDbPool(mysqlConf, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
