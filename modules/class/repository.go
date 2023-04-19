package class

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllClass() ([]Class, error) {
	var class []Class
	result := repo.DB.Find(&class)
	return class, result.Error
}

func (repo Repository) GetClassById(id string) (*Class, error) {
	var class *Class
	result := repo.DB.First(&class, id)
	return class, result.Error
}

func (repo Repository) CreateClass(class Class) error {
	result := repo.DB.Create(&class)
	return result.Error
}

func (repo Repository) UpdateClassById(id string, class Class) error {
	result := repo.DB.Model(&Class{}).Where("id = ?", id).Updates(class)
	return result.Error
}

func (repo Repository) DeleteClassById(id string) error {
	result := repo.DB.Delete(&Class{}, id)
	return result.Error
}