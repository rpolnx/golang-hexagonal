package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	c "rpolnx.com.br/golang-hex/application/config"
	ce "rpolnx.com.br/golang-hex/application/error"
	u "rpolnx.com.br/golang-hex/domain/model/user"
	"rpolnx.com.br/golang-hex/domain/ports/out"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func InitializeRepo(mongo c.Mongo) (u out.MongoUserRepository, err error) {
	repo, err := newMongoRepository(mongo.Uri, mongo.Db, mongo.Timeout)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func newMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (out.MongoUserRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, errors.Wrap(err, "repository.newMongoClient")
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "repository.newMongoClient")
	}
	return client, nil
}

func (r *mongoRepository) FindAll() ([]u.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("users")

	users := []u.User{}

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user u.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *mongoRepository) Find(name string) (*u.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	user := &u.User{}
	collection := r.client.Database(r.database).Collection("users")
	filter := bson.M{"name": name}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(ce.ErrNotFound, "repository.user.Find")
		}
		return nil, errors.Wrap(err, "repository.user.Find")
	}
	return user, nil
}

func (r *mongoRepository) Post(user *u.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("users")
	_, err := collection.InsertOne(
		ctx,
		user,
	)
	if err != nil {
		return errors.Wrap(err, "repository.user.Store")
	}
	return nil
}

func (r *mongoRepository) Delete(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("users")

	filter := bson.M{"name": name}

	res, err := collection.DeleteOne(ctx, filter)

	if res.DeletedCount == 0 {
		return errors.Wrap(ce.ErrNotFound, "repository.user.Delete")
	}

	if err != nil {
		return errors.Wrap(err, "repository.user.Delete")
	}

	return nil
}
