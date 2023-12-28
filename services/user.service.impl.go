package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/kunle001/gogingonic/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{
		Key:   "name",
		Value: name,
	}}

	if err := u.usercollection.FindOne(u.ctx, query).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	filter := bson.D{{}}
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(u.ctx)

	for cursor.Next(u.ctx) {
		var user models.User

		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("no documents")
	}

	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{
		bson.E{
			Key:   "user_name",
			Value: user.Name,
		},
	}

	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{Key: "user_name", Value: user.Name},
				bson.E{Key: "age", Value: user.Age},
				bson.E{Key: "address", Value: user.Address},
			},
		},
	}

	result, _ := u.usercollection.UpdateByID(u.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no existing user with this name")
	}

	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{
		bson.E{
			Key:   "user_name",
			Value: name,
		},
	}
	user, err := u.usercollection.DeleteOne(u.ctx, filter)

	if user.DeletedCount != 1 {
		return errors.New("no existing user with this name")
	}
	return err
}
