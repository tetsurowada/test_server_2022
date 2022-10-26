package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const prettyIndent = "  "

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func GetOwnedItems(c echo.Context) error {
	param := GetOwnedItemsInput{}
	if err := c.Bind(&param); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("something went wrong:%+v", err))
	}
	resp, err := getOwnedItems(param)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("something went wrong:%+v", err))
	}
	return c.JSONPretty(http.StatusOK, resp, prettyIndent)
}
