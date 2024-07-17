package routes

import (
	"bentol/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// ユーザー関連エンドポイント
	r.POST("/registration", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// session check
	user := r.Group("/").Use(controllers.UserAuthRequired())
	{
		// user.PUT("/user/:id/update")
		user.GET("/store", controllers.GetStores)
		user.GET("/store/:id", controllers.GetStoreMenus)

		// 予約関連エンドポイント
		user.POST("/payment", controllers.MakeReservation)
		// user.POST("/payment/:id/cancel", controllers.CancelReservation)

		// 予約確認エンドポイント
		user.GET("/user/:id/reservation", controllers.GetUserReservation)
	}
	r.POST("/logout", controllers.LogoutUser)

	// 店舗関連エンドポイント
	r.POST("/store/register", controllers.RegisterStore)
	r.POST("/store/login", controllers.LoginStore)

	// session check
	store := r.Group("/").Use(controllers.StoreAuthRequired())
	{
		// store.GET("/store/:id/info")
		store.PUT("/store/:id/update", controllers.UpdateStorePolicy)
		store.POST("/store/:id/policy", controllers.SetSpecificPolicy)

		// メニュー関連エンドポイント
		// store.GET("/menue/:id")
		store.POST("/menue/create", controllers.AddMenue)
		store.PUT("/menue/:id/update", controllers.UpdateMenue)
		// store.DELETE("/menue/:id/delete")

		// 予約確認エンドポイント
		store.GET("/store/reservation", controllers.CheckStoreReservation)
		// store.GET("/store/reservation/:id")
		store.DELETE("/store/reservation/delete", controllers.ShipReservation)
	}
	r.POST("/store/logout", controllers.LogoutStore)

}
