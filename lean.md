创建网络请求



# Go 网络请求库、适用场景、选型 + GET / PUT 示例

## 一、Go 可用的网络请求方案

### 1. 标准库 `net/http` ✅（**内置，不用额外下载**）

Go 自带，不需要`go get`，**绝大多数场景首选**

- 优点：标准库、稳定、无第三方依赖、性能好
- 缺点：API 偏底层；Cookie、重定向、请求重试、超时需要自己写；JSON 解析要手动处理

### 2. 第三方库 `github.com/go-resty/resty/v2`（最常用简易 HTTP 客户端）

> 
> 封装 net/http，简化代码，自动处理 JSON、超时、重试、Cookie

- 优点：代码简洁，直接传结构体序列化 json，开发速度快
- 缺点：第三方依赖，简单小 demo 推荐

### 3. `github.com/valyala/fasthttp`

高性能 http 库，**适合高并发压测、网关、爬虫**

> 
> 注意：API 和标准库不一样，内存复用，上手稍复杂；普通业务不推荐

### 4. `github.com/httprunner/httprunner` 更多是测试工具，业务开发不用

---

## 二、怎么选（一句话选型）

1. **普通业务、项目不想引入额外依赖** → 用标准库 `net/http`（优先学习这个）
2. **快速开发，大量 JSON 接口、不想写重复代码** → resty
3. **超高并发、百万级请求、爬虫服务** → fasthttp

> 
> 学习阶段优先掌握标准库 `net/http`，看懂底层，再用 resty 简化开发。







# Go struct 结构体，对比 Java Object

一句话：**Go 的 `struct`（结构体）≈ Java 的 Class（实体类 / POJO），但不是完全等同。Go 没有 class，没有继承，struct 是用来自定义复合数据类型的载体。**

> 
> 不要直接理解成 `object`，Java Object 是所有类的父类；Go 没有统一根对象。

## 1. struct 作用

把**多个不同类型的变量打包在一起，组成一个新类型**。
比如用户信息：账号、密码、年龄，可以封装到一个结构体里。
表格

| 特性 | Go struct | Java Class |
| --- | --- | --- |
| 打包字段 | ✅支持 | ✅支持 |
| 方法 | ✅可以绑定方法（接收器） | ✅成员方法 |
| 继承 | ❌没有继承，用**组合**代替 | ✅支持继承 |
| 构造函数 | ❌没有内置构造函数，自己写工厂函数 | ✅构造器 |
| 所有类型父类 | ❌Go 没有统一根 Object | ✅所有类默认继承 Object |
| 指针 | ✅结构体指针 | 引用类型 |

> 
> Go 结构体可以绑定方法，写法：
> 效果类似 Java 对象调用方法，但底层原理不一样。

## 3. 两种实例：值类型 vs 指针结构体（重点）

```
// 值类型：拷贝一份数据
u1 := User{}
// 指针：引用，操作同一块内存（类似Java对象引用）
u2 := &User{}
```

- `User{}`：值类型，赋值的时候会完整复制所有字段
- `&User{}`：结构体指针，不会拷贝，函数传参推荐用指针，节省内存

## 4. 常用场景

1. 数据库映射：SQLite/MySQL 查询出来的数据，映射到 struct（对应 Java Entity）

```
type LoginRecord struct {
	Id       int
	Username string
	LoginAt  time.Time
}
```

2. HTTP JSON 序列化：`json.Unmarshal` 把接口返回 JSON 转成 struct

```
type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
```

> 
> `json:"code"` 这个叫标签 tag，struct 的附加元信息，反射的时候读取。

## 5. 关键区别：不是 Object

Java：所有类都继承 `Object`，都有 `equals()`、`toString()`。
Go：**没有统一父类**。`struct` 只是自定义类型。空结构体 `struct{}` 也不是 Object。

## 一句话总结

Go 的 `struct` = **可以包含字段、绑定方法的自定义复合数据类型**，功能对标 Java 的实体类，**但是没有 class、没有继承，不是 Object**。

如果你想，我写一段：**struct + json 解析 + sqlite 查询** 的完整小例子，刚好衔接前面 sqlite 的代码。
# Go 结构体指针 & 值类型 详细展开

核心一句话：

- `User{}` → **结构体值类型**：赋值 / 传参时，完整复制一份结构体所有字段数据
- `&User{}` → **结构体指针（*User）**：只复制内存地址，指向**同一份原始结构体数据**，类似 Java 对象引用

> 
> 类型区分：
> `User`：结构体值类型
> `*User`：结构体指针类型

