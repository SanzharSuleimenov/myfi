package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type payment struct {
	ID            string    `json:"id"`
	Source        string    `json:"source"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Date          time.Time `json:"date"`
	PaymentMethod string    `json:"payment_method"`
}

// slice, not an array. will serve as a database
var payments = []payment{
	{
		ID:            "1",
		Source:        "HSBC",
		Amount:        36.7,
		Currency:      "AED",
		Date:          time.Now(),
		PaymentMethod: "Card",
	},
	{
		ID:            "2",
		Source:        "HSBC",
		Amount:        12.95,
		Currency:      "AED",
		Date:          time.Now(),
		PaymentMethod: "Card",
	},
}

func main() {
	// Default Gin Engine instance, that's used to handle requests and responses
	var router *gin.Engine = gin.Default()
	// request handlers
	router.GET("/payment", getPayments)
	router.POST("/payment", postPayments)
	router.GET("/payment/:id", getPaymentById)
	router.GET("/payment/total", getTotal)
	// listen at localhost:8080
	router.Run(":8080")
}

func getPayments(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, payments)
}

func postPayments(context *gin.Context) {
	var newPayment payment

	if err := context.BindJSON(&newPayment); err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	payments = append(payments, newPayment)
	context.IndentedJSON(http.StatusCreated, newPayment)
}

func getPaymentById(context *gin.Context) {
	var id string = context.Param("id")
	for _, p := range payments {
		if p.ID == id {
			context.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "pament not found",
	})
}

func getTotal(context *gin.Context) {
	var total float64
	for _, p := range payments {
		total += p.Amount
	}
	context.IndentedJSON(http.StatusOK, gin.H{
		"total": total,
	})
}
