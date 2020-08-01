package services

import (
	"fmt"
	"simpleapi/models"
	"simpleapi/repositories"
	"simpleapi/helpers"
	// "log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

// MessageService handles CRUID operations of a message model,
// it depends on a message repository for its actions.
type MessageService interface {
	
	GetMessageList(inputPagination InputPagination) ([]models.Message, map[string]interface{}, error)
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

func (service *messageService) GetMessageList(inputPagination InputPagination) ([]models.Message, map[string]interface{}, error) {
	return service.repo.All(func(db *gorm.DB) (messages []models.Message, paginator map[string]interface{} , err error) {
		dbCon := db.Debug()

		// Process query condition and its bind parameters
		paginator, _ = helpers.GetPagination(dbCon, inputPagination.Limit, inputPagination.Page)

		// Set offset
		offset := 0
		if inputPagination.Page > 1 {
			offset = (inputPagination.Page - 1) * inputPagination.Limit
		}
		dbWithConfig := db.Debug().Limit(inputPagination.Limit).Offset(offset).Order(inputPagination.OrderBy)

		dbWithConfig.Find(&messages)

		return messages, paginator, err
	})
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