package router

import (
	"context"
	"github.com/dy-dayan/community-api-info/form"
	"github.com/dy-dayan/community-api-info/idl"
	info "github.com/dy-dayan/community-api-info/idl/dayan/community/srv-info"
	"github.com/dy-dayan/community-api-info/util"
	"github.com/dy-gopkg/kit/micro"
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/runc/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

//
func convertCommunity(c *info.Community) *form.Community {
	return &form.Community{
		ID:           c.Common.Id,
		Name:         c.Common.Name,
		SerialNumber: c.Common.SerialNumber,
		Provinces:    c.Common.Province,
		City:         c.Common.City,
		Region:       c.Common.Region,
		Street:       c.Common.Street,
		OrgID:        c.Common.OrgID,
		HouseCount:   c.Common.HouseCount,
		CheckInCount: c.Common.CheckInCount,
		BuildingArea: c.Common.BuildingArea,
		GreeningArea: c.Common.GreeningArea,
		SealedState:  c.Common.SealedState,
		Loc:          c.Common.Loc,
		State:        c.Common.State,
		OperatorID:   c.Common.OperatorID,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

//
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
	idStr := ctx.Param("id")
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

//GeCommunity
func GetCommunity(ctx *gin.Context) {
	limitStr := ctx.Query("limit")
	offsetStr := ctx.Query("offset")
	if limitStr == "" ||
		offsetStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  "param not correct",
		})
		return
	}
	limit := util.Str2Int32(limitStr)
	offset := util.Str2Int32(offsetStr)
	locStr := ctx.QueryArray("loc")
	distance, flag := ctx.GetQuery("distance")
	client := info.NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	//带地理位置查询
	if flag {
		loc := []float32{}
		for _, item := range locStr {
			tmp := util.Str2Float32(item)
			loc = append(loc, tmp)
		}
		infoReq := info.GetCommunityByLocReq{
			Limit:    limit,
			Offset:   offset,
			Loc:      loc,
			Distance: util.Str2Float32(distance),
		}

		//服务异常
		infoResp, err := client.GetCommunityByLoc(context.Background(), &infoReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": base.CODE_SERVICE_EXCEPTION,
				"msg":  err.Error(),
			})
		}
		//数据异常
		if infoResp.BaseResp.Code != int32(base.CODE_OK) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": infoResp.BaseResp.Code,
				"msg":  infoResp.BaseResp.Msg,
			})
		}
		data := []form.Community{}
		for _, item := range infoResp.Communitys {
			tmpItem := item
			tmp := convertCommunity(tmpItem)
			data = append(data, *tmp)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    data,
		})
		return
	}

	//不带地理位置查询
	infoReq := info.GetCommunityReq{
		Limit:  limit,
		Offset: offset,
	}

	//服务异常
	infoResp, err := client.GetCommunity(context.Background(), &infoReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  err.Error(),
		})
	}
	//数据异常
	if infoResp.BaseResp.Code != int32(base.CODE_OK) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": infoResp.BaseResp.Code,
			"msg":  infoResp.BaseResp.Msg,
		})
	}
	data := []form.Community{}
	for _, item := range infoResp.Communitys {
		tmpItem := item
		tmp := convertCommunity(tmpItem)
		data = append(data, *tmp)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
	return
}

//GetCommunityByID 查询具体社区信息
func GetCommunityByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  "param not correct",
		})
		return
	}
	client := info.
		NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	infoReq := info.GetCommunityByIDReq{
		CommunityID: id,
	}
	infoResp, err := client.GetCommunityByID(context.Background(), &infoReq)
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

	data := convertCommunity(infoResp.Community)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
