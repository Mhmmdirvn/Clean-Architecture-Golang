package class

type Class struct {
	Id        int    `gorm:"primarykey" json:"id"`
	ClassName string `gorm:"varchar(150)" json:"class_name"`
}