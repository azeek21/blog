package repository

import (
	"fmt"

	"github.com/azeek21/blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (bool, error)
	AddRoles(user *models.User, roles []string) (*models.User, error)
	RemoveRoles(user *models.User, roles []string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type RoleRepository interface {
	UpdateRole(role *models.Role) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
	DeleteRole(role *models.Role) (bool, error)
	CreateRole(role *models.Role) (*models.Role, error)
}

type ArticleRepository interface {
}

type Repository struct {
	UserRepository
	RoleRepository
	ArticleRepository
}

func NewRepositroy(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		RoleRepository: NewRolesRepository(db),
	}
}

type PostgresConnectionConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PWD"`
	Dbname   string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

func connectDb(config PostgresConnectionConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(`host=%v user=%v password=%v dbname=%v port=%v sslmode=disable`,
		config.Host, config.User, config.Password, config.Dbname, config.Port)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

func CreateDb(config PostgresConnectionConfig) (*gorm.DB, error) {
	//	config := PostgresConnectionConfig{
	//		Host:     viper.GetString("DB_HOST"),
	//		User:     viper.GetString("db_user"),
	//		Password: viper.GetString("DB_PWD"),
	//		Dbname:   viper.GetString("db_name"),
	//		Port:     viper.GetString("db_port"),
	//	}
	fmt.Printf("PG CONFIG: %+v", config)

	return connectDb(config)
}
