package service

import (
	"github.com/m1kr0b/message-processing/internal/kafka"
	"github.com/m1kr0b/message-processing/internal/repository"
	Intern "github.com/m1kr0b/message-processing/model/message"
)

type MessageService struct {
	repos    repository.MessageRepository
	producer *kafka.Producer
}

func NewMessageService(repos *repository.Repository, producer *kafka.Producer) *MessageService {
	return &MessageService{
		repos:    repos.MessageRepository,
		producer: producer,
	}
}

func (s *MessageService) CreateMessage(message Intern.Message) (int, error) {
	id, err := s.repos.CreateMessage(message)
	message.Id = id
	s.producer.SendMessage(message)
	return id, err
}

func (s *MessageService) ProcessMessage(id int) {

	s.repos.ProcessMessage(id)
}

func (s *MessageService) GetMessageById(id int) (Intern.Message, error) {
	return s.repos.GetMessageById(id)
}

func (s *MessageService) GetStats() (int, error) {
	return s.repos.GetStats()
}
