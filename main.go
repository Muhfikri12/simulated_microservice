package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	secured := r.Group("/payment")
	{
		secured.POST("/", CreatePayment)
		secured.GET("/", ListPayment)
		secured.PUT("/:id", UpdateStatus)
	}

	r.Run(":8088")
}

type InputRequest struct {
	Amount  float64 `json:"amount" binding:"required"`
	Payment string  `json:"payment" binding:"required"`
	OrderID int     `json:"order_id" binding:"required"`
	Status  string  `json:"status"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildResponse(status int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func CreatePayment(c *gin.Context) {

	var input InputRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, BuildResponse(
			http.StatusBadRequest,
			"Invalid input data",
			gin.H{"error": err.Error()},
		))
		return
	}

	status := "pending"

	response := BuildResponse(
		http.StatusOK,
		"Payment created successfully",
		gin.H{
			"order_id": input.OrderID,
			"amount":   input.Amount,
			"payment":  input.Payment,
			"status":   status,
		},
	)

	c.JSON(http.StatusOK, response)

}

func ListPayment(c *gin.Context) {

	dummyData := []InputRequest{
		{Amount: 150000, Payment: "credit_card", OrderID: 1, Status: "Pending"},
		{Amount: 250000, Payment: "bank_transfer", OrderID: 2, Status: "Success"},
		{Amount: 100000, Payment: "e_wallet", OrderID: 3, Status: "success"},
	}

	c.JSON(http.StatusOK, BuildResponse(
		http.StatusOK,
		"Successfully retrieved payments",
		dummyData,
	))
}

func UpdateStatus(c *gin.Context) {

	id := c.Param("id")

	status := "Success"

	response := BuildResponse(
		http.StatusOK,
		"Status updated successfully",
		gin.H{
			"id":     id,
			"status": status,
		},
	)

	c.JSON(http.StatusOK, response)
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token == "" {
// 			c.JSON(http.StatusUnauthorized, BuildResponse(
// 				http.StatusUnauthorized,
// 				"Unauthorized: Missing token",
// 				nil,
// 			))
// 			c.Abort()
// 			return
// 		}
//
// 		req, err := http.NewRequest("POST", "http://localhost:8081/authentication", nil)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, BuildResponse(
// 				http.StatusInternalServerError,
// 				"Internal server error",
// 				nil,
// 			))
// 			c.Abort()
// 			return
// 		}
//
// 		req.Header.Set("Authorization", token)
// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil || resp.StatusCode != http.StatusOK {
// 			c.JSON(http.StatusUnauthorized, BuildResponse(
// 				http.StatusUnauthorized,
// 				"Unauthorized: Invalid token",
// 				nil,
// 			))
// 			c.Abort()
// 			return
// 		}
//
// 		c.Next()
// 	}
// }
