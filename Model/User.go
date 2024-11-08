package Model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	GetUserDetails(ctx context.Context, id int) (User, error)
	CreateUser(ctx context.Context, user User) error
	FetchUser(ctx context.Context, email string) User
	GetUserName(ctx context.Context) (map[string]string, error)
}

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	UserName     string             `json:"username"`
	Dob          time.Duration      `json:"dob"`
	PasswordHash string             `json:"passwordHash"`
	Gender       string             `json:"gender"`
}
