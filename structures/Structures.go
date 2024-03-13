package structures

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
	Segment         uint64 `json:"segment"`
	LocationId      uint64 `json:"location_id"`
	MicroCategoryId uint64 `json:"microcategory_id"`
	Price           uint64 `json:"price"`
}
