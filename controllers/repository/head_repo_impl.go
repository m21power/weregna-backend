package repository

import (
	"database/sql"
	"weregna-backend/domain"
)

type HeadRepoImpl struct {
	db *sql.DB
}

func NewHeadRepoImpl(db *sql.DB) *HeadRepoImpl {
	return &HeadRepoImpl{
		db: db,
	}
}
func (r *HeadRepoImpl) CreateHead(head *domain.Head) error {
	query := "INSERT INTO head (email, password, name, role) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, head.Email, head.Password, head.Name, head.Role)
	if err != nil {
		return err
	}
	return nil
}
func (r *HeadRepoImpl) FindHeadByEmail(email string) (*domain.Head, error) {
	var head domain.Head
	err := r.db.QueryRow("SELECT * FROM head WHERE email = $1", email).Scan(&head.ID, &head.Email, &head.Password, &head.ProfilePic, &head.Name, &head.Role, &head.CreatedAt, &head.TelegramUsername)
	if err != nil {
		return nil, err
	}
	return &head, nil
}
func (r *HeadRepoImpl) FindHeadByID(id int) (*domain.Head, error) {
	var head domain.Head
	err := r.db.QueryRow("SELECT * FROM head WHERE id = $1", id).Scan(&head.ID, &head.Email, &head.Password, &head.ProfilePic, &head.Name, &head.Role, &head.CreatedAt, &head.TelegramUsername)
	if err != nil {
		return nil, err
	}
	return &head, nil
}
func (r *HeadRepoImpl) UpdateHead(head *domain.Head) error {
	query := "UPDATE head SET email = $1, password = $2,  name = $3, role = $4, profile_pic = $5,telegram_username = $6 WHERE id = $7"
	_, err := r.db.Exec(query, head.Email, head.Password, head.Name, head.Role, head.ProfilePic, head.TelegramUsername, head.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *HeadRepoImpl) DeleteHead(id int) error {
	query := "DELETE FROM head WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *HeadRepoImpl) GetHeads() ([]*domain.Head, error) {
	rows, err := r.db.Query("SELECT * FROM head")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var heads []*domain.Head
	for rows.Next() {
		var head domain.Head
		err := rows.Scan(&head.ID, &head.Email, &head.Password, &head.ProfilePic, &head.Name, &head.Role, &head.CreatedAt, &head.TelegramUsername)
		if err != nil {
			return nil, err
		}
		heads = append(heads, &head)
	}
	return heads, nil
}

func (s *HeadRepoImpl) AddMyStudent(studentID int, headID int) error {
	_, err := s.db.Exec("UPDATE student SET head_id = $1 WHERE id = $2", headID, studentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *HeadRepoImpl) GetMyStudents(headID int) ([]*domain.StudentModel, error) {
	var students []*domain.StudentModel
	rows, err := s.db.Query("SELECT * FROM student WHERE head_id = $1", headID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
