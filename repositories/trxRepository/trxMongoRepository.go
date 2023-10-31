package trxRepository

import (
	"m2ps/models"
	"m2ps/repositories"

	"go.mongodb.org/mongo-driver/mongo/options"
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

	options := options.Find().SetSkip(0).SetLimit(1)

	data, err := ctx.RepoDB.MongoDB.Collection("Trx").Find(ctx.RepoDB.Context, filter, options)
	if err != nil {
		return nil, false, err
	}

	for data.Next(ctx.RepoDB.Context) {
		var val models.Trx
		if err = data.Decode(&val); err != nil {
			return nil, false, err
		}

		trxlist = append(trxlist, &val)
		// log.Println("datanya", utils.ToString(&val))
	}
	data.Close(ctx.RepoDB.Context)

	if len(trxlist) == 0 {
		return nil, false, nil
	}
	// log.Println("datanya", utils.ToString(trxlist))

	return trxlist, true, nil
}
