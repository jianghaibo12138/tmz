package snowflake

import (
	"sync"
	"testing"
	"time"
)

func TestSnowflake_NextVal(t *testing.T) {
	type fields struct {
		Mutex        sync.Mutex
		timestamp    int64
		workerId     int64
		dataCenterId int64
		sequence     int64
	}
	f := fields{
		timestamp:    time.Now().UnixMilli(),
		workerId:     0,
		dataCenterId: 0,
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "c1", fields: f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snowflake{
				timestamp:    tt.fields.timestamp,
				workerId:     tt.fields.workerId,
				dataCenterId: tt.fields.dataCenterId,
				sequence:     tt.fields.sequence,
			}
			for i := 0; i < 100; i++ {
				go func() {
					ret := s.NextVal()
					t.Log(ret)
				}()
			}
		})
	}
}

func TestInitSnowFlake(t *testing.T) {
	type args struct {
		dataCenterId int64
	}
	var c1 = args{
		dataCenterId: 0,
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "c1", args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := InitSnowFlake(tt.args.dataCenterId)
			t.Logf("%+v", sf)
		})
	}
}
