package db

import (
	"fmt"
	"github.com/xuanxiaox/ginx/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"time"
)

func NewDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=Local",
		global.DBConfig.GetString("mysql.user_name"),
		global.DBConfig.GetString("mysql.password"),
		global.DBConfig.GetString("mysql.host"),
		global.DBConfig.GetString("mysql.db_name"),
		"UTF8",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: newLogger(),
		NowFunc: func() time.Time {
			return time.Now().Local() //更改创建时间使用的函数
		},
		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建数据库外键约束
	})
	if err != nil {
		global.Log.Fatalf("NewDb gorm.Open err:%v", err)
	}

	//err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
	//if err != nil {
	//	global.Log.Fatalf("AutoMigrate err:%v", err)
	//}

	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		global.Log.Fatalf("NewDb db.DB() err:%v", err)
	}

	//设置连接池最大空闲连接数
	sqlDB.SetMaxIdleConns(global.DBConfig.GetInt("mysql.max_idle_conns"))
	//设置打开的最大连接数
	sqlDB.SetMaxOpenConns(global.DBConfig.GetInt("mysql.max_open_conns"))
	//设置连接的最大重用时间
	sqlDB.SetConnMaxLifetime(time.Duration(global.DBConfig.GetInt64("mysql.conn_max_lifetime")) * time.Second)

	return db
}

func newLogger() logger.Interface {
	filePath := global.DBConfig.GetString("mysql.log_dir") + global.DBConfig.GetString("mysql.log_filename")
	fileExt := global.DBConfig.GetString("mysql.log_ext")
	return logger.New(
		log.New(getWriterByLumberjack(filePath, fileExt), "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold: time.Duration(global.DBConfig.GetInt64("mysql.log_slow_threshold")), //慢SQL阀值
			Colorful:      false,                                                               //禁止彩色打印
			LogLevel:      getLogLevel(global.DBConfig.GetString("mysql.log_level")),           //log level
		},
	)
}

//根据lumberjack实现日志分割
func getWriterByLumberjack(filePath string, fileExt string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filePath + fileExt,
		MaxSize:    global.DBConfig.GetInt("mysql.log_max_size"),    //M
		MaxAge:     global.DBConfig.GetInt("mysql.log_max_age"),     //days 保留旧日志文件的最大天数
		MaxBackups: global.DBConfig.GetInt("mysql.log_max_backups"), //要保留的最大旧日志文件数
		LocalTime:  global.DBConfig.GetBool("mysql.log_local_time"), //是否格式化时间戳的时间
		Compress:   global.DBConfig.GetBool("mysql.log_compress"),   //是否使用gzip压缩， 默认不压缩
	}
}

//获取对应的sql log level
func getLogLevel(level string) (l logger.LogLevel) {
	switch level {
	case "silent":
		l = logger.Silent
	case "error":
		l = logger.Error
	case "warn":
		l = logger.Warn
	case "info":
		l = logger.Info
	}
	return
}
