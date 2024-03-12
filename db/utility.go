package db

import (
	"purple_hack_tree/service"
	"sort"
)

func GetPrice(request service.Request) service.Response {
	discountIds := GetSegmentsByUserID(request.UserId)
	sort.Slice(discountIds, func(i, j int) bool { return discountIds[i] > discountIds[j] })
	var response
	if discountIds != nil {
		for _, discountId := range discountIds {
			// Todo search in storage of discount
			// FindInDiscount(&response, discountId, request.LocationId, request.MicroCategoryId)
			// if response != nil {
			//	return response
			//}

		}
	}

	// Todo search in storage of baseline
	//FindInBaseline(&response, request.LocationId, request.MicroCategoryId)

	return response
}
