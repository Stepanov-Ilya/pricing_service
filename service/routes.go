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

	price, category, location, matrixId, UserSegmentId := db.GetPrice(request)

	response := structures.Response{
		Price:         price,
		LocationId:    category,
		MicroCategory: location,
		MatrixId:      matrixId,
		UserSegmentId: UserSegmentId}

	return c.JSON(http.StatusOK, response)
}

func GetData(c echo.Context) error {

	category := db.ReturnJSONForSelector_Category()
	location := db.ReturnJSONForSelector_Location()

	selectors := structures.Selectors{
		Category: category,
		Location: location,
	}
	return c.JSON(http.StatusOK, selectors)
}

func AddBaseline(c echo.Context) error {
	var matrix structures.Matrix
	if err := c.Bind(&matrix); err != nil {
		return c.String(http.StatusOK, "Invalid data")
	}

	db.AddProcessBaseline(matrix.MicroCategoryId, matrix.LocationId, matrix.Price)

	return c.String(http.StatusOK, "Success add to baseline")
}

func AddDiscounts(c echo.Context) error {
	var discounts structures.Discounts
	if err := c.Bind(&discounts); err != nil {
		return c.String(http.StatusOK, "Invalid data")
	}

	db.AddProcessDiscounts(discounts.Segment, discounts.MicroCategoryId, discounts.LocationId, discounts.Price)

	return c.String(http.StatusOK, "Success add to discounts")
}

func UpdateStorage(c echo.Context) error {
	db.CleanMongoCollections()
	db.NewMongoBaseline()

	db.UpdateStorage()

	db.ClearBaseline()
	db.ClearDiscounts()
	db.ClearDiscountsSegments()
	//GetArrayOfBaseline()
	//GetArrayOfDiscount()
	return c.String(http.StatusOK, "Success update storage")
}
