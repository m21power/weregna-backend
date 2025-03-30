package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weregna-backend/domain"
	"weregna-backend/usecases"
	"weregna-backend/utils"

	"github.com/gorilla/mux"
)

type HeadHandler struct {
	headUsecases *usecases.HeadUsecases
}

func NewHeadHandler(headUsecases *usecases.HeadUsecases) *HeadHandler {
	return &HeadHandler{
		headUsecases: headUsecases,
	}
}
func (h *HeadHandler) CreateHead(w http.ResponseWriter, r *http.Request) {
	var head domain.Head
	err := json.NewDecoder(r.Body).Decode(&head)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	err = h.headUsecases.CreateHead(&head)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, head)
}

func (h *HeadHandler) GetHeadByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	head, err := h.headUsecases.FindHeadByEmail(email)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, head)
}

func (h *HeadHandler) GetHeadByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	head, err := h.headUsecases.FindHeadByID(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, head)
}

func (h *HeadHandler) UpdateHead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	var head domain.Head
	err := json.NewDecoder(r.Body).Decode(&head)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	oldHead, err := h.headUsecases.FindHeadByEmail(email)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	updatedHead := changeToNewHead(&head, oldHead)
	err = h.headUsecases.UpdateHead(updatedHead)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, head)
}
func (h *HeadHandler) DeleteHead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	err = h.headUsecases.DeleteHead(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Deleted")
}
func (h *HeadHandler) GetHeads(w http.ResponseWriter, r *http.Request) {
	heads, err := h.headUsecases.GetHeads()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, heads)
}
func (h *HeadHandler) AddMyStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	headID, err := strconv.Atoi(vars["headID"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	var payload struct {
		StudentID int `json:"studentId"`
	}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	studentid := payload.StudentID

	err = h.headUsecases.AddMyStudent(headID, studentid)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Added")
}

func (h *HeadHandler) GetMyStudents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	headID, err := strconv.Atoi(vars["headID"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	students, err := h.headUsecases.GetMyStudents(headID)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, students)
}
func changeToNewHead(newHead *domain.Head, oldHead *domain.Head) *domain.Head {
	if newHead.Email != "" {
		oldHead.Email = newHead.Email
	}
	if newHead.Password != "" {
		oldHead.Password = newHead.Password
	}
	if newHead.Name != "" {
		oldHead.Name = newHead.Name
	}
	if newHead.Role != "" {
		oldHead.Role = newHead.Role
	}
	if newHead.ProfilePic != nil {
		oldHead.ProfilePic = newHead.ProfilePic
	}
	if newHead.TelegramUsername != nil {
		oldHead.TelegramUsername = newHead.TelegramUsername
	}
	return oldHead

}
