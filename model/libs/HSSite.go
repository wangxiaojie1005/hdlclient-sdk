package libs

import (
	"encoding/json"
	"fmt"
	"github.com/wangxiaojie1005/hdlclient-sdk/units"
)

// Data  HS_SITE
type HSSite struct {
	DataFormatVersion    uint16
	MajorProtocolVersion uint8
	MinorProtocolVersion uint8
	SerialNumber         uint16
	PrimaryMask          uint8
	HashOption           uint8
	HashFilter           string
	Attributes           []HSSiteServerAttribute
	Servers              []HSSiteServer
}

// Data  HS_SITE Server
type HSSiteServer struct {
	ServerID  uint32
	IpAddress string
	PublicKey HSPubkey
	Interface []HSSiteServerInterface
}

// Data  HS_SITE Attribute
type HSSiteServerAttribute struct {
	Name  string
	Value string
}

// Data  HS_SITE Server Interface
type HSSiteServerInterface struct {
	IType     string
	Iprotocol string
	IPort     uint32
}

func (site *HSSite) ToObj(data []byte) (datalen int) {
	offset := 0
	site.DataFormatVersion = units.BytesToUint16(data[offset : offset+2])
	offset += 2
	site.MajorProtocolVersion = units.BytesToUint8(data[offset : offset+1])
	offset += 1
	site.MinorProtocolVersion = units.BytesToUint8(data[offset : offset+1])
	offset += 1
	site.SerialNumber = units.BytesToUint16(data[offset : offset+2])
	offset += 2
	site.PrimaryMask = units.BytesToUint8(data[offset : offset+1])
	offset += 1
	site.HashOption = units.BytesToUint8(data[offset : offset+1])
	offset += 1

	hfLen := int(units.BytesToUint32(data[offset : offset+4]))
	offset += 4
	if hfLen > 0 {
		site.HashFilter = units.Byte2string(data[offset : offset+hfLen])
		offset += hfLen
	}

	attsLen := int(units.BytesToUint32(data[offset : offset+4]))
	offset += 4
	var attributes []HSSiteServerAttribute = make([]HSSiteServerAttribute, attsLen)
	if attsLen > 0 {
		for i := 0; i < attsLen; i++ {
			att := HSSiteServerAttribute{}
			attNameLen := int(units.BytesToUint32(data[offset : offset+4]))
			offset += 4
			attName := units.Byte2string(data[offset : offset+attNameLen])
			att.Name = attName
			offset += attNameLen

			attValueLen := int(units.BytesToUint32(data[offset : offset+4]))
			offset += 4
			attValue := units.Byte2string(data[offset : offset+attValueLen])
			att.Value = attValue
			offset += attValueLen
			attributes[i] = att
		}
	}
	site.Attributes = attributes

	serverNum := int(units.BytesToUint32(data[offset : offset+4]))
	offset += 4
	servers := make([]HSSiteServer, serverNum)
	if serverNum > 0 {
		for i := 0; i < serverNum; i++ {
			server := HSSiteServer{}
			server.ServerID = units.BytesToUint32(data[offset : offset+4])
			offset += 4
			server.IpAddress = units.ByteToIPDaddress(data[offset : offset+16])
			offset += 16
			pubkeyLen := int(units.BytesToUint32(data[offset : offset+4]))
			offset += 4
			pk := HSPubkey{}
			pk.GetPubKey(data[offset : offset+pubkeyLen])
			server.PublicKey = pk
			offset += pubkeyLen
			interfaceLen := int(units.BytesToUint32(data[offset : offset+4]))
			offset += 4
			interfaces := make([]HSSiteServerInterface, interfaceLen)
			for nn := 0; nn < interfaceLen; nn++ {
				interfaceItem := HSSiteServerInterface{}
				iType := units.BytesToUint8(data[offset : offset+1])
				interfaceItem.IType = units.GetSTValue(iType)
				offset += 1
				iprotocol := units.BytesToUint8(data[offset : offset+1])
				interfaceItem.Iprotocol = units.GetSPValue(iprotocol)
				offset += 1
				iPort := units.BytesToUint32(data[offset : offset+4])
				interfaceItem.IPort = iPort
				offset += 4

				interfaces[nn] = interfaceItem
			}
			server.Interface = interfaces
			servers[i] = server
		}
	}
	site.Servers = servers
	return offset
}

func (site *HSSite) ToString() {
	data, _ := json.MarshalIndent(site, "\t", "")
	fmt.Printf("Recv MessageEnvelope:%s\n", data)
}
