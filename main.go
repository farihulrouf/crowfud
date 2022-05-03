package main
import (
	"fmt"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"crowdgo/user"
)

func main() {	
    db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	fmt.Println("Database Connection successed")

	var users []user.User
	length := len(users)

	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}
}