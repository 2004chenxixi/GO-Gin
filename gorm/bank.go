package gorm

type Bank struct {
	Id       int
	Username string
	Money    int
}

func (Bank) TableName() string {
	return "bank"
}
