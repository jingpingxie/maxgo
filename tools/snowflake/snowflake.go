package snowflake

import (
	"fmt"
	logs "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	epoch = int64(1659196800000)
	//epoch             = int64(1659283200000)                           // 设置起始时间(时间戳/毫秒)：2022-08-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	datacenterIdBits  = uint(5)                                        // 数据中心id所占位数
	workerIdBits      = uint(5)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	datacenterIdMax   = int64(-1 ^ (-1 << datacenterIdBits))           // 支持的最大数据中心id数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workerIdShift     = sequenceBits                                   // 机器id左移位数
	datacenterIdShift = sequenceBits + workerIdBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + datacenterIdBits // 时间戳左移位数
)

//
// @Title:Snowflake
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:35:33
//
type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerId     int64
	datacenterId int64
	sequence     int64
}

//
// @Title:NewSnowflake
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:36:04
// @Param:datacenterId
// @Param:workerId
// @Return:*Snowflake
// @Return:error
//
func NewSnowflake(datacenterId, workerId int64) (*Snowflake, error) {
	if datacenterId < 0 || datacenterId > datacenterIdMax {
		return nil, fmt.Errorf("datacenterid must be between 0 and %d", datacenterIdMax-1)
	}
	if workerId < 0 || workerId > workerIdMax {
		return nil, fmt.Errorf("workerid must be between 0 and %d", workerIdMax-1)
	}
	return &Snowflake{
		timestamp:    0,
		datacenterId: datacenterId,
		workerId:     workerId,
		sequence:     0,
	}, nil
}

//
// @Title:NextVal
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:36:00
// @Receiver:s
// @Return:uint64
//
func (s *Snowflake) NextVal() uint64 {
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
		logs.Errorf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	r := (uint64)((t)<<timestampShift | (s.datacenterId << datacenterIdShift) | (s.workerId << workerIdShift) | (s.sequence))
	s.Unlock()
	return r
}

//
// @Title:GenerateSnowflakeId
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:34:38
// @Return:uint64
// @Return:error
//
func GenerateSnowflakeId() (uint64, error) {
	s, err := NewSnowflake(int64(0), int64(0))
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	//(int64): unique id
	id := s.NextVal()
	return id, nil
}
