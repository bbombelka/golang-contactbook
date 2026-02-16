package main

import (
	"context"

	"github.com/bbombelka/contactbook/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const DATABASE_NAME = "placeholder"
const COLLECION_NAME = "users"

func findUser(id string) (types.FullUser, bool) {
	var result types.FullUser
	var usersCollection = mongoClient.Database(DATABASE_NAME).Collection(COLLECION_NAME)

	err := usersCollection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return types.FullUser{}, false
	}

	if err != nil {
		panic(err)
	}

	return result, true
}

func findUsers(limit int64, skip int64) []types.ShortUser {
	var results []types.ShortUser
	var usersCollection = mongoClient.Database(DATABASE_NAME).Collection(COLLECION_NAME)

	cursor, error := usersCollection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.D{{"id", 1}, {"name", 1}}).SetLimit(limit).SetSkip(skip))

	if error != nil {
		panic(error)
	}
	defer cursor.Close(context.TODO())

	cursorError := cursor.All(context.TODO(), &results)

	if cursorError != nil {
		panic(cursorError)
	}

	if len(results) == 0 {
		return []types.ShortUser{}
	}

	return results
}

func deleteUser(userId string) (*mongo.DeleteResult, error) {
	result, error := mongoClient.Database(DATABASE_NAME).Collection(COLLECION_NAME).DeleteOne(context.TODO(), bson.D{{"id", userId}})

	return result, error
}

func addUser(addUserBody types.AddUserRequestBody, newUserId string) (*mongo.InsertOneResult, error) {
	usersCollection := mongoClient.Database(DATABASE_NAME).Collection(COLLECION_NAME)

	newUser := types.FullUser{Id: newUserId, Name: *addUserBody.Name, Username: *addUserBody.Username, Address: *addUserBody.Address, PhoneNumber: *addUserBody.PhoneNumber}
	result, error := usersCollection.InsertOne(context.TODO(), newUser)

	return result, error
}

func updateUser(id string, updateUserData types.UpdateUserRequestBody) (*mongo.UpdateResult, error) {
	filter := bson.D{{"id", id}}
	updatedFields := getUpdateFields(updateUserData)
	update := bson.D{{"$set", updatedFields}}

	result, error := mongoClient.Database(DATABASE_NAME).Collection(COLLECION_NAME).UpdateOne(context.TODO(), filter, update)

	return result, error
}