## 1. 代码直观对比，看两者差异

```
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 值接收者：传入的是副本，修改不会影响外面原始变量
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收者：传入地址，直接修改原始结构体
func (u *User) ChangeNamePtr(newName string) {
	u.Name = newName
}

func main() {
	// 1. 值类型实例
	u1 := User{Name: "张三", Age: 20}
	u1.ChangeName("李四")
	fmt.Println("u1 值类型修改后：", u1) // 仍然是张三！修改的是副本

	// 2. 指针实例
	u2 := &User{Name: "张三", Age: 20}
	u2.ChangeNamePtr("李四")
	fmt.Println("u2 指针修改后：", u2) // 变成李四，修改原始数据
}
```

输出：

```
u1 值类型修改后： {张三 20}
u2 指针修改后： &{李四 20}
```
`json:"city"` 不是**指定类型**，这个叫 **struct tag（结构体标签）**，是附加在字段上的**元信息**，专门给 `encoding/json` 包反射解析用的。

> 
> 你写的两段分开看：

```
City string `json:"city"`
```

和

```
CityInfo struct { ... } `json:"cityInfo"`
```

## 1. 含义拆解

```
type Area struct {
    City string `json:"city"`
}
```

- `City`：Go 结构体字段名（Go 代码里用 `area.City` 访问）
- `string`：**字段本身的类型**（这个才是类型！）
- `\`json:"city"\``：标签
  - 含义：JSON 序列化 / 反序列化时，**JSON 里面 key 名字叫 `city`**
  - 把 JSON 的`"city"`的值映射到 Go 的`City`字段

### 嵌套结构体例子，对应你提到的 `json:"cityInfo"`







# Windows

## VS Code（写 Go 推荐）

- **Shift + Alt + F**：格式化整个文档 ✅
- **Ctrl + K 再按 Ctrl + F**：只格式化选中一段代码

> 
> 配置：安装 Go 官方插件，开启 `editor.formatOnSave`，按 `Ctrl+S` 自动格式化

## GoLand / IDEA

- **Ctrl + Alt + L**：格式化代码（Go 最常用）JetBrains

# Mac

## VS Code

- **Option + Shift + F** 格式化文档

## GoLand

- **Cmd + Option + L**

# 命令行（任何编辑器都能用，原生 go 工具）

```
gofmt -w weather.go
```

`-w` 直接覆盖文件，自动修正缩进、大括号（Go 强制格式）






# 拆解 resp（*http.Response）里面的内容

`resp` 是 `*http.Response` 指针结构体，`client.Get()` 请求完之后，所有返回信息都装在 `resp` 里面。

```
resp, err := client.Get(url)
```

先看结构体定义（简化版）

```
type Response struct {
	Status     string // 状态文本，例如 "200 OK"
	StatusCode int    // 状态码，200成功 / 404找不到 /500错误
	Header     http.Header // 响应头，字典格式
	Body       io.ReadCloser // 响应体，就是我们的JSON数据，必须Close
	// ...其他字段
}
```
resp, err := client.Get(url)
if err != nil {
	fmt.Printf("请求失败：%s\n", err)
	return
}
defer resp.Body.Close() // Body一定要关闭

// 1. HTTP状态码，最常用
fmt.Println("状态码：", resp.StatusCode)
fmt.Println("状态文本：", resp.Status)

// 2. 读取响应头 Header
fmt.Println("服务器：", resp.Header.Get("Server"))
fmt.Println("内容类型：", resp.Header.Get("Content-Type"))

// 3. 读取Body（接口返回的JSON字符串，重点！）
bodyBytes, err := io.ReadAll(resp.Body)
if err != nil {
	fmt.Println("读取body失败", err)
	return
}
fmt.Println("原始返回字符串：", string(bodyBytes))

// 4. 把字节数组bodyBytes 解析到结构体WeatherResp
var weather WeatherResp
err = json.Unmarshal(bodyBytes, &weather)




html 的页面如何展示

场景 A：Go 程序**请求网页拿到 HTML 源码**（就像刚才拿天气 JSON 一样）
场景 B：Go 写一个 Web 服务，**把 HTML 页面返回给浏览器展示**


package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w 就是写给浏览器的输出流
	html := `
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>Go返回HTML页面</title>
	</head>
	<body>
		<h1>天气查询页面</h1>
		<p>辽宁朝阳天气</p>
	</body>
	</html>
	`
	// 设置响应头，告诉浏览器返回的是html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 把html字符串写给浏览器
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("服务启动，访问：http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}



