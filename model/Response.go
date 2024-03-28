package model

import (
	"bytes"
	"encoding/binary"
)

type ErrBody struct {
	ErrorMessage UTF8String
}

func (errBody *ErrBody) ErrToBytes() (msg []byte) {
	var byteBuf bytes.Buffer
	binary.Write(&byteBuf, binary.BigEndian, errBody.ErrorMessage.Length)
	byteBuf.Write([]byte(errBody.ErrorMessage.Content))
	return byteBuf.Bytes()
}

func DecodeErrBodyFromBytes(data []byte) ErrBody {
	var errBody ErrBody
	// 使用 binary 包解码字节流
	var nameLength uint32
	byteReader := bytes.NewReader(data)
	// 读取 Name.Length
	binary.Read(byteReader, binary.BigEndian, &nameLength)
	// 读取 Name.Content
	nameContent := make([]byte, nameLength)
	byteReader.Read(nameContent)
	// 设置 MsgDeleteID 对象的属性
	errBody.ErrorMessage.Length = nameLength
	errBody.ErrorMessage.Content = string(nameContent)
	return errBody
}
