package main

import (
	"log"

	"github.com/azeek21/blog/pkg/handler"
	"github.com/azeek21/blog/pkg/repository"
	"github.com/azeek21/blog/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	HtmlFormat = "html"
	JsonFormat = "json"
	FormatKey  = "format"
	IdKey      = "id"
)

func main() {
	err := godotenv.Load()
	engine := gin.Default()
	engine.LoadHTMLGlob("views/**/*")
	db, err := repository.CreateDb()

	if err != nil {
		log.Fatal(err.Error())
	}
	repo := repository.NewRepositroy(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	handler.RegisterHandlers(engine)
	engine.Run(":8888")
}
