package libs

//Data  HS_VLIST
type HSVlist struct {
	ValueRefNumber uint32
	ValueRef       map[uint32]ReferenceModel `json:ValueRef`
}
