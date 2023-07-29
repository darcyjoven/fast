package cmd

// 引入相关的包
import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fast/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化日志
func initLogger() {
	// 初始化全局日志
	err := initLogGlobal()
	// 如果初始化失败，抛出错误
	if err != nil {
		panic(err)
	}
	// 创建一个新的原子级别
	atom := zap.NewAtomicLevel()
	// 设置日志级别为Debug
	atom.SetLevel(zap.DebugLevel)
	// 配置日志
	cfg := zap.Config{
		Level:            atom,
		Encoding:         "console",
		OutputPaths:      []string{"stdout", filepath.Join(global.LogPath, global.LogName)},
		ErrorOutputPaths: []string{"stderr"},
	}
	// 配置编码器
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 构建日志
	logger, err := cfg.Build(
		zap.Fields(zap.Int("pid", os.Getpid())),
		zap.AddCaller(),
		zap.Development(),
	)
	// 如果构建失败，抛出错误
	if err != nil {
		panic(err)
	}
	// 设置全局日志
	global.L = logger

}

// 自定义时间编码器
func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	// 添加时间字符串
	encoder.AppendString("[fast]" + t.Format("2006/01/02 - 15:04:05.000"))
}

// 初始化全局日志
func initLogGlobal() (err error) {
	// 从配置中获取日志目录，如果没有则使用默认值
	global.LogPath = viper.GetString("logdir")
	if global.LogPath == "" {
		global.LogPath = "temp/"
	}

	// 如果目录不存在，则创建
	if _, err := os.Stat(global.LogPath); os.IsNotExist(err) {
		if err = os.MkdirAll(global.LogPath, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// 从配置中获取日志名称，如果没有则使用默认值
	global.LogName = viper.GetString("logname")
	if global.LogName == "" {
		global.LogName = "run"
	}

	// 添加日志扩展名
	global.LogName += logExt()

	return nil
}

// 获取日志名后缀
func logExt() string {
	// 获取上海时间
	local, _ := time.LoadLocation("Asia/Shanghai")
	date := time.Now().In(local)
	// 根据时间间隔设置日志名称
	switch viper.GetString("loginterval") {
	case "one":
		return ".log"
	case "every":
		return fmt.Sprintf(".%s.%d.log", date.Format("2006-01-02_19.54.000"), os.Getpid())
	case "year":
		return fmt.Sprintf(".%s.log", date.Format("2006"))
	case "month":
		return fmt.Sprintf(".%s.log", date.Format("2006-01"))
	case "week":
		return fmt.Sprintf(".%s.log", date.Format("2006-01-Feb"))
	case "day":
		return fmt.Sprintf(".%s.log", date.Format("2006-01-02"))
	default:
		return fmt.Sprintf(".%s.log", date.Format("2006-01-02"))
	}
}
