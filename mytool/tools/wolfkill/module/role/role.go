package role

const (
	RoleId          = "role_id"
	RoleName        = "role_name"
	RoleSkill       = "role_skill"
	RoleExplain     = "role_explain"
	RoleType        = "role_type"
	RoleOrderNumber = "role_order_number"
	RoleOpenNumber  = "role_open_number"
)

type Role struct {
	RoleId          string `json:"role_id" xorm:"pk unique notnull"`
	RoleName        string `json:"role_name" xorm:"unique notnull"`
	RoleChinaName   string `json:"role_china_name" xorm:"unique notnull"`
	RoleSkill       string `json:"role_skill" xorm:"notnull"`
	RoleExplain     string `json:"role_explain" xorm:"notnull"`
	RoleType        string `json:"role_type" xorm:"notnull"`
	RoleOrderNumber int    `json:"role_order_number" xorm:"unique notnull"`
	RoleOpenNumber  int    `json:"role_open_number" xorm:"unique notnull"`
}

type RoleJson struct {
	RoleChineseNames map[string]string
	RoleInfos        []*RoleInfo        `json:"roleInfos"`
}

type RoleInfo struct {
	Name        string `json:"name"`
	ChinaName   string `json:"chinaName"`
	Skill       string `json:"skill"`
	Types       string `json:"types"`
	OpenNumber  int    `json:"openNumber"`
	OrderNumber int    `json:"orderNumber"`
	Explain     string `json:"explain"`
}

type FindTypeIn struct {
	RoleType string `json:"role_type"`
	Page int `json:"page"`
	Perpage int `json:"perpage"`
}

type ListIn struct {
	Page int `json:"page"`
	Perpage int `json:"perpage"`
	RoleType string `json:"role_type"`
}