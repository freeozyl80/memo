# golang 一些简答问题：
- slice 是引用类型

# golang 实例化：
```golang
t :=new(T)
var t T
t := &T{}
```





# 权限控制: 

- rbac(role-based-access-control): casbin 库
- abac(attribue-based-access-control): 

[参考链接](https://freedium.cfd/https%3A%2F%2Fmedium.com%2F%40abhinavv.singh%2Fa-comprehensive-guide-to-authentication-and-authorization-in-go-golang-6f783b4cea18)

# 框架选择

- gin：
``` golang
r.GET("/ping", func(c *gin.Context) {
    result := make(chan gin.H)
    go func(context *gin.Context) {
        if !ok {
            ok = true
            time.Sleep(5 * time.Second) // 5 Seconds
        }

        result <- gin.H{
            "message": "pong",
            "requestedPath": context.Request.URL.Path,
        }
    }(c.Copy())
    c.JSON(http.StatusOK, <-result)
})
```
## 关于gin async 的问题解决思路： https://github.com/gin-gonic/gin/issues/2499

## 一个比较好的 对 golang err,success 包装的例子（泛型的例子）

- (跳转链接)[https://towardsdev.com/error-handling-in-go-with-the-result-pattern-c1017ed1a272]

## golang 闭包
```golang
// 方案1
for _, val := range values {
    go func(val interface{}) {
    fmt.Println(val)
    }(val)
}

// 方案2
for _, val := range values {
    val := val
    go func() {
    fmt.Println(val)
    }()
}

```
## 一个简单的消费者例子，lua in redis， watch

- (链接)[https://blog.stackademic.com/solving-concurrency-problems-with-redis-and-golang-720e68b2b95a]

## jsonmarshal

```golang
package main

import (
 "encoding/json"
 "fmt"
)

func main() {
 // JSON data as a byte slice
 jsonData := []byte(`{
  "name": "John Doe",
  "age": 30,
  "city": "New York",
  "hasCar": true,
  "languages": ["Go", "JavaScript"]
 }`)

 // Parse JSON into an empty interface
 var result interface{}
 err := json.Unmarshal(jsonData, &result)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }

 // Accessing dynamic JSON fields
 dataMap, ok := result.(map[string]interface{})
 if !ok {
  fmt.Println("Invalid JSON structure")
  return
 }

 // Accessing specific fields
 name, nameExists := dataMap["name"].(string)
 age, ageExists := dataMap["age"].(float64)
 city, cityExists := dataMap["city"].(string)
 hasCar, hasCarExists := dataMap["hasCar"].(bool)
 languages, languagesExists := dataMap["languages"].([]interface{})

 // Displaying parsed data
 if nameExists {
  fmt.Println("Name:", name)
 }

 if ageExists {
  fmt.Println("Age:", int(age))
 }

 if cityExists {
  fmt.Println("City:", city)
 }

 if hasCarExists {
  fmt.Println("Has Car:", hasCar)
 }

 if languagesExists {
  fmt.Println("Languages:")
  for _, lang := range languages {
   fmt.Println(" -", lang)
  }
 }
}
```

## interface

## mode：

- 工厂模式即

```golang
type IShirt interface {
    genTshirt() string
}
```
- 适配器模式

``` golang
type windowAdaptor struct {
    windowMachine *window
}
func (m *windowAdaptor) convertUSBtoLightining {
    m.windowMachine.insertIntoUSBPort()
}

```
- 桥接模式
抽象不同的层
- 装饰器模式
不断加工
- 享元模式
提出公共属性，不依靠实例创建，核心还是依靠单例模式
- 责任链模式
很像链表
