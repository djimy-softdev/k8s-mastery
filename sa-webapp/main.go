package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var SaUrl = os.Getenv("SA_URL")

type SA struct {
	Sentence string `json:"sentence" binding:"required"`
}

type SaResponse struct {
	Polarity float64 `json:"polarity"`
	Sentence string  `json:"sentence"`
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/sentiment", func(ctx *gin.Context) {
		var sa SA
		if err := ctx.ShouldBindJSON(&sa); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := getSA(SaUrl, sa)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, resp)
	})

	er := router.Run()
	if er != nil {
		log.Error("Fail to start webapp server.")
	}
}

func getSA(url string, sa SA) (interface{}, error) {
	bodyBuff, _ := json.Marshal(sa)
	resp, err := http.Post(url+"/analyse/sentiment", "application/json", bytes.NewBuffer(bodyBuff))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Internal server error " + resp.Status)
	}

	var saResp SaResponse
	err = json.NewDecoder(resp.Body).Decode(&saResp)
	if err != nil {
		return nil, err
	}

	return saResp, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
