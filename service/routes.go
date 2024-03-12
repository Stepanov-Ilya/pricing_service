package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
