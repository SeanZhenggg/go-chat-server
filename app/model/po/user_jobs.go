package po

type UserJob struct {
	BaseId
	UserId string `gorm:"column:user_id"`
	Job    string `gorm:"column:job"`
	BaseTimeColumns
}

func (UserJob) TableName() string {
	return "user_jobs"
}
