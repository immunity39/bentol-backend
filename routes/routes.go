package routes

import (
	"bentol/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// session, cookie setting
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// ユーザー関連エンドポイント
	r.POST("/registration", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	// r.GET("/logout", controllers.LogoutUser)
	// r.PUT("/user/:id/update")
	r.GET("/store", controllers.GetStores)
	r.GET("/store/:id", controllers.GetStoreMenus)

	// 店舗関連エンドポイント
	r.POST("/store/register", controllers.RegisterStore)
	r.POST("/store/login", controllers.LoginStore)
	// r.GET("/store/logout", controllers.LogoutStore)
	// r.GET("/store/:id/info")
	r.PUT("/store/:id/update", controllers.UpdateStorePolicy)
	r.POST("/store/:id/policy", controllers.SetSpecificPolicy)

	// メニュー関連エンドポイント
	// r.GET("/menue/:id")
	r.POST("/menue/create", controllers.AddMenue)
	r.PUT("/menue/:id/update", controllers.UpdateMenue)
	// r.DELETE("/menue/:id/delete")

	// 予約関連エンドポイント
	r.POST("/payment", controllers.MakeReservation)
	// r.PUT("/payment/:id/update")
	// r.DELETE("/payment/:id/delete")

	// 予約確認エンドポイント
	r.GET("/store/reservation", controllers.CheckStoreReservation)
	// r.GET("/store/reservation/:id")
	// r.DELETE("/store/reservation/:id/delete")
}
