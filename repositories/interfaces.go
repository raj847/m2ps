package repositories

import (
	"database/sql"
	"m2ps/models"
)

type TrxRepository interface {
	CreateTrxInquiry(trx *models.InquryTrx, tx *sql.Tx) (id int, err error)
	CreateTrxExt(trx *models.TrxExt, tx *sql.Tx) (id int, err error)
	CreateTrxOu(trx *models.TrxOu, tx *sql.Tx) (id int, err error)
	GetTrx(trx models.TrxFilter) ([]models.TrxResponseData, error)
	GetTrxSummariesAdvance(trx models.TrxFilter) (models.TrxResponseSummaries, error)
}

type TrxMongoRepository interface {
	GetData(start string, end string) (trxlist []*models.Trx, exists bool, err error)
}
