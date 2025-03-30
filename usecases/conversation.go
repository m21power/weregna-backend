package usecases

import "weregna-backend/domain"

type ConversationUsecases struct {
	conversationRepo domain.ConversationRepository
}

func NewConversationUsecases(cr domain.ConversationRepository) *ConversationUsecases {
	return &ConversationUsecases{
		conversationRepo: cr,
	}
}
func (cu *ConversationUsecases) CreateConversation(conversation *domain.Conversation) (string, error) {
	return cu.conversationRepo.CreateConversation(conversation)
}
func (cu *ConversationUsecases) GetConversationsByUserID(userID int) ([]*domain.Conversation, error) {
	return cu.conversationRepo.GetConversationsByUserID(userID)
}
func (cu *ConversationUsecases) GetConversationByID(id string) (*domain.Conversation, error) {
	return cu.conversationRepo.GetConversationByID(id)
}

func (cu *ConversationUsecases) UpdateConversation(conversation *domain.Conversation) error {
	return cu.conversationRepo.UpdateConversation(conversation)
}
func (cu *ConversationUsecases) DeleteConversation(id string) error {
	return cu.conversationRepo.DeleteConversation(id)
}
func (cu *ConversationUsecases) GetConversationByUsers(user1ID int, user2ID int) (*domain.Conversation, error) {
	return cu.conversationRepo.GetConversationByUsers(user1ID, user2ID)
}

func (cu *ConversationUsecases) GetConversationId(user1ID int, user2ID int) (string, error) {
	return cu.conversationRepo.GetConversationId(user1ID, user2ID)
}
