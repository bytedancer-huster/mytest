package handler

import (
	"mytest/constdef"
	"mytest/dao"
	"mytest/model"
	"mytest/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createReq struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type createResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CreateAccount(c *gin.Context) {
	//初始化一个返回值
	resp := &createResp{
		Code: constdef.RespCodeOk,
	}
	defer func() {
		c.String(http.StatusOK, util.Struct2Json(resp))
	}()
	var req createReq
	//校验输入参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Code, resp.Message = constdef.BuildRequestParamErr()
		return
	}
	//判断用户是否已经存在
	users, err := dao.NewUserLoginTable().WithUserName(req.UserName).All()
	if err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	if len(users) > 0 {
		resp.Code, resp.Message = constdef.BuildUserRepeatErr()
		return
	}

	//创建账户
	err = dao.NewUserLoginTable().Create(&model.UserLogin{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		resp.Code, resp.Message = constdef.BuildDBErr()
		return
	}
	return
}
