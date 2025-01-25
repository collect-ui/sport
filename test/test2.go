package main

import (
	"github.com/gin-gonic/gin"
)

func main2() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, HTTPS!")
	})
	domain := "sport.lassion.com"
	cert := "/etc/letsencrypt/live/" + domain + "/fullchain.pem"
	key := "/etc/letsencrypt/live/" + domain + "/privkey.pem"
	if err := r.RunTLS(":8011", cert, key); err != nil {
		panic(err)
	}
}
