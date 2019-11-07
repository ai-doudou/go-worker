/**
 * @Author : jinchunguang
 * @Date : 19-11-7 下午4:11
 * @Project : go-worker
 */
package main

import (
    "github.com/gin-gonic/gin"
    "go-worker/utils"
    "go-worker/worker"
    "time"
)

func main() {
    // gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())



    // 监测
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // 启用协程池
    wk := worker.NewWorker()
    // 任务发送测试
    r.GET("/send_task", func(c *gin.Context) {
        wk.SendTask(utils.GenerateMd5(time.Now().String()))
        c.JSON(200, gin.H{
            "message": "success",
        })
    })

    r.Run(":6600")

}
