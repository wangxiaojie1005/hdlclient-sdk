package handler

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"github.com/wangxiaojie1005/hdlclient-sdk/operate"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
	"math/rand"
)

func CreateID(id model.MsgCreateID, protocolType string) (res interface{}, err error) {
	//登录
	//var userBody []byte = userInfo.UserToBytes()
	//fmt.Println("r", userBody)
	//userMsgHeader := model.MessageHeader{}
	//var userHeader []byte = userMsgHeader.UserToBytes(uint32(len(userBody)))
	//userEnvelope := model.MessageEnvelope{}
	//var userLenth uint32 = uint32(len(userBody)) + uint32(len(userHeader))
	//userEnv := userEnvelope.EnvelopeToBytes(0, 0, userLenth)
	//var userMsg []byte
	//userMsg = append(userMsg, userEnv...)
	//userMsg = append(userMsg, userHeader...)
	//userMsg = append(userMsg, userBody...)
	//userReq := operate.Request{}
	//userReq.SendUdpMsg(userMsg)

	// 使用 encoding/json 库进行序列化
	msgBody := id.CreateToBytes()
	// 打印转换后的字节切片
	logrus.Info(msgBody)
	// 创建标识消息头结构
	msgHeader := model.MessageHeader{
		OpCode:               u.OC_CREATE,
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
	if protocolType == "udp" {
		res, err := req.SendUdpMsg(msg)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return "创建成功", nil
		}
		return res, nil
	} else if protocolType == "tcp" {
		res, err := req.SendTcpMsg(msg)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return "创建成功", nil
		}
		return res, nil
	} else {
		return res, errors.New("不支持该协议类型")
	}
}
