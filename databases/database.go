package databases

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() error {
	var err error
	mdb, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_addr))
	if err != nil {
		return err
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		mdb.Disconnect(context.TODO())
		return err
	}
	return nil
}

func Close() error {
	mdb.Disconnect(context.TODO())
	rdb.Close()
	return nil
}

func GetIndex(opt string) (int, int) {
	vl := fmt.Sprintf("%s.latest", opt)
	vs := fmt.Sprintf("%s.sum", opt)
	l, err := rdb.Get(context.TODO(), vl).Int()
	if err != nil {
		panic(err)
	}
	r, err := rdb.Get(context.TODO(), vs).Int()
	if err != nil {
		panic(err)
	}
	return l, r
}

func GetDesc(opt string, page int) []bson.M {
	c := mdb.Database(descs).Collection(opt)
	var option options.FindOptions
	option.SetProjection(pro)
	option.SetSort(sor)
	cur, err := c.Find(context.TODO(), bson.D{{"author", opt}, {"page", page}}, &option)
	if err != nil {
		panic(err)
	}
	var res []bson.M
	err = cur.All(context.TODO(), &res)
	if err != nil {
		panic(err)
	}
	return res
}

func GetContent(opt string, hash string) bson.M {

	c := mdb.Database(contents).Collection(opt)

	var option options.FindOneOptions
	var res bson.M
	{
		option.SetProjection(pro)
		err := c.FindOne(context.TODO(), bson.D{{"sha256", hash}}, &option).Decode(&res)
		if err != nil {
			panic(err)
		}
	}

	return res
}

func GetLatest() []bson.M {

	db := mdb.Database(descs)
	s, _ := db.ListCollectionNames(context.TODO(), bson.D{})
	c := db.Collection(s[0])

	var res []bson.M
	var cur *mongo.Cursor
	var err error
	if len(s) == 1 {
		var option options.FindOptions
		option.SetProjection(pro)
		option.SetSort(sor)
		option.SetLimit(lim)
		cur, err = c.Find(context.TODO(), bson.D{}, &option)
		if err != nil {
			panic(err)
		}
	} else {
		var stage mongo.Pipeline
		for i := 1; i < len(s); i++ {
			stage = append(stage, bson.D{{"$unionWith", s[i]}})
		}
		stage = append(stage, bson.D{{"$sort", sor}})
		stage = append(stage, bson.D{{"$limit", lim}})
		cur, err = c.Aggregate(context.TODO(), stage)
		if err != nil {
			panic(err)
		}
	}

	err = cur.All(context.TODO(), &res)

	if err != nil {
		panic(err)
	}

	return res
}
