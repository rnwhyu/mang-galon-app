package order

import (
	"fmt"
	"galon-app/controller"
	"strconv"

	//"galon-app/controller/galon"
	"galon-app/enum"
	"galon-app/models"
	token "galon-app/utils"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
	//jwt "github.com/golang-jwt/jwt/request"
)

func Make(c *gin.Context) {
	var req OrderMakeRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	claims, err := token.ExtractTokenClaims(c)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(claims)

	roleID, _ := claims["role_id"].(float64)
	userID, _ := claims["user_id"].(float64)
	if int(roleID) != enum.BUYER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.UserID = int(userID)
	order.GalonID = req.GalonID
	order.TotalOrder = req.TotalOrder
	order.Status = "Confirmed"
	err = order.MakeOrder()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	galon := models.Galon{}
	err = galon.GetById()
	galon.Stock -= order.TotalOrder
	galon.UpdateStock()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, "Success", order)
}
func Update(c *gin.Context) {
	var orderId OrderFindReq
	var req OrderUpdateReq
	err := c.ShouldBindUri(&orderId)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	req.ID = orderId.ID
	err = c.ShouldBindJSON(&req)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	claims, err := token.ExtractTokenClaims(c)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(claims)

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.ID = id
	order.Status = req.Status
	err = order.UpdateStatus()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Status #%d updated successfully", order.ID), order)

}
