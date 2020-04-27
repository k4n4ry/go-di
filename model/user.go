package model

type User struct {
	Id   int64  `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
}

func (User) TableName() string {
	return "m_user"
}
