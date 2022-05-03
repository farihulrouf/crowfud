package main
import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"crowdgo/user"
	"fmt"
)

func main() {	
    db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

    if err != nil {
    	log.Fatal(err.Error())
    }
    fmt.Println("Connection Success")

    userRepository := user.NewRepository(db)
    user := user.User{
    	Name: "Test Save",
    }
    userRepository.Save(user)
}