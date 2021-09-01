package mongo

import (
	"context"
	"go_news/pkg/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client *mongo.Client
}

const (
	dbName         = "news"  // имя БД
	collectionName = "posts" // имя коллекции в БД
)

// New - конструктор, conn - строка подключения к БД
func New(conn string) (*Storage, error) {
	// подключение к СУБД Mongo
	mongoOpts := options.Client().ApplyURI(conn)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	s := Storage{
		Client: client,
	}

	return &s, nil
}

// Close - закрывает соединение к базе данных
func (s *Storage) Close() {
	s.Client.Disconnect(context.Background())
}

// Posts - получает список всех публикаций из БД
func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.Client.Database(dbName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var p storage.Post
		err := cur.Decode(&p)

		if err != nil {
			return nil, err
		}

		data = append(data, p)
	}

	return data, cur.Err()
}

// AddPost - добавляет публикацию в БД
func (s *Storage) AddPost(post storage.Post) error {
	collection := s.Client.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), post)

	if err != nil {
		return err
	}
	return nil
}

// UpdatePost - обновляет публикацию в БД
func (s *Storage) UpdatePost(post storage.Post) error {
	collection := s.Client.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": post.ID}
	update := bson.M{"$set": post}
	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

// DeletePost - удаляет публикацию из БД
func (s *Storage) DeletePost(post storage.Post) error {
	collection := s.Client.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": post.ID}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}
	return nil
}
