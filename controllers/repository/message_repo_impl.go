package repository

import (
	"database/sql"
	"weregna-backend/domain"
)

type MessageRepoImpl struct {
	db *sql.DB
}

func NewMessageRepoImpl(db *sql.DB) *MessageRepoImpl {
	return &MessageRepoImpl{
		db: db,
	}
}

func (r *MessageRepoImpl) CreateMessage(message *domain.Message) error {
	query := "INSERT INTO messages (conversation_id, sender_id, content) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, message.ConversationID, message.SenderID, message.Content)
	if err != nil {
		return err
	}
	return nil
}
func (r *MessageRepoImpl) GetMessagesByConversationID(conversationID string) ([]*domain.Message, error) {
	rows, err := r.db.Query("SELECT * FROM messages WHERE conversation_id = $1", conversationID)
	if err != nil {
		return nil, err
	}
	var messages []*domain.Message
	for rows.Next() {
		var message domain.Message
		err := rows.Scan(&message.ID, &message.ConversationID, &message.SenderID, &message.Content, &message.Status, &message.SentAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}

func (r *MessageRepoImpl) GetMessageByID(id int) (*domain.Message, error) {
	var message domain.Message
	err := r.db.QueryRow("SELECT * FROM messages WHERE id = $1", id).Scan(&message.ID, &message.ConversationID, &message.SenderID, &message.Content, &message.Status, &message.SentAt)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
func (r *MessageRepoImpl) UpdateMessage(message *domain.Message) error {
	query := "UPDATE messages SET conversation_id = $1, sender_id = $2, content = $3, status = $4, sent_at = $5 WHERE id = $6"
	_, err := r.db.Exec(query, message.ConversationID, message.SenderID, message.Content, message.Status, message.SentAt, message.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *MessageRepoImpl) DeleteMessage(id int) error {
	query := "DELETE FROM messages WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepoImpl) MarkMessageAsRead(convId string, userId string) error {
	query := "UPDATE messages SET status = 'read' WHERE conversation_id = $1 AND sender_id != $2"
	_, err := r.db.Exec(query, convId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepoImpl) UnreadMessagesCount(convId string, userId string) (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM messages WHERE conversation_id = $1 AND sender_id != $2 AND status = 'unseen'", convId, userId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
