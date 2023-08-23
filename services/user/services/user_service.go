package services

import (
	"fmt"
	pagination "github.com/gobeam/mongo-go-pagination"
	"github.com/JubaerHossain/gomd/services/user/validation"
	"github.com/JubaerHossain/gomd/services/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
    "time"
)

func AllUser(requestFilter map[string]interface{}) ([]models.User, pagination.PaginationData) {
	var users []models.User

	filter := bson.M{}

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	page, _ := strconv.ParseInt(requestFilter["page"].(string), 10, 64)
	limit, _ := strconv.ParseInt(requestFilter["limit"].(string), 10, 64)

	paginatedData, err := pagination.New(models.UserCollection.Collection).
		Page(page).
		Limit(limit).
		Sort("created_at", -1).
		Decode(&users).
		Filter(filter).
		Find()

	if err != nil {
		panic(err)
	}
	return users, paginatedData.Pagination
}

func CreateAUser(createUser validation.CreateUserRequest) models.User {
	user := models.User{
		Id:                 primitive.NewObjectID(),
		Task:               createUser.Task,
		Status:             createUser.Status,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	result, err := models.UserCollection.InsertOne(user)
	if err != nil || result == nil {
		panic(err)
	}

	return user
}

func UpdateAUser(userId string, updateUser validation.UpdateUserRequest) (models.User, error) {

	objId, _ := primitive.ObjectIDFromHex(userId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.UserCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"task":                updateUser.Task,
				"status":              updateUser.Status,
				"updated_at":          time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.User{}, result.Err()
	}

	var user models.User
	if err := result.Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func AUser(userId string) models.User {
	var user models.User

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := models.UserCollection.FindOne(bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return user
}

func DeleteAUser(userId string) bool {
	objId, _ := primitive.ObjectIDFromHex(userId)

	result := models.UserCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}