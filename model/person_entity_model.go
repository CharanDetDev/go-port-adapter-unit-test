package model

type Person struct {
	PersonID  int    `gorm:"column:PersonID"`
	LastName  string `gorm:"column:LastName"`
	FirstName string `gorm:"column:FirstName"`
	Address   string `gorm:"column:Address"`
	City      string `gorm:"column:City"`
}

func (m *Person) TableName() string {
	return "Persons"
}
