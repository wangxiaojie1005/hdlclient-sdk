package model

import (
	"encoding/json"
	"github.com/wangxiaojie1005/hdlclient-sdk/model/libs"
	"github.com/wangxiaojie1005/hdlclient-sdk/units"
)

// 消息体
type MessageBody struct {
	Name      string
	ValueList interface{}
}

func (m *MessageBody) ToObject(data []byte, opCode uint32) {
	switch opCode {
	case units.OC_GET_SITE_INFO:
		m.Name = "HS_SITE"
		body := new(libs.HSSite)
		body.ToObj(data)
		m.ValueList = body
	default:
		m.dataToObject(data)
	}
}

func (m *MessageBody) dataToObject(data []byte) {
	/////NAME
	off := uint32(0)
	p_name_length := units.BytesToUint32(data[off:4])
	off = off + 4
	p_name_content := units.Byte2string(data[off : off+p_name_length])
	off = off + p_name_length

	/////ValueList ValueNumber
	p_valueNumber := units.BytesToUint32(data[off : off+4])
	off = off + 4

	valueMap := make([]ItemData, p_valueNumber)
	for ii := 0; ii < int(p_valueNumber); ii++ {

		// value Index
		p_valueIndex := units.BytesToUint32(data[off : off+4])
		off = off + 4
		//value timeStam
		p_timeStamp := units.BytesToUint32(data[off : off+4])
		off = off + 4
		//value TTLType
		p_ttlType := units.BytesToUint8(data[off : off+1])
		off = off + 1
		//value TTL
		p_ttl := units.BytesToUint32(data[off : off+4])
		off = off + 4
		//value Permission
		p_permission := units.BytesToUint8(data[off : off+1])
		off = off + 1
		//value Type
		p_type_length := units.BytesToUint32(data[off : off+4])
		off = off + 4
		p_type_content := units.Byte2string(data[off : off+p_type_length])
		off = off + p_type_length

		//value Data
		p_data_length := units.BytesToUint32(data[off : off+4])
		off = off + 4
		p_data_content := data[off : off+p_data_length]
		off = off + p_data_length

		//value reference Number
		p_ref_num := units.BytesToUint32(data[off : off+4])
		off = off + 4
		var ref libs.ReferencesModel = libs.ReferencesModel{}
		if p_ref_num > 0 {

			referenceMap := make(map[uint32]libs.ReferenceModel)
			for i := 0; i < int(p_ref_num); i++ {
				p_ref_name_len := units.BytesToUint32(data[off : off+4])
				off = off + 4
				p_ref_name := units.Byte2string(data[off : off+p_ref_name_len])
				off = off + p_ref_name_len
				p_ref_index := units.BytesToUint32(data[off : off+4])
				off = off + 4
				refmodel := libs.ReferenceModel{
					RefName:  p_ref_name,
					RefIndex: p_ref_index,
				}
				referenceMap[p_ref_index] = refmodel
			}

			if p_ref_num > 0 {
				ref.ReferenceNumber = p_ref_num
				ref.Reference = referenceMap
			}
		}

		var item ItemData = ItemData{
			Index:      p_valueIndex,
			TimeStamp:  p_timeStamp,
			TTLType:    p_ttlType,
			TTL:        p_ttl,
			Permission: p_permission,
			Type:       p_type_content,
			//Data:       p_data_content,
			References: ref,
		}

		switch item.Type {
		case "HS_ADMIN":
			admin := libs.HSAdmin{}
			admin.AdminPermission = units.BytesToUint16(p_data_content[:2])
			model := libs.ReferenceModel{}
			refNameLen := units.BytesToUint32(p_data_content[2:6])
			if refNameLen > 0 {
				model.RefName = units.Byte2string(p_data_content[6 : 6+refNameLen])
				model.RefIndex = units.BytesToUint32(p_data_content[6+refNameLen : 6+refNameLen+4])
			}
			admin.AdminRef = model
			item.Data = admin

		case "HS_SITE":
			hssite := libs.HSSite{}
			hssite.ToObj(p_data_content)
			item.Data = hssite
		case "HS_SITE.PREFIX":
			println("")
		case "HS_VLIST":
			vlist := &libs.HSVlist{}
			vlist.ValueRefNumber = units.BytesToUint32(p_data_content[0:4])
			if vlist.ValueRefNumber > 0 {
				referenceMap := make(map[uint32]libs.ReferenceModel)
				admin_off := 4
				for i := 0; i < int(vlist.ValueRefNumber); i++ {
					model := libs.ReferenceModel{}
					refNameLen := units.BytesToUint32(p_data_content[4 : admin_off+4])
					admin_off = admin_off + 4
					model.RefName = units.Byte2string(p_data_content[4 : admin_off+int(refNameLen)])
					admin_off = admin_off + 4
					model.RefIndex = units.BytesToUint32(p_data_content[4 : admin_off+4])
					admin_off = admin_off + 4
					referenceMap[model.RefIndex] = model
				}
				vlist.ValueRef = referenceMap
			}
		case "HS_PUBKEY":
			pk := libs.HSPubkey{}
			pk.GetPubKey(p_data_content)
			item.Data = pk

		case "HS_CERT":
			item.Data = p_data_content
		case "HS_SIGNATURE":
			item.Data = p_data_content
		default:
			item.Data = string(p_data_content)

		}
		valueMap[ii] = item
	}

	valuelist := ValueListModel{
		ValueNumber: p_valueNumber,
		Value:       valueMap,
	}
	m.Name = p_name_content
	m.ValueList = valuelist
	//m.ToString()
}

func (m *MessageBody) ToString() string {
	d, _ := json.MarshalIndent(m, "", "\t")
	//fmt.Printf("whole MessageBody:\n%s\n", d)
	return string(d)
}

// 消息体值数据
type ValueListModel struct {
	ValueNumber uint32
	Value       []ItemData
}

// 标识值结构模型
type ItemData struct {
	Index      uint32
	TimeStamp  uint32
	TTLType    uint8
	TTL        uint32
	Permission uint8
	Type       string
	Data       interface{}
	References libs.ReferencesModel
}

//Data  HS_SITE.PREFIX
