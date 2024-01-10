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
