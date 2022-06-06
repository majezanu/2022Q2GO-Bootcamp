package utils

import "sync"

type WorkFunc interface {
	Run()
}

type Work struct {
	fn WorkFunc
}

type Pool struct {
	queue chan Work
	wg    sync.WaitGroup
}

func NewPool(workerSize int) *Pool {
	gp := &Pool{queue: make(chan Work)}

	gp.AddWorkers(workerSize)

	return gp
}

func (gp *Pool) Close() {
	close(gp.queue)
	gp.wg.Wait()
}

func (gp *Pool) ScheduleWork(fn WorkFunc) {
	gp.queue <- Work{fn}
}

func (gp *Pool) AddWorkers(workerSize int) {
	gp.wg.Add(workerSize)
	for i := 0; i < workerSize; i++ {
		go func(workerId int) {
			for job := range gp.queue {
				job.fn.Run()
			}
			gp.wg.Done()
		}(i)
	}
}
