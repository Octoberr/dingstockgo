package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoE(idata []interface{}) {
	mongourl := ""
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongourl))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("haitao_data").Collection("ht_product")
	res2, err := collection.InsertMany(ctx, idata)
	if res2 != nil {
		fmt.Printf("Inserted %v documents into episode collection!\n", len(res2.InsertedIDs))
	}
	//res, err := collection.InsertOne(ctx, idata)
	//if res != nil {
	//	id := res.InsertedID
	//	fmt.Println(id)
	//}
	//res2, err := collection.InsertOne(ctx, bson.D{
	//	{Key: "title", Value: "The Polyglot Developer Podcast"},
	//	{Key: "author", Value: "Nic Raboy"},
	//	{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	//})
	//if res2 != nil{
	//	id2 := res2.InsertedID
	//	fmt.Println(id2)
	//}
	//episodeResult, err := collection.InsertMany(ctx, []interface{}{
	//	bson.D{
	//		{"podcast", res.InsertedID},
	//		{"title", "GraphQL for API Development"},
	//		{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
	//		{"duration", 25},
	//	},
	//	bson.D{
	//		{"podcast", res2.InsertedID},
	//		{"title", "Progressive Web Application Development"},
	//		{"description", "Learn about PWA development with Tara Manicsic."},
	//		{"duration", 32},
	//	},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))
}
