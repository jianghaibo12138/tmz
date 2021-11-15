package tmz_logger

import (
	"fmt"
	"github.com/jianghaibo12138/TMZ/configs"
	"github.com/jianghaibo12138/TMZ/pkg/tools"
	"github.com/labstack/gommon/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type LoggerRus struct{}

var logRus *logrus.Logger

func InitLogrus(moduleName string, debug bool) *logrus.Logger {
	logDirPath := path.Join(configs.GetHomePath(), "logs")
	if !tools.IsDir(logDirPath) {
		err := os.Mkdir(logDirPath, 0777)
		if err != nil {
			panic(err)
		}
	}
	baseLogPath := path.Join(logDirPath, fmt.Sprintf("%s.log", moduleName))

	logRus = logrus.New()
	logRus.Formatter = new(logrus.JSONFormatter)
	// logRus.Formatter = new(logrus.TextFormatter)                     //default
	// logRus.Formatter.(*logrus.TextFormatter).ForceColors = true    // remove colors
	// logRus.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	// logRus.SetReportCaller(true)

	logRus.Level = logrus.DebugLevel
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	  `WithMaxAge` 设置文件清理前的最长保存时间
	  `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(90*24*time.Hour), // 文件最大保存时间
		// rotatelogs.WithRotationCount(5),
		// rotatelogs.WithRotationTime(7*time.Hour), // 日志切割时间间隔
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}
	logRus.SetOutput(writer)
	return logRus
}

func (logger *LoggerRus) Debug(message string) {
	fmt.Println(message)
	logRus.Debug(message)
}

func (logger *LoggerRus) Info(message string) {
	logRus.Info(message)
}

func (logger *LoggerRus) Warning(message string) {
	logRus.Warn(message)
}

func (logger *LoggerRus) Error(message string) {
	logRus.Error(message)
}

func (logger *LoggerRus) Fatal(message string) {
	logRus.Fatal(message)
}

func (logger *LoggerRus) Panic(message string) {
	logRus.Panic(message)
}
