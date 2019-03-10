package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Group("/community/")
	{

		r.POST("community", AddCommunity)
		r.DELETE("community/:id", DelCommunity)
		r.GET("community/:id", GetCommunity)

		r.POST("asset", AddAsset)
		r.DELETE("asset/:id", DelAsset)
		r.GET("asset/:id", GetAsset)

		r.POST("building", AddBuilding)
		r.DELETE("building/:id", DelBuilding)
		r.GET("building/:id", GetBuilding)

		r.POST("house", AddHouse)
		r.DELETE("house/:id", DelHouse)
		r.GET("house/:id", GetHouse)
	}

	return r
}
