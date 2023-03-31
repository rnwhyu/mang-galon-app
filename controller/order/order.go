package order

import (
	"fmt"
	"galon-app/controller"
	"galon-app/enum"
	"galon-app/models"
	token "galon-app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	galon := models.Galon{ID: order.GalonID}
	err = galon.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if galon.Stock <= 0 {
		controller.ApiErrorResponse(c, http.StatusBadRequest, "There arn't any stock available")
		return
	}

	err = order.MakeOrder()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	galon.Stock -= order.TotalOrder
	err = galon.UpdateStock()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ApiResponse(c, http.StatusOK, "Success", order)
}
func UpdateProcessing(c *gin.Context) {
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

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.ID = id

	err = order.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if order.Status != "Confirmed" {
		controller.ApiErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid Update Status: Must be 'Confirmed' but got '%s' instead", order.Status))
		return
	}

	order.Status = "Processed"
	err = order.UpdateStatus()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Status #%d updated successfully", order.ID), order)

}
func UpdateDelivery(c *gin.Context) {
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

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.ID = id

	err = order.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if order.Status != "Processed" {
		controller.ApiErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid Update Status: Must be 'Processed' but got '%s' instead", order.Status))
		return
	}

	order.Status = "On Delivery"
	err = order.UpdateStatus()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Status #%d updated successfully", order.ID), order)

}
func UpdateDelivered(c *gin.Context) {
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

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.ID = id

	err = order.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if order.Status != "On Delivery" {
		controller.ApiErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid Update Status: Must be 'On Delivery' but got '%s' instead", order.Status))
		return
	}

	order.Status = "Delivered"
	err = order.UpdateStatus()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Status #%d updated successfully", order.ID), order)
}
func UpdateCompleted(c *gin.Context) {
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

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.ID = id

	err = order.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if order.Status != "Delivered" {
		controller.ApiErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid Update Status: Must be 'Delivered' but got '%s' instead", order.Status))
		return
	}

	order.Status = "Completed"
	err = order.UpdateStatus()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Status #%d updated successfully", order.ID), order)
}
func UpdateCanceled(c *gin.Context) {
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

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.BUYER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}
	order := models.Order{}
	order.ID = id

	err = order.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if order.Status != "Confirmed" {
		controller.ApiErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid Update Status: Must be 'Confirmed' but got '%s' instead", order.Status))
		return
	}

	order.Status = "Canceled"
	err = order.UpdateStatus()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	galon := models.Galon{ID: order.GalonID}
	err = galon.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	galon.Stock += order.TotalOrder
	err = galon.UpdateStock()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Status #%d updated successfully", order.ID), order)
}

func GetAll(c *gin.Context) {
	claims, err := token.ExtractTokenClaims(c)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	roleID, _ := claims["role_id"].(float64)
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}

	orders := models.Orders{}
	err = orders.GetAll()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, "Success", orders)
}

func GetByUserId(c *gin.Context) {
	claims, err := token.ExtractTokenClaims(c)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, _ := claims["user_id"].(float64)

	orders := models.Orders{}
	err = orders.GetByUserId(int(userID))
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, "Success", orders)
}
