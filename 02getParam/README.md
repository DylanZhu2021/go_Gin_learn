## 1.Restful风格的API
### you can learn how to get param from URL in this block
*gin支持Restful风格的API*<br>
*Representational State Transfer的缩写。直接翻译的意思互联网应是"表现层状态转化"，是一种用程序的API设计理念：URL定位资源，用HTTP描述操作。
如：<br>
 1.获取文章 /blog/getXxx Get blog/Xxx<br>
 2.添加 /blog/addXxx POST blog/Xxx<br>
 3.修改 /blog/updateXxx PUT blog/Xxx<br>
 4.删除 /blog/delXxxx DELETE blog/Xxx<br>

## 2 API参数
可以通过Context的Param方法获取API参数<br>
localhost:8080/xxx/you

## 3 URL参数
URL参数可以通过DefaultQuery()或Query()方法获取<br>
DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串

## 4 表单参数
#### 表单传输为post请求，http常见的传输格式为四种：<br>

application/json<br>
application/x-www-form-urlencoded<br>
application/xml<br>
multipart/form-data<br>
表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

## 5.路由拆分成单独的文件或包

* ### 1. 基本的路由注册

  这是最基本的路由注册，适合小项目，或者项目demo

  ```go
  func helloHandler(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
          "message": "Hello www.topgoer.com!",
      })
  }
  
  func main() {
      r := gin.Default()
      r.GET("/topgoer", helloHandler)
      if err := r.Run(); err != nil {
          fmt.Println("startup service failed, err:%v\n", err)
      }
  }
  ```

* ### 2.路由拆分成单独的文件或包

  ```go
  package main
  
  import (
      "github.com/gin-gonic/gin"
      "net/http"
  )
  
  func helloHandler(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
          "message": "Hello Gin",
      })
  }
  
  func setupRouter() *gin.Engine {
      r := gin.Default()
      r.GET("/user", helloHandler)
      return r
  }
  ```

  此时main.go中调用上面定义好的setupRouter函数：

  ```go
  func main() {
      r := setupRouter()
      if err := r.Run(); err != nil {
          fmt.Println("startup service failed, err:%v\n", err)
      }
  }
  ```

  ```go
  LearnGin
  ├── go.mod
  ├── go.sum
  ├── main.go
  └── routers.go
  ```

  也可以把routers.go单独封装为一个包，此时需要主义函数名的大小写。

* ### 3.路由拆分成多个文件

  当项目更大的时候，可以把很多很多的路由，才分为多个文件，不用把路由注册写在一个文件中，需要注意的是：怎么调用？以及函数名大小写的问题，否则调用不成功！