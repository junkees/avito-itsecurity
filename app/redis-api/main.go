package main

import (
	"app/redisClient"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type PostParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

	ctx := context.Background()

	r := gin.Default()

	r.GET("/get_key", getKey)
	r.POST("/set_key", setKey)
	r.POST("/del_key", delKey)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(403, gin.H{
			"error": "Forbidden",
		})
	})

	//r.RunTLS(":"+os.Getenv("PORT"), os.Getenv("clientCertFile"), os.Getenv("clientKeyFile"))
	r.Run(":" + os.Getenv("PORT"))

	red := redisClient.GetConnection()
	pong, err := red.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pong)
}

func getKey(c *gin.Context) {

	key := c.Query("key")
	red := redisClient.GetConnection()

	value, _ := red.Get(c, key).Result()

	if len(value) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "key not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key": value,
	})
	return
}

func setKey(c *gin.Context) {
	var params map[string]string
	err := c.ShouldBindJSON(&params)
	if err != nil {
		return
	}

	red := redisClient.GetConnection()

	for key, value := range params {
		if len(key) == 0 || len(value) == 0 {
			c.JSON(403, gin.H{
				"error": "Incorrect params!",
			})
			return
		}
		result := red.Set(c, key, value, time.Hour*24)

		fmt.Println(result)
	}

	c.JSON(200, gin.H{
		"status": "Key set!",
	})
	return
}

func delKey(c *gin.Context) {
	var params map[string]string

	err := c.ShouldBindJSON(&params)
	if err != nil {
		return
	}

	red := redisClient.GetConnection()
	delkey := params["key"]

	if len(delkey) == 0 {
		return
	}

	res, _ := red.Del(c, delkey).Result()
	if res == 0 {
		c.JSON(404, gin.H{
			"status": "Key not found!",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "Key deleted!",
	})
	return
}
