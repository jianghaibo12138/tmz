package statistic

import (
	"github.com/go-mysql-org/go-mysql/canal"
)

type SyncEventHandler struct {
	canal.DummyEventHandler
}

func (h *SyncEventHandler) OnRow(e *canal.RowsEvent) error {
	return nil
}

func (h *SyncEventHandler) String() string {
	return "SyncEventHandler"
}
