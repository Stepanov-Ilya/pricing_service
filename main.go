package main

import "purple_hack_tree/db"

// Набросок деревьев и сегментации пользователей для реализации платформы ценообразования
// Деревья можно использовать as-is или переделать под себя, наработки даны просто как пример
func main() {
	client := db.Open_bd()

	db.Start_database(client)

	//service.Run()
	//db.CreateProcessBd()

	db.Close_db(client)
}
