package tests

import (
	"fmt"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	"github.com/wangxiaojie1005/hdlclient-sdk/model/libs"
	"reflect"
	"testing"
	"time"
)

func TestFindBytes(t *testing.T) {
	tt := model.MsgFindID{
		Name: model.UTF8String{
			Length:  uint32(len("TestFindToBytes")),
			Content: "TestFindToBytes",
		},
	}
	var body []byte = tt.FindToBytes()
	fmt.Println(body)
	dec := model.DecodeMsgFindIDFromBytes(body)
	fmt.Println(dec)
	if !reflect.DeepEqual(dec, tt) {
		t.Error("Decoded does not match the original")
	}
}

func TestCreateBytes(t *testing.T) {
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
	vlistValue := model.BodyData{
		Index:      7890,
		TimeStamp:  uint32(time.Now().UnixNano() / int64(time.Millisecond)),
		TTLType:    0,
		TTL:        7200,
		Permission: 0,
		Type: model.UTF8String{
			Length:  uint32(len("HS_VLIST")),
			Content: "HS_VLIST",
		},
		Data: libs.HSVlist{
			ValueRefNumber: 2,
			ValueRef: map[uint32]libs.ReferenceModel{
				1: {RefName: "Ref1", RefIndex: 10},
				2: {RefName: "Ref2", RefIndex: 20},
			},
		},
		References: model.ReferencesModule{
			ReferenceNumber: 0,
		},
	}
	tt := model.MsgCreateID{
		Name: model.UTF8String{
			Length:  uint32(len("10.1016/j.elerap.2011.05.005")),
			Content: "10.1016/j.elerap.2011.05.005",
		},
		ValueList: model.ValueListModule{
			ValueNumber: 2,
			Value: []model.BodyData{
				emailValue, vlistValue,
			},
		},
	}
	var body []byte = tt.CreateToBytes()
	fmt.Println(body)
	dec := model.DecodeMsgCreateIDFromBytes(body)
	fmt.Println(dec)
	if !reflect.DeepEqual(dec, tt) {
		t.Error("Decoded does not match the original")
	}
}

func TestDeleteBytes(t *testing.T) {
	tt := model.MsgDeleteID{
		Name: model.UTF8String{
			Length:  uint32(len("10.1016/j.elerap.2011.05.005")),
			Content: "10.1016/j.elerap.2011.05.005",
		},
	}
	var body []byte = tt.DeleteToBytes()
	fmt.Println(body)
	dec := model.DecodeMsgDeleteIDFromBytes(body)
	fmt.Println(dec.Name.Length)
	if !reflect.DeepEqual(dec, tt) {
		t.Error("Decoded does not match the original")
	}
}

func TestDeleteValBytes(t *testing.T) {
	tt := model.MsgDeleteIdValue{
		Name: model.UTF8String{
			Length:  uint32(len("10.1016/j.elerap.2011.05.005")),
			Content: "10.1016/j.elerap.2011.05.005",
		},
		IndexList: model.IndexListModule{
			IndexNumber: 3,
			Index:       []uint32{1, 2, 3},
		},
	}
	var body []byte = tt.DeleteValToBytes()
	fmt.Println(body)
	dec := model.DecodeMsgDeleteIdValueFromBytes(body)
	fmt.Println(dec)
	if !reflect.DeepEqual(dec, tt) {
		t.Error("Decoded does not match the original")
	}
}
