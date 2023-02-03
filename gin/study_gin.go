package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 自定义中间件 (相当于拦截器)
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 可以通过if进行一些判断
		// 通过set方法可以在context中加入一些值
		// 通过Next方法放行，Abort方法拦截
		context.Set("usersession", "1")
		context.Next()
	}
}
func main() {

	// 创建一个服务
	ginServer := gin.Default()
	// 注册中间件 (如果没在接口中指定会全局使用)
	ginServer.Use(myHandler())

	// 接口
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "hello gin"}) // 返回JSON
	})

	// 接收前端通过url形式传递的参数
	//  1.  usr?userid=1&username=zs形式
	ginServer.GET("/usr", func(context *gin.Context) {
		// 代码中依靠Query方法获取
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	//  2. usr/1/zs形式
	ginServer.GET("/usr/:userid/:username", func(context *gin.Context) {
		// url中指明对应位置的含义，然后在代码中依靠Query方法获取
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 接收传递的json
	ginServer.POST("/json", func(context *gin.Context) {
		//request body
		data, _ := context.GetRawData()
		// 解析JSON
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)

		context.JSON(http.StatusOK, m)
	})

	// 路由
	ginServer.GET("/bili", func(context *gin.Context) {
		// 重定向 301
		context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com/")
	})

	// 路由组
	userGroup := ginServer.Group("/usr")
	{
		// /usr/add
		userGroup.GET("/add", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "添加用户"})
		})
		// /usr/login
		userGroup.GET("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "用户登录"})
		})
	}

	// 中间件的使用
	ginServer.GET("/info", myHandler(), func(context *gin.Context) {
		// 取出中间件中的值
		s := context.MustGet("usersession").(string)
		log.Println("=====> ", s)
		context.JSON(http.StatusOK, gin.H{"mag": "ok"})
	})

	// 服务端口 (注意有 : )
	err := ginServer.Run(":8081")
	if err != nil {
		fmt.Printf("[服务器异常] %+v", err)
		return
	}

}
