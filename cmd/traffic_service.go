package main

import (
	"fmt"
	"github.com/jianghaibo12138/TMZ/pkg/mariadb"
	"github.com/jianghaibo12138/TMZ/pkg/tmz_logger"
	"github.com/jianghaibo12138/TMZ/services/statistic"
	"os"
	"os/signal"
)

func main() {
	tmz_logger.InitLogrus("tmz", true)
	var logger = tmz_logger.LoggerRus{}
	mariadb.GetConnect()
	defer func() {
		err := mariadb.CloseConnect()
		if err != nil {
			logger.Fatal(fmt.Sprintf("[TMZ] got error in disconnecting from DB: %s", err.Error()))
			return
		}
		logger.Debug("[TMZ] disconnect from DB")
	}()

	go func() {
		statistic.BinlogSlaveRegister()
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	//logger.Debug("[traffic service main] Server exiting")
}
