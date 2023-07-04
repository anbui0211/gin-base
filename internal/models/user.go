package pgmodel

type User struct {
	PgModel  `gorm:",inline"`
	UserID   string `gorm:"column:user_id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email;unique"`
	Phone    string `gorm:"column:phone"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}
