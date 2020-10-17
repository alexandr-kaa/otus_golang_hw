package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	// Place your code here 
	var currentErrorCount int
	chanTask := make(chan Task, len(Tasks))
	done:=make(chan interface{},1)
	var waitGr sync.WaitGroup
	waitGr.Add(N)
	localLock:=sync.Mutex
	for i:=0;i<N;i++{
		go func(){
		 defer waitGr.Done()
		 for{
		 select{
		   case <-done:return
		   default:{}
		 }
		 localLock.Lock()
		 if currentErrorCount >M-1{
			 localLock.Unlock()
			 return
		 }
		 localLock.Unlock()
		 select{
		 case task<-chanTask:
			 err:=task()
			 if err != nil{
			 localLock.Lock()
			 currentErrorCount++
			 localLock.Unlock()
			 }
	         default : return
		 }
	 }
		 
	 }()
	 for _,task:= range(tasks){
		 localLock.Lock()
		 if currentErrorCount > M-1{
			 close (done)
			 close (chanTask)
			 waitGr.Wait()
			 return ErrErrorsLimitExceeded
		 }
		 chanTask<-task
	 }

	}
	waitGr.Wait()
	return nil
}
