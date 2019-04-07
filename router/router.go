package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.Group("/community/")
	{

		r.POST("community", AddCommunity)
		r.DELETE("community/:id", DelCommunity)
		r.GET("community/:id", GetCommunityByID)
		r.GET("community", GetCommunity)

		r.POST("asset", AddAsset)
		r.DELETE("asset/:id", DelAsset)
		r.GET("asset/:id", GetAssetByID)
		r.GET("asset", GetAsset)

		r.POST("building", AddBuilding)
		r.DELETE("building/:id", DelBuilding)
		r.GET("building/:id", GetBuilding)

		r.POST("house", AddHouse)
		r.DELETE("house/:id", DelHouse)
		r.GET("house/:id", GetHouse)
	}

	return r
}
