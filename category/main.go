package main

import (
	"./common"
	"./domain/repository"
	service2 "./domain/service"
	"./handler"
	category "./proto/category"
	"github.com/asim/go-micro"
	"github.com/asim/go-micro/plugins/registry/consul"
	"github.com/asim/go-micro/registry"
	"github.com/asim/go-micro/util/log"
	"github.com/jinzhu/gorm"
)

func main() {
	// config center
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// registry center
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{
				"127.0.0.1:8500",
			}
		})
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
	)

	//getting mysql config
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	//create db
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	//no db single table
	db.SingularTable(true)

	rp := repository.NewCategoryRepository(db)
	rp.InitTable()
	// Initialise service
	service.Init()

	// service
	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{
		CategoryDataService:categoryDataService})
	if err != nil {
		log.Error(err)
	}
	// Register Handler
	category.RegisterCategoryHandler(service.Server(), new(handler.Category))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
