package utils

import (
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	
)

type Collection struct {
	*mongo.Collection
}

// Insert a document
func (c *Collection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, document)
}

// Find a document
func (c *Collection) Find(ctx context.Context, filter interface{}, result interface{}) error {
	cursor, err := c.Collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	
	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}

// Update a document
func (c *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateOne(ctx, filter, update)
}

// Delete a document
func (c *Collection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return c.Collection.DeleteOne(ctx, filter)
}