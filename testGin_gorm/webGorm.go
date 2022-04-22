package testGin_gorm

type Web struct {
	Id       string `gorm:"column:Id;Primary_key""`
	Account  string
	Password string
}

func (Web) TableName() string {
	return "web"
}
