package router

import (
	"github.com/dy-dayan/community-api-info/idl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func FailedByParam(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": base.CODE_INVALID_PARAMETER,
		"msg":  "param not correct",
		"data": nil,
	})
}

func FailedByInternal(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": base.CODE_SERVICE_EXCEPTION,
		"msg":  "Internal server error",
		"data": nil,
	})
}

func FailedByNotFind(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": base.CODE_DATA_EXCEPTION,
		"msg":  "data not found",
		"data": nil,
	})
}
