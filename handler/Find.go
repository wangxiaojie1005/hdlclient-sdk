package handler

import (
	"errors"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"github.com/wangxiaojie1005/hdlclient-sdk/operate"
	u "github.com/wangxiaojie1005/hdlclient-sdk/units"
	"math/rand"
)

func FindID(id string, protocolType string) (interface{}, error) {
	findId := model.MsgFindID{
		Name: model.UTF8String{
			Length:  uint32(len(id)),
			Content: id,
		},
		IndexList: model.IndexListModule{},
		TypeList:  model.TypeListModule{},
	}
	var body []byte = findId.FindToBytes()
	// 查询标识消息头结构
	msgHeader := model.MessageHeader{
		OpCode:               u.OC_RESOLUTION,
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
		return res, nil
	} else if protocolType == "tcp" {
		res, err := req.SendTcpMsg(msg)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		return nil, errors.New("不支持该协议类型")
	}
	//创建标识
	//	jsonData := `{
	//			"value": [{
	//			   "index": "1",
	//			   "type": "URL",
	//			   "data": {
	//				  "format": "string",
	//				  "value": " www.test.com"
	//			   },
	//			   "ttl": "86400",
	//			   "ttlType": "0",
	//			   "timestamp": "0",
	//			   "adminRead": "1",
	//			   "adminWrite": "1",
	//			   "publicRead": "1",
	//			   "publicWrite": "0"
	//			}]
	//		}`
	//	var createIdParams httpModel.CreateId
	//	err := json.Unmarshal([]byte(jsonData), &createIdParams)
	//	if err != nil {
	//		fmt.Println("解析 JSON 数据失败:", err)
	//		return nil, err
	//	}
	//	// 输出解析后的结构体
	//	fmt.Println(createIdParams)
	//
	//	//添加标识值 修改标识值
	//	//"action": "addHandleValue","action": "modifyValue"
	//	addIdValuejsonData := `{
	//       "action": "modifyHandleValue",
	//    "value": [{
	//       "index": "3",
	//       "type": "HS_PRIMARY",
	//       "data": {
	//          "format": "string",
	//          "value": "2222"
	//       },
	//       "ttl": "86400",
	//       "ttlType": "0",
	//       "timestamp": "0",
	//       "adminRead": "1",
	//       "adminWrite": "1",
	//       "publicRead": "1",
	//       "publicWrite": "0"
	//    }]
	//}`
	//	var addIdValuejsonDataParams httpModel.AddIdValue
	//	err2 := json.Unmarshal([]byte(addIdValuejsonData), &addIdValuejsonDataParams)
	//	if err2 != nil {
	//		fmt.Println("解析 JSON 数据失败:", err)
	//		return nil, err
	//	}
	//	// 输出解析后的结构体
	//	fmt.Println(addIdValuejsonDataParams)
	//
	//	//添加标识值 修改标识值
	//	//"action": "addHandleValue","action": "modifyValue"
	//	removeIdValuejsonData := `{
	//       "action": "removeHandleValue",
	//			  "index": [
	//				 "3"
	//			   ]
	//		}`
	//	var removeIdValuejsonDataParams httpModel.RemoveIdValue
	//	err3 := json.Unmarshal([]byte(removeIdValuejsonData), &removeIdValuejsonDataParams)
	//	if err3 != nil {
	//		fmt.Println("解析 JSON 数据失败:", err)
	//		return nil, err
	//	}
	//	// 输出解析后的结构体
	//	fmt.Println(removeIdValuejsonDataParams)
	//
	//	userLoginValueData := `{
	//    "user":"88.1000.1/abc",
	//    "index":1
	//    }`
	//	var userLoginValueDataParams httpModel.UserLoginValue
	//	err4 := json.Unmarshal([]byte(userLoginValueData), &userLoginValueDataParams)
	//	if err4 != nil {
	//		fmt.Println("解析 JSON 数据失败:", err)
	//		return nil, err
	//	}
	//	// 输出解析后的结构体
	//	fmt.Println(userLoginValueDataParams)
	//	//req := operate.Request{}
	//	//req.SendHttpMsg("88.1000", "Delete")
	//
	//	req4 := operate.Request{}
	//	//查询标识
	//	//req.SendHttpMsg("88.9527", "Get", nil)
	//	req4.SendHttpMsg("88.1000", "Get", nil)
	//查询标识,条件查询
	//	//req.SendHttpMsg("88.1000", "Post", params)
	//	//创建标识
	//	//req.SendHttpMsg("88.9527", "Post", createIdParams)
	//	//fmt.Println("msgSend-------------", msg)
	//	//删除标识
	//	//req.SendHttpMsg("88.9527", "Delete", nil)
	//	//标识值添加
	//	//req.SendHttpMsg("88.9527", "Put", addIdValuejsonDataParams)
	//	//标识值修改
	//	//req.SendHttpMsg("88.9527", "Put", addIdValuejsonDataParams)
	//	//标识值删除
	//	//req.SendHttpMsg("88.9527", "Put", removeIdValuejsonDataParams)
	//	//分布式登录
	//	//req.SendHttpMsg("authentication", "Post", userLoginValueDataParams)
	//	return nil, err
}
