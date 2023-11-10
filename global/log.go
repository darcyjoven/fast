package global

import (
	"go.uber.org/zap"
)

var (
	// L 当执行CMD的时候可以用来记录日期
	L *zap.Logger
	// LogName 日志的名称
	LogName string
	// LogPath 日志的目录
	LogPath string
)
