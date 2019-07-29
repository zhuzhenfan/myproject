package room

const (
	RoomId         = "room_id"
	RoomCode       = "room_code"
	RoomOwner      = "room_owner"
	RoomPassWord   = "room_pass_word"
	RoomPlayerNum  = "room_player_num"
	RoomGameStatus = "room_game_status"
)
const (
	WaitingStr = "waiting"
	GamingStr  = "gaming"
	OverStr    = "over"
	Waiting = "等待中"
	Gaming  = "游戏中"
	Over    = "已结束"
)
func GetStatus(status string)string{
	if status == WaitingStr{
		return Waiting
	}
	if status == GamingStr{
		return Gaming
	}
	if status == OverStr{
		return Over
	}
	return ""
}

type Room struct {
	RoomId         string `json:"room_id" xorm:"pk unique notnull"`
	RoomCode       string `json:"room_code" xorm:"notnull autoincr"`
	RoomOwner      string `json:"room_owner" xorm:"notnull"`
	RoomPassWord   string `json:"room_pass_word"`
	RoomPlayerNum  int    `json:"room_player_num" xorm:"notnull"`
	RoomGameStatus string `json:"room_game_status"`
}

type FindIn struct {
	Room
	Page    int `json:"page"`
	Perpage int `json:"perpage"`
}

type InsertIn struct {
	RoomOwner     string `json:"room_owner"`
	RoomPassWord  string `json:"room_pass_word"`
	RoomPlayerNum int    `json:"room_player_num"`
}

type UpdateGameStatusIn struct {
	RoomId string `json:"room_id"`
	RoomGameStatus string `json:"room_game_status"`
}