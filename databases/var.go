package databases

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	mongo_addr = "mongodb://127.0.0.1:27017/"
	redis_addr = "127.0.0.1:6379"
	website    = "website"
	descs      = "descs"
	contents   = "contents"
	lim        = 10
)

var mdb *mongo.Client
var rdb *redis.Client

var pro = bson.D{{"_id", 0}}
var sor = bson.D{{"year", -1}, {"month", -1}, {"day", -1}}
