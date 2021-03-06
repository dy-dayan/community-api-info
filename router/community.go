package router

import (
	"context"
	_ "github.com/dy-dayan/community-api-info/docs"
	"github.com/dy-dayan/community-api-info/form"
	"github.com/dy-dayan/community-api-info/idl"
	info "github.com/dy-dayan/community-api-info/idl/dayan/community/srv-info"
	"github.com/dy-dayan/community-api-info/util"
	"github.com/dy-gopkg/kit/micro"
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/runc/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"strconv"
)

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

//AddCommunity 增加一个社区信息
//
// @title addCommunity
// @version 1.0
// @description 添加社区信息.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func AddCommunity(ctx *gin.Context) {
	req := form.Community{}
	err := ctx.BindJSON(&req)
	if err != nil {
		FailedByParam(ctx)
		return
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
		FailedByInternal(ctx)
		return
	}

	Success(ctx, infoResp.CommunityID)
}

func DelCommunity(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		FailedByParam(ctx)
		return
	}
	client := info.
		NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	infoReq := info.DelCommunityReq{
		CommunityID: id,
	}
	resp, err := client.DelCommunity(context.Background(), &infoReq)
	//todo:分开判断错误原因
	if err != nil {
		logrus.Errorf("DelCommunity failed server error [%v]", err)
		FailedByInternal(ctx)
		return
	}
	if resp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("DelCommunity failed, code[%d]",
			resp.BaseResp.Code)
		FailedByNotFind(ctx)
		return
	}

	Success(ctx, nil)
}

//GeCommunity
func GetCommunity(ctx *gin.Context) {
	limitStr := ctx.Query("limit")
	offsetStr := ctx.Query("offset")
	if limitStr == "" ||
		offsetStr == "" {
		FailedByParam(ctx)
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
			logrus.Errorf("GetCommunityByLoc failed server error [%v]", err)
			FailedByInternal(ctx)
			return
		}
		//数据异常
		if infoResp.BaseResp.Code != int32(base.CODE_OK) {
			logrus.Errorf("Get community Failed code error [%v]", infoResp.BaseResp.Code)
			FailedByInternal(ctx)
			return
		}
		data := []form.Community{}
		for _, item := range infoResp.Communitys {
			tmpItem := item
			tmp := convertCommunity(tmpItem)
			data = append(data, *tmp)
		}

		Success(ctx, data)
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
		logrus.Error("GetCommunity failed server error [%v]", err)
		FailedByInternal(ctx)
		return
	}
	//数据异常
	if infoResp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("GetCommunity failed code error [%v]", infoResp.BaseResp.Code)
		FailedByInternal(ctx)
		return
	}
	data := []form.Community{}
	for _, item := range infoResp.Communitys {
		tmpItem := item
		tmp := convertCommunity(tmpItem)
		data = append(data, *tmp)
	}

	Success(ctx, data)
	return
}

//GetCommunityByID 查询具体社区信息
func GetCommunityByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		FailedByParam(ctx)
		return
	}
	client := info.
		NewCommunityInfoService("dayan.community.srv.info", micro.Client())
	infoReq := info.GetCommunityByIDReq{
		CommunityID: id,
	}
	infoResp, err := client.GetCommunityByID(context.Background(), &infoReq)
	if err != nil {
		logrus.Errorf("GetCommunityByID  server error [%v]", err)
		FailedByInternal(ctx)
		return
	}

	if infoResp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Warnf("GetCommunityByID code error [%v]")
		FailedByNotFind(ctx)
		return
	}

	data := convertCommunity(infoResp.Community)

	Success(ctx, data)
	return
}
