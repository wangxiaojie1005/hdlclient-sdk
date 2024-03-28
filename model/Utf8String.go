package model

import (
	"encoding/json"
	"fmt"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
)

// UTF8-String 类型
type UTF8String struct {
	Length  uint32
	Content string
}

func (utfString *UTF8String) ToBytes(c string) (msg []byte) {
	tmp_content := u.String2byte(c)
	utfString.Length = uint32(len(tmp_content))
	var content []byte
	//bytesBuffer := bytes.NewBuffer([]byte{})
	//binary.Write(bytesBuffer, binary.BigEndian, uint32(utfString.Length))
	//content = append(content, bytesBuffer.Bytes()...)
	ccc := u.UintToBytes(int(utfString.Length), 32)
	content = append(content, ccc...)
	fmt.Printf("%032b\n", ccc)
	content = append(content, tmp_content...)
	return content
}

func (m *UTF8String) ToObject(msgByte []byte) {
	m.Length = u.BytesToUint32(msgByte[44:48])
	m.Content = string(msgByte[48:(48 + m.Length)])
	data, _ := json.Marshal(m)
	fmt.Printf("Recv BodyData:%s\n", data)
}
