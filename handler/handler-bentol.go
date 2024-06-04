package handler

import (
	"bentol/domain/model"
	"bentol/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUsecase *usecase.LoginUsecase
}

func NewLoginHandler(lu *usecase.LoginUsecase) *LoginHandler {
	return &LoginHandler{LoginUsecase: lu}
}

func (lh *LoginHandler) Login(c *gin.Context) {
	var loginData struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	user, err := lh.LoginUsecase.Login(loginData.Name, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

type StoreHandler struct {
	StoreUsecase *usecase.StoreUsecase
}

func NewStoreHandler(su *usecase.StoreUsecase) *StoreHandler {
	return &StoreHandler{StoreUsecase: su}
}

func (sh *StoreHandler) GetAllStores(c *gin.Context) {
	stores, err := sh.StoreUsecase.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stores)
}

type MenueHandler struct {
	MenueUsecase *usecase.MenueUsecase
}

func NewMenueHandler(mu *usecase.MenueUsecase) *MenueHandler {
	return &MenueHandler{MenueUsecase: mu}
}

func (mh *MenueHandler) GetMenuesByStoreID(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	menues, err := mh.MenueUsecase.GetMenuesByStoreID(uint(storeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menues)
}

func (mh *MenueHandler) GetMenueByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menue ID"})
		return
	}

	menue, err := mh.MenueUsecase.GetMenueByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menue)
}

type UserReservationHandler struct {
	UserReservationUsecase *usecase.UserReservationUsecase
}

func NewUserReservationHandler(uru *usecase.UserReservationUsecase) *UserReservationHandler {
	return &UserReservationHandler{UserReservationUsecase: uru}
}

func (urh *UserReservationHandler) CreateReservation(c *gin.Context) {
	var reservationData struct {
		UserID     uint   `json:"user_id" binding:"required"`
		StoreID    uint   `json:"store_id" binding:"required"`
		MenueID    uint   `json:"menue_id" binding:"required"`
		ReservTime string `json:"reserv_time" binding:"required"`
		ReservCnt  int    `json:"reserv_cnt" binding:"required"`
	}
	if err := c.ShouldBindJSON(&reservationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	reservation := &model.UserReservation{
		UserID:     reservationData.UserID,
		StoreID:    reservationData.StoreID,
		MenueID:    reservationData.MenueID,
		ReservTime: reservationData.ReservTime,
		ReservCnt:  reservationData.ReservCnt,
		IsRecipt:   false,
	}

	if err := urh.UserReservationUsecase.CreateReservation(reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation created successfully"})
}

func (urh *UserReservationHandler) DeleteReservation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	if err := urh.UserReservationUsecase.DeleteReservationByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
}
