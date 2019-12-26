package log

import (
	"fmt"

	"github.com/bndr/gotabulate"
	"github.com/ltto/T/gobox/str"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		//EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeLevel: zapcore.CapitalColorLevelEncoder, //这里可以指定颜色
		//EncodeLevel:    zapcore.LowercaseColorLevelEncoder, //这里可以指定颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	config := zap.Config{
		//Encoding:         "json",                                              // 输出格式 console 或 json
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),                 // 日志级别
		Development:      true,                                                 // 开发模式，堆栈跟踪
		Encoding:         "console",                                            // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                        // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "wisdom_park"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout",},                                  // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	// 构建日志
	ZapLog, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}

	row_1 := []interface{}{"john", 20, "ready"}
	row_2 := []interface{}{"bndr", 23, "ready"}

	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row_1, row_2})

	// Set the Headers (optional)
	t.SetHeaders([]string{"age", "status"})

	// Set the Empty String (optional)
	t.SetEmptyString("None")

	// Set Align (Optional)
	t.SetAlign("right")

	ZapLog.Debug("DEBUG", zap.String("s", "s"))
	ZapLog.Info("DEBUG", zap.String("s", "s"))
	ZapLog.Warn("DEBUG", zap.String("s", "s"))
	ZapLog.Error("DEBUG", zap.String("s", "s"))
}

func logFmt(logStr string, params ...interface{}) string {
	index := 0
	return str.ExpandS(logStr, func(s string) string {
		if index >= len(params) {
			return fmt.Sprintf("{%v}", s)
		} else {
			return fmt.Sprintf("%v", params[index])
		}
	})
}
func Debug(logStr string, params ...interface{}) {
	zapLog.Debug(logFmt(logStr, params), )
}
