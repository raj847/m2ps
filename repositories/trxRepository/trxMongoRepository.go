package trxRepository

import (
	"m2ps/models"
	"m2ps/repositories"

	"gopkg.in/mgo.v2/bson"
)

type trxMongoRepository struct {
	RepoDB repositories.Repository
}

func NewTrxMongoRepository(repoDB repositories.Repository) trxMongoRepository {
	return trxMongoRepository{
		RepoDB: repoDB,
	}
}

func (ctx trxMongoRepository) GetData(start string, end string) (trxlist []*models.Trx, exists bool, err error) {

	filter := bson.M{
		"docDate": bson.M{"$gte": start, "$lte": end},
	}

	data, err := ctx.RepoDB.MongoDB.Collection("Trx").Find(ctx.RepoDB.Context, filter)
	if err != nil {
		return nil, false, err
	}

	for data.Next(ctx.RepoDB.Context) {
		var val models.Trx
		if err = data.Decode(&val); err != nil {
			return nil, false, err
		}

		trxlist = append(trxlist, &val)
	}
	data.Close(ctx.RepoDB.Context)

	if len(trxlist) == 0 {
		return nil, false, nil
	}

	return trxlist, true, nil
}
