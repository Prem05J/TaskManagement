package db

import (
	"context"
	"crypto/tls"

	"github.com/taskManagement/Utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Connecting with database collection
func GetDBCollection(coll string) *mongo.Collection {
	return db.Collection(coll)
}

// DB Connection Establisment
func InitializeDatabase() (*mongo.Database, error) {
	URI := Utils.GetEnv("MOGODB_URI", "mongodb+srv://admin:root@cluster0.w4lxs.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI).SetTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		return nil, err
	}
	db = client.Database(Utils.GetEnv("DATABASE_NAME", "taskManagement"))
	return db, nil

}
