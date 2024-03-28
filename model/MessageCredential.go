package model

//////////////////////////////////////////消息凭据数据模型///////////////////////////////////
//消息凭据
type MessageCredential struct {
	CredentialLength uint32
	Version          uint8
	Reserved         uint8
	Options          uint16
	Signer           SignerModel
	Type             UTF8String
	SignerInfo       SignerInfoMode
}

//消息凭据 - 签名者
type SignerModel struct {
	Name  UTF8String
	Index uint32
}

//消息凭据 - 签名信息-签名内容
type SignerDataMode struct {
	Length    uint32
	Signature string
}

//消息凭据 - 签名信息
type SignerInfoMode struct {
	length          uint32
	DigestAlgorithm UTF8String
	SignerData      SignerDataMode
}
