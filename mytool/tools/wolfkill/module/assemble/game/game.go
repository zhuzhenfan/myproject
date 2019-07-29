package game

import (
	"wolfkill/wolfkill/module/role"
	"wolfkill/wolfkill/module/room"
	"wolfkill/wolfkill/module/user"
)

const (
	UserId     = "user_id"
	RoleId     = "role_id"
	RoomId     = "room_id"
	Status     = "status"
	DeadWay    = "dead_way"
	Select     = "select"
	GameId     = "game_id"
	UserNumber = "user_number"
	Exist      = "exist"
)

const (
	FirstGetRole = "first_get_role"
	JoinGame     = "join_game"
	CloseRoom    = "close_room"
)

type Game struct {
	GameId     string `json:"game_id" xorm:"pk notnull unique"`
	RoomId     string `json:"room_id" xorm:"notnull"`
	RoleId     string `json:"role_id" xorm:"notnull"`
	UserId     string `json:"user_id"`
	UserNumber int    `json:"user_number"`
	Status     bool   `json:"status"`
	DeadWay    string `json:"dead_way"`
	Select     bool   `json:"select"`
}

type CreateGameIn struct {
	room.Room
	RoleIds []string `json:"role_ids"`
}

type CreateGameOut struct {
	RoomId   string `json:"room_id"`
	RoomCode string `json:"room_code"`
}

type JoinGameIn struct {
	UserId       string `json:"user_id"`
	RoomCode     string `json:"room_code"`
	RoomPassWord string `json:"room_pass_word"`
}

type ListPlayerOut struct {
	user.User
	role.Role
}

type GameStatus struct {
	RoomId         string `json:"room_id"`
	RoomGameStatus string `json:"room_game_status"`
}

