package statistic

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/schema"
	"github.com/jianghaibo12138/TMZ/configs"
	"github.com/jianghaibo12138/TMZ/pkg/mariadb"
	"github.com/jianghaibo12138/TMZ/pkg/tmz_logger"
)

var logger = tmz_logger.LoggerRus{}

type BaseStatistic interface {
	ActionDispatch(table *schema.Table, action string, rows [][]interface{})
	ParseMainKeys(rows [][]interface{}) ([]string, []string)
	HandlerInsert(companyCodes []string, keys []string)
	HandlerUpdate(companyCodes []string, keys []string)
	HandlerDelete(companyCodes []string, keys []string)
}

func DispatchStatisticCall(bs BaseStatistic, table *schema.Table, action string, rows [][]interface{}) {
	bs.ActionDispatch(table, action, rows)
}

func BinlogSlaveRegister() {
	pos := logPosParser()
	c, err := binLogConn()
	if err != nil {
		panic(err)
	}
	c.SetEventHandler(&SyncEventHandler{})
	startPos := mysql.Position{
		Name: pos.File,
		Pos:  pos.Position,
	}
	_ = c.RunFrom(startPos)
}

func binLogConn() (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%d", configs.Settings.Mysql.Host, configs.Settings.Mysql.Port)
	cfg.User = configs.Settings.Mysql.User
	cfg.Password = configs.Settings.Mysql.Password
	cfg.Flavor = configs.Settings.Mysql.Flavor
	cfg.Dump.ExecutionPath = ""
	cfg.Dump.TableDB = configs.Settings.Mysql.Database
	cfg.Dump.Tables = configs.Settings.Mysql.ListenTables
	return canal.NewCanal(cfg)
}

func logPosParser() *LogPos {
	session := mariadb.GetConnect()
	var pos LogPos
	result := session.Raw("SHOW MASTER STATUS").Scan(&pos)
	if result.Error != nil {
		logger.Debug(fmt.Sprintf("[logPosParser] scan MASTER STATUS err: %s", result.Error.Error()))
		return nil
	}
	logger.Debug(fmt.Sprintf("[logPosParser] scan MASTER STATUS success, pos: %+v", pos))
	return &pos
}
