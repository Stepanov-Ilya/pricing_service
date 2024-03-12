package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetCategoriesTree(collection mongo.Collection) *CategoryNode {
	// Создаем корневую категорию - ROOT
	rootNode := NewCategory("ROOT")

	for category, subCategories := range RawCategories {
		categoryNode := NewCategory(category)

		for _, subCategory := range subCategories {
			subCategoryNode := NewCategory(subCategory)
			categoryNode.AddChild(subCategoryNode, collection)
		}

		rootNode.AddChild(categoryNode, collection)
	}

	insertResult, err := collection.InsertOne(context.TODO(), rootNode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

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
func (l *CategoryNode) AddChild(child *CategoryNode, collection mongo.Collection) {
	l.Children = append(l.Children, child)

	insertResult, err := collection.InsertOne(context.TODO(), child)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
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

var RawCategories = map[string][]string{
	"Бытовая электроника":           {"Товары для компьютера", "Фототехника", "Телефоны", "Планшеты и электронные книги", "Оргтехника и расходники", "Ноутбуки", "Настольные компьютеры", "Игры, приставки и программы", "Аудио и видео"},
	"Готовый бизнес и оборудование": {"Готовый бизнес", "Оборудование для бизнеса"},
	"Для дома и дачи":               {"Мебель и интерьер", "Ремонт и строительство", "Продукты питания", "Растения", "Бытовая техника", "Посуда и товары для кухни"},
	"Животные":                      {"Другие животные", "Товары для животных", "Птицы", "Аквариум", "Кошки", "Собаки"},
	"Личные вещи":                   {"Детская одежда и обувь", "Одежда, обувь, аксессуары", "Товары для детей и игрушки", "Часы и украшения", "Красота и здоровье"},
	"Недвижимость":                  {"Недвижимость за рубежом", "Квартиры", "Коммерческая недвижимость", "Гаражи и машиноместа", "Земельные участки", "Дома, дачи, коттеджи", "Комнаты"},
	"Работа":                        {"Резюме", "Вакансии"},
	"Транспорт":                     {"Автомобили", "Запчасти и аксессуары", "Грузовики и спецтехника", "Водный транспорт", "Мотоциклы и мототехника"},
	"Услуги":                        {"Предложения услуг"},
	"Хобби и отдых":                 {"Охота и рыбалка", "Спорт и отдых", "Коллекционирование", "Книги и журналы", "Велосипеды", "Музыкальные инструменты", "Билеты и путешествия"},
}
