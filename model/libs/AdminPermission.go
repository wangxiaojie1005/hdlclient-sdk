package libs

import "fmt"

type P int

const (
	AP_Add_Identifier    P = 1
	AP_Delete_Identifier   = 2
	AP_Add_NA              = 3
	AP_Delete_NA           = 4
	AP_Modify_Value        = 5
	AP_Delete_Value        = 6
	AP_Add_Value           = 7
	AP_Modify_Admin        = 8
	AP_Remove_Admin        = 9
	AP_Add_Admin           = 10
	AP_Authorized_Read     = 11
	AP_LIST_Identifier     = 12
)

func SetAdminPermission(p ...int) {
	var c = 0
	for _, o := range p {
		c = c | (1 << (o - 1))
	}
	//fmt.Printf("%b----%d\n", c, c)
	GetAdminPermissionToString(c)
}

func GetAdminPermissionToString(p int) string {
	e := fmt.Sprintf("%b", p)
	//println("AdminPermissionToString:", e)
	//println("......", e)
	//for i := 0; i < len(e); i++ {
	//	println(len(e)-i, "......", e[i:i+1])
	//}
	return e
}
