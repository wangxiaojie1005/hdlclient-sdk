package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wangxiaojie1005/hdlclient-sdk/handler"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"time"
)

func main() {
	// 格式化日志记录
	logrus.SetFormatter(&logrus.JSONFormatter{})
	find()
	//create()
	//delete()
	//createMeta()
	//delVal()
	//addValues()
}
func create() {
	//userInfo := model.MsgUserinfo{
	//	UserName: model.UTF8String{
	//		Length:  uint32(len("wxj")),
	//		Content: "wxj",
	//	},
	//	PassWord: model.UTF8String{
	//		Length:  uint32(len("1234")),
	//		Content: "1234",
	//	},
	//}
	emailValue := model.BodyData{
		Index:      5678,
		TimeStamp:  uint32(time.Now().UnixNano() / int64(time.Millisecond)),
		TTLType:    0,
		TTL:        7200,
		Permission: 0,
		Type: model.UTF8String{
			Length:  uint32(len("EMAIL")),
			Content: "EMAIL",
		},
		Data: model.UTF8String{
			Length:  uint32(len("wxj@gmail.com")),
			Content: "wxj@gmail.com",
		},
		References: model.ReferencesModule{
			ReferenceNumber: 0,
		},
	}
	msgCreateID := model.MsgCreateID{
		Name: model.UTF8String{
			Length:  uint32(len("88.102")),
			Content: "88.102",
		},
		ValueList: model.ValueListModule{
			ValueNumber: 1,
			Value:       []model.BodyData{emailValue},
		},
	}
	res, err := handler.CreateID(msgCreateID, "tcp")
	if err != nil {
		fmt.Println(err)
		return
	}
	d, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println("create values", res)
	fmt.Println("create values", string(d))
}

func addValues() {
	emailValue := model.BodyData{
		Index:      5678,
		TimeStamp:  uint32(time.Now().UnixNano() / int64(time.Millisecond)),
		TTLType:    0,
		TTL:        7200,
		Permission: 0,
		Type: model.UTF8String{
			Length:  uint32(len("Email")),
			Content: "Email",
		},
		Data: model.UTF8String{
			Length:  uint32(len("wxj@gmail.com")),
			Content: "wxj@gmail.com",
		},
		References: model.ReferencesModule{
			ReferenceNumber: 0,
		},
	}
	ValueList := model.ValueListModule{
		ValueNumber: 1,
		Value:       []model.BodyData{emailValue},
	}
	//msgCreateID := model.MsgCreateID{
	//	Name: model.UTF8String{
	//		Length:  uint32(len("88.102")),
	//		Content: "88.102",
	//	},
	//	ValueList: model.ValueListModule{
	//		ValueNumber: 1,
	//		Value: []model.BodyData {emailValue},
	//	},
	//}
	res, err := handler.AddIdentifierValues("88.102", ValueList, "tcp")
	if err != nil {
		fmt.Println(err)
		return
	}
	d, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println("create", res)
	fmt.Println("create", string(d))
}

func createMeta() {
	msg := model.MsgCreatMeta{
		MetaID:      model.UTF8String{Length: uint32(len("meta01")), Content: "meta01"},
		MetaName:    model.UTF8String{Length: uint32(len("Example")), Content: "Example"},
		MetaVersion: 1,
		IDCodeRule:  model.UTF8String{Length: uint32(len("Rule001")), Content: "Rule001"},
		MetaDesc:    model.UTF8String{Length: uint32(len("Description")), Content: "Description"},
		AttrList: model.AttrListModule{
			AttrNumber: 1,
			AttrValue: []model.AttrValueModule{
				{
					Name:        model.UTF8String{Length: uint32(len("Attr1")), Content: "Attr1"},
					Alias:       model.UTF8String{Length: uint32(len("Alias1")), Content: "Alias1"},
					Type:        model.UTF8String{Length: uint32(len("Type1")), Content: "Type1"},
					ReadLevel:   1,
					WriteLevel:  2,
					RWLevelDesv: model.UTF8String{Length: uint32(len("Description")), Content: "Description"},
				},
			},
		},
		MetaRefList: model.MetaRefListModule{
			MetaRefNum: 1,
			DataList: []model.DataListModule{
				{
					MetaID:  model.UTF8String{Length: uint32(len("meta02")), Content: "meta02"},
					Version: 2,
				},
			},
		},
	}
	var body []byte = msg.MsgCreateMetaToBytes()
	fmt.Println(body)
	mod := model.DecodeMsgCreatMetaFromBytes(body)
	fmt.Println(mod)
}

func find() {
	//res, err := handler.FindID("10.1016/j.elerap.2011.05.004", "tcp")
	res, err := handler.FindID("88.102", "tcp")
	if err != nil {
		fmt.Println(err)
		return
	}
	d, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println("find", res)
	fmt.Println("find", string(d))
}

func delete() {
	logrus.Info("执行删除")
	res, err := handler.DeleteID("88.102", "tcp")
	if err != nil {
		fmt.Println(err)
		return
	}
	d, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println("delete", res)
	fmt.Println("delete", string(d))
}

func delVal() {
	msg := model.MsgDeleteIdValue{
		Name: model.UTF8String{
			Length:  uint32(len("10.1016/j.elerap.2011.05.005")),
			Content: "10.1016/j.elerap.2011.05.005",
		},
		IndexList: model.IndexListModule{
			IndexNumber: 3,
			Index:       []uint32{1, 2, 3},
		},
	}
	var body []byte = msg.DeleteValToBytes()
	fmt.Println(body)
	dec := model.DecodeMsgDeleteIdValueFromBytes(body)
	fmt.Println(dec)
}
