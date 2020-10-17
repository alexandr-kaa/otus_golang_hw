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
	currentErrorCount *int
}

func worker(data Data, chanTask <-chan Task, done <-chan struct{}) {
	defer data.pWaitGr.Done()
	for {
		select {
		case <-done:
			return
		default:
			{
			}
		}
		data.localLock.RLock()
		if data.M > 0 && *data.currentErrorCount > data.M-1 {
			data.localLock.RUnlock()
			return
		}
		data.localLock.RUnlock()
		select {
		case task := <-chanTask:
			err := task()
			if err != nil {
				data.localLock.Lock()
				*data.currentErrorCount++
				data.localLock.Unlock()
			}
		default:
			return
		}
	}
}
func Run(tasks []Task, N int, M int) error {
	// Place your code here
	var currentErrorCount int
	chanTask := make(chan Task, len(tasks))
	done := make(chan struct{}, 1)
	var waitGr sync.WaitGroup

	var localLock sync.RWMutex
	data := Data{&waitGr, &localLock, N, M, &currentErrorCount}
	for i := 0; i < data.N; i++ {
		waitGr.Add(1)
		go worker(data, chanTask, done)
	}
	for _, task := range tasks {
		localLock.RLock()
		if M > 0 && currentErrorCount > M-1 {
			close(done)
			close(chanTask)
			localLock.RUnlock()
			waitGr.Wait()
			return ErrErrorsLimitExceeded
		}
		localLock.RUnlock()
		chanTask <- task
	}
	waitGr.Wait()
	if M > 0 && currentErrorCount > M-1 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
