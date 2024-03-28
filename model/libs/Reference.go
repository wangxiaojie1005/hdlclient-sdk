package libs

// 引用模型
type ReferencesModel struct {
	ReferenceNumber uint32
	Reference       map[uint32]ReferenceModel
}

// 引用值
type ReferenceModel struct {
	RefName  string
	RefIndex uint32
}

type Email struct {
	RefName  string
	RefIndex uint32
}
