package model

import (
	"bytes"
	"encoding/binary"
)

// 解码 TCP 字节流为 MsgCreatMeta 对象
func DecodeMsgCreatMetaFromBytes(data []byte) MsgCreatMeta {
	var msg MsgCreatMeta

	byteReader := bytes.NewReader(data)

	// 解码 MetaID
	var metaIDLength uint32
	// 读取 MetaID.Length
	binary.Read(byteReader, binary.BigEndian, &metaIDLength)
	// 读取 MetaID.Content
	metaIDContent := make([]byte, metaIDLength)
	byteReader.Read(metaIDContent)
	// 设置 MetaID 对象的属性
	msg.MetaID.Length = metaIDLength
	msg.MetaID.Content = string(metaIDContent)

	// 解码 MetaName
	var metaNameLength uint32
	// 读取 MetaName.Length
	binary.Read(byteReader, binary.BigEndian, &metaNameLength)
	// 读取 MetaName.Content
	metaNameContent := make([]byte, metaNameLength)
	byteReader.Read(metaNameContent)
	// 设置 MetaName 对象的属性
	msg.MetaName.Length = metaNameLength
	msg.MetaName.Content = string(metaNameContent)

	// 解码 MetaVersion
	binary.Read(byteReader, binary.BigEndian, &msg.MetaVersion)

	// 解码 IDCodeRule
	var IDCodeRuleLength uint32
	// 读取 IDCodeRule.Length
	binary.Read(byteReader, binary.BigEndian, &IDCodeRuleLength)
	// 读取 IDCodeRule.Content
	IDCodeRuleContent := make([]byte, IDCodeRuleLength)
	byteReader.Read(IDCodeRuleContent)
	// 设置 IDCodeRule 对象的属性
	msg.IDCodeRule.Length = IDCodeRuleLength
	msg.IDCodeRule.Content = string(IDCodeRuleContent)

	// 解码 MetaDesc
	var metaDescLength uint32
	// 读取 IDCodeRule.Length
	binary.Read(byteReader, binary.BigEndian, &metaDescLength)
	// 读取 IDCodeRule.Content
	metaDescContent := make([]byte, metaDescLength)
	byteReader.Read(metaDescContent)
	// 设置 IDCodeRule 对象的属性
	msg.MetaDesc.Length = metaDescLength
	msg.MetaDesc.Content = string(metaDescContent)

	// 解码 AttrList
	binary.Read(byteReader, binary.BigEndian, &msg.AttrList.AttrNumber)
	msg.AttrList.AttrValue = make([]AttrValueModule, msg.AttrList.AttrNumber)
	for i := uint32(0); i < msg.AttrList.AttrNumber; i++ {
		// 解码name
		var nameLength uint32
		binary.Read(byteReader, binary.BigEndian, &nameLength)
		nameContent := make([]byte, nameLength)
		byteReader.Read(nameContent)
		msg.AttrList.AttrValue[i].Name.Length = nameLength
		msg.AttrList.AttrValue[i].Name.Content = string(nameContent)

		// 解码Alias
		var aliasLength uint32
		binary.Read(byteReader, binary.BigEndian, &aliasLength)
		aliasContent := make([]byte, aliasLength)
		byteReader.Read(aliasContent)
		msg.AttrList.AttrValue[i].Alias.Length = aliasLength
		msg.AttrList.AttrValue[i].Alias.Content = string(aliasContent)

		// 解码type
		var typeLength uint32
		binary.Read(byteReader, binary.BigEndian, &typeLength)
		typeContent := make([]byte, typeLength)
		byteReader.Read(typeContent)
		msg.AttrList.AttrValue[i].Type.Length = typeLength
		msg.AttrList.AttrValue[i].Type.Content = string(typeContent)

		binary.Read(byteReader, binary.BigEndian, &msg.AttrList.AttrValue[i].ReadLevel)
		binary.Read(byteReader, binary.BigEndian, &msg.AttrList.AttrValue[i].WriteLevel)

		// 解码RWLevelDesv
		var RWLevelDesvLength uint32
		binary.Read(byteReader, binary.BigEndian, &RWLevelDesvLength)
		RWLevelDesvContent := make([]byte, RWLevelDesvLength)
		byteReader.Read(RWLevelDesvContent)
		msg.AttrList.AttrValue[i].RWLevelDesv.Length = RWLevelDesvLength
		msg.AttrList.AttrValue[i].RWLevelDesv.Content = string(RWLevelDesvContent)
	}

	// 解码 MetaRefList
	binary.Read(byteReader, binary.BigEndian, &msg.MetaRefList.MetaRefNum)

	msg.MetaRefList.DataList = make([]DataListModule, msg.MetaRefList.MetaRefNum)
	for i := uint8(0); i < msg.MetaRefList.MetaRefNum; i++ {
		// 解码MetaID
		var MetaIDLength uint32
		binary.Read(byteReader, binary.BigEndian, &MetaIDLength)
		MetaIDContent := make([]byte, MetaIDLength)
		byteReader.Read(MetaIDContent)
		msg.MetaRefList.DataList[i].MetaID.Length = MetaIDLength
		msg.MetaRefList.DataList[i].MetaID.Content = string(MetaIDContent)
		binary.Read(byteReader, binary.BigEndian, &msg.MetaRefList.DataList[i].Version)
	}

	return msg
}

