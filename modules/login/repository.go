package login

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) Login(users User) error {
	var result User
	res := repo.DB.Where("username = ?", users.Username,).Where("password = ?", users.Password).First(&result)
	return res.Error
}