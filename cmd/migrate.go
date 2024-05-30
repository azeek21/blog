package main

import (
	"log"
	"os"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
	"github.com/joho/godotenv"
)

func must(x error) {
	if x != nil {
		log.Fatalf(x.Error())
		os.Exit(1)
	}
}

var superUserrole = "super"

func main() {
	err := godotenv.Load()
	must(err)

	db, err := repository.CreateDb()
	must(err)

	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Article{})
	must(err)

	repo := repository.NewRepositroy(db)
	adminRole, err := repo.RoleRepository.CreateRole(&models.Role{
		Code:        superUserrole,
		Name:        "Admin",
		Description: "Administrator over everything",
	})
	must(err)
	user, err := repo.CreateUser(&models.User{
		FullName: "malton Aka Azeek",
		Username: "azeek",
		Email:    "pymanuz@gmail.com",
		Roles: []models.Role{
			*adminRole,
		},
	})
	must(err)
	log.Printf("Migration SUCCESS\nCreated super user:\nName: %v, Email: %v, Roles: %v\n", user.FullName, user.Email, user.Roles)

}
