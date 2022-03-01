package entity

type User struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-" faker:"-"`
	Name     string `gorm:"type:varchar(100)" json:"-" faker:"name"`
	Email    string `gorm:"type:varchar(100);unique;" json:"-" faker:"email"`
	Password string `gorm:"type:varchar(100)" json:"-" faker:"password"`
}
