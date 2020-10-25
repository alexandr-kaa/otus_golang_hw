package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

/* Run starts tasks in N goroutines and stops its work when receiving M errors from them.... */
type TaskContext struct {
	pWaitGr           *sync.WaitGroup
	localLock         *sync.RWMutex
	N                 int
	M                 int
	currentErrorCount *int32
	number            int // Используется для отладки
	chanTask          <-chan Task
}

func (context TaskContext) ErrLimWasExceed() bool {
	(context.localLock).RLock()
	defer (context.localLock).RUnlock()
	return (context.M > 0 && *context.currentErrorCount > int32(context.M-1))
}

func (context TaskContext) IncreaseCounter() {
	context.localLock.Lock()
	defer context.localLock.Unlock()
	(*context.currentErrorCount)++
}

func worker(taskContext TaskContext) {
	defer taskContext.pWaitGr.Done()
	for {
		if taskContext.ErrLimWasExceed() {
			return
		}
		task, ok := <-taskContext.chanTask
		if !ok {
			return
		}
		err := task()
		if err != nil {
			taskContext.IncreaseCounter()
		}
	}
}
func Run(tasks []Task, n int, m int) error {
	// Place your code here
	var currentErrorCount int32
	chanTask := make(chan Task, len(tasks))

	var waitGr sync.WaitGroup

	var localLock sync.RWMutex
	taskContext := TaskContext{&waitGr, &localLock, n, m, &currentErrorCount, 0, chanTask}
	for i := 0; i < n; i++ {
		waitGr.Add(1)
		taskContext.number = i
		go worker(taskContext)
	}
	for _, task := range tasks {
		if taskContext.ErrLimWasExceed() {
			close(chanTask)
			waitGr.Wait()
			return ErrErrorsLimitExceeded
		}
		chanTask <- task
	}
	close(chanTask)
	waitGr.Wait()
	if m > 0 && int(currentErrorCount) > m-1 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
