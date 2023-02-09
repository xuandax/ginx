package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xuandax/ginx/g"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

// SugarLogger 性能不错，可结构化和调用printf样式的API（建议优先使用）
func NewZapSugarLogger() (sugar *zap.SugaredLogger) {
	logger := newSugarLogger()
	defer logger.Sync()
	sugar = logger.Sugar()
	return
}

// 当性能和类型安全至关重要时使用，它比SugaredLogger更快，分配的数量更少，但仅支持结构化日志记录（不建议）
func NewZapLogger() (logger *zap.Logger) {
	logger = newLogger()
	defer logger.Sync()
	return
}

// 创建Logger
func newLogger() (logger *zap.Logger) {
	//设置info日志级别处理逻辑
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})
	//设置error日志级别处理逻辑
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	//设置writer
	infoWriter := getWriterByLumberjack(g.ServerConfig.GetString("log.dir")+g.ServerConfig.GetString("log.info_filename"), g.ServerConfig.GetString("log.ext"))
	errorWriter := getWriterByLumberjack(g.ServerConfig.GetString("log.dir")+g.ServerConfig.GetString("log.error_filename"), g.ServerConfig.GetString("log.ext"))

	//设置编码配置
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, os.Stdout, infoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	logger = zap.New(core, zap.AddCaller())
	return logger
}

// 创建SugarLogger
func newSugarLogger() (logger *zap.Logger) {
	//设置info日志级别处理逻辑
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})
	//设置error日志级别处理逻辑
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	//设置writer
	infoWriter := getWriterByLumberjack(g.ServerConfig.GetString("log.sugar_dir")+g.ServerConfig.GetString("log.sugar_info_filename"), g.ServerConfig.GetString("log.sugar_ext"))
	errorWriter := getWriterByLumberjack(g.ServerConfig.GetString("log.sugar_dir")+g.ServerConfig.GetString("log.sugar_error_filename"), g.ServerConfig.GetString("log.sugar_ext"))

	//设置编码配置
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, os.Stdout, infoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	logger = zap.New(core, zap.AddCaller())
	return logger
}

// 根据lumberjack实现日志分割
func getWriterByLumberjack(filePath string, fileExt string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filePath + fileExt,
		MaxSize:    g.ServerConfig.GetInt("log.max_size"),    //M
		MaxAge:     g.ServerConfig.GetInt("log.max_age"),     //days 保留旧日志文件的最大天数
		MaxBackups: g.ServerConfig.GetInt("log.max_backups"), //要保留的最大旧日志文件数
		LocalTime:  g.ServerConfig.GetBool("log.local_time"), //是否格式化时间戳的时间
		Compress:   g.ServerConfig.GetBool("log.compress"),   //是否使用gzip压缩， 默认不压缩
	}
}

// 根据rotalogs实现日志分割
func getWriterByRotatelogs(filePath string, fileExt string) io.Writer {
	logs, err := rotatelogs.New(
		filePath+"%Y%m%d%H%M"+fileExt,
		//rotatelogs.WithLinkName(filename),
		rotatelogs.WithRotationSize(1*1024*1024), //单位 byte
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(10*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	return logs
}
