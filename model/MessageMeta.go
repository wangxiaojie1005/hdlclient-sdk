package model

import (
	"bytes"
	"encoding/binary"
)

// 创建标识元数据

type MsgCreatMeta struct {
	MetaID      UTF8String
	MetaName    UTF8String
	MetaVersion uint8
	IDCodeRule  UTF8String
	MetaDesc    UTF8String
	AttrList    AttrListModule
	MetaRefList MetaRefListModule
}

type AttrListModule struct {
	AttrNumber uint32
	AttrValue  []AttrValueModule
}

type AttrValueModule struct {
	Name        UTF8String
	Alias       UTF8String
	Type        UTF8String
	ReadLevel   uint8
	WriteLevel  uint8
	RWLevelDesv UTF8String
}

type MetaRefListModule struct {
	MetaRefNum uint8
	DataList   []DataListModule
}

type DataListModule struct {
	MetaID  UTF8String
	Version uint8
}

// 查询标识元数据

type FindMetaReq struct {
	MetaID UTF8String
}

type FindMetaRes struct {
	DataLen  uint8
	DataList []MsgCreatMeta
}

func (msg *MsgCreatMeta) MsgCreateMetaToBytes() []byte {
	var bytebuf bytes.Buffer

	// 编码 MetaID, MetaName, MetaVersion, IDCodeRule, MetaDesc
	encodeUTF8String(&bytebuf, msg.MetaID)
	encodeUTF8String(&bytebuf, msg.MetaName)
	binary.Write(&bytebuf, binary.BigEndian, msg.MetaVersion)
	encodeUTF8String(&bytebuf, msg.IDCodeRule)
	encodeUTF8String(&bytebuf, msg.MetaDesc)
	// 编码 AttrList
	binary.Write(&bytebuf, binary.BigEndian, msg.AttrList.AttrNumber)
	for _, attrValue := range msg.AttrList.AttrValue {
		encodeUTF8String(&bytebuf, attrValue.Name)
		encodeUTF8String(&bytebuf, attrValue.Alias)
		encodeUTF8String(&bytebuf, attrValue.Type)
		binary.Write(&bytebuf, binary.BigEndian, attrValue.ReadLevel)
		binary.Write(&bytebuf, binary.BigEndian, attrValue.WriteLevel)
		encodeUTF8String(&bytebuf, attrValue.RWLevelDesv)
	}
	// 编码 MetaRefList
	binary.Write(&bytebuf, binary.BigEndian, msg.MetaRefList.MetaRefNum)
	for _, dataList := range msg.MetaRefList.DataList {
		encodeUTF8String(&bytebuf, dataList.MetaID)
		binary.Write(&bytebuf, binary.BigEndian, dataList.Version)
	}
	return bytebuf.Bytes()
}

// 编码 UTF8String 到字节流
func encodeUTF8String(buf *bytes.Buffer, str UTF8String) {
	binary.Write(buf, binary.BigEndian, str.Length)
	buf.Write([]byte(str.Content))
}
