package main

import (
	"fmt"
	"time"
)

type worker struct {
	id  int
	err error
}

type WorkerManager struct {
	// 監控 worker 是否死亡的緩衝 channel
	workerChan chan *worker
	// 總共要監控的 worker 數量
	nWorkers int
}

// 創建一個 WorkerManager
func NewWorkerManager(nworkers int) *WorkerManager {
	return &WorkerManager{
		nWorkers:   nworkers,
		workerChan: make(chan *worker, nworkers),
	}
}

// 啟動 worker 池, 並分配 ID 讓 worker 進行工作
func (wm *WorkerManager) StartWorkerPool() {
	for i := 0; i < wm.nWorkers; i++ {
		i := i
		wk := &worker{id: i}
		go wk.work(wm.workerChan)
	}

	// 啟動保持存活監控
	wm.KeepLiveWorkers()
}

// 保持存活監控 workers
func (wm *WorkerManager) KeepLiveWorkers() {
	for wk := range wm.workerChan {
		// print error
		fmt.Printf("Worker %d stopped with err: [%v] \n", wk.id, wk.err)
		// reset error
		wk.err = nil
		// wk死亡, 重新啟動worker
		go wk.work(wm.workerChan)
	}
}

func (wk *worker) work(workerChan chan<- *worker) (err error) {
	// 任何 Goroutine 只要退出都會調用 defer, 所以在這裡利用 WorkerManager 的 WorkChan 發送通知
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("Panic happened with [%v]", r)
			}
		} else {
			wk.err = err
		}

		// 通知主 Goroutine, 當前 Goroutine 已經死亡
		workerChan <- wk
	}()

	// do something
	fmt.Println("Start Worker...ID = ", wk.id)

	// 每個 worker 睡眠一段時間之後, panic退出或 go exit()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
	}

	panic("worker panic...")
	// runtime.GoExit()

	return err
}

func main() {
	wm := NewWorkerManager(10)

	wm.StartWorkerPool()
}
