package handler

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"github.com/wangxiaojie1005/hdlclient-sdk/operate"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
	"math/rand"
)

func DeleteID(id string, protocolType string) (res interface{}, err error) {
	deleteId := model.MsgDeleteID{
		Name: model.UTF8String{
			Length:  uint32(len(id)),
			Content: id,
		},
	}
	var body []byte = deleteId.DeleteToBytes()
	// 删除标识消息头结构
	msgHeader := model.MessageHeader{
		OpCode:               u.OC_DELETE,
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
	logrus.Info("完成消息封装")

	req := operate.Request{}
	if protocolType == "udp" {
		res, err := req.SendUdpMsg(msg)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return "删除成功", nil
		}
		return res, nil
	} else if protocolType == "tcp" {
		res, err := req.SendTcpMsg(msg)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return "删除成功", nil
		}
		return res, nil
	} else {
		return res, errors.New("不支持该协议类型")
	}
}
