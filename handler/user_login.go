package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mytest/constdef"
	"mytest/dao"
	"mytest/util"
	"net/http"
)

type loginReq struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type loginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func UserLogin(c *gin.Context) {
	resp := &loginResp{
		Code: 0,
	}
	defer func() {
		c.String(http.StatusOK, util.Struct2Json(resp))
	}()
	var req loginReq
	//校验输入参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Code, resp.Message = constdef.BuildRequestParamErr()
		return
	}
	users, err := dao.NewUserLoginTable().WithUserName(req.UserName).WithPassword(req.Password).All()
	if err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	if len(users) == 0 {
		resp.Code, resp.Message = constdef.BuildUserNotExist()
		return
	}
	tokenModel := util.PackTokenModel(req.UserName, req.Password)
	resp.Token = util.GenToken(tokenModel)
	err = dao.GetRedis().Set(resp.Token, util.Struct2Json(tokenModel), constdef.TokenExpireTime).Err()
	fmt.Println(err)
	if err != nil {
		resp.Code, resp.Message = constdef.BuildServerErr()
		return
	}
	return
}
