package main
import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
    "github.com/gin-gonic/gin"
    "crowdgo/handler"
	"crowdgo/user"
    "crowdgo/todo"
	"fmt"
)

func main() {	
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

    if err != nil {
    	log.Fatal(err.Error())
    }
    fmt.Println("Connection Success")
    //db.AutoMigrate(&user.User{}) //create Migrate
    //db.AutoMigrate(&todo.Todo{}) //create Migrate
    

    userRepository := user.NewRepository(db)
    userService := user.NewService(userRepository)

    userHandler := handler.NewUserHandler(userService)


    todoRepository := todo.NewRepository(db)
    todoService := todo.NewService(todoRepository)
    todoHandler  := handler.NewTodoHandler(todoService)
    
    router := gin.Default()
    
    api := router.Group("api/v1")
    api.POST("/users", userHandler.RegisterUser)
    api.POST("/sessions", userHandler.Login)
    api.POST("/todo", todoHandler.CreateTodo)
    
    router.Run(":8111")

    /*
    userInput := user.RegisterUserInput{}
    userInput.Name = "Test simpan dari service"
    userInput.Email = "contoh@gmail.com"
    userInput.Occupation = "Anak aband"
    userInput.Password = "password"

    userService.RegisterUser(userInput)
    */
}