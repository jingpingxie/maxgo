//
// @File:number
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/10 16:00
//
package number

import "encoding/binary"

//
// @Title:Int64ToBytes
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:13
// @Param:i
// @Return:[]byte
//
func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//
// @Title:BytesToInt64
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:11
// @Param:buf
// @Return:int64
//
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
