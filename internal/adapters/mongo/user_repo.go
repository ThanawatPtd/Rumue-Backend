package mongo

import (
	"context"
	"fmt"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Client *mongo.Client
}

func ProvideMongoUserRepository(client *mongo.Client) repositories.UserRepository {
	return &MongoUserRepository{
		Client: client,
	}
}

// Save implements repositories.UserRepository.
func (m *MongoUserRepository) Save(c context.Context, user *entities.User) error {
	document := map[string]interface{}{
		"email":       user.Email,
		"fname":       user.Fname,
		"lname":       user.Lname,
		"password":    user.Password,
		"phoneNumber": user.PhoneNumber,
		"address":     user.Address,
		"nationality": user.Nationality,
		"citizenID":   user.CitizenID,
		"birthDate":   user.BirthDate,
	}
	collection := m.Client.Database("local").Collection("user")
	_, err := collection.InsertOne(context.TODO(), document)
	return err
}

// Delete implements repositories.UserRepository.
func (m *MongoUserRepository) Delete(c context.Context, id string) error {
	panic("unimplemented")
}

// GetByID implements repositories.UserRepository.
func (m *MongoUserRepository) GetByID(c context.Context, id string) (*entities.User, error) {
	panic("unimplemented")
}

// GetIDPasswordByEmail implements repositories.UserRepository.
func (m *MongoUserRepository) GetIDPasswordByEmail(c context.Context, email string) (*entities.User, error) {
	emailFilter := bson.M{
		"email": email,
	}
	var result bson.M
	err := m.Client.Database("local").Collection("user").FindOne(context.TODO(), emailFilter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with the given filter.")
			return nil, nil
		} else {
			fmt.Println("Error finding document:", err)
			return nil, err
		}
	}

	var user entities.User
	if result != nil {
		if email, ok := result["email"].(string); ok {
			user.Email = email
		}
	}
	return &user, nil
}

// GetIDPasswordByID implements repositories.UserRepository.
func (m *MongoUserRepository) GetIDPasswordByID(c context.Context, id string) (*entities.User, error) {
	panic("unimplemented")
}

// GetUserProfileByID implements repositories.UserRepository.
func (m *MongoUserRepository) GetUserProfileByID(c context.Context, id string) (*entities.UserProfile, error) {
	panic("unimplemented")
}

// ListAll implements repositories.UserRepository.
func (m *MongoUserRepository) ListAll(c context.Context) ([]entities.User, error) {
	panic("unimplemented")
}

// Update implements repositories.UserRepository.
func (m *MongoUserRepository) Update(c context.Context, user *entities.User) (*entities.User, error) {
	panic("unimplemented")
}

// UpdatePassword implements repositories.UserRepository.
func (m *MongoUserRepository) UpdatePassword(c context.Context, user *entities.User) error {
	panic("unimplemented")
}
