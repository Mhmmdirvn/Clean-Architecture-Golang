package students

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllStudents() ([]Student, error) {
	var students []Student
	result := repo.DB.Preload("Class").Find(&students)
	return students, result.Error
}

func (repo Repository) GetStudentById(id string) (Student, error) {
	var student Student
	result := repo.DB.Preload("Class").First(&student, id)
	return student, result.Error
}

func (repo Repository) CreateStudent(student Student) error {
	result := repo.DB.Create(&student)
	return result.Error
}

func (repo Repository) UpdateStudentById(id string, student Student) error {
	result := repo.DB.Model(&Student{}).Where("id = ?", id).Updates(student)
	return result.Error
} 

func (repo Repository) DeleteStudentById(id string) error {
	result := repo.DB.Delete(&Student{}, id)
	return result.Error
}