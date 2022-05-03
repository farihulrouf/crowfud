package main
import (
	"crowdgo/user"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	/*	
    db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	fmt.Println("Database Connection successed")

	var users []user.User


	db.Find(&users)


	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}
	*/
	/*
	router := gin.Default()
	router.GET("/handler", Handler)
	router.Run(":8111")
	*/

	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	user := user.User{
		Name: "Test Simpan",
	}

	userRepo.save(user)

	/*
	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "Test Simpan",
	}

	userRepository.save(user)
	*/
}
/*
func Handler(c *gin.Context) {
    db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
    if err != nil {
    	log.Fatal("Db connection error")
    }

    //fm.Println("Database Connection successed")

    var users []user.User
    db.Find(&users)
    c.JSON(http.StatusOK, users)

    // input
    // handler
    // service
    // repository
    // db


}
*/

