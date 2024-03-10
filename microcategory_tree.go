package main

import (
	"fmt"
	"purple_hack_tree/database"
)

func GetCategoriesTree() *CategoryNode {
	// Создаем корневую категорию - ROOT
	rootNode := NewCategory("ROOT")

	for category, subCategories := range database.GetRawCategories() {
		categoryNode := NewCategory(category)

		for _, subCategory := range subCategories {
			subCategoryNode := NewCategory(subCategory)
			categoryNode.AddChild(subCategoryNode)
		}

		rootNode.AddChild(categoryNode)
	}

	return rootNode
}

var categoryID int64

// CategoryNode представляет собой узел дерева локаций
type CategoryNode struct {
	ID       int64
	Name     string
	Children []*CategoryNode
}

// NewCategory Создает новый узел локации
func NewCategory(name string) *CategoryNode {
	categoryID++
	return &CategoryNode{
		ID:       categoryID,
		Name:     name,
		Children: []*CategoryNode{},
	}
}

// AddChild Добавляет дочернюю локацию к родительской категории
func (l *CategoryNode) AddChild(child *CategoryNode) {
	l.Children = append(l.Children, child)
}

// PrintTree Рекурсивно выводит дерево категорий
func (l *CategoryNode) PrintTree(indent int) {
	fmt.Printf("%s%d - %s\n", generateCategoryIndent(indent), l.ID, l.Name)
	for _, child := range l.Children {
		child.PrintTree(indent + 2)
	}
}

// Генерирует отступ для вывода
func generateCategoryIndent(indent int) string {
	result := ""
	for i := 0; i < indent; i++ {
		result += " "
	}
	return result
}
