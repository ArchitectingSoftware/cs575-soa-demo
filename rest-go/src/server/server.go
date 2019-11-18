package server

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"strconv"
	"net/http"
)

func Boot() {
	r := gin.Default()			//bootstrap web service famework

	pDB, err := PubDB()
	if err != nil {
		fmt.Println(err);
		os.Exit(0)
	}


	var pubMap = make(map[int]Pub)
	for _, pub := range pDB {
		pubMap[pub.Id] = pub
	}


	r.GET("/publications", func(c *gin.Context) {
		c.JSON(http.StatusOK, pDB)
	})

	r.GET("/publications/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, _ := strconv.Atoi(id)
		pub := pubMap[i]
		if (pub == Pub{}){
			c.JSON(http.StatusNotFound,"NOT_FOUND")
		} else {
			c.JSON(http.StatusOK, pub)
		}
	})

	r.Run("0.0.0.0:9090")
}
