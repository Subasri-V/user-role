package interfaces

import (
	"github.com/Subasri-V/user-role.git/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRole interface {
	AddUser(*models.User) (*models.DBResponse, error)
	EnableUser(string) (*mongo.UpdateResult, error)
	DisableUser(string)(*mongo.UpdateResult,error)
	UpdateRole(string,[]string) (*mongo.UpdateResult,error)
	AssociateRole(string,string)  (*mongo.UpdateResult,error)
	Remove(string,string)  (*mongo.UpdateResult,error)
	AppendArray(string,[]string) (*mongo.UpdateResult,error)
	ListFeatures(string)(error )
	ListFeaturesInPostman(role *models.Role) (*models.Role,error)
}
