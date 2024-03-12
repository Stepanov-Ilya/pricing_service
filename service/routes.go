package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"purple_hack_tree/db"
)

type Request struct {
	LocationId      uint64 `json:"location_id"`
	MicroCategoryId uint64 `json:"microcategory_id"`
	UserId          uint64 `json:"user_id"`
}

type Response struct {
	Price         uint64 `json:"price"`
	LocationId    uint64 `json:"location_id"`
	MicroCategory uint64 `json:"microcategory_id"`
	MatrixId      uint64 `json:"matrix_id"`
	UserSegmentId uint64 `json:"user_segment_id"`
}

type Storage struct {
	Baseline  []Line            `json:"baseline"`
	Discounts map[uint64][]Line `json:"discounts"`
}

type Line struct {
	LocationId      uint64 `json:"location_id"`
	MicroCategoryId uint64 `json:"microcategory_id"`
	Price           uint64 `json:"price"`
}

func GetPrice(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, "Invalid data")
	}

	// TODO : GetPrice(request)

	response := Response{}
	// GetPrice(&response)

	return c.JSON(http.StatusOK, response)
}

func GetCategory(c echo.Context) error {
	category := c.QueryParam("category")
	return c.Render(http.StatusOK, "template_category.html", category)
}

func GetLocation(c echo.Context) error {
	location := c.QueryParam("location")
	return c.Render(http.StatusOK, "template_location.html", location)
}

func GetStorage(c echo.Context) error {
	var storage Storage
	if err := c.Bind(&storage); err != nil {
		return c.String(http.StatusOK, "Invalid data")
	}

	db.AddStorage(storage)

	return c.String(http.StatusOK, "Success add storage")
}
