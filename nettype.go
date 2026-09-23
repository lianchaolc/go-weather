package main

//  天气 例子  网络请求http
// package main`：声明这是**可执行程序**，不是库文件。必须带`func main()`入口，能够`go run`编译运行
//1. **package 前面有不可见字符、中文空格、多余符号**（最常见）
//2. 代码前面有别的代码片段，`package main` 不在**文件第一行**
//> Go 强制规则：**整个 go 源码文件，第一行有效代码必须是 `package xxx`**，不能前面放别的内容
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"time"
)

// ## 两个核心函数（就是咱们天气代码里用的）

// 1. **`json.Marshal()`**：Go 结构体 → JSON 字符串（编码）
// 2. **`json.Unmarshal()`**：JSON 字符串 → Go 结构体（解码，**天气代码核心**）
// `encoding/json` 就是 Go 官方自带工具包，专门用来**处理 JSON 格式数据**

// 当前文件夹缺少 `go.mod` 文件（最主要）
// import 里面的缩进是**中文全角空格**，Go 不认

type WeatherResp struct {
	CityInfo struct {
		City string `json:"city"`
	} `json:"cityInfo"`
	Data struct {
		Forecast []struct {
			Date string `json:"date"`
			High string `json:"high"`
			Low  string `json:"low"`
			Wea  string `json:"wea"`
		} `json:"forecast"`
		Shidu string `json:"shidu"`
	} `json:"data"`
}

// 	Go 规则：

// >
// > `package main` 的程序，**必须要有 `func main(){}`**，这是程序启动入口。
// > 你现在文件里面只有 `package main` 和 import，**缺少 `func main()`，还有结构体、getCityWeather 函数**。

// ## 解决：把下面完整代码全部覆盖替换 nettype.go

// 获取城市天气函数
func getCityWeather(cityCode string) {
	client := http.Client{Timeout: 10 * time.Second}
	url := "http://t.weather.sojson.com/api/weather/city/" + cityCode


	// `resp` **需要声明**，但是在这一行同时完成了「声明 + 赋值」：

// ```
// resp, err := client.Get(url)
// ```

// `:=` 短变量声明，Go 自动创建 `resp` 和 `err` 两个变量。
// - `client.Get(url)` 函数返回两个值：`*http.Response` 和 `error`
// - `resp`：接收响应对象（类型 `*http.Response`）
// - `err`：接收错误（类型 `error`）
// - `:=`：**短变量声明**，只能在函数内部使用，自动推断类型，不需要写 `var`

// 等价的 `var` 写法（长写法，效果完全一样）

// ```
// var resp *http.Response
// var err error
// resp, err = client.Get(url)
// ```

⚠️ 区分两个符号：

1. `:=`：**声明并赋值**，变量必须是新定义的
2. `=`：**单纯赋值**，变量必须已经提前用 var 声明好
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("请求失败：%s\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}

	// json解析到结构体
	var weather WeatherResp
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("JSON解析失败：%s\n", err)
		return
	}

	// 取预报第0条=今天
	today := weather.Data.Forecast[0]
	fmt.Printf("【%s】 日期:%s 天气:%s 最高:%s 最低:%s 湿度:%s\n",
		weather.CityInfo.City, today.Date, today.Wea, today.High, today.Low, weather.Data.Shidu)
}

func main() {
	lstCityCode := []string{
		"101070101", //沈阳
		"101070201", //大连
		"101070301", //鞍山
		"101070401", //抚顺
		"101070501", //本溪
		"101070601", //丹东
		"101070701", //锦州
		"101070801", //朝阳
	}
	fmt.Println("===== 辽宁省各城市今日天气 =====")
	for _, code := range lstCityCode {
		getCityWeather(code)
		time.Sleep(500 * time.Millisecond)
	}
}
