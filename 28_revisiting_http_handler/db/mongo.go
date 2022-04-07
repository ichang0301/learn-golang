package mongo

import (
	revisiting "github.com/ichang0301/learn-golang/28_revisiting_http_handler"
)

type MongoUserService struct {
}

func NewMongoUserService() *MongoUserService {
	//todo: pass in DB URL as argument to this function
	//todo: connect to db, create a connection pool
	return &MongoUserService{}
}

func (m MongoUserService) Register(user revisiting.User) (insertedID string, err error) {
	// use m.mongoConnection to perform queries
	panic("implement me")
}
