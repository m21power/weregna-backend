package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"weregna-backend/domain"
	"weregna-backend/usecases"
	"weregna-backend/utils"

	"github.com/gorilla/mux"
)

type MessageHandler struct {
	messageUsecases     *usecases.MessageUsecases
	conversationUsecase *usecases.ConversationUsecases
}

func NewMessageHandler(messageUsecases *usecases.MessageUsecases, conversationUsecases *usecases.ConversationUsecases) *MessageHandler {
	return &MessageHandler{
		messageUsecases:     messageUsecases,
		conversationUsecase: conversationUsecases,
	}
}
func (m *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	// Define payload structure
	type createMessagePayload struct {
		SenderID       int     `json:"sender_id"`
		ReceiverID     int     `json:"receiver_id"`
		Content        string  `json:"content"`
		ConversationID *string `json:"conversation_id"`
	}

	// Decode request body once
	var payload createMessagePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	// Log the payload
	fmt.Println(payload)

	// Check if ConversationID exists
	if payload.ConversationID != nil {
		message := domain.Message{
			ConversationID: *payload.ConversationID,
			Content:        payload.Content,
			SenderID:       payload.SenderID,
		}

		err := m.messageUsecases.CreateMessage(&message)
		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, http.StatusCreated, message)
		return
	}

	// If no conversation exists, create a new one
	conver := domain.Conversation{
		User1ID:         payload.SenderID,
		User2ID:         payload.ReceiverID,
		LastMessageText: payload.Content,
		LastMessageAt:   time.Now(),
	}

	convId, err := m.conversationUsecase.CreateConversation(&conver)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Println("Conversation ID: ", convId)

	// Create message with new conversation ID
	message := domain.Message{
		ConversationID: convId,
		Content:        payload.Content,
		SenderID:       payload.SenderID,
	}

	err = m.messageUsecases.CreateMessage(&message)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (m *MessageHandler) GetMessagesByConversationID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conversationID := vars["conversationID"]

	messages, err := m.messageUsecases.GetMessagesByConversationID(conversationID)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, messages)
}

func (m *MessageHandler) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	Id, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	message, err := m.messageUsecases.GetMessageByID(Id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, message)
}
func (m *MessageHandler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Content string `json:"content"`
	}
	var content payload
	vars := mux.Vars(r)
	id := vars["id"]
	Id, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	message, err := m.messageUsecases.GetMessageByID(Id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	message.Content = content.Content
	err = m.messageUsecases.UpdateMessage(message)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, message)
}
func (m *MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	Id, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	err = m.messageUsecases.DeleteMessage(Id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusNoContent, nil)
}
func (m *MessageHandler) MarkMessageAsRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	convId := vars["conversationID"]
	userId := vars["userID"]
	err := m.messageUsecases.MarkMessageAsRead(convId, userId)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
func (m *MessageHandler) UnreadMessagesCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	convId := vars["conversationID"]
	userId := vars["userID"]
	count, err := m.messageUsecases.UnreadMessagesCount(convId, userId)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, count)
}
