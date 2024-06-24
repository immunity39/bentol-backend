package routes

import (
	"bentol/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// ユーザー関連エンドポイント
	r.POST("/registration", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// 店舗関連エンドポイント
	r.GET("/store", controllers.GetStores)
	r.GET("/store/:id", controllers.GetStoreMenus)
	r.POST("/store/register", controllers.RegisterStore)
	r.POST("/store/login", controllers.LoginStore)
	r.PUT("/store/:id/update", controllers.UpdateStorePolicy)
	r.POST("/store/:id/policy", controllers.SetSpecificPolicy)

	// メニュー関連エンドポイント
	r.POST("/menue/add", controllers.AddMenue)
	r.PUT("/menue/:id/update", controllers.UpdateMenue)

	// 予約関連エンドポイント
	r.POST("/payment", controllers.MakeReservation)
}
