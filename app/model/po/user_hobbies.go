package po

type UserHobbies struct {
	BaseId
	UserId string `gorm:"column:user_id"`
	Hobby  string `gorm:"column:hobby"`
	BaseTimeColumns
}

func (UserHobbies) TableName() string {
	return "user_hobbies"
}
