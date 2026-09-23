package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// 和之前天气代码一样的结构体
type WeatherResp struct {
	CityInfo struct {
		City string `json:"city"`
	} `json:"cityInfo"`
	Data struct {
		Forecast []struct {
			Date string `json:"date"`
			High string `json:"high"`
			Low  string `json:"low"`
			Type string `json:"type"`
		} `json:"forecast"`
	} `json:"data"`
}

// 判断是否下雨
func isRaining(weatherType string) bool {
	rainKeywords := []string{"雨", "阵雨", "雷阵雨", "暴雨", "小雨", "中雨", "大雨"}
	for _, keyword := range rainKeywords {
		if strings.Contains(weatherType, keyword) {
			return true
		}
	}
	return false
}

// 全局单例client，底层http.Transport连接池配置
func newRestyClient() *resty.Client {
	// 原生net/http的Transport，控制长连接池
	transport := &http.Transport{
		MaxIdleConns:        200,
		MaxIdleConnsPerHost: 50, // 单个域名最大空闲长连接
		IdleConnTimeout:     90 * time.Second,
	}

	client := resty.New()
	client.SetTransport(transport) // 把transport注入resty

	// 超时
	client.SetTimeout(15 * time.Second)
	// 自动重试，网络抖动重试3次
	client.SetRetryCount(3)
	client.SetRetryWaitTime(1 * time.Second)
	client.SetRetryMaxWaitTime(3 * time.Second)

	return client
}

func main() {
	client := newRestyClient()

	// 辽宁朝阳天气接口
	url := "http://t.weather.itboy.net/api/weather/city/101071201"

	var weather WeatherResp
	// 链式调用 SetResult自动解析json
	resp, err := client.R().
		SetResult(&weather).
		SetHeader("User-Agent", "Mozilla/5.0").
		SetHeader("Cache-Control", "no-cache, no-store, must-revalidate").
		SetHeader("Pragma", "no-cache").
		// 增加随机时间戳，URL每次不一样，彻底绕过CDN缓存
		SetQueryParam("_t", fmt.Sprintf("%d", time.Now().UnixMilli())).
		Get(url)

	if err != nil {
		fmt.Printf("请求发生错误：%v\n", err)
		return
	}

	if resp.StatusCode() != 200 {
		fmt.Printf("接口状态码异常：%d\n", resp.StatusCode())
		return
	}

	fmt.Printf("城市：%s\n", weather.CityInfo.City)
	fmt.Println("========================================")
	for _, day := range weather.Data.Forecast {
		rainTip := "✅ 无雨"
		if isRaining(day.Type) {
			rainTip = "🌧️ 有雨，记得带伞"
		}
		fmt.Printf("日期:%s 天气:%s 温度:%s ~ %s | %s\n", day.Date, day.Type, day.Low, day.High, rainTip)
	}

	fmt.Printf("\n本次请求耗时：%s\n", resp.Time())
	fmt.Printf("HTTP状态：%s\n", resp.Status())
}
