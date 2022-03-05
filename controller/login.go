package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tafaquh/crud-example/services"
	"github.com/tafaquh/crud-example/utils/errors"

	loginDomain "github.com/tafaquh/crud-example/domain/login"
	"github.com/tafaquh/crud-example/domain/response"
)

func Login(c *gin.Context) {
	var requestBody loginDomain.LoginRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errLogin := services.LoginService.Login(&requestBody)
	if errLogin != nil {
		c.JSON(errLogin.Status, &errors.RestErr{
			Message: errLogin.Message,
			Status:  errLogin.Status,
			Error:   errLogin.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully login"))
}
