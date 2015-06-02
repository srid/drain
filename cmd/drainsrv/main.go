package main

import (
	"bufio"
	"fmt"
	"github.com/bmizerany/lpx"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewGinRouter() gin.Accounts {
	return gin.Accounts{
		config.Username: config.Password,
	}
}

func main() {
	var routerGroup *gin.RouterGroup
	router := gin.Default()

	username, password, authorize = config.GetUserPass()
	if authorize {
		routerGroup = router.Group("/",
			gin.BasicAuth(gin.Accounts{username: password}))
	} else {
		routerGroup = router
	}

	routerGroup := router.Group("/", gin.BasicAuth(makeGinAccounts()))
	routerGroup.POST("/logs", logsReceived)
	log.Printf("Running drain server at port %v\b", config.Port)
	router.Run(":" + config.Port)
}

func logsReceived(c *gin.Context) {
	defer c.Request.Body.Close()
	handleLog(c.Request)
	c.JSON(200, gin.H{})
}

func handleLog(r *http.Request) {
	lp := lpx.NewReader(bufio.NewReader(r.Body))

	for lp.Next() {
		header := lp.Header()
		data := lp.Bytes()
		line := fmt.Sprintf("LOG=> %s %s %s %s %s %s %s\n",
			string(header.PrivalVersion),
			string(header.Time),
			string(header.Hostname),
			string(header.Name),
			string(header.Procid),
			string(header.Msgid),
			string(data))

		fmt.Printf("%s", line)
	}
}
