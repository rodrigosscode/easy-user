package db

type (
	User struct {
		ID    int    `gorm:"type:int;primaryKey;autoIncrement"`
		Name  string `gorm:"type:varchar(100);not null"`
		Email string `gorm:"type:varchar(100);unique;not null"`
		Age   int    `gorm:"type:int;not null"`
	}
)
