package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	Setup code, DAOs, sample data, TODO: extract code
*/

type answer struct {
	ID         string `bson:"id,omitempty"`
	QuestionID string `bson:"questionID,omitempty"`
	Text       string `bson:"text,omitempty"`
	IsTrue     bool   `bson:"isTrue,omitempty"`
}

type indexCard struct {
	ID       string   `bson:"id,omitempty"`
	Question string   `bson:"question,omitempty"`
	Answers  []answer `bson:"answers,omitempty"`
}

var cards = []indexCard{
	{ID: "1", Question: "Test question", Answers: []answer{{ID: "1", QuestionID: "1", Text: "first answer", IsTrue: true}, {ID: "2", QuestionID: "1", Text: "second answer", IsTrue: false}}},
	{ID: "2", Question: "Test question two", Answers: []answer{{ID: "3", QuestionID: "2", Text: "new first answer", IsTrue: true}, {ID: "4", QuestionID: "2", Text: "new second answer", IsTrue: true}}},
}

/*
	Server functionality, TODO: extract code
*/

func getCards(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, cards)
}

func main() {
	/*
	   Database connection, TODO: extract code
	*/
	var dbURI = "mongodb://admin:password@0.0.0.0:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// // begin insertOne and create testDB database
	coll := client.Database("testDB").Collection("quizCards")

	result, err := coll.InsertOne(ctx, bson.D{{Key: "id", Value: "1"}, {Key: "question", Value: "Test question"}, {Key: "answers", Value: "test answers"}})

	// // end insertOne

	// // When you run this file, it should print:
	// // Document inserted with ID: ObjectID("...")
	fmt.Printf("Document inserted with ID: %s\n", result)

	router := gin.Default()
	router.GET("/cards", getCards)
	router.Run("0.0.0.0:9999")
}
