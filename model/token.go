package model

type TokenModel struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	TimeStamp int64  `json:"time_stamp"`
}
