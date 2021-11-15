// Package mariadb
/**
 * @Time: 2021/8/11 10:58 上午
 * @Author: Haibo Jiang
 * @Email: haibojiang@bitorobotics.ltd
 * @File: sql_logger.go
 * @Version: V0.0.1
 * @license: (C) Copyright 2017-2030, Bito Robotics Co.Ltd.
 * @desc:
**/
package mariadb

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"stp_go/universal/stp_logger"
)

type SQLLogger struct {
	logger *logrus.Logger
}

// Printf
//实现gorm/logger.Writer接口
func (s *SQLLogger) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	s.logger.Debug(logStr)
}
func NewSQLLogger() *SQLLogger {
	log := stp_logger.InitLogrus("sql", false)
	return &SQLLogger{
		logger: log,
	}
}
