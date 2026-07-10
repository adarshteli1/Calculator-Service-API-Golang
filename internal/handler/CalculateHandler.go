package handler

import (
	"Calculator/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
	Operation string  `json:"operation"`
}
type Response struct {
	Operand1      float64 `json:"operand1"`
	Operand2      float64 `json:"operand2"`
	Operation     string  `json:"operation"`
	Answer        float64 `json:"answer"`
	ExecutionTime string  `json:"executionTime"`
}

func CalculateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var calculate Request
		start := time.Now()

		if err := c.BindJSON(&calculate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		resp, err := service.Calculate(calculate.Operand1, calculate.Operand2, calculate.Operation)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		t := time.Now()
		executime := t.Sub(start)

		answer := Response{
			Operand1:      calculate.Operand1,
			Operand2:      calculate.Operand2,
			Operation:     calculate.Operation,
			Answer:        resp,
			ExecutionTime: executime.String(),
		}

		c.JSON(http.StatusOK, answer)
	}
}
