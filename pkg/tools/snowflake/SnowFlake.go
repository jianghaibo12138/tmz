package snowflake

import (
	"fmt"
	"stp_go/universal/stp_logger"
	"stp_go/universal/tools"
	"sync"
	"time"
)

type Snowflake struct {
	sync.Mutex         // 锁
	timestamp    int64 // 时间戳 ，毫秒
	workerId     int64 // 工作节点
	dataCenterId int64 // 数据中心机房id
	sequence     int64 // 序列号
}

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	dataCenterIdBits  = uint(2)                                        // 数据中心id所占位数
	workerIdBits      = uint(7)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	dataCenterIdMax   = int64(-1 ^ (-1 << dataCenterIdBits))           // 支持的最大数据中心id数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workerIdShift     = sequenceBits                                   // 机器id左移位数
	dataCenterIdShift = sequenceBits + workerIdBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + dataCenterIdBits // 时间戳左移位数
)

var sf Snowflake

var logger = stp_logger.LoggerRus{}

var ipv4Addr = ""

func init() {
	ipv4, err := tools.GetIpv4Address()
	if err != nil {
		logger.Error(fmt.Sprintf("[SnowFlake] meet error in getting ipv4 addr, err: %s, use default.", err.Error()))
		ipv4 = "123.123.123.123"
	}
	ipv4Addr = ipv4
}

//
// InitSnowFlake
// @Description: 初始化一个snowflake实例
// @param dataCenterId
// @return *Snowflake
//
func InitSnowFlake(dataCenterId int64) *Snowflake {
	stp_logger.InitLogrus("SnowFlake", true)
	sf.timestamp = time.Now().UnixNano() / 1000000
	sf.workerId = tools.Ipv42Int(ipv4Addr)
	sf.dataCenterId = dataCenterId
	return &sf
}

//
// NextVal
// @Description: 获取snowflake数
// @receiver s
// @return int64
//
func (s *Snowflake) NextVal() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		fmt.Printf(fmt.Sprintf("epoch must be between 0 and %d", timestampMax-1))
		return 0
	}
	s.timestamp = now
	r := int64((t)<<timestampShift | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | (s.sequence))
	s.Unlock()
	return r
}
