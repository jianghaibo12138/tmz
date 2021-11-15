package statistic

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
)

type SyncEventHandler struct {
	canal.DummyEventHandler
}

func (h *SyncEventHandler) OnRow(e *canal.RowsEvent) error {
	logger.Debug(fmt.Sprintf("[SyncEventHandler OnRow] Recv data trigger, action: %s database: %s, table: %s, row data: %+v", e.Action, e.Table.Schema, e.Table.Name, e.Rows))
	return nil
}

func (h *SyncEventHandler) String() string {
	return "SyncEventHandler"
}