// 解码查询元数据
func DecodeFindMetaResFromBytes(data []byte) FindMetaRes {
	var metaRes FindMetaRes
	byteReader := bytes.NewReader(data)
	binary.Read(byteReader, binary.BigEndian, &metaRes.DataLen)
	metaRes.DataList = make([]MsgCreatMeta, metaRes.DataLen)
	for i := uint8(0); i < metaRes.DataLen; i++ {
		// 解码 MetaID
		var metaIDLength uint32
		// 读取 MetaID.Length
		binary.Read(byteReader, binary.BigEndian, &metaIDLength)
		// 读取 MetaID.Content
		metaIDContent := make([]byte, metaIDLength)
		byteReader.Read(metaIDContent)
		// 设置 MetaID 对象的属性
		metaRes.DataList[i].MetaID.Length = metaIDLength
		metaRes.DataList[i].MetaID.Content = string(metaIDContent)

		// 解码 MetaName
		var metaNameLength uint32
		// 读取 MetaName.Length
		binary.Read(byteReader, binary.BigEndian, &metaNameLength)
		// 读取 MetaName.Content
		metaNameContent := make([]byte, metaNameLength)
		byteReader.Read(metaNameContent)
		// 设置 MetaName 对象的属性
		metaRes.DataList[i].MetaName.Length = metaNameLength
		metaRes.DataList[i].MetaName.Content = string(metaNameContent)

		// 解码 MetaVersion
		binary.Read(byteReader, binary.BigEndian, &metaRes.DataList[i].MetaVersion)

		// 解码 IDCodeRule
		var IDCodeRuleLength uint32
		// 读取 IDCodeRule.Length
		binary.Read(byteReader, binary.BigEndian, &IDCodeRuleLength)
		// 读取 IDCodeRule.Content
		IDCodeRuleContent := make([]byte, IDCodeRuleLength)
		byteReader.Read(IDCodeRuleContent)
		// 设置 IDCodeRule 对象的属性
		metaRes.DataList[i].IDCodeRule.Length = IDCodeRuleLength
		metaRes.DataList[i].IDCodeRule.Content = string(IDCodeRuleContent)

		// 解码 MetaDesc
		var metaDescLength uint32
		// 读取 IDCodeRule.Length
		binary.Read(byteReader, binary.BigEndian, &metaDescLength)
		// 读取 IDCodeRule.Content
		metaDescContent := make([]byte, metaDescLength)
		byteReader.Read(metaDescContent)
		// 设置 IDCodeRule 对象的属性
		metaRes.DataList[i].MetaDesc.Length = metaDescLength
		metaRes.DataList[i].MetaDesc.Content = string(metaDescContent)

		// 解码 AttrList
		binary.Read(byteReader, binary.BigEndian, &metaRes.DataList[i].AttrList.AttrNumber)
		metaRes.DataList[i].AttrList.AttrValue = make([]AttrValueModule, metaRes.DataList[i].AttrList.AttrNumber)
		for j := uint32(0); j < metaRes.DataList[i].AttrList.AttrNumber; j++ {
			// 解码name
			var nameLength uint32
			binary.Read(byteReader, binary.BigEndian, &nameLength)
			nameContent := make([]byte, nameLength)
			byteReader.Read(nameContent)
			metaRes.DataList[i].AttrList.AttrValue[j].Name.Length = nameLength
			metaRes.DataList[i].AttrList.AttrValue[j].Name.Content = string(nameContent)

			// 解码Alias
			var aliasLength uint32
			binary.Read(byteReader, binary.BigEndian, &aliasLength)
			aliasContent := make([]byte, aliasLength)
			byteReader.Read(aliasContent)
			metaRes.DataList[i].AttrList.AttrValue[j].Alias.Length = aliasLength
			metaRes.DataList[i].AttrList.AttrValue[j].Alias.Content = string(aliasContent)

			// 解码type
			var typeLength uint32
			binary.Read(byteReader, binary.BigEndian, &typeLength)
			typeContent := make([]byte, typeLength)
			byteReader.Read(typeContent)
			metaRes.DataList[i].AttrList.AttrValue[j].Type.Length = typeLength
			metaRes.DataList[i].AttrList.AttrValue[j].Type.Content = string(typeContent)

			binary.Read(byteReader, binary.BigEndian, &metaRes.DataList[i].AttrList.AttrValue[j].ReadLevel)
			binary.Read(byteReader, binary.BigEndian, &metaRes.DataList[i].AttrList.AttrValue[j].WriteLevel)

			// 解码RWLevelDesv
			var RWLevelDesvLength uint32
			binary.Read(byteReader, binary.BigEndian, &RWLevelDesvLength)
			RWLevelDesvContent := make([]byte, RWLevelDesvLength)
			byteReader.Read(RWLevelDesvContent)
			metaRes.DataList[i].AttrList.AttrValue[j].RWLevelDesv.Length = RWLevelDesvLength
			metaRes.DataList[i].AttrList.AttrValue[j].RWLevelDesv.Content = string(RWLevelDesvContent)
		}

		// 解码 MetaRefList
		binary.Read(byteReader, binary.BigEndian, &metaRes.DataList[i].MetaRefList.MetaRefNum)

		metaRes.DataList[i].MetaRefList.DataList = make([]DataListModule, metaRes.DataList[i].MetaRefList.MetaRefNum)
		for k := uint8(0); k < metaRes.DataList[i].MetaRefList.MetaRefNum; k++ {
			// 解码MetaID
			var MetaIDLength uint32
			binary.Read(byteReader, binary.BigEndian, &MetaIDLength)
			MetaIDContent := make([]byte, MetaIDLength)
			byteReader.Read(MetaIDContent)
			metaRes.DataList[i].MetaRefList.DataList[k].MetaID.Length = MetaIDLength
			metaRes.DataList[i].MetaRefList.DataList[k].MetaID.Content = string(MetaIDContent)
			binary.Read(byteReader, binary.BigEndian, &metaRes.DataList[i].MetaRefList.DataList[k].Version)
		}
	}
	return metaRes
}
