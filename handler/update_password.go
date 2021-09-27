package handler

import (
	"github.com/gin-gonic/gin"
	"mytest/constdef"
	"mytest/dao"
	"mytest/util"
	"net/http"
)

type updateReq struct {
	UserName       string `form:"user_name" binding:"required"`
	OriginPassword string `form:"origin_password" binding:"required"`
	UpdatePassword string `form:"update_password" binding:"required"`
}

type updateResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UpdatePassword(c *gin.Context) {
	//初始化一个返回
	resp := &updateResp{
		Code: constdef.RespCodeOk,
	}
	defer func() {
		c.String(http.StatusOK, util.Struct2Json(resp))
	}()
	var req updateReq
	//提取参数，校验参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Code, resp.Message = constdef.BuildRequestParamErr()
		return
	}
	//判断用户名和密码是否正确
	users, err := dao.NewUserLoginTable().WithUserName(req.UserName).WithPassword(req.OriginPassword).All()
	if err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	if len(users) == 0 {
		resp.Code, resp.Message = constdef.BuildUserNotExist()
		return
	}
	//更新密码
	err = dao.NewUserLoginTable().WithUserName(req.UserName).Update(map[string]interface{}{"password": req.UpdatePassword})
	if err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	//将token设置为无效，如果设置失败了，怎么办？
	dao.GetRedis().Del(util.GenToken(util.PackTokenModel(req.UserName, req.OriginPassword)))
	return
}
