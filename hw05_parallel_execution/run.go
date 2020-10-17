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
	chanTask := make(chan Task, len(tasks))
	done:=make(chan interface{},1)
	var waitGr sync.WaitGroup
	var localLock sync.Mutex
	for i:=0;i<N;i++{
        	waitGr.Add(1)
        	go func(){
		 defer waitGr.Done()
		 for{
		 select{
		   case <-done:return
		   default:{}
		 }
		 localLock.Lock()
		 if M>0 && currentErrorCount > M-1{
		 	 localLock.Unlock()
			 return
		 }
		 localLock.Unlock()
		 select{
		 case task:=<-chanTask:
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
	 }
	 for _,task:= range(tasks){
		 localLock.Lock()
		 if M>0 && currentErrorCount > M-1{
			 close (done)
			 close (chanTask)
			 localLock.Unlock()
			 waitGr.Wait()
			 return ErrErrorsLimitExceeded
		 }
		 localLock.Unlock()
		 chanTask<-task
	 }
	waitGr.Wait()
	if M>0 && currentErrorCount > M-1{
	 return ErrErrorsLimitExceeded
	}
	return nil
}
