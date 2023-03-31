package galon

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

func Add(c *gin.Context) {
	var req GalonAddReq

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
	if int(roleID) != enum.SELLER {
		controller.ApiErrorResponse(c, http.StatusForbidden, "Invalid Access")
		return
	}

	galon := models.Galon{}
	galon.Brandname = req.Brandname
	galon.Stock = req.Stock
	err = galon.AddStock()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusCreated, fmt.Sprintf("Galon #%d added successfully", galon.ID),
		galon)
}
func Update(c *gin.Context) {
	var galonId GalonFindReq
	var request GalonUpdateReq

	err := c.ShouldBindUri(&galonId)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	request.ID = galonId.ID
	err = c.ShouldBindJSON(&request)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(request.ID)
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
	galon := models.Galon{}
	galon.ID = id
	galon.Stock = request.Stock
	err = galon.UpdateStock()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Galon #%d updated successfully", galon.ID), galon)
}
func Delete(c *gin.Context) {
	var req GalonFindReq

	err := c.ShouldBindUri(&req)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	galon := models.Galon{ID: id}
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
	err = galon.DeleteGalon()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, fmt.Sprintf("Galon #%d deleted successfully", galon.ID))
}
func GetAll(c *gin.Context) {
	galon := models.Gallons{}
	err := galon.GetAll()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, "Success", galon)
}
func FindById(c *gin.Context) {
	var galonId GalonFindReq

	err := c.ShouldBindUri(&galonId)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(galonId.ID)
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	galon := models.Galon{ID: id}
	err = galon.GetById()
	if err != nil {
		controller.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ApiResponse(c, http.StatusOK, "Success", galon)

}
