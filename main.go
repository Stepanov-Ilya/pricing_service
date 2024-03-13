package main

import (
	"purple_hack_tree/db"
	"purple_hack_tree/service"
)

// Набросок деревьев и сегментации пользователей для реализации платформы ценообразования
// Деревья можно использовать as-is или переделать под себя, наработки даны просто как пример
func main() {
	//db.Start_database()
	db.CreateProcessBd()
	//db.CreateProcessBd()
	db.GetCurrentStorage()
	service.Run()

}
