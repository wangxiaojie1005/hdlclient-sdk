package units

import (
	"fmt"
)

// OPCODE
const (
	OC_RESOLUTION            uint32 = 1
	OC_GET_SITE_INFO         uint32 = 2
	OC_CREATE                uint32 = 100
	OC_DELETE                uint32 = 101
	OC_ADD_VALUE             uint32 = 102
	OC_REMOVE_VALUE          uint32 = 103
	OC_MODIFY_VALUE          uint32 = 104
	OC_LIST_IDENTIFIERS      uint32 = 105
	OC_RESPONSE_TO_CHALLENGE uint32 = 200
	OC_INCR_UPDATE           uint32 = 1001
	OC_FULL_UPDATE           uint32 = 1002
	OC_ACTIVE_UPDATE_NOTIFY  uint32 = 1004
	OC_LOGIN                 uint32 = 2000
	OC_META_CREATE           uint32 = 606
	OC_META_REMOVE           uint32 = 607
	OC_META_MODIFY           uint32 = 608
	OC_META_QUERYT           uint32 = 609
)

const (
	RC_RESERVED                 uint32 = 0
	RC_SUCCESS                  uint32 = 1
	RC_ERROR                    uint32 = 2
	RC_SERVER_BUSY              uint32 = 3
	RC_PROTOCOL_ERROR           uint32 = 4
	RC_OPERATION_DENIED         uint32 = 5
	RC_RECUR_LIMIT_EXCEEDED     uint32 = 6
	RC_IDENTIFIER_NOT_FOUND     uint32 = 100
	RC_IDENTIFIER_ALREADY_EXIST uint32 = 101
	RC_INVALID_IDENTIFIER       uint32 = 102
	RC_VALUE_NOT_FOUND          uint32 = 200
	RC_VALUE_ALREADY_EXIST      uint32 = 201
	RC_VALUE_INVALID            uint32 = 202
	RC_EXPIRED_SITE_INFO        uint32 = 300
	RC_SERVER_NOT_RESP          uint32 = 301
	RC_SERVICE_REFERRAL         uint32 = 302
	RC_NA_DELEGATE              uint32 = 303
	RC_NOT_AUTHORIZED           uint32 = 400
	RC_ACCESS_DENIED            uint32 = 401
	RC_AUTHEN_NEEDED            uint32 = 402
	RC_AUTHEN_FAILED            uint32 = 403
	RC_INVALID_CREDENTIAL       uint32 = 404
	RC_AUTHEN_TIMEOUT           uint32 = 405
	RC_UNABLE_TO_AUTHEN         uint32 = 406
	RC_SESSION_TIMEOUT          uint32 = 500
	RC_SESSION_FAILED           uint32 = 501
	RC_NO_SESSION_KEY           uint32 = 502
	RC_SESSION_NO_SUPPORT       uint32 = 503
	RC_SESSION_KEY_INVALID      uint32 = 504
	RC_SESSION_MESSAGE_REJECTED uint32 = 505
)

// MessageFlag
const (
	CP = 1
	EC = 2
	TC = 3
)

// SERVICE TYPE
const (
	ST_OUT_OF_SERVICE  uint8 = 0
	ST_ADMIN           uint8 = 1
	ST_QUERY           uint8 = 2
	ST_ADMIN_AND_QUERY uint8 = 3

	SP_HDL_UDP   uint8 = 0
	SP_HDL_TCP   uint8 = 1
	SP_HDL_HTTP  uint8 = 2
	SP_HDL_HTTPS uint8 = 3
)

func GetSTValue(v uint8) string {
	switch v {
	case ST_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	case ST_ADMIN:
		return "ADMIN"
	case ST_QUERY:
		return "QUERY"
	case ST_ADMIN_AND_QUERY:
		return "ADMIN_AND_QUERY"
	default:
		return "UNKNOWN"
	}
}

func GetSPValue(v uint8) string {
	switch v {
	case SP_HDL_UDP:
		return "UDP"
	case SP_HDL_TCP:
		return "TCP"
	case SP_HDL_HTTP:
		return "HTTP"
	case SP_HDL_HTTPS:
		return "HTTPS"
	default:
		return "UNKNOWN"
	}
}

func SetMValue(p []int) uint16 {
	c := uint16(0)
	for _, o := range p {
		c = c | (1 << (16 - o))
	}
	return c
}

func getMFValue(value uint16) []string {
	v := fmt.Sprintf("%016b", value)

	vl := len(v)
	of := make([]string, vl)
	for i := 0; i < vl; i++ {
		of[i] = string(v[i])
	}
	return of
}

// OPFLAG
const (
	AT  = 1  //权威位：请求中设置此位为1表示请求应该直接发到主服务站点（而不是镜像站点）。应答消息设置此位为1，表示消息来自主服务器
	CT  = 2  //身份证明位：CT被置位,表示要求服务器对其应答进行签名。CT被置位的应答表示消息已被签名。如果请求时CT被置位，服务器必须对应答数据签名。如果服务器不能在应答中提供有效的签名，客户端应该丢弃应答并视此请求失败
	ENC = 3  //加密位： 表示需要服务器使用会话密钥对其应答加密。
	REC = 4  //递归。 请求中REC置位，要求服务器以客户端身份转发（如果需要）此请求到其它服务器。
	CA  = 5  //cache 验证位 请求中CA置位，要求缓存服务器以客户端的身份对于所有的服务器应答进行验证（比如验证服务器签名）。应答中CA置位，表示应答数据已经缓存服务器验证通过
	CN  = 6  //连续位 消息中CN置位，告诉消息接收者属于此请求（或应答）的更多的消息随后就到。这发生在请求（或应答）数据太大，一条消息装不下，从而必须分片成多条消息的情况下
	KC  = 7  //保持连接位 消息中KC置位，要求消息接收者保持TCP连接打开（在应答数据发回之后）。这使得同一TCP连接可用于多个标识操作
	PO  = 8  //公共查询位 仅用于询问操作。询问请求中PO置位，表示客户端只是想询问有PUB_READ许可的标识值。请求PO位设置为0，就会询问所有的标识值，而不管它们的读权限。如果标识值需要ADMIN_READ权限，服务器必须验证客户端是管理员
	RD  = 9  //请求摘要位 要求服务器在其应答中包含消息摘要（message digest）。应答消息中RD置位，表示消息体（Message Body）的第一个字段包含了原始请求的消息摘要。消息摘要可用于检查服务器应答数据的完整性
	DT  = 31 //可信解析验证位 0，表示不需要对解析结果做认证 1，表示需要对解析结果做认证。
	TR  = 32 //可信解析验证结果位 0，可信认证失败 1，可信认证成功
)

func SetOFValue(p []int) uint32 {
	c := uint32(0)
	for _, o := range p {
		c = c | (1 << (32 - o))
	}
	return c
}

func getOFValue(value uint32) []string {
	v := fmt.Sprintf("%032b", value)

	vl := len(v)
	of := make([]string, vl)
	for i := 0; i < vl; i++ {
		of[i] = string(v[i])
	}
	return of
}
