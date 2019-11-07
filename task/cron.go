/**
 * @Author : jinchunguang
 * @Date : 19-11-7 下午4:00
 * @Project : go-worker
 */
package task

import (
    "github.com/robfig/cron"
    "log"
    "time"
)

// 任务启用开关
const task_enable = 1

type Task struct {
    cron *cron.Cron
}

func NewTask() *Task {

    task := new(Task)
    task.cron = cron.New()
    task.loader()
    task.cron.Start()
    return task
}

func (t *Task) loader() {
    if task_enable != 1 {
        return
    }
    t.cron.AddFunc("*/10 * * * * *", func() {
        log.Println("Run ...", time.Now())
    })
}
