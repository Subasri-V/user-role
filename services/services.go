package services

import (
	"context"
	"fmt"

	"github.com/Subasri-V/user-role.git/interfaces"
	"github.com/Subasri-V/user-role.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRoleService struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
	client          *mongo.Client
	roleCollection  *mongo.Collection
}

func InitUserRoleService(ctx context.Context, mongoCollection *mongo.Collection, client *mongo.Client, roleCollection *mongo.Collection) interfaces.IUserRole {
	return &UserRoleService{ctx, mongoCollection, client, roleCollection}
}

// AddUser implements interfaces.IUserRole.
func (c *UserRoleService) AddUser(user *models.User) (*models.DBResponse, error) {
	res, err := c.mongoCollection.InsertOne(c.ctx, &user)
	if err != nil {
		return nil, err
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.mongoCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println(newUser)
	return newUser, nil
}
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// DisableUser implements interfaces.IUserRole.
func (c *UserRoleService) DisableUser(name string) (*mongo.UpdateResult, error) {
	fmt.Println("services", name)
	iv := bson.M{"name": name}
	fv := bson.M{"$set": bson.M{"status": "disabled"}}
	res, err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableUser implements interfaces.IUserRole.
func (c *UserRoleService) EnableUser(name string) (*mongo.UpdateResult, error) {
	iv := bson.M{"name": name}
	fv := bson.M{"$set": bson.M{"status": "enabled"}}
	res, err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *UserRoleService) UpdateRole(name string, roles []string) (*mongo.UpdateResult, error) {
	iv := bson.M{"name": name}
	updatedDoc := c.mongoCollection.FindOne(c.ctx, iv)
	if updatedDoc.Err() != nil {
		return nil, updatedDoc.Err()
	}
	var result struct {
		Status string `bson:"status"`
	}
	if err := updatedDoc.Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "enabled" {
		return nil, fmt.Errorf("status is not enabled")
	} else {
		fv := bson.M{"$set": bson.M{"role": roles}}
		res, err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

}

func (c *UserRoleService) AssociateRole(name string, role string) (*mongo.UpdateResult, error) {
	iv := bson.M{"name": name}
	updatedDoc := c.mongoCollection.FindOne(c.ctx, iv)
	if updatedDoc.Err() != nil {
		return nil, updatedDoc.Err()
	}
	var result struct {
		Status string `bson:"status"`
	}
	if err := updatedDoc.Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "enabled" {
		return nil, fmt.Errorf("status is not enabled")
	} else {
		update := bson.M{
			"$push": bson.M{
				"role": role,
			},
		}
		res, err := c.mongoCollection.UpdateOne(c.ctx, iv, update)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

}

func (c *UserRoleService) Remove(name string, role string) (*mongo.UpdateResult, error) {
	iv := bson.M{"name": name}

	updatedDoc := c.mongoCollection.FindOne(c.ctx, iv)
	if updatedDoc.Err() != nil {
		return nil, updatedDoc.Err()
	}
	var result struct {
		Status string `bson:"status"`
	}
	if err := updatedDoc.Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "enabled" {
		return nil, fmt.Errorf("status is not enabled")
	} else {
		update := bson.M{
			"$pull": bson.M{
				"role": role,
			},
		}
		res, err := c.mongoCollection.UpdateOne(c.ctx, iv, update)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (c *UserRoleService) AppendArray(name string, roles []string) (*mongo.UpdateResult, error) {
	iv := bson.M{"name": name}

	updatedDoc := c.mongoCollection.FindOne(c.ctx, iv)
	if updatedDoc.Err() != nil {
		return nil, updatedDoc.Err()
	}
	var result struct {
		Status string `bson:"status"`
	}
	if err := updatedDoc.Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "enabled" {
		return nil, fmt.Errorf("status is not enabled")
	} else {
		update := bson.M{
			"$push": bson.M{
				"role": bson.M{
					"$each": roles,
				},
			},
		}
		res, err := c.mongoCollection.UpdateOne(c.ctx, iv, update)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (c *UserRoleService) ListFeatures(role string) (error ){
	iv := bson.M{"role": role}

	// Find the document
	doc := c.roleCollection.FindOne(c.ctx, iv)
	if doc.Err() != nil {
		return doc.Err()
	}

	// Decode the document into a map to access its fields
	var result map[string]interface{}
	if err := doc.Decode(&result); err != nil {
		return err
	}

	// Iterate through the fields and print their names and values
	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}

	return nil
}

func (c*UserRoleService) ListFeaturesInPostman(role *models.Role) (*models.Role,error){
	filter := bson.D{{Key: "role", Value: role.Role}}
	var result *models.Role
	err := c.roleCollection.FindOne(c.ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}