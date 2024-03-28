package model

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
)

// 消息头
type MessageHeader struct {
	OpCode               uint32 //操作码
	ResponseCode         uint32 //响应码
	OpFlag               uint32 //操作标志位
	SiteInfoSerialNumber uint16 //站点信息序列号
	RecursionCount       uint8  //服务的递归数目
	Reseved              uint8  //保留字段，未使用
	ExpirationTime       uint32 //消息过期时间
	BodyLength           uint32 //消息体长度
}

func (h *MessageHeader) HeaderToBytes() (msg []byte) {
	var bytebuf *bytes.Buffer = bytes.NewBuffer([]byte{})

	binary.Write(bytebuf, binary.BigEndian, uint32(h.OpCode))
	binary.Write(bytebuf, binary.BigEndian, uint32(h.ResponseCode))
	binary.Write(bytebuf, binary.BigEndian, uint32(h.OpFlag))
	binary.Write(bytebuf, binary.BigEndian, uint16(h.SiteInfoSerialNumber))
	binary.Write(bytebuf, binary.BigEndian, uint8(h.RecursionCount))
	binary.Write(bytebuf, binary.BigEndian, uint8(h.Reseved))
	binary.Write(bytebuf, binary.BigEndian, uint32(h.ExpirationTime))
	binary.Write(bytebuf, binary.BigEndian, uint32(h.BodyLength))
	data, _ := json.Marshal(h)
	fmt.Printf("Send MessageHeader:%s\n", data)

	return bytebuf.Bytes()
}

func (m *MessageHeader) ToObject(msgByte []byte) {
	m.OpCode = u.BytesToUint32(msgByte[20:24])
	m.ResponseCode = u.BytesToUint32(msgByte[24:28])
	m.OpFlag = u.BytesToUint32(msgByte[28:32])
	m.SiteInfoSerialNumber = u.BytesToUint16(msgByte[32:34])
	m.RecursionCount = u.BytesToUint8(msgByte[34:35])
	m.Reseved = u.BytesToUint8(msgByte[35:36])
	m.ExpirationTime = u.BytesToUint32(msgByte[36:40])
	m.BodyLength = u.BytesToUint32(msgByte[40:44])
	data, _ := json.Marshal(m)
	fmt.Printf("Recv MessageHeader:%s\n", data)
}
