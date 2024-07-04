package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowAction(ginCtx *gin.Context) {
	var tReq service.FollowRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	FollowService := ginCtx.Keys["follow"].(service.FollowServiceClient)
	TaskResp, err := FollowService.FollowAction(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
