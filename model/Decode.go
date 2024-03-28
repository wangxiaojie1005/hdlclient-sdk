package model

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/wangxiaojie1005/hdlclient-sdk/model/libs"
)

func DecodeMsgFindIDFromBytes(data []byte) MsgFindID {
	var msg MsgFindID
	// 使用 binary 包解码字节流
	var nameLength uint32
	byteReader := bytes.NewReader(data)
	// 读取 Name.Length
	binary.Read(byteReader, binary.BigEndian, &nameLength)
	// 读取 Name.Content
	nameContent := make([]byte, nameLength)
	byteReader.Read(nameContent)
	// 设置 MsgDeleteID 对象的属性
	msg.Name.Length = nameLength
	msg.Name.Content = string(nameContent)
	return msg
}

func DecodeMsgDeleteIDFromBytes(data []byte) MsgDeleteID {
	var msg MsgDeleteID
	// 使用 binary 包解码字节流
	var nameLength uint32
	byteReader := bytes.NewReader(data)
	// 读取 Name.Length
	binary.Read(byteReader, binary.BigEndian, &nameLength)
	// 读取 Name.Content
	nameContent := make([]byte, nameLength)
	byteReader.Read(nameContent)
	// 设置 MsgDeleteID 对象的属性
	msg.Name.Length = nameLength
	msg.Name.Content = string(nameContent)
	return msg
}

func DecodeMsgDeleteIdValueFromBytes(data []byte) MsgDeleteIdValue {
	var msg MsgDeleteIdValue

	// 使用 binary 包解码字节流
	byteReader := bytes.NewReader(data)

	// 读取 Name.Length
	var nameLength uint32
	binary.Read(byteReader, binary.BigEndian, &nameLength)
	// 读取 Name.Content
	nameContent := make([]byte, nameLength)
	byteReader.Read(nameContent)

	// 设置 MsgDeleteIdValue 对象的 Name 属性
	msg.Name.Length = nameLength
	msg.Name.Content = string(nameContent)

	// 读取 IndexList.IndexNumber
	binary.Read(byteReader, binary.BigEndian, &msg.IndexList.IndexNumber)

	// 读取 IndexList.Index 切片
	indexListLength := int(msg.IndexList.IndexNumber)
	msg.IndexList.Index = make([]uint32, indexListLength)
	for i := 0; i < indexListLength; i++ {
		binary.Read(byteReader, binary.BigEndian, &msg.IndexList.Index[i])
	}

	return msg
}

func DecodeMsgCreateIDFromBytes(data []byte) MsgCreateID {
	var msg MsgCreateID

	// 使用 binary 包解码字节流
	byteReader := bytes.NewReader(data)

	// 读取 Name.Length
	var nameLength uint32
	binary.Read(byteReader, binary.BigEndian, &nameLength)

	// 读取 Name.Content
	nameContent := make([]byte, nameLength)
	byteReader.Read(nameContent)

	// 设置 MsgCreateID 对象的 Name 属性
	msg.Name.Length = nameLength
	msg.Name.Content = string(nameContent)

	// 读取 ValueList.ValueNumber
	binary.Read(byteReader, binary.BigEndian, &msg.ValueList.ValueNumber)

	// 循环读取 ValueList.Value
	for i := 0; i < int(msg.ValueList.ValueNumber); i++ {
		var value BodyData

		// 读取 Index 和 TimeStamp
		binary.Read(byteReader, binary.BigEndian, &value.Index)
		binary.Read(byteReader, binary.BigEndian, &value.TimeStamp)

		// 读取 TTLType 和 TTL
		binary.Read(byteReader, binary.BigEndian, &value.TTLType)
		binary.Read(byteReader, binary.BigEndian, &value.TTL)

		// 读取 Permission
		permission, err := byteReader.ReadByte()
		if err != nil {
			fmt.Println(err)
		}
		value.Permission = permission

		// 读取 Type.Length 和 Type.Content
		var typeLength uint32
		binary.Read(byteReader, binary.BigEndian, &typeLength)
		typeContent := make([]byte, typeLength)
		byteReader.Read(typeContent)
		value.Type.Length = typeLength
		value.Type.Content = string(typeContent)

		// 根据类型读取数据
		switch value.Type.Content {
		case "EMAIL":
			// 读取 Email 类型的数据
			var email UTF8String
			binary.Read(byteReader, binary.BigEndian, &email.Length)
			emailContent := make([]byte, email.Length)
			byteReader.Read(emailContent)
			email.Content = string(emailContent)
			value.Data = email
		case "HS_VLIST":
			// 读取 HS_VLIST 类型的数据
			var vlist libs.HSVlist
			binary.Read(byteReader, binary.BigEndian, &vlist.ValueRefNumber)
			var numElements uint32
			binary.Read(byteReader, binary.BigEndian, &numElements)
			vlist.ValueRef = make(map[uint32]libs.ReferenceModel)
			// 读取 ValueRef 的每个元素
			for i := uint32(0); i < numElements; i++ {
				var refIndex uint32
				binary.Read(byteReader, binary.BigEndian, &refIndex)

				var nameLen uint32
				binary.Read(byteReader, binary.BigEndian, &nameLen)

				nameBuf := make([]byte, nameLen)
				byteReader.Read(nameBuf)
				refName := string(nameBuf)

				var refModel libs.ReferenceModel
				refModel.RefName = refName

				binary.Read(byteReader, binary.BigEndian, &refModel.RefIndex)

				vlist.ValueRef[refIndex] = refModel
			}
			value.Data = vlist
		}

		// 读取 References.ReferenceNumber
		binary.Read(byteReader, binary.BigEndian, &value.References.ReferenceNumber)

		// 读取 References.Reference.RefName.Length 和 References.Reference.RefName.Content
		if value.References.ReferenceNumber != 0 {
			for j := 0; j < int(value.References.ReferenceNumber); j++ {
				var refval ReferenceModule
				var refNameLength uint32
				binary.Read(byteReader, binary.BigEndian, &refNameLength)
				refvalContent := make([]byte, refNameLength)
				byteReader.Read(refvalContent)
				refval.RefName.Length = refNameLength
				refval.RefName.Content = string(refvalContent)
				binary.Read(byteReader, binary.BigEndian, &refval.RefIndex)
				value.References.Reference = append(value.References.Reference, refval)
			}
		}

		// 添加读取的值到 ValueList
		msg.ValueList.Value = append(msg.ValueList.Value, value)
	}

	return msg
}
