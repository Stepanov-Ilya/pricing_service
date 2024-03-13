package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
)

func Update_pids_in_mong(cat_col mongo.Collection, loc_col monogo.Collection, matrix [][3]int64) {
	// Заполнить корневые каталоги
	var category_id, location_id int64

	sort.Slice(matrix, func(i, j int) bool {
		return matrix[0][i] < matrix[0][j]
	})

	// Заполняем категории
	for _, line := range matrix {
		category_id = line[0]
		node := Find_node_in_mongo(category_id, cat_col)
		BFS(*node, cat_col)
	}

	// Заполняем локации
	sort.Slice(matrix, func(i, j int) bool {
		return matrix[1][i] < matrix[1][j]
	})
	for _, line := range matrix {
		location_id = line[1]
		node := Find_node_in_mongo(location_id, loc_col)
		BFS(*node, loc_col)
	}
}

type Queue []Node

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(node Node) {
	*q = append(*q, node)
}

func (q *Queue) Dequeue() (Node, bool) {
	if q.IsEmpty() {
		return *NewCategory(""), false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

// Пустая структура
type void struct{}

var vp void

func BFS(node Node, collection mongo.Collection) {
	visited := make(map[int64]bool)

	queue := Queue{}

	queue.Enqueue(node)

	for !queue.IsEmpty() {
		cur, _ := queue.Dequeue()
		// Добавляем всех потомков курсора
		for _, curChild := range cur.Children {
			if !visited[curChild.ID] {
				// Если у какого-то потомка уже проставлен PID, то
				// он уже нас не интересует, потомков текущего потомка
				// мы тоже пропускаем, поэтому пропускаем эту ноду
				if curChild != nil && curChild.PID == 0 {
					queue.Enqueue(*curChild)
					visited[curChild.ID] = true
				}
			}
		}

		cur.PID = node.ID
		change_pid(collection, cur.ID, node.PID)
	}

}

func change_pid(collection mongo.Collection, id int64, pid int64) {
	update := bson.D{{"$set", bson.D{{"pid", pid}}}}
	filter := bson.D{{"id", id}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func GetData(category_id int64, location_id int64, discount_segments []int) int64 {

}
