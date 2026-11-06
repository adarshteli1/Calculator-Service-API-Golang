package handler

import (
	"Calculator/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Operand1  *float64 `json:"operand1" binding:"required"`
	Operand2  *float64 `json:"operand2" binding:"required"`
	Operation string  `json:"operation" binding:"required"`
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

		if err := c.ShouldBindJSON(&calculate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		resp, err := service.Calculate(*calculate.Operand1, *calculate.Operand2, calculate.Operation)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		t := time.Now()
		executime := t.Sub(start)  // Or could have written as executiontime:=time.Now().Sub(start)

		answer := Response{
			Operand1:      *calculate.Operand1,
			Operand2:      *calculate.Operand2,
			Operation:     calculate.Operation,
			Answer:        resp,
			ExecutionTime: executime.String(),
		}

		c.JSON(http.StatusOK, answer)
	}
}
