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

type Matrix struct {
	LocationId      uint64 `json:"location_id"`
	MicroCategoryId uint64 `json:"microcategory_id"`
	Price           uint64 `json:"price"`
}

type Discounts struct {
	Segment string `json:"segment"`
	Matrix  Matrix `json:"matrix"`
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

func GetData(c echo.Context) error {
	// Todo Get category from database
	category := db.GetData()

	return c.JSON(http.StatusOK, category)
}

func AddBaseline(c echo.Context) error {
	var matrix Matrix
	if err := c.Bind(&matrix); err != nil {
		return c.String(http.StatusOK, "Invalid data")
	}

	//Todo add to local database

	return c.String(http.StatusOK, "Success add category")
}

func AddDiscounts(c echo.Context) error {
	var discounts Discounts
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
