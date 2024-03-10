package main

import (
	"fmt"
)

func GetLocationsTree() *LocationNode {
	// Создаем корневую локацию - Все регионы
	allRegions := NewLocation("Все регионы")

	for region, cities := range rawLocations {
		regionNode := NewLocation(region)

		for _, city := range cities {
			cityNode := NewLocation(city)
			regionNode.AddChild(cityNode)
		}
		allRegions.AddChild(regionNode)
	}

	return allRegions
}

var locationID int64

// LocationNode представляет собой узел дерева локаций
type LocationNode struct {
	ID       int64
	Name     string
	Children []*LocationNode
}

// NewLocation Создает новый узел локации
func NewLocation(name string) *LocationNode {
	locationID++
	return &LocationNode{
		ID:       locationID,
		Name:     name,
		Children: []*LocationNode{},
	}
}

// AddChild Добавляет дочернюю локацию к родительской локации
func (l *LocationNode) AddChild(child *LocationNode) {
	l.Children = append(l.Children, child)
}

// PrintTree Рекурсивно выводит дерево локаций
func (l *LocationNode) PrintTree(indent int) {
	fmt.Printf("%s%d - %s\n", generateLocationIndent(indent), l.ID, l.Name)
	for _, child := range l.Children {
		child.PrintTree(indent + 2)
	}
}

// Генерирует отступ для вывода
func generateLocationIndent(indent int) string {
	result := ""
	for i := 0; i < indent; i++ {
		result += " "
	}
	return result
}
