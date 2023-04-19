package students

import "Clean-Architecture/modules/class"

type Gender string

const (
	Male   Gender = "Laki-Laki"
	Female Gender = "Perempuan"
)

type Student struct {
	Id       int    `gorm:"primarykey" json:"id"`
	Name     string `gorm:"varchar(200)" json:"student_name"`
	Gender   Gender `gorm:"type:enum('Laki-Laki', 'Perempuan')" json:"gender"`
	Class_Id int    `gorm:"foreignKey:Class_Id" json:"class_id"`
	Class    class.Class
}