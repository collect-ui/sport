package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 配置 CORS 中间件
	config := cors.DefaultConfig()
	//config.AllowAllOrigins = true // 允许所有来源
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"} // 添加 "client-mac" 和 "client-ip" 头
	r.Use(cors.New(config))

	// 定义根路由的 POST 处理程序
	r.POST("/", func(c *gin.Context) {
		// 获取 POST 请求的 JSON 数据
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 返回接收到的 JSON 数据
		c.JSON(http.StatusOK, gin.H{
			"received": jsonData,
		})
	})

	// 定义一个简单的路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 定义一个 POST 路由
	r.POST("/post", func(c *gin.Context) {
		// 获取 POST 请求的 JSON 数据
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 返回接收到的 JSON 数据
		c.JSON(http.StatusOK, gin.H{
			"received": jsonData,
		})
	})

	// 处理 OPTIONS 请求
	r.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*") // 添加 "client-mac" 和 "client-ip" 头
		c.Status(http.StatusNoContent)
	})

	// 启动服务器并监听 8848 端口
	r.Run(":8848")
}
