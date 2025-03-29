package domain

type StudentModel struct {
	ID               int     `json:"id" db:"id"`
	Email            string  `json:"email" db:"email"`
	Password         string  `json:"password" db:"password"`
	ProfilePic       *string `json:"profile_pic,omitempty" db:"profile_pic"`
	Name             string  `json:"name" db:"name"`
	TelegramUsername *string `json:"telegram_username,omitempty" db:"telegram_username"`
	HeadID           *int    `json:"head_id,omitempty" db:"head_id"`
	TotalDuration    int     `json:"total_duration" db:"total_duration"`
	TotalActiveDays  int     `json:"total_active_days" db:"total_active_days"`
}

type StudentRepository interface {
	CreateStudent(student *StudentModel) error
	GetStudentByEmail(email string) (*StudentModel, error)
	GetStudentByID(id int) (*StudentModel, error)
	GetStudents() ([]*StudentModel, error)
	UpdateStudent(student *StudentModel) error
	DeleteStudent(id int) error
}
