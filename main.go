package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	templateService "github.com/collect-ui/collect/src/collect/service_imp"
	collect "github.com/collect-ui/collect/src/collect/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"moon/model"
	"moon/plugins"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"strings"
	"time"
)

// HmacWithShaTobase64 生成 HMAC-SHA256 签名并编码为 Base64
func HmacWithShaTobase64(message string, sign string, secret string) string {
	// 创建一个新的 HMAC 实例，使用 SHA256 算法
	h := hmac.New(sha256.New, []byte(secret))

	// 将消息和签名拼接在一起
	combinedMessage := message + sign

	// 将拼接后的消息写入 HMAC 实例
	h.Write([]byte(combinedMessage))

	// 计算 HMAC 签名
	signature := h.Sum(nil)

	// 将签名编码为 Base64
	base64Signature := base64.StdEncoding.EncodeToString(signature)

	return base64Signature
}

// @hosturl :  like  wss://tts-api.xfyun.cn/v2/tts
// @apikey : apiKey
// @apiSecret : apiSecret
func assembleAuthUrl(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("api_key=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))
	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}
func main1() {
	fmt.Println(assembleAuthUrl("wss://tts-api.xfyun.cn/v2/tts", "128a773f8ae545652713e4c0e3b3dba9", "MDU5ODExZDJjNzI4ODRkOGU2ZDgyM2Ji"))
}
func serveStatic(urlPrefix, root string, cache bool) gin.HandlerFunc {
	fs := http.Dir(root)
	//fileServer := http.FileServer(fs)

	return func(c *gin.Context) {
		p := c.Request.URL.Path
		if cache {
			c.Header("Cache-Control", "public, max-age=5184000") // 60天 = 60 * 24 * 60 * 60 秒
			// 设置Expires头，缓存60天
			expires := time.Now().Add(60 * 24 * time.Hour)
			c.Header("Expires", expires.Format(http.TimeFormat))
		}
		// 设置Cache-Control头，缓存60天

		if !strings.HasPrefix(p, urlPrefix) {
			c.Next()
			return
		}

		// 去掉 URL 前缀
		p = strings.TrimPrefix(p, urlPrefix)

		// 尝试查找文件
		f, err := fs.Open(p)
		if err != nil {
			// 如果文件不存在，尝试查找 index.html
			f, err = fs.Open("index.html")
			if err != nil {
				c.Next()
				return
			}
		}
		defer f.Close()

		info, err := f.Stat()
		if err != nil || info.IsDir() {
			// 如果文件是目录或无法获取文件信息，尝试查找 index.html
			f, err = fs.Open("index.html")
			if err != nil {
				c.Next()
				return
			}
			defer f.Close()
			info, _ = f.Stat()
		}

		// 提供文件服务
		http.ServeContent(c.Writer, c.Request, info.Name(), info.ModTime(), f)
		c.Abort()
	}
}

func main() {
	// todo go profile 使用
	//gin.SetMode(gin.ReleaseMode)
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	// 全局设置跨域头
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// 生成cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.Static("/static", "./static")

	//r.Use(serveStatic("/data", "./frontend/data", false))
	//r.Use(serveStatic("/files", "./file_data/files", true))
	//r.Use(serveStatic("/collect-ui", "./frontend/collect-ui", false))
	//r.Use(serveStatic("/dashboard", "./frontend/dashboard", false))
	//r.Use(serveStatic("/training", "./frontend/dashboard-training", false))
	dirStr := collect.GetAppKey("dirList")
	dirList := strings.Split(dirStr, ";")
	for _, file := range dirList {
		if collect.IsValueEmpty(file) {
			continue
		}
		fileInfo := strings.Split(file, ",")
		cache := false
		if fileInfo[2] == "true" {
			cache = true
		}
		r.Use(serveStatic(fileInfo[0], fileInfo[1], cache))
	}

	fileStr := collect.GetAppKey("fileList")
	fileList := strings.Split(fileStr, ";")
	for _, file := range fileList {
		if collect.IsValueEmpty(file) {
			continue
		}
		fileInfo := strings.Split(file, ",")
		r.StaticFile(fileInfo[0], fileInfo[1])
	}
	// 设置数据库
	templateService.SetDatabaseModel(&model.TableData{})
	// 设置外部处理器
	templateService.SetOuterModuleRegister(plugins.GetRegisterList())
	// 添加定时任务
	templateService.RunScheduleService()
	// 添加启动服务
	templateService.RunStartupService()
	//r.GET("/", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/ssh")
	//})
	r.POST("/template_data/data", func(c *gin.Context) {
		templateService.HandlerRequest(c)
	})

	r.GET("/template_data/ws/:token", func(context *gin.Context) {
		templateService.HandlerWsRequest(context)
	})

	serverPort := collect.GetAppKey("server_port")
	isHttps := collect.GetAppKey("is_https")
	domains := strings.Split(collect.GetAppKey("domain"), ",")

	if isHttps == "true" {
		tlsConfig := &tls.Config{}
		for _, domain := range domains {
			certFile := "/etc/letsencrypt/live/" + domain + "/fullchain.pem"
			keyFile := "/etc/letsencrypt/live/" + domain + "/privkey.pem"
			cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			if err != nil {
				panic(err)
			}
			tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
		}
		server := &http.Server{
			Addr:      ":" + serverPort,
			Handler:   r,
			TLSConfig: tlsConfig,
		}
		if err := server.ListenAndServeTLS("", ""); err != nil {
			panic(err)
		}
	} else {
		r.Run(":" + serverPort) // listen and serve on 0.0.0.0:8080
	}
}
