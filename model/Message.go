package model

import (
	"bytes"
	"encoding/binary"
	"github.com/wangxiaojie1005/hdlclient-sdk/model/libs"
)

// /////////////////////////消息体////////////////////////////////////////////////////////
// 标识Value定义
type BodyData struct {
	Index      uint32
	TimeStamp  uint32
	TTLType    uint8
	TTL        uint32
	Permission uint8
	Type       UTF8String
	Data       interface{}
	References ReferencesModule
}

// 引用
type ReferencesModule struct {
	ReferenceNumber uint32
	Reference       []ReferenceModule
}

// 引用值
type ReferenceModule struct {
	RefName  UTF8String
	RefIndex uint32
}

// 用户名密码
type MsgUserinfo struct {
	UserName UTF8String //标识名称
	PassWord UTF8String // 标识值
}

// 创建标识
type MsgCreateID struct {
	Name      UTF8String      //标识名称
	ValueList ValueListModule // 标识值
}

// 标识值 - 数组
type ValueListModule struct {
	ValueNumber uint32
	Value       []BodyData
}

// 删除标识
type MsgDeleteID struct {
	Name UTF8String
}

// 查询标识
type MsgFindID struct {
	Name      UTF8String
	IndexList IndexListModule
	TypeList  TypeListModule
}

type IndexListModule struct {
	IndexNumber uint32
	Index       []uint32
}

type TypeListModule struct {
	TypeNumber uint32
	Type       []UTF8String
}

// 增加标识值
type MsgAddIdValue struct {
	Name      UTF8String
	ValueList ValueListModule
}

// 修改标识值
type MsgModifyIdValue struct {
	Name      UTF8String
	ValueList ValueListModule
}

// 删除标识值
type MsgDeleteIdValue struct {
	Name      UTF8String
	IndexList IndexListModule
}

// 应答消息(成功状态，无消息提)
type MsgReponse struct {
	ErrorMessage UTF8String
}

func (find *MsgFindID) FindToBytes() (msg []byte) {
	var bytebuf bytes.Buffer
	// 遍历 ValueList，写入每个标识值的信息
	binary.Write(&bytebuf, binary.BigEndian, find.Name.Length)
	bytebuf.Write([]byte(find.Name.Content))
	binary.Write(&bytebuf, binary.BigEndian, find.IndexList.IndexNumber)
	for _, v := range find.IndexList.Index {
		binary.Write(&bytebuf, binary.BigEndian, v)
	}
	binary.Write(&bytebuf, binary.BigEndian, find.TypeList.TypeNumber)
	for _, v := range find.TypeList.Type {
		binary.Write(&bytebuf, binary.BigEndian, v.Length)
		bytebuf.Write([]byte(find.Name.Content))
	}
	//name := new(UTF8String)
	//name.Content = c
	//name.Length = uint32(len(c))
	//return name.ToBytes(c)
	return bytebuf.Bytes()
}

func (user *MsgUserinfo) UserToBytes() (msg []byte) {
	var bytebuf bytes.Buffer
	binary.Write(&bytebuf, binary.BigEndian, user.UserName.Length)
	bytebuf.Write([]byte(user.UserName.Content))
	binary.Write(&bytebuf, binary.BigEndian, 0)
	binary.Write(&bytebuf, binary.BigEndian, 1)
	binary.Write(&bytebuf, binary.BigEndian, user.PassWord.Length)
	bytebuf.Write([]byte(user.PassWord.Content))
	return bytebuf.Bytes()
}

func (create *MsgCreateID) CreateToBytes() (msg []byte) {
	var bytebuf bytes.Buffer
	// 遍历 ValueList，写入每个标识值的信息
	binary.Write(&bytebuf, binary.BigEndian, create.Name.Length)
	bytebuf.Write([]byte(create.Name.Content))
	binary.Write(&bytebuf, binary.BigEndian, create.ValueList.ValueNumber)
	for _, v := range create.ValueList.Value {
		binary.Write(&bytebuf, binary.BigEndian, v.Index)
		binary.Write(&bytebuf, binary.BigEndian, v.TimeStamp)
		bytebuf.WriteByte(v.TTLType)
		binary.Write(&bytebuf, binary.BigEndian, v.TTL)
		bytebuf.WriteByte(v.Permission)
		binary.Write(&bytebuf, binary.BigEndian, v.Type.Length)
		bytebuf.Write([]byte(v.Type.Content))
		// 根据类型写入数据
		switch v.Type.Content {
		case "EMAIL":
			// 编码 Email 类型的数据
			email := v.Data.(UTF8String)
			binary.Write(&bytebuf, binary.BigEndian, email.Length)
			bytebuf.Write([]byte(email.Content))
		case "HS_VLIST":
			// 编码 HS_VLIST 类型的数据
			HS_VLISTData := v.Data.(libs.HSVlist)
			binary.Write(&bytebuf, binary.BigEndian, HS_VLISTData.ValueRefNumber)
			binary.Write(&bytebuf, binary.BigEndian, uint32(len(HS_VLISTData.ValueRef)))
			for key, value := range HS_VLISTData.ValueRef {
				binary.Write(&bytebuf, binary.BigEndian, key)
				encodeUTF8String(&bytebuf, UTF8String{Length: uint32(len(value.RefName)), Content: value.RefName})
				binary.Write(&bytebuf, binary.BigEndian, value.RefIndex)
			}
		}
		// 编码引用值
		binary.Write(&bytebuf, binary.BigEndian, v.References.ReferenceNumber)
		if v.References.ReferenceNumber > 0 {
			for _, m := range v.References.Reference {
				binary.Write(&bytebuf, binary.BigEndian, m.RefName.Length)
				bytebuf.Write([]byte(m.RefName.Content))
				binary.Write(&bytebuf, binary.BigEndian, m.RefIndex)
			}
		}
	}
	return bytebuf.Bytes()
}

func (delete *MsgDeleteID) DeleteToBytes() (msg []byte) {
	var bytebuf bytes.Buffer
	binary.Write(&bytebuf, binary.BigEndian, delete.Name.Length)
	bytebuf.Write([]byte(delete.Name.Content))
	return bytebuf.Bytes()
}

func (deleteval *MsgDeleteIdValue) DeleteValToBytes() (msg []byte) {
	var bytebuf bytes.Buffer
	binary.Write(&bytebuf, binary.BigEndian, deleteval.Name.Length)
	bytebuf.Write([]byte(deleteval.Name.Content))
	binary.Write(&bytebuf, binary.BigEndian, deleteval.IndexList.IndexNumber)
	for _, v := range deleteval.IndexList.Index {
		binary.Write(&bytebuf, binary.BigEndian, v)
	}
	return bytebuf.Bytes()
}
