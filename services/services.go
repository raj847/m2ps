package services

import (
	"database/sql"
	"m2ps/repositories"
)

type UsecaseService struct {
	RepoDB       *sql.DB
	TrxRepo      repositories.TrxRepository
	TrxMongoRepo repositories.TrxMongoRepository
}

func NewUsecaseService(
	RepoDB *sql.DB,
	TrxRepo repositories.TrxRepository,
	TrxMongoRepo repositories.TrxMongoRepository,
) UsecaseService {
	return UsecaseService{
		RepoDB:       RepoDB,
		TrxRepo:      TrxRepo,
		TrxMongoRepo: TrxMongoRepo,
	}
}
