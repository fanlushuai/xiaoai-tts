package main

import (
	"github.com/gin-gonic/gin"
	xiaoaitts "github.com/qfyang-cn/xiaoai-tts"
	"net/http"
	"os"
)

func main() {
	//系统环境变量中读取
	//todo 可能会session失效，日后再说
	m := &xiaoaitts.MiAccount{
		User: os.Getenv("User"),
		Pwd:  os.Getenv("Pwd"),
	}
	x := xiaoaitts.NewXiaoAi(m)

	r := gin.Default()

	r.POST("/xiaoai/speak", func(c *gin.Context) {
		msg := c.PostForm("msg")

		x.Say(msg)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run(":8848")
}
