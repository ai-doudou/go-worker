/**
 * @Author : jinchunguang
 * @Date : 19-11-7 下午5:46
 * @Project : go-worker
 */
package worker

import (
    "fmt"
)

// 执行任务
func (w *Worker) ExecuteTask(request string) {
    fmt.Println("[Execute Task] Queue msg = ", request, " is started.")
}
