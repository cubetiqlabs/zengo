package app

import (
	"fmt"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cubetiq/zengo/config"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		t := time.Now().Unix()
		c.JSON(http.StatusOK, gin.H{
			"timestamp": t,
			"msg":       "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"cubetiq": "cubetiq",
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

type App struct {
	Router *gin.Engine
}

func (a *App) Run(config *config.Config) error {
	r := setupRouter()
	a.Router = r
	return a.Router.Run(fmt.Sprintf("%s:%d", config.App.Addr, config.App.Port))
}
