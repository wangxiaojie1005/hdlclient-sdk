package handler

import (
	"errors"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"github.com/wangxiaojie1005/hdlclient-sdk/operate"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
	"math/rand"
)

func AddIdentifierValues(identifier string, identifierValues model.ValueListModule, protocolType string) (res interface{}, err error) {
	addIdentifierValues := model.MsgCreateID{
		Name: model.UTF8String{
			Length:  uint32(len(identifier)),
			Content: identifier,
		},
		ValueList: identifierValues,
	}

	var body = addIdentifierValues.CreateToBytes()
	// 创建标识值消息头结构
	msgHeader := model.MessageHeader{
		OpCode:               u.OC_ADD_VALUE,
		ResponseCode:         0,
		OpFlag:               u.SetOFValue([]int{}),
		SiteInfoSerialNumber: 1,
		RecursionCount:       0,
		Reseved:              0,
		ExpirationTime:       0,
		BodyLength:           uint32(len(body)),
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
		MessageLen:     uint32(len(body)) + uint32(len(header)),
	}
	env := envelope.EnvelopeToBytes()

	var msg []byte
	msg = append(msg, env...)
	msg = append(msg, header...)
	msg = append(msg, body...)

	req := operate.Request{}
	if protocolType == "udp" {
		res, err := req.SendUdpMsg(msg)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return "添加标识值成功", nil
		}
		return res, nil
	} else if protocolType == "tcp" {
		res, err := req.SendTcpMsg(msg)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return "添加标识值成功", nil
		}
		return res, nil
	} else {
		return res, errors.New("不支持该协议类型")
	}
}
