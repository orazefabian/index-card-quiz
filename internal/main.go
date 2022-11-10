package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type answer struct {
	ID         string
	QuestionID string
	Text       string
	IsTrue     bool
}

type indexCard struct {
	ID       string
	Question string
	Answers  []answer
}

var cards = []indexCard{
	{ID: "1", Question: "Test question", Answers: []answer{{ID: "1", QuestionID: "1", Text: "first answer", IsTrue: true}, {ID: "2", QuestionID: "1", Text: "second answer", IsTrue: false}}},
	{ID: "2", Question: "Test question two", Answers: []answer{{ID: "3", QuestionID: "2", Text: "new first answer", IsTrue: true}, {ID: "4", QuestionID: "2", Text: "new second answer", IsTrue: true}}},
}

func getCards(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, cards)
}

func main() {
	router := gin.Default()
	router.GET("/cards", getCards)
	router.Run("localhost:9999")
}
