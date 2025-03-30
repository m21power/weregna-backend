package usecases

import "weregna-backend/domain"

type HeadUsecases struct {
	headRepo domain.HeadRepository
}

func NewHeadUsecases(headRepo domain.HeadRepository) *HeadUsecases {
	return &HeadUsecases{
		headRepo: headRepo,
	}
}

func (h *HeadUsecases) CreateHead(head *domain.Head) error {
	return h.headRepo.CreateHead(head)
}
func (h *HeadUsecases) FindHeadByEmail(email string) (*domain.Head, error) {
	return h.headRepo.FindHeadByEmail(email)
}
func (h *HeadUsecases) FindHeadByID(id int) (*domain.Head, error) {
	return h.headRepo.FindHeadByID(id)
}
func (h *HeadUsecases) UpdateHead(head *domain.Head) error {
	return h.headRepo.UpdateHead(head)
}
func (h *HeadUsecases) DeleteHead(id int) error {
	return h.headRepo.DeleteHead(id)
}
func (h *HeadUsecases) GetHeads() ([]*domain.Head, error) {
	return h.headRepo.GetHeads()
}
func (h *HeadUsecases) AddMyStudent(headID int, studentID int) error {
	return h.headRepo.AddMyStudent(headID, studentID)
}
func (h *HeadUsecases) GetMyStudents(headID int) ([]*domain.StudentModel, error) {
	return h.headRepo.GetMyStudents(headID)
}
