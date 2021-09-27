package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"mytest/model"
	"time"
)

func Struct2Json(i interface{}) string {
	bs, _ := json.Marshal(i)
	return string(bs)
}

func PackTokenModel(userName string, password string) *model.TokenModel {
	return &model.TokenModel{
		UserName:  userName,
		Password:  password,
		TimeStamp: time.Now().Unix(),
	}
}

func GenToken(model *model.TokenModel) string {
	bs, _ := json.Marshal(model)
	tokenByte := md5.Sum(bs)
	return hex.EncodeToString(tokenByte[:])
}
