package routes

import (
	"bentol/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// ユーザー関連エンドポイント
	r.POST("/registration", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	// r.PUT("/user/:id/update")
	r.GET("/store", controllers.GetStores)
	r.GET("/store/:id", controllers.GetStoreMenus)
	
	// 店舗関連エンドポイント
	r.POST("/store/register", controllers.RegisterStore)
	r.POST("/store/login", controllers.LoginStore)
	// r.GET("/store/:id/info")
	// r.POST("/store/:id/create")
	r.PUT("/store/:id/update", controllers.UpdateStorePolicy)
	r.POST("/store/:id/policy", controllers.SetSpecificPolicy)

	// メニュー関連エンドポイント
	// r.GET("/menue/:id")
	r.POST("/menue/create", controllers.AddMenue)
	r.PUT("/menue/:id/update", controllers.UpdateMenue)
	r.DELETE("/menue/:id/delete")

	// 予約関連エンドポイント
	r.POST("/payment", controllers.MakeReservation)
	// r.PUT("/payment/:id/update")
	// r.DELETE("/payment/:id/delete")

	// paypay api関連エンドポイント
	// r.POST("/pay")

	// 予約確認エンドポイント
	// r.GET("/store/reservation") // cron jobで数分おきに実行
	// r.GET("/store/reservation/:id")
	// r.DELETE("/store/reservation/:id/delete")
}
