/**
 * @Author: cloudintheking
 * @Description:
 * @File: logger
 * @Version: 1.0.0
 * @Date: 2021/10/22 15:41
 */
package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	Env         string            `yaml:"env"`         //环境 develop、production
	Path        string            `yaml:"path"`        //存放路径
	MaxSize     int               `yaml:"maxSize"`     //最大存储容量,单位M
	MaxBackup   int               `yaml:"maxBackup"`   //最大备份数量
	MaxAge      int               `yaml:"maxAge"`      //保留天数
	Compress    bool              `yaml:"compress"`    //是否压缩
	Level       string            `yaml:"level"`       //日志等级
	Encoding    string            `yaml:"encoding"`    //输出格式 console、json
	ExtraFields map[string]string `yaml:"extraFields"` //额外字段
}

/**
 *  @Description: 初始化日志器
 *  @param config 日志配置
 *  @return *zap.Logger
 */
func InitLogger(config *Config) *zap.Logger {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   config.Path,      // 日志文件路径，默认 os.TempDir()
		MaxSize:    config.MaxSize,   // 每个日志文件保存10M，默认 100M
		MaxBackups: config.MaxBackup, // 保留30个备份，默认不限
		MaxAge:     config.MaxAge,    // 保留7天，默认不限
		Compress:   config.Compress,  // 是否压缩，默认不压缩
	}
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "panic":
		level = zap.DPanicLevel
	default:
		level = zap.InfoLevel
	}
	var encodeConfig zapcore.EncoderConfig
	switch config.Env {
	case "develop":
		encodeConfig = zap.NewDevelopmentEncoderConfig()
	case "production":
		encodeConfig = zap.NewProductionEncoderConfig()
	default:
		encodeConfig = zap.NewDevelopmentEncoderConfig()
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	writer := zapcore.AddSync(&hook)
	var encoder zapcore.Encoder
	switch config.Encoding {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encodeConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encodeConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encodeConfig)
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)), // 打印到控制台和文件
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	fields := make([]zap.Field, 0)
	if config.ExtraFields != nil {
		for k, v := range config.ExtraFields {
			fields = append(fields, zap.String(k, v))
		}
	}
	// 构造日志
	logger := zap.New(core).WithOptions(caller, development, zap.AddStacktrace(zapcore.ErrorLevel), zap.Fields(fields...))
	logger.Info("DefaultLogger init success")
	return logger
}
