package models

type User struct {
	ID        int64  `gorm:"id" json:"ID"`
	Login     string `gorm:"login" json:"login"`
	Phone     string `gorm:"phone" json:"phone"`
	Password  string `gorm:"password" json:"password"`
	FirstName string `gorm:"first_name" json:"firstName"`
	LastName  string `gorm:"last_name" json:"lastName"`
	Type      string `gorm:"type" json:"type"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) FillDefaultValue() {
	u.ID = 0
}

func (u *User) ConvertToDto() (result map[string]interface{}) {
	return
}
