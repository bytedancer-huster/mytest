package handler

import (
	"github.com/gin-gonic/gin"
	"mytest/constdef"
	"mytest/dao"
	"mytest/util"
	"net/http"
)

type userOutResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UserOut(c *gin.Context) {
	resp := &userOutResp{
		Code: constdef.RespCodeOk,
	}
	defer func() {
		c.String(http.StatusOK, util.Struct2Json(resp))
	}()
	token := c.GetHeader("token")
	if token == "" {
		resp.Code, resp.Message = constdef.BuildNoLoginErr()
		return
	}
	if err := dao.GetRedis().Del(token).Err(); err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	return
}
