package pgmodel

//type User struct {
//	PgModel `gorm:",inline"`
//	Status  string `gorm:"column:status;not null"`
//	Name    string `gorm:"column:name"`
//	Email   string `gorm:"column:email;unique"`
//}

type User struct {
	PgModel  `gorm:",inline"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email;unique"`
	Phone    string `gorm:"column:phone"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}
