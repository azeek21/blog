package repository

import (
	"github.com/azeek21/blog/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (r UserRepo) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	err := r.db.Model(&models.User{}).Preload("Roles").Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepo) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r UserRepo) GetUserById(id uint) (*models.User, error) {
	user := &models.User{}
	user.ID = id
	err := r.db.Model(user).Preload("Roles").Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r UserRepo) UpdateUser(user *models.User) (*models.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r UserRepo) DeleteUser(user *models.User) (bool, error) {
	user.DeletedAt = gorm.DeletedAt{}
	err := r.db.Delete(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r UserRepo) AddRoles(user *models.User, roles []string) (*models.User, error) {
	for _, code := range roles {
		user.Roles = append(user.Roles, models.Role{
			Code: code,
		})
	}

	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r UserRepo) RemoveRoles(user *models.User, roles []string) (*models.User, error) {
	newRoles := []models.Role{}

	for _, userRole := range user.Roles {
		for _, role := range roles {
			if role == userRole.Code {
				continue
			}
		}
		newRoles = append(newRoles, models.Role{
			Code: userRole.Code,
		})
	}

	user.Roles = newRoles
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
