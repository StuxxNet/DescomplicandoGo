package main

import (
	"fmt"

	"github.com/StuxxNet/DescomplicandoGo/CRUD/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	var usersList []users.User

	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate DB
	db.AutoMigrate(&users.User{})

	id := users.Create(db, "Joaquim")
	usersList = users.List(db)
	for _, user := range usersList {
		fmt.Println(user.Name)
	}

	users.Update(db, id, "Ramon")

	usersList = users.List(db)
	for _, user := range usersList {
		fmt.Println(user.Name)
	}

	users.Delete(db, id)
	users.List(db)
}
