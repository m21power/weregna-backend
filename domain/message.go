package domain

import "time"

type Message struct {
	ID             int       `json:"id" db:"id"`
	ConversationID string    `json:"conversation_id" db:"conversation_id"`
	SenderID       int       `json:"sender_id" db:"sender_id"`
	Content        string    `json:"content" db:"content"`
	Status         string    `json:"status" db:"status"`
	SentAt         time.Time `json:"sent_at" db:"sent_at"`
}

type MessageRepository interface {
	CreateMessage(message *Message) error
	GetMessagesByConversationID(conversationID string) ([]*Message, error)
	GetMessageByID(id int) (*Message, error)
	UpdateMessage(message *Message) error
	DeleteMessage(id int) error
	MarkMessageAsRead(convId string, userId string) error
	UnreadMessagesCount(convId string, userId string) (int, error)
}
