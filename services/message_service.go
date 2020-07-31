package services

import (
	"fmt"
	"simpleapi/models"
	"simpleapi/repositories"
	// "simpleapi/helpers"
	// "log"

	// npq "github.com/Knetic/go-namedParameterQuery"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

// MessageService handles CRUID operations of a message model,
// it depends on a message repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type MessageService interface {
	
	CreateMessage(messageModel models.Message, ctx iris.Context) (models.Message, error)
}
// NewMessageService returns the default message service.
func NewMessageService(repo repositories.MessageRepository) MessageService {
	return &messageService{
		repo: repo,
	}
}

type messageService struct {
	repo repositories.MessageRepository
}

func (service *messageService) CreateMessage(messageModel models.Message, ctx iris.Context) (models.Message, error) {
	return service.repo.CreateMessage(func(db *gorm.DB) (message models.Message, b bool, err error) {

		tx := db.Begin()
		err = tx.Create(&messageModel).Error
		if err != nil {
			tx.Rollback()
			return models.Message{}, false, fmt.Errorf("error create message, err := %s", err.Error())
		}

		tx.Commit()
		return messageModel, true, err
	})
}