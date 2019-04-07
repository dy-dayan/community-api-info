package router

import (
	"context"
	"github.com/dy-dayan/community-api-info/form"
	"github.com/dy-dayan/community-api-info/idl"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	info "github.com/dy-dayan/community-api-info/idl/dayan/community/srv-info"
)

func ConvertAsset(asset *info.Asset) *form.Asset {
	return &form.Asset{
		ID:           asset.Common.Id,
		SerialNumber: asset.Common.SerialNumber,
		Category:     asset.Common.Category,
		Loc:          asset.Common.Loc,
		State:        asset.Common.State,
		CommunityID:  asset.Common.CommunityID,
		Brand:        asset.Common.Brand,
		Desc:         asset.Common.Desc,
		CreatedAt:    asset.CreatedAt,
		UpdatedAt:    asset.UpdatedAt,
		OperatorID:   asset.Common.OperatorID,
	}
}

//AddAsset 增加一个资产信息
func AddAsset(ctx *gin.Context) {
	assert := form.Asset{}
	err := ctx.BindJSON(&assert)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  err.Error(),
		})
		return
	}
	client := getClient()
	req := info.AddAssetReq{
		Asset: &info.AssetCommon{
			Id:           0,
			SerialNumber: assert.SerialNumber,
			Category:     assert.Category,
			State:        assert.State,
			CommunityID:  assert.CommunityID,
			Loc:          assert.Loc,
			Brand:        assert.Brand,
			Desc:         assert.Desc,
			OperatorID:   assert.OperatorID,
		},
	}

	resp, err := client.AddAsset(context.Background(), &req)
	if err != nil {
		logrus.Errorf("AddAsset service error [%v]", err)
		FailedByParam(ctx)
		return
	}

	if resp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("AddAsset code error [%v]", resp.BaseResp.Code)
		FailedByInternal(ctx)
		return
	}

	Success(ctx, resp.AssetID)

}

//DelAsset
func DelAsset(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		FailedByParam(ctx)
		return
	}
	req := info.DelAssetReq{
		AssetID: id,
	}
	client := getClient()
	resp, err := client.DelAsset(context.Background(), &req)
	if err != nil {
		logrus.Errorf("DelAsset service error [%v]", err)
		FailedByInternal(ctx)
		return
	}

	if resp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("DelAsset code error [%v]", resp.BaseResp.Code)
		FailedByNotFind(ctx)
		return
	}

	Success(ctx, nil)
	return

}

//GetAssetByID
func GetAssetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		FailedByParam(ctx)
		return
	}

	req := info.GetAssetByIDReq{
		AssetID: id,
	}
	client := getClient()
	resp, err := client.GetAssetByID(context.Background(), &req)
	if err != nil {
		logrus.Errorf("GetAssetByID service error %v", err)
		FailedByInternal(ctx)
		return
	}
	if resp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("GetAssetByID code error [%v]", resp.BaseResp.Code)
		FailedByNotFind(ctx)
		return
	}

	Success(ctx, *ConvertAsset(resp.Asset))

}

//GetAsset
func GetAsset(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")
	communityIdStr := ctx.DefaultQuery("communityID", "0")

	limit, errLimit := strconv.Atoi(limitStr)
	offset, errOffset := strconv.Atoi(offsetStr)
	communityId, errCommunityId := strconv.ParseInt(communityIdStr, 10, 64)
	if errLimit != nil || errOffset != nil || errCommunityId != nil {
		FailedByParam(ctx)
		return
	}

	req := info.GetAssetReq{
		Limit:       int32(limit),
		Offset:      int32(offset),
		CommunityID: communityId,
	}

	client := getClient()

	resp, err := client.GetAsset(context.Background(), &req)
	if err != nil {
		logrus.Errorf("GetAsset server error [%v]", err)
		FailedByInternal(ctx)
		return
	}

	if resp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Errorf("GetAsset code error [%v]", resp.BaseResp.Code)
		return
	}

	data := []form.Asset{}
	for _, item := range resp.Assets {
		tmpItem := *item
		data = append(data, *ConvertAsset(&tmpItem))
	}

	Success(ctx, data)
}
