package gomd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JubaerHossain/gomd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *MongoDB

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Ctx      context.Context
}

type MongoCollection struct {
	*mongo.Collection
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

func (mongoCollection MongoCollection) WithContext() MongoCollection {
	mongoCollection.Ctx, mongoCollection.CancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
	return mongoCollection
}

func NewNoSqlDB() *MongoDB {
	port := config.Config.GetString(config.Config.NoSqlConfig + ".Port")
	var noSqlProtocol, host string

	if port != "" {
		noSqlProtocol = "mongodb://"
		host = config.Config.GetString(config.Config.NoSqlConfig+".Host") + ":" + port
	} else {
		noSqlProtocol = "mongodb+srv://"
		host = config.Config.GetString(config.Config.NoSqlConfig + ".Host")
	}

	password := config.Config.GetString(config.Config.NoSqlConfig + ".Password")
	noSqlUserInfo := ""
	if password != "" {
		noSqlUserInfo = config.Config.GetString(config.Config.NoSqlConfig+".Username") + ":" + password + "@"
	}

	noSqlUri := noSqlProtocol + noSqlUserInfo + host + "/" + config.Config.GetString(config.Config.NoSqlConfig+".Database") + "?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(noSqlUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[C-Log] Connected to MongoDB")

	database := client.Database(config.Config.GetString(config.Config.NoSqlConfig + ".Database"))

	return &MongoDB{Client: client, Database: database, Ctx: ctx}
}

func (mongodb *MongoDB) Collection(name string) MongoCollection {
	return MongoCollection{Collection: mongodb.Database.Collection(name)}
}

func (collection MongoCollection) Find(filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error, context.Context) {

	collection = collection.WithContext()
	defer collection.CancelFunc()

	cursor, err := collection.Collection.Find(collection.Ctx, filter, opts...)
	return cursor, err, collection.Ctx
}

func (collection MongoCollection) InsertOne(document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {

	collection = collection.WithContext()
	defer collection.CancelFunc()

	return collection.Collection.InsertOne(collection.Ctx, document, opts...)
}

func (collection MongoCollection) FindOneAndUpdate(filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {

	collection = collection.WithContext()
	defer collection.CancelFunc()

	return collection.Collection.FindOneAndUpdate(collection.Ctx, filter, update, opts...)
}

func (collection MongoCollection) FindOne(filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {

	collection = collection.WithContext()
	defer collection.CancelFunc()

	result := collection.Collection.FindOne(collection.Ctx, filter, opts...)

	return result
}

func (collection MongoCollection) FindOneAndDelete(filter interface{},
	opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {

	collection = collection.WithContext()
	defer collection.CancelFunc()

	return collection.Collection.FindOneAndDelete(collection.Ctx, filter, opts...)
}
