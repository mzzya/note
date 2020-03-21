package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var cli *mongo.Client

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	cli = client
}

func main() {
	db := cli.Database("config_center")

	mongoTest()
	jsonTest()
}
func saveSchema() {

}

func jsonTest() {
	bts, err := ioutil.ReadFile("./schema.json")
	if err != nil {
		panic(err)
	}
	value := gjson.GetBytes(bts, "type")
	println(value.String())

	value2, _ := jsonparser.GetString(bts, "type")
	println(value2)
}

func mongoTest() {
	db := cli.Database("config_center")
	couponCollection := db.Collection("coupon", &options.CollectionOptions{})
	data := make([]interface{}, 0)
	result, err := couponCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		panic(errors.WithMessage(err, "DeleteMany"))
	}
	fmt.Printf("%d\n", result.DeletedCount)
	for i := 0; i < 1000000; i++ {
		data = append(data, bson.M{"batch_id": i, "qty": i, "status": i%2 == 0, "create_time": time.Now(), "update_time": time.Now()})
	}

	now := time.Now()
	resutl, err := couponCollection.InsertMany(context.Background(), data)
	if err != nil {
		panic(errors.WithMessage(err, "InsertMany"))
	}
	t := time.Since(now)
	fmt.Printf("time:%d毫秒\tcount:%d\n", t.Milliseconds(), len(resutl.InsertedIDs))
	// result, err := collection.InsertOne(context.Background(), bson.M{"key": "pi", "value": 3.14159})
	// if err != nil {
	// 	panic(err)
	// }
	// r, err := collection.UpdateMany(context.Background(), bson.M{"key": "pi"}, bson.M{
	// 	// "$set": bson.M{
	// 	// 	"name": nil,
	// 	// 	"key":  "pi",
	// 	// },
	// 	"$rename": bson.M{
	// 		"name": "name2",
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// bts, _ := json.Marshal(collection.Indexes())
	// println(result, r, string(bts))
}
