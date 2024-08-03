package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Cors())

	router.GET("/", func(c *gin.Context) {
		https := c.Query("https")
		srv := c.Query("srv")
		srvs := strings.Split(srv, ".")
		if len(srvs) <= 3 {
			c.String(http.StatusInternalServerError, "srv is required, and it's format must as _sip._tcp.example.com.")
		}
		service, _ := strings.CutPrefix(srvs[0], "_")
		protocol, _ := strings.CutPrefix(srvs[1], "_")
		new, err := resolve(service, protocol, strings.Join(srvs[2:], "."))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		if https == "" {
			https = "http"
		}
		if b, e := strconv.ParseBool(https); b && e != nil {
			https = "https"
		}
		c.Redirect(http.StatusMovedPermanently, https+"://"+new)
	})

	router.GET("/help", func(c *gin.Context) {
		c.String(http.StatusOK, `
本服务用于解析并301跳转srv dns纪录。
用法:
  - 浏览器打开 /srv/${SRV_RECORD}  =>  301 跳转到 http://${SRV_RECORD}
  - 浏览器打开 /srv/${SRV_RECORD}?https=true  => 301 跳转到 https://${SRV_RECORD}

This service is used to parse and 301 redirect srv DNS records.
Usage: 
  - Open /srv/${SRV_RECORD} in a browser => 301 redirect to http://${SRV_RECORD}
  - Open /srv/${SRV_RECORD}?https=true in a browser => 301 redirect to https://${SRV_RECORD}

about dns srv record：
  - https://zh.wikipedia.org/zh-hans/Template:SRV
  - https://baike.baidu.com/item/SRV%E8%AE%B0%E5%BD%95/10637211
`)
	})

	// Run the server on port 8080
	router.Run(":8080")
}

func resolve(service string, protocol string, domain string) (string, error) {
	_, srvs, err := net.LookupSRV(service, protocol, domain)
	if err != nil {
		return "", err
	}
	srv := srvs[0]

	fmt.Println(srv)
	host, _ := strings.CutSuffix(srv.Target, ".")
	return fmt.Sprintf("%s:%d", host, srv.Port), nil
}
