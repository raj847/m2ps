package repositories

import "m2ps/models"

type TrxRepository interface {
	CreateTrxInquiry(trx *models.InquryTrx) (id int, err error)
	CreateTrxExt(trx *models.TrxExt) (id int, err error)
	CreateTrxOu(trx *models.TrxOu) (id int, err error)
}

type TrxMongoRepository interface {
	GetData(start string, end string) (trxlist []*models.Trx, exists bool, err error)
}
