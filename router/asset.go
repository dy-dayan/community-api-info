package router

import (
	"context"
	"github.com/dy-dayan/community-api-info/form"
	"github.com/dy-dayan/community-api-info/idl"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  err.Error(),
		})
		return
	}

	if resp.BaseResp.Code != int32(base.CODE_OK) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  "internal service error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": base.CODE_OK,
		"msg":  "success",
		"data": resp.AssetID,
	})

}

//DelAsset
func DelAsset(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  err.Error(),
		})
		return
	}
	req := info.DelAssetReq{
		AssetID: id,
	}
	client := getClient()
	resp, err := client.DelAsset(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_DATA_EXCEPTION,
			"msg":  err.Error(),
		})
		return
	}

	if resp.BaseResp.Code != int32(base.CODE_OK) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_DATA_EXCEPTION,
			"msg":  "not found data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
	return

}

//GetAssetByID
func GetAssetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  err.Error(),
		})
		return
	}

	req := info.GetAssetByIDReq{
		AssetID: id,
	}
	client := getClient()
	resp, err := client.GetAssetByID(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  err.Error(),
		})
		return
	}
	if resp.BaseResp.Code != int32(base.CODE_OK) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  "service exception",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": base.CODE_OK,
		"msg":  "success",
		"data": *ConvertAsset(resp.Asset),
	})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_INVALID_PARAMETER,
			"msg":  "param not correct",
		})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  err.Error(),
		})
		return
	}

	if resp.BaseResp.Code != int32(base.CODE_OK) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": base.CODE_SERVICE_EXCEPTION,
			"msg":  "internal service exception",
		})
		return
	}

	data := []form.Asset{}
	for _, item := range resp.Assets {
		tmpItem := *item
		data = append(data, *ConvertAsset(&tmpItem))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
