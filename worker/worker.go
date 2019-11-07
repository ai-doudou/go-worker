/**
 * @Author : jinchunguang
 * @Date : 19-11-7 下午4:48
 * @Project : go-worker
 */
package worker

import (
    "fmt"
    "time"
)

type Worker struct {
    Queue chan string
}

func NewWorker() *Worker {
    w := new(Worker)
    w.StartWorkerPool()
    return w
}

// 启用一个任务
func (w *Worker) StartOneWorker(workerID int) {
    for {
        select {
        case request := <-w.Queue:
            fmt.Println("[Wait Task] WorkerID  = ", workerID, " is started.")
            w.ExecuteTask(request)
        default:
            time.Sleep(500*time.Millisecond)
        }
    }
}

// 启动协程池
func (w *Worker) StartWorkerPool() {
    for i := 0; i < 10; i++ {
        w.Queue = make(chan string, 100)
        go w.StartOneWorker(i)
    }
}

