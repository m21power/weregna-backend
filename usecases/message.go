package usecases

import "weregna-backend/domain"

type MessageUsecases struct {
	messageRepo domain.MessageRepository
}

func NewMessageUsecases(mr domain.MessageRepository) *MessageUsecases {
	return &MessageUsecases{
		messageRepo: mr,
	}
}

func (mu *MessageUsecases) CreateMessage(message *domain.Message) error {
	return mu.messageRepo.CreateMessage(message)
}
func (mu *MessageUsecases) GetMessagesByConversationID(conversationID string) ([]*domain.Message, error) {
	return mu.messageRepo.GetMessagesByConversationID(conversationID)
}

func (mu *MessageUsecases) GetMessageByID(id int) (*domain.Message, error) {
	return mu.messageRepo.GetMessageByID(id)
}
func (mu *MessageUsecases) UpdateMessage(message *domain.Message) error {
	return mu.messageRepo.UpdateMessage(message)
}
func (mu *MessageUsecases) DeleteMessage(id int) error {
	return mu.messageRepo.DeleteMessage(id)
}

func (mu *MessageUsecases) MarkMessageAsRead(convId string, userId string) error {
	return mu.messageRepo.MarkMessageAsRead(convId, userId)
}
func (mu *MessageUsecases) UnreadMessagesCount(convId string, userId string) (int, error) {
	return mu.messageRepo.UnreadMessagesCount(convId, userId)
}
