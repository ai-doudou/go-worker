/**
 * @Author : jinchunguang
 * @Date : 19-11-7 下午5:46
 * @Project : go-worker
 */
package worker

import (
    "fmt"
)

func (w *Worker) SendTask(request string) {
    fmt.Println("[Send Task] Queue Msg = ", request)
    w.Queue <- request
}
