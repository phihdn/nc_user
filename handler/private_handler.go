package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/phihdn/nc_user/db"
	"github.com/phihdn/nc_user/models"
)

func UpdateUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.UserClaims)

	var req models.UserUpdateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	req.ID = claims.UserID

	res, err := db.UpdateUser(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}
