package domain

import "time"

type Conversation struct {
	ID              string    `json:"id" db:"id"`
	User1ID         int       `json:"user1_id" db:"user1_id"`
	User2ID         int       `json:"user2_id" db:"user2_id"`
	LastMessageText string    `json:"last_message_text" db:"last_message_text"`
	LastMessageAt   time.Time `json:"last_message_at" db:"last_message_at"`
}

type ConversationRepository interface {
	CreateConversation(conversation *Conversation) (string, error)
	GetConversationsByUserID(userID int) ([]*Conversation, error)
	GetConversationByID(id string) (*Conversation, error)
	GetConversationByUsers(user1ID int, user2ID int) (*Conversation, error)
	UpdateConversation(conversation *Conversation) error
	DeleteConversation(id string) error
	GetConversationId(user1ID int, user2ID int) (string, error)
}
