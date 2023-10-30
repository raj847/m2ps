package repositories

import (
	"context"
	"database/sql"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB       *sql.DB
	MongoDB  *mongo.Database
	Context  context.Context
	DynamoDB *dynamodb.DynamoDB
}

func NewRepository(conn *sql.DB, MongoDB *mongo.Database, ctx context.Context, dynamodb *dynamodb.DynamoDB) Repository {
	return Repository{
		DB:       conn,
		MongoDB:  MongoDB,
		Context:  ctx,
		DynamoDB: dynamodb,
	}
}
