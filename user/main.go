package main

import (
	"../user/handler"
	"./domain/repository"
	service1 "./domain/service"
	user "./proto/user"
	"fmt"
	"github.com/asim/go-micro"
	"github.com/jinzhu/gorm"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	srv.Init()
	//init db
	db, err := gorm.Open("mysql", "root:123456@micro?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//table only initial once
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	userDataService := service1.NewUserDataService(repository.NewUserRepository(db))

	// Register handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
