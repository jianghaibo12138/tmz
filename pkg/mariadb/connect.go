package mariadb

import (
	"database/sql"
	"fmt"
	"github.com/jianghaibo12138/TMZ/configs"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var (
	DbConnect *gorm.DB
	DbErr     error
	SqlDb     *sql.DB
)

func init() {
	NewConnect(false)
}

func GetConnect() *gorm.DB {
	// 协程安全, https://gorm.io/zh_CN/docs/method_chaining.html#goroutine_safe
	// 开启预编译, https://gorm.io/zh_CN/docs/performance.html
	session := DbConnect.Session(&gorm.Session{PrepareStmt: false})
	if configs.Settings.Mysql.Debug {
		return session.Debug()
	} else {
		return session
	}
}

func NewConnect(skipDefaultTransaction bool) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		configs.Settings.Mysql.User, configs.Settings.Mysql.Password, configs.Settings.Mysql.Host,
		configs.Settings.Mysql.Port, configs.Settings.Mysql.Database)
	fmt.Println("New connect to mysql master node", dsn)

	slowLogger := gormLogger.New(
		// 设置Logger
		NewSQLLogger(),
		gormLogger.Config{
			// 慢SQL阈值
			SlowThreshold: time.Millisecond,
			// 设置日志级别，只有Warn以上才会打印sql
			LogLevel: gormLogger.Info,
		},
	)
	conf := &gorm.Config{
		SkipDefaultTransaction: skipDefaultTransaction,
		// Logger:                 slowLogger,
	}
	if !configs.Settings.Mysql.Debug {
		conf.Logger = slowLogger
	}
	DbConnect, DbErr = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), conf)
	if DbErr != nil {
		panic(DbErr)
	}
	sqlDb, err := DbConnect.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDb.SetMaxIdleConns(100)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(2000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(60 * time.Second)

	SqlDb = sqlDb

	// CallBackRegister(DbConnect)

	return DbConnect
}

func CloseConnect() error {
	return SqlDb.Close()
}
