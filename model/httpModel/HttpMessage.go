package httpModel

// 查询标识 条件查询
type Params struct {
	Action string   `json:"action"`
	Index  []int    `json:"index"`
	Type   []string `json:"type"`
}

// 创建标识
type Data struct {
	Format string `json:"format"`
	Value  string `json:"value"`
}

type Reference struct {
	Identifier string `json:"identifier"`
	Index      string `json:"index"`
}

type Value struct {
	Index       string      `json:"index"`
	Type        string      `json:"type"`
	Data        Data        `json:"data"`
	TTL         string      `json:"ttl"`
	TTLType     string      `json:"ttlType"`
	Timestamp   string      `json:"timestamp"`
	References  []Reference `json:"references"`
	AdminRead   string      `json:"adminRead"`
	AdminWrite  string      `json:"adminWrite"`
	PublicRead  string      `json:"publicRead"`
	PublicWrite string      `json:"publicWrite"`
}

type CreateId struct {
	Value []Value `json:"value"`
}

// 标识值添加、标识值修改
type IdData struct {
	Format string `json:"format"`
	Value  string `json:"value"`
}

type IdReference struct {
	Identifier string `json:"identifier"`
	Index      string `json:"index"`
}

type IdValue struct {
	Index       string      `json:"index"`
	Type        string      `json:"type"`
	Data        Data        `json:"data"`
	TTLType     string      `json:"ttlType"`
	TTL         string      `json:"ttl"`
	Timestamp   string      `json:"timestamp"`
	References  []Reference `json:"references"`
	AdminRead   string      `json:"adminRead"`
	AdminWrite  string      `json:"adminWrite"`
	PublicRead  string      `json:"publicRead"`
	PublicWrite string      `json:"publicWrite"`
}

type AddIdValue struct {
	Action string  `json:"action"`
	Value  []Value `json:"value"`
}

// 标识值删除
type RemoveIdValue struct {
	Action string   `json:"action"`
	Index  []string `json:"index"`
}

// 用户登录  分布式登录
type UserLoginValue struct {
	User  string
	Index int
}
