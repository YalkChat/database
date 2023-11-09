package database

import (
	"github.com/AleRosmo/engine/models/db"
	"github.com/AleRosmo/engine/models/models"
)

// TODO: Evaluate strings arguments used
// TODO: We could pass the type instance with the ID property set and use .Find() on it
// TODO: Decide and choose if these operations must have the db model passed as arg or a string based on the example above.
// TODO: Also need to decide the returns of these functions

type DatabaseOperations interface {
	SaveMessage(*models.Message) (*db.Message, error)
	GetMessage(uint) (*db.Message, error)
	GetUsersByChatId(uint) ([]*db.User, error)
	GetUserByID(uint) (*db.User, error)
	GetUserByUsername(username string) (user *db.User, err error)
	NewUser(*db.User) (*db.User, error)
	NewChat(*models.Chat) (*db.Chat, error)
	NewChatType(*models.ChatType) (*db.ChatType, error)
	NewUserWithPassword(*models.UserCreationEvent) (*db.User, error)
	IsServerInitialized() (bool, error)
	SaveServerSettings(*models.ServerSettings) error
}
