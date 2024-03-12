package db

type CurrentStorage struct {
	Baseline  uint64
	Discounts map[uint64]uint64
}
