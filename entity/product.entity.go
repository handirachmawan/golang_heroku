package entity

type Product struct {
	ID     int64  `gorm:"primary_key:auto_increment" json:"-" faker:"-"`
	Name   string `gorm:"type:varchar(100)" json:"-" faker:"word"`
	Price  uint64 `gorm:"type:bigint" json:"-" faker:"-"`
	UserID int64  `gorm:"not null" json:"-" faker:"-"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-" faker:"-"`
}
