package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

/* Run starts tasks in N goroutines and stops its work when receiving M errors from them.... */
type Data struct {
	pWaitGr           *sync.WaitGroup
	localLock         *sync.RWMutex
	N                 int
	M                 int
	currentErrorCount *int32
	number            int
}

func worker(data Data, chanTask <-chan Task) {
	defer data.pWaitGr.Done()
	for {

		data.localLock.RLock()
		if data.M > 0 && int(*data.currentErrorCount) > data.M-1 {
			data.localLock.RUnlock()
			return
		}
		data.localLock.RUnlock()
		select {
		case task, ok := <-chanTask:
			if !ok {
				return
			}
			err := task()
			data.localLock.Lock()
			if err != nil {
				(*data.currentErrorCount)++
			}
			data.localLock.Unlock()
		}
	}
}
func Run(tasks []Task, n int, m int) error {
	// Place your code here
	var currentErrorCount int32
	chanTask := make(chan Task, len(tasks))

	var waitGr sync.WaitGroup

	var localLock sync.RWMutex
	data := Data{&waitGr, &localLock, n, m, &currentErrorCount, 0}
	for i := 0; i < data.N; i++ {
		waitGr.Add(1)
		data.number = i
		go worker(data, chanTask)
	}
	for _, task := range tasks {
		localLock.RLock()
		if m > 0 && int(currentErrorCount) > m-1 {
			close(chanTask)
			localLock.RUnlock()
			waitGr.Wait()
			return ErrErrorsLimitExceeded
		}
		localLock.RUnlock()
		chanTask <- task
	}
	close(chanTask)
	waitGr.Wait()
	if m > 0 && int(currentErrorCount) > m-1 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
