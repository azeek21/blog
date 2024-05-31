package main

import (
	"log"
	"os"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/spf13/viper"
)

func must(x error) {
	if x != nil {
		log.Fatalf(x.Error())
		os.Exit(1)
	}
}

func main() {
	err := utils.InitConfig("dev")
	must(err)

	dbConf := repository.PostgresConnectionConfig{}
	dbConf, err = utils.LoadConfig(dbConf)
	must(err)
	db, err := repository.CreateDb(dbConf)
	must(err)

	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Article{})
	must(err)

	repo := repository.NewRepositroy(db)
	roles := viper.GetStringSlice("ROLES")
	for _, role := range roles {
		repo.CreateRole(&models.Role{
			Code: role,
		})
	}

	superUser := &models.User{
		Email:    viper.GetString("SU_EMAIL"),
		FullName: viper.GetString("SU_FULL_NAME"),
		Username: viper.GetString("SU_USERNAME"),
		Password: viper.GetString("SU_PWD"),
		Roles:    []models.Role{{Code: viper.GetString("SU_ROLE")}},
	}

	_, err = repo.CreateUser(superUser)
	if err != nil {
		repo.UpdateUser(superUser)
	}

	log.Printf("Migration SUCCESS\nCreated super user:\nName: %v, Email: %v, Roles: %v\n", superUser.FullName, superUser.Email, superUser.Roles)
}
