package libs

import "github.com/wangxiaojie1005/hdlclient-sdk/units"

// Data  HS_ADMIN
type HSAdmin struct {
	AdminPermission uint16
	AdminRef        ReferenceModel `json:AdminRef`
}

// 设置AdminPermission
func (h *HSAdmin) SetAdminPermission(args ...int) {
	var n = "0000000000000000"
	n2 := []byte(n)
	for _, arg := range args {
		if arg > 0 {
			n2[arg-1] = '1'
		}
	}
	decimal, _ := units.BinaryToUint16(string(n2))
	h.AdminPermission = decimal
}

// 获取AdminPermission
func (h *HSAdmin) GetAdminPermissionToBinary() []int {
	l := []int{}
	binary, _ := units.DecimalToBinary(int(h.AdminPermission))
	for i, v := range binary {
		if string(v) == `1` {
			l = append(l, i+1)
		}
	}
	return l
}
