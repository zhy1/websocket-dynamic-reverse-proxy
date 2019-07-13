package custom

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"websocket-dynamic-reverse-proxy/proxy"
)

type ProxyConfigObject struct {
	ListenerPath string `json:"listenerPath"`
	UpstreamUrl  string `json:"upstream"`
	Fragment     string `json:"fragment"`
	RawQuery     string `json:"rawQuery"`
	Path         string `json:"path"`
}

//
// @Summary AddProxy
// @Description add
// @Accept  application/json
// @Produce  application/json
// @Failure 400 {string}  status:bad    "parse json failed"
// @Success 200 {string}  success:true
// @Router /add [POST]
func AddProxy(context *gin.Context) {
	var proxyConfig ProxyConfigObject
	byteData, _ := context.GetRawData()
	err := json.Unmarshal(byteData, &proxyConfig)
	if err != nil {
		context.JSON(400, map[string]interface{}{
			"success":        false,
			"ProxyConfigMap": proxy.ProxyConfigMap,
			"err":            err.Error(),
		})
		return
	}
	proxy.ProxyConfigMap[proxyConfig.ListenerPath] = proxyConfig.UpstreamUrl
	context.JSON(200, map[string]interface{}{
		"success":        true,
		"ProxyConfigMap": proxy.ProxyConfigMap,
	})
	return
}

//
// @Summary HealthCheckOrigin
// @Description HealthCheckOrigin
// @Produce  application/json
// @Success 200 {string}  success:true
// @Router / [GET]
func HealthCheckOrigin(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"success":        true,
		"ProxyConfigMap": proxy.ProxyConfigMap,
	})
	return
}

func BootstrapHTTP() {
	router := gin.Default()
	corsDefault := cors.DefaultConfig()
	corsDefault.AllowCredentials = true
	corsDefault.AllowHeaders = []string{"Origin", "token", "Content-Length", "Content-Type", "session", "DNT", "content-type", "s", "timezone", "tz", "specify", "order"}
	corsDefault.AllowAllOrigins = true
	router.Use(cors.New(corsDefault))
	router.POST("/add", AddProxy)
	router.GET("/", HealthCheckOrigin)
	addr := "0.0.0.0:6661"
	err := router.Run(addr)
	if err != nil {
		fmt.Println("error :", err)
	}
}
