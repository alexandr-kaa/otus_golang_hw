package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

type Worker struct {
	in    In
	done  In
	stage Stage
}

func (w Worker) execStage() Out {
	innerOut := make(Bi)
	go func() {
		defer close(innerOut)
		outchan := w.stage(w.in)
		for {
			select {
			case <-w.done:
				return
			default:
			}
			select {
			case s, ok := <-outchan:
				if !ok {
					return
				}
				innerOut <- s
			case <-w.done:
				return
			}
		}
	}()
	return innerOut
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here
	innerChannel := in
	out := make(Bi)
	if in == nil {
		close(out)
		return out
	}
	go func() {
		defer close(out)
		for _, stage := range stages {
			actual := stage
			worker := Worker{in: innerChannel,
				done:  done,
				stage: actual}
			innerChannel = worker.execStage()
		}
		for data := range innerChannel {
			out <- data
		}
	}()
	return out
}
