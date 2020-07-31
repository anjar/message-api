package repositories

import (
	"simpleapi/models"
	"github.com/jinzhu/gorm"
)

// Repository handles the basic operations of a message entity/model.
// It's an interface in order to be testable, i.e a memory message repository or
// a connected to an sql database.
type QueryMessage func(db *gorm.DB) (models.Message, bool, error)
type QueryMessages func(db *gorm.DB) ([]models.Message, map[string]interface{}, error)

type MessageRepository interface {
	All(query QueryMessages) (messages []models.Message, paginator map[string]interface{}, err error)
	CreateMessage(query QueryMessage) (messages models.Message, err error)
}

// NewMessageRepository returns a new message memory-based repository,
// the one and only repository type in our example.
func NewMessageRepository() MessageRepository {
	baseRepo := NewBaseRepositories()

	return &messageRepository{
		BaseRepository: baseRepo,
	}
}

// messageRepository is a "MessageRepository"
// which manages the Dependency Injector.
type messageRepository struct {
	BaseRepository *BaseRepositories
}

func (repository *messageRepository) All(query QueryMessages) (messages []models.Message, paginator map[string]interface{}, err error) {
	return query(repository.BaseRepository.Database)
}

func (repository *messageRepository) CreateMessage(query QueryMessage) (message models.Message, err error) {
	message, _, err = query(repository.BaseRepository.Database)
	return
}