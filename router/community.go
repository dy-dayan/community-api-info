package router

import (
	"context"
	"github.com/dy-dayan/community-api-info/form"
	"github.com/dy-dayan/community-api-info/idl"
	info "github.com/dy-dayan/community-api-info/idl/dayan/community/srv-info"
	"github.com/dy-gopkg/kit/micro"
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/runc/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

func AddCommunity(ctx *gin.Context) {
	req := form.Community{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    base.CODE_INVALID_PARAMETER,
			"message": err.Error(),
		})
	}

	client := info.
		NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	infoReq := info.AddCommunityReq{
		Community: &info.CommunityCommon{
			Name:         req.Name,
			Province:     req.Provinces,
			City:         req.City,
			Region:       req.Region,
			Street:       req.Street,
			OrgID:        req.OrgID,
			HouseCount:   req.HouseCount,
			CheckInCount: req.CheckInCount,
			BuildingArea: req.BuildingArea,
			GreeningArea: req.GreeningArea,
			Loc:          req.Loc,
			State:        req.State,
			OperatorID:   req.OperatorID,
		},
	}

	infoResp, err := client.AddCommunity(context.Background(), &infoReq)
	if err != nil || infoResp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("add community code [%d] error:[%v], msg [%s]",
			infoResp.BaseResp.Code, err, infoResp.BaseResp.Msg)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func DelCommunity(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  "param not correct",
		})
		return
	}
	client := info.
		NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	infoReq := info.DelCommunityReq{
		CommunityID: id,
	}
	resp, err := client.DelCommunity(context.Background(), &infoReq)
	//todo:分开判断错误原因
	if err != nil ||
		resp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("Del community failed, code[%d],  msg[%s],error[%v]",
			resp.BaseResp.Code, resp.BaseResp.Msg, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func GetCommunity(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  "param not correct",
		})
		return
	}
	client := info.
		NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	infoReq := info.GetCommunityReq{
		CommunityID: id,
	}
	infoResp, err := client.GetCommunity(context.Background(), &infoReq)
	if err != nil {
		logrus.Errorf("Get community error [%v]", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  err.Error(),
		})
	}

	if infoResp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Warnf("Get community error code [%v]")
		ctx.JSON(http.StatusOK, gin.H{
			"code": base.CODE_DATA_EXCEPTION,
			"msg":  "not found",
		})
	}

	data := form.Community{
		Name:         infoResp.Community.Name,
		SerialNumber: infoResp.Community.SerialNumber,
		Provinces:    infoResp.Community.Province,
		City:         infoResp.Community.City,
		Region:       infoResp.Community.Region,
		Street:       infoResp.Community.Street,
		OrgID:        infoResp.Community.OrgID,
		HouseCount:   infoResp.Community.HouseCount,
		CheckInCount: infoResp.Community.CheckInCount,
		BuildingArea: infoResp.Community.BuildingArea,
		GreeningArea: infoResp.Community.GreeningArea,
		SealedState:  infoResp.Community.SealedState,
		Loc:          infoResp.Community.Loc,
		State:        infoResp.Community.State,
		OperatorID:   infoResp.Community.OperatorID,
		CreatedAt:    infoResp.CreatedAt,
		UpdatedAt:    infoResp.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}
