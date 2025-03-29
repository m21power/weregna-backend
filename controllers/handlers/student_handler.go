package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weregna-backend/domain"
	"weregna-backend/usecases"
	"weregna-backend/utils"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type StudentHandler struct {
	studentUsecases *usecases.StudentUsecases
}

func NewStudentHandler(studentUsecases *usecases.StudentUsecases) *StudentHandler {
	return &StudentHandler{
		studentUsecases: studentUsecases,
	}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student domain.StudentModel
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	student.Password = string(hash)
	err = h.studentUsecases.CreateStudent(&student)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, student)

}
func (h *StudentHandler) GetStudentByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	student, err := h.studentUsecases.GetStudentByEmail(email)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, student)
}
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	var student domain.StudentModel
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	oldStudent, err := h.studentUsecases.GetStudentByEmail(email)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	newStudent := changeToNewStudent(&student, oldStudent)
	if newStudent.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(newStudent.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}
		newStudent.Password = string(hash)
	}

	err = h.studentUsecases.UpdateStudent(newStudent)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, newStudent)

}
func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]
	id, err := strconv.Atoi(Id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	student, err := h.studentUsecases.GetStudentByID(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, student)
}
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]
	id, err := strconv.Atoi(Id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	err = h.studentUsecases.DeleteStudent(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Student deleted successfully", nil, http.StatusOK)
}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.studentUsecases.GetStudents()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, students)
}

func changeToNewStudent(newStudent *domain.StudentModel, oldStudent *domain.StudentModel) *domain.StudentModel {
	if newStudent.Email != "" {
		oldStudent.Email = newStudent.Email
	}
	if newStudent.Password != "" {
		oldStudent.Password = newStudent.Password
	}
	if newStudent.Name != "" {
		oldStudent.Name = newStudent.Name
	}
	if newStudent.ProfilePic != nil {
		oldStudent.ProfilePic = newStudent.ProfilePic
	}
	if newStudent.TelegramUsername != nil {
		oldStudent.TelegramUsername = newStudent.TelegramUsername
	}
	if newStudent.HeadID != nil {
		oldStudent.HeadID = newStudent.HeadID
	}
	if newStudent.TotalDuration != 0 {
		oldStudent.TotalDuration = newStudent.TotalDuration
	}
	if newStudent.TotalActiveDays != 0 {
		oldStudent.TotalActiveDays = newStudent.TotalActiveDays
	}
	return oldStudent
}
