package repository

import (
	"database/sql"
	"weregna-backend/domain"
)

type ConversationRepoImpl struct {
	db *sql.DB
}

func NewConversationRepoImpl(db *sql.DB) *ConversationRepoImpl {
	return &ConversationRepoImpl{
		db: db,
	}
}

func (r *ConversationRepoImpl) CreateConversation(conversation *domain.Conversation) (string, error) {
	// Check if the conversation already exists
	query := "SELECT id FROM conversations WHERE user1_id = $1 AND user2_id = $2 OR user1_id = $2 AND user2_id = $1"
	var id string
	err := r.db.QueryRow(query, conversation.User1ID, conversation.User2ID).Scan(&id)
	if err == sql.ErrNoRows {
		// If it doesn't exist, create a new conversation
		insertQuery := "INSERT INTO conversations (user1_id, user2_id, last_message_text, last_message_at) VALUES ($1, $2, $3, $4) RETURNING id"
		err = r.db.QueryRow(insertQuery, conversation.User1ID, conversation.User2ID, conversation.LastMessageText, conversation.LastMessageAt).Scan(&id)
		if err != nil {
			return "", err
		}
	} else if err == nil {
		// If it exists, update the last message and timestamp
		updateQuery := "UPDATE conversations SET last_message_text = $1, last_message_at = $2 WHERE id = $3"
		_, err = r.db.Exec(updateQuery, conversation.LastMessageText, conversation.LastMessageAt, id)
		if err != nil {
			return "", err
		}
	} else {
		// Handle other errors
		return "", err
	}
	return id, nil
}

func (r *ConversationRepoImpl) GetConversationsByUserID(userID int) ([]*domain.Conversation, error) {
	rows, err := r.db.Query("SELECT * FROM conversations WHERE user1_id = $1 OR user2_id = $1", userID)
	if err != nil {
		return nil, err
	}
	var conversations []*domain.Conversation
	for rows.Next() {
		var conversation domain.Conversation
		err := rows.Scan(&conversation.ID, &conversation.User1ID, &conversation.User2ID, &conversation.LastMessageText, &conversation.LastMessageAt)
		if err != nil {
			return nil, err
		}
		conversations = append(conversations, &conversation)
	}
	return conversations, nil
}

func (r *ConversationRepoImpl) GetConversationByID(id string) (*domain.Conversation, error) {
	var conversation domain.Conversation
	err := r.db.QueryRow("SELECT * FROM conversations WHERE id = $1", id).Scan(&conversation.ID, &conversation.User1ID, &conversation.User2ID, &conversation.LastMessageText, &conversation.LastMessageAt)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepoImpl) UpdateConversation(conversation *domain.Conversation) error {
	query := "UPDATE conversations SET last_message_text = $1, last_message_at = $2 WHERE id = $3"
	_, err := r.db.Exec(query, conversation.LastMessageText, conversation.LastMessageAt, conversation.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ConversationRepoImpl) DeleteConversation(id string) error {
	query := "DELETE FROM conversations WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *ConversationRepoImpl) GetConversationByUsers(user1ID int, user2ID int) (*domain.Conversation, error) {
	var conversation domain.Conversation
	err := r.db.QueryRow("SELECT * FROM conversations WHERE user1_id = $1 AND user2_id = $2 OR user1_id = $2 AND user2_id = $1", user1ID, user2ID).Scan(&conversation.ID, &conversation.User1ID, &conversation.User2ID, &conversation.LastMessageText, &conversation.LastMessageAt)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepoImpl) GetConversationId(user1ID int, user2ID int) (string, error) {
	query := "SELECT id FROM conversations WHERE user1_id = $1 AND user2_id = $2 OR user1_id = $2 AND user2_id = $1"
	var id string
	err := r.db.QueryRow(query, user1ID, user2ID).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
