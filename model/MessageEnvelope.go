package model

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
)

// 消息信封
type MessageEnvelope struct {
	MajorVersion   uint8  //主版本
	MinorVersion   uint8  //次版本
	MessageFlag    uint16 //消息标志位
	SessionId      uint32 //会话ID
	RequestId      uint32 //请求ID
	SequenceNumber uint32 //序列号，在消息截断的情况下计数
	MessageLen     uint32 //消息长度，这里指除消息信封以外的长度，包含消息头、消息体和消息凭据。
}

func (msgEnv *MessageEnvelope) New() *MessageEnvelope {
	return &MessageEnvelope{}
}
func (msgEnv *MessageEnvelope) EnvelopeToBytes() (msg []byte) {
	var bytebuf *bytes.Buffer = bytes.NewBuffer([]byte{})

	binary.Write(bytebuf, binary.BigEndian, uint8(msgEnv.MajorVersion))
	binary.Write(bytebuf, binary.BigEndian, uint8(msgEnv.MinorVersion))
	binary.Write(bytebuf, binary.BigEndian, uint16(msgEnv.MessageFlag))
	binary.Write(bytebuf, binary.BigEndian, uint32(msgEnv.SessionId))
	binary.Write(bytebuf, binary.BigEndian, uint32(msgEnv.RequestId))
	binary.Write(bytebuf, binary.BigEndian, uint32(msgEnv.SequenceNumber))
	binary.Write(bytebuf, binary.BigEndian, uint32(msgEnv.MessageLen))
	return bytebuf.Bytes()
}

func (m *MessageEnvelope) ToObject(msgByte []byte) {
	m.MajorVersion = u.BytesToUint8(msgByte[:1])
	m.MinorVersion = u.BytesToUint8(msgByte[1:2])
	m.MessageFlag = u.BytesToUint16(msgByte[2:4])
	m.SessionId = u.BytesToUint32(msgByte[4:8])
	m.RequestId = u.BytesToUint32(msgByte[8:12])
	m.SequenceNumber = u.BytesToUint32(msgByte[12:16])
	m.MessageLen = u.BytesToUint32(msgByte[16:20])
	data, _ := json.Marshal(m)
	fmt.Printf("Recv MessageEnvelope:%s\n", data)
}

// 设置2个字节位，转为Uint16类型
// 可用于Envelope opFlag
func SetEnvelopeOpflag(args ...int) uint16 {

	var n string = "0000000000000000"
	n2 := []byte(n)
	for _, arg := range args {
		if arg > 0 {
			n2[arg-1] = '1'
		}

	}
	decimal, _ := u.BinaryToUint16(string(n2))
	return decimal
}

// 解析消息信封Flag
func GetMessageFlagValue(flag uint16) []int {
	l := []int{}
	binary, _ := u.DecimalToBinary(int(flag))
	//println(binary)
	for i, v := range binary {
		if string(v) == `1` {
			l = append(l, i+1)
		}
	}
	return l
	//33664  1000001110000000
}
