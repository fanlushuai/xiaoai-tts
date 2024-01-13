package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	xiaoaitts "github.com/qfyang-cn/xiaoai-tts"
	"net/http"
	"os"
	"time"
)

var x xiaoaitts.XiaoAiFunc

func refreshXiaoAiSessionTicker(){
	//系统环境变量中读取
	m := &xiaoaitts.MiAccount{
		User: os.Getenv("User"),
		Pwd:  os.Getenv("Pwd"),
	}
	x = xiaoaitts.NewXiaoAi(m)

	//1 个小时更新一次
	ticker := time.NewTicker(1*time.Hour)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"))
			x = xiaoaitts.NewXiaoAi(m)
		}
	}()

}


func main() {
	refreshXiaoAiSessionTicker()
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
