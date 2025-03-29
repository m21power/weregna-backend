package usecases

import "weregna-backend/domain"

type StudentUsecases struct {
	studentRepo domain.StudentRepository
}

func NewStudentUsecases(studentRepo domain.StudentRepository) *StudentUsecases {
	return &StudentUsecases{
		studentRepo: studentRepo,
	}
}

func (s *StudentUsecases) CreateStudent(student *domain.StudentModel) error {
	return s.studentRepo.CreateStudent(student)
}
func (s *StudentUsecases) GetStudentByEmail(email string) (*domain.StudentModel, error) {
	return s.studentRepo.GetStudentByEmail(email)
}
func (s *StudentUsecases) GetStudentByID(id int) (*domain.StudentModel, error) {
	return s.studentRepo.GetStudentByID(id)
}
func (s *StudentUsecases) GetStudents() ([]*domain.StudentModel, error) {
	return s.studentRepo.GetStudents()
}
func (s *StudentUsecases) UpdateStudent(student *domain.StudentModel) error {
	return s.studentRepo.UpdateStudent(student)
}
func (s *StudentUsecases) DeleteStudent(id int) error {
	return s.studentRepo.DeleteStudent(id)
}
