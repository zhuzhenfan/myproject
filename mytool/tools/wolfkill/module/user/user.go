// user
package user

const (
	UserId     = "user_id"
	UserName   = "user_name"
	UserOnline = "user_online"
	PassWord   = "pass_word"
	CreateTime = "create_time"
	LoginTime  = "login_time"
)

type User struct {
	UserId     string `json:"user_id" xorm:"pk unique notnull"`
	UserName   string `json:"user_name" xorm:"unique notnull"`
	PassWord   string `json:"pass_word" xorm:"notnull"`
	UserOnline bool   `json:"user_online"`
	CreateTime int64  `json:"create_time" xorm:"notnull"`
	LoginTime  int64  `json:"login_time"`
}

type FindIn struct {
	Page     int    `json:"page"`
	Perpage  int    `json:"perpage"`
	UserName string `json:"user_name"`
}

type LoginIn struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type LoginOut struct {
	*User
	Token string `json:"token"`
}
