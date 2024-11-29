package product_mng_repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"graphql/config"
	"graphql/internal/db_store"
)

type ProductMngRepository struct {
	cfg   *config.Config
	mngDB *mongo.Database
}

func NewProductMngRepository(cfg *config.Config, mngDB *mongo.Database) *ProductMngRepository {
	return &ProductMngRepository{
		cfg:   cfg,
		mngDB: mngDB,
	}
}

func (r *ProductMngRepository) AddProductAttributes(ctx context.Context, attributes map[string]interface{}) (string, error) {
	collection := r.mngDB.Collection(db_store.AttributesCollectionName)
	insertResult, err := collection.InsertOne(ctx, attributes)
	if err != nil {
		return "", err
	}

	value, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("impossible to convert mngdb id")
	}

	return value.Hex(), nil
}

func (r *ProductMngRepository) GetProductAttributesByFilter(ctx context.Context, attributesFilter *AttributesFilter) ([]ProductAttributes, error) {
	collection := r.mngDB.Collection(db_store.AttributesCollectionName)

	filter, err := getFindParams(attributesFilter)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var productAttributes []ProductAttributes

	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		attributes := ProductAttributes{Attributes: make(map[string]interface{})}

		for key, value := range result {
			if key != "_id" {
				attributes.Attributes[key] = value
				continue
			}

			if id, ok := value.(primitive.ObjectID); ok {
				attributes.ProductId = id.Hex()
			}
		}

		productAttributes = append(productAttributes, attributes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return productAttributes, nil
}

func getFindParams(attributesFilter *AttributesFilter) (bson.M, error) {
	filter := bson.M{}

	if len(attributesFilter.AttributesIds) != 0 {
		var ids []primitive.ObjectID
		for _, idStr := range attributesFilter.AttributesIds {
			id, err := primitive.ObjectIDFromHex(idStr)
			if err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}

		filter[db_store.AttributesIdColumnName] = bson.M{"$in": ids}
	}

	if len(attributesFilter.Attributes) != 0 {
		for key, value := range attributesFilter.Attributes {
			filter[key] = value
		}
	}

	return filter, nil
}
