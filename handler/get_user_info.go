package handler

import (
	"encoding/json"
	"mytest/constdef"
	"mytest/dao"
	"mytest/model"
	"mytest/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type getUserResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	UserName string `json:"user_name"`
}

func GetUserInfo(c *gin.Context) {
	resp := &getUserResp{
		Code: constdef.RespCodeOk,
	}
	defer func() {
		c.String(http.StatusOK, util.Struct2Json(resp))
	}()
	token := c.GetHeader("token")
	if token == "" {
		resp.Code, resp.Message = constdef.BuildRequestParamErr()
		return
	}
	val, err := dao.GetRedis().Get(token).Result()
	if err == redis.Nil {
		resp.Code, resp.Message = constdef.BuildNoLoginErr()
		return
	}
	if err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	var tokenModel model.TokenModel
	if err := json.Unmarshal([]byte(val), &tokenModel); err != nil {
		resp.Code, resp.Message = constdef.BuildServerErr()
		return
	}
	resp.UserName = tokenModel.UserName
	return
}
