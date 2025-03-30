package repository

import (
	"database/sql"
	"weregna-backend/domain"
)

type StudentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepoImpl(db *sql.DB) *StudentRepoImpl {
	return &StudentRepoImpl{
		db: db,
	}
}

func (s *StudentRepoImpl) CreateStudent(student *domain.StudentModel) error {
	_, err := s.db.Exec("INSERT INTO student (email, password, name, total_duration, total_active_days) VALUES ($1, $2, $3, $4, $5)",
		student.Email, student.Password, student.Name, student.TotalDuration, student.TotalActiveDays)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentRepoImpl) GetStudentByEmail(email string) (*domain.StudentModel, error) {
	var student domain.StudentModel
	err := s.db.QueryRow("SELECT * FROM student WHERE email = $1", email).Scan(&student.ID, &student.Email, &student.Password, &student.ProfilePic, &student.Name, &student.TelegramUsername, &student.HeadID, &student.TotalDuration, &student.TotalActiveDays, &student.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepoImpl) GetStudentByID(id int) (*domain.StudentModel, error) {
	print(id)
	var student domain.StudentModel
	err := s.db.QueryRow("SELECT * FROM student WHERE id = $1", id).Scan(&student.ID, &student.Email, &student.Password, &student.ProfilePic, &student.Name, &student.TelegramUsername, &student.HeadID, &student.TotalDuration, &student.TotalActiveDays, &student.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepoImpl) GetStudents() ([]*domain.StudentModel, error) {
	var students []*domain.StudentModel
	rows, err := s.db.Query("SELECT * FROM student")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student domain.StudentModel
		err = rows.Scan(&student.ID, &student.Email, &student.Password, &student.ProfilePic, &student.Name, &student.TelegramUsername, &student.HeadID, &student.TotalDuration, &student.TotalActiveDays, &student.CreatedAt)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}
func (s *StudentRepoImpl) UpdateStudent(student *domain.StudentModel) error {
	_, err := s.db.Exec("UPDATE student SET email = $1, password = $2, profile_pic = $3, name = $4, telegram_username = $5, head_id = $6, total_duration = $7, total_active_days = $8 WHERE id = $9",
		student.Email, student.Password, student.ProfilePic, student.Name, student.TelegramUsername, student.HeadID, student.TotalDuration, student.TotalActiveDays, student.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *StudentRepoImpl) DeleteStudent(id int) error {
	_, err := s.db.Exec("DELETE FROM student WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
