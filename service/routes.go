package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"purple_hack_tree/db"
	"purple_hack_tree/structures"
)

func GetPrice(c echo.Context) error {
	var request structures.Request
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, "Invalid data")
	}

	// TODO : GetPrice(request)

	response := structures.Response{}
	// GetPrice(&response)

	return c.JSON(http.StatusOK, response)
}

func GetData(c echo.Context) error {
	// Todo Get category from database
	category := db.GetData()

	return c.JSON(http.StatusOK, category)
}

func AddBaseline(c echo.Context) error {
	var matrix structures.Matrix
	if err := c.Bind(&matrix); err != nil {
		return c.String(http.StatusOK, "Invalid data")
	}

	//Todo add to local database

	return c.String(http.StatusOK, "Success add category")
}

func AddDiscounts(c echo.Context) error {
	var discounts structures.Discounts
	if err := c.Bind(&discounts); err != nil {
		return c.String(http.StatusOK, "Invalid data")
	}

	//Todo add to local database

	return c.String(http.StatusOK, "Success add location")
}

func UpdateStorage(c echo.Context) error {
	// Todo relocate data from temporarily database to public matrix

	// Drop temporarily databases

	return c.String(http.StatusOK, "Success update storage")
}
