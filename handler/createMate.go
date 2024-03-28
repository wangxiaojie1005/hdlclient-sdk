package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"github.com/wangxiaojie1005/hdlclient-sdk/operate"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
	"math/rand"
)

func CreateMate(mate model.MsgCreatMeta, protocolType string) (response model.MessageBody, err error) {
	// 使用 encoding/json 库进行序列化
	msgBody := mate.MsgCreateMetaToBytes()
	// 打印转换后的字节切片
	logrus.Info(msgBody)
	// 创建标识消息头结构
	msgHeader := model.MessageHeader{
		OpCode:               u.OC_META_CREATE,
		ResponseCode:         0,
		OpFlag:               u.SetOFValue([]int{}),
		SiteInfoSerialNumber: 1,
		RecursionCount:       0,
		Reseved:              0,
		ExpirationTime:       0,
		BodyLength:           uint32(len(msgBody)),
	}
	header := msgHeader.HeaderToBytes()
	// 消息信封结构
	envelope := model.MessageEnvelope{
		MajorVersion:   2,
		MinorVersion:   5,
		MessageFlag:    u.SetMValue([]int{}),
		SessionId:      0,
		RequestId:      rand.Uint32(),
		SequenceNumber: 0,
		MessageLen:     uint32(len(msgBody)) + uint32(len(header)),
	}
	env := envelope.EnvelopeToBytes()

	var msg []byte
	msg = append(msg, env...)
	msg = append(msg, header...)
	msg = append(msg, msgBody...)

	req := operate.Request{}
	req.SendUdpMsg(msg)
	//if protocolType == "tcp" {
	//	response = req.SendTcpMsg(msg)
	//} else if protocolType == "udp" {
	//	response = req.SendUdpMsg(msg)
	//} else if protocolType == "http" {
	//	response = req.SendHttpMsg(id)
	//} else {
	//	return response, errors.New("协议类型错误！")
	//}
	return response, nil
}
