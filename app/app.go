package app

import (
	"database/sql"
	"m2ps/repositories"
	"m2ps/repositories/trxRepository"
	"m2ps/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repository) services.UsecaseService {
	// Create table for dynamo db
	//CreateTableMovies(svc)

	// Services
	trxRepo := trxRepository.NewTrxRepository(repo)
	trxMongoRepo := trxRepository.NewTrxMongoRepository(repo)

	usecaseSvc := services.NewUsecaseService(DB, trxRepo, trxMongoRepo)

	return usecaseSvc
}
