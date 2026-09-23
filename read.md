# ✅ 核心结论

**推荐：每一个独立 Go 程序（可执行程序，`package main`），单独新建一个文件夹。**
这样多个项目完全隔离，互不干扰。

> 
> 你现在的结构问题：
> `D:\default\gostudynew` 是你的记账项目文件夹。
> D 盘 default 根目录放了 `go.mod`，这是**模块放错位置**，容易造成模块混乱、包互相干扰。

## 1. 两种项目类型，分开理解

### 类型 A：独立可执行程序（你写的记账程序、网络请求 demo）

> 
> 特点：`package main`，带`func main()`，可以`go run`直接跑，编译出 exe
> ✅ **建议：每一个单独程序，单独一个文件夹，每个文件夹自己一套 go.mod**
> 举例：

```
D:\default\
├─gostudynew\        # 项目1：SQLite记账本
│   ├─go.mod
│   ├─go.sum
│   ├─account.go
│   ├─.gitignore
│   └─account.db
├─httpdemo\          # 项目2：网络请求demo（单独文件夹）
│   ├─go.mod
│   └─main.go
└─calculator\        # 项目3：计算器小程序
    ├─go.mod
    └─main.go
```

每个文件夹独立模块，`go mod init`各自初始化，**包、依赖完全隔离，互不影响**。

### 类型 B：公共库（被别的项目调用，不能直接运行）

> 
> 没有`func main()`，`package xxx`，只写函数、结构体。
> 可以多个库放在同一个大项目内，作为子包。**适合大型项目，新手写小 demo 不用这个。**

## 2. 为什么不建议把多个独立程序放同一个文件夹？

1. **main 函数冲突**：同一个文件夹下多个`package main`的 go 文件，不能同时存在多个`func main()`，就是你刚才遇到的报错。
2. **go.mod 冲突**：一个目录只能有 1 个 go.mod，依赖全部共用。A 项目装的包，B 项目也会读到，版本冲突。
3. **git 管理混乱**：多个程序代码混在一起，提交代码分不清。

创建网络请求
# Go 语言完整学习路线（序号版，循序渐进，贴合你现在 http / 网络编程练习）

> 
> 定位：后端开发方向，从零基础到能写项目；**当前进度：基础语法 → net/http 客户端网络请求**

## 阶段 1：前置准备 & Go 环境（1～3）

1. 安装 Go SDK，配置 GOPATH/GOMODULE，掌握 `go version`、`go env`；熟悉 VSCode + Go 插件，`Shift+Alt+F`格式化、gofmt 命令。
2. 理解 Go 工程基础：`package`、`import`、`func main()`入口；区分**可执行程序 package main** 和**库文件 package xxx**。
3. Go 模块管理：`go mod init`、`go mod tidy`，理解 go.mod、go.sum 作用，解决包导入问题。

## 阶段 2：基础语法（核心，4～14）

4. 变量与常量：`var`声明、短变量`:=`、常量`const`；类型推断；变量作用域（包级别 / 函数内）。
5. 基础数据类型：int、float、bool、string；字符串操作，`string<->[]byte`转换。
6. 运算符：算术、比较、逻辑；注意 Go 没有 while，只有`for`。
7. 流程控制：`if else`、`for`（普通 for、for-range）、`switch`（自带 break，fallthrough）、`goto`（了解即可，尽量不用）。
8. 复合类型：数组 array、切片 slice（重点！append、len、cap、底层数组原理）。
9. 复合类型：map 字典，初始化、增删查，注意 map 零值 panic 问题。
10. 函数：函数定义、多返回值（Go 标志性特性）、参数、返回值；匿名函数、闭包。
11. 指针：`*`、`&`；指针基础，区分值传递 vs 指针传递（重中之重）。
12. 结构体 struct：结构体定义、实例化；结构体方法（值接收者 / 指针接收者）。
13. 结构体标签 tag：`json:"xxx"`，就是你天气代码里面用到的，配合`encoding/json`序列化 / 反序列化。
14. 错误处理：`error`接口；`if err != nil`范式；`errors.New()`；了解 panic/recover（不滥用）。

## 阶段 3：Go 核心特性（15～19，Go 语言灵魂）

15. 接口 interface：鸭子类型，空接口`interface{}`（万能类型）；类型断言。
16. 并发基础（Go 最大亮点）：goroutine，`go`关键字启动协程；理解协程轻量级。
17. channel：通道，用于 goroutine 之间通信；无缓冲、有缓冲 channel；`range`、`close`。
18. 同步包 sync：sync.WaitGroup、sync.Mutex 互斥锁、sync.Once；解决并发竞争。
19. select 多路监听；time 包定时器、超时控制。

## 阶段 4：标准库实战（20～26，你正在这个阶段）

> 
> 标准库是 Go 生产力核心，优先学这些

20. `fmt`：格式化输入输出。
21. `time`：时间处理、休眠、定时。
22. `encoding/json`：JSON 序列化 / 反序列化（你的天气项目）。
23. `net/http`
   - 客户端：`http.Client`、`Get/Post`请求，读取 resp、resp.Body，就是你当前写的天气爬虫。
   - 服务端：`http.HandleFunc`、`http.ResponseWriter`，搭建 web 页面，返回 HTML/JSON。
24. `io / io.Reader / io.Writer`：IO 流概念，`io.ReadAll`。
25. `os`：文件读写、创建文件、读取命令行参数。
26. `strconv`、`strings`字符串和类型转换。

## 阶段 5：项目开发 & 常用第三方包（27～31）

27. 工程化规范：代码分层，配置读取，日志打印。
28. Web 框架入门：Gin（最流行），路由、参数绑定、中间件。
29. 数据库操作：database/sql；mysql 驱动；CRUD；ORM 框架 GORM。
30. 其他常用包：viper 配置、zap 日志。
31. 爬虫进阶：http 请求头、cookie、代理；了解 colly 爬虫库。

## 阶段 6：进阶底层 & 性能调优（32～36，工作进阶）

32. 内存模型：slice 底层、map 底层、逃逸分析。
33. 并发模型：CSP 模型，原子操作 sync/atomic。
34. 单元测试：`_test.go`，`go test`。
35. pprof 性能分析，内存泄漏排查。
36. 编译、交叉编译，打包部署。

## 阶段 7：实战项目（由小到大，建议穿插在学习中）

1. 小项目 1：天气查询工具（**你现在正在写的这个**）
2. 小项目 2：HTTP 静态网页服务器，读取本地 html 文件返回浏览器
3. 小项目 3：简易爬虫，抓取网页保存 html
4. 项目 4：Gin 开发简单 Web 接口，返回 JSON 数据
5. 项目 5：带数据库的后端接口服务（用户增删改查）
6. 项目 6：并发爬虫

---

# ✅ 给你的学习建议

你现在处于【阶段 4，23 net/http 客户端】
👉 下一步顺序：

1. 写完天气程序，吃透 resp、body、json 解析
2. 练习：http 请求抓取网页，保存 html 到本地文件
3. 学习：http 服务端，写一个 web 页面返回 HTML
4. 回头巩固：结构体、指针、error 错误处理
5. 再进入 goroutine 并发