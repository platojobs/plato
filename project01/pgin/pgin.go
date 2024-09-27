package pgin

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/goccy/go-json"
)
var router = gin.Default()

func PingHandler() {
	
    router.GET("/login",handle);
	router.POST("/userinfo",handle12);
	router.PUT("/user",handle);
	router.DELETE("/user/:id",handle);
	router.POST("/user",handle2);
	router.POST("/user/login",handler123);

    router.Run(":8089")

}

func handle(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func handle12(c *gin.Context) {
	 phone := c.Query("phone")
	 password := c.Query("password")
	 fmt.Println("passwod:",password)
	 c.JSON(http.StatusOK, gin.H{
		 "code":    200,
		 "message": "pong",
		 "data":   map[string]any{
			 "id": 1021,
			 "token":"kxweqj$jksjkjsksnnsnnblghjk",
			 "user": map[string]any{
				"name": "zhangsan",
				 "age":  18,
				 "phone": phone,
				 "sex": 1,
			 },
		 },
	 })
}


func handle2(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if c.Request.Header.Get("token") != "kxweqj$jksjkjsksnnsnnblghjk" {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "token is error",
			"data":   map[string]any{},
		})
		return
	}
	c.Header("token", "kxweqj$jksjkjsksnnsnnblghjk")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "pong",
		"data":   map[string]any{
			"name": "zhangsan",
			"age":  18,
		},
	})
}



func handler123(c *gin.Context){
	var person Person
    if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "参数错误",
			"data":   map[string]any{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功了",
		"data":  map[string]any{
			"user": &person},
	})
}