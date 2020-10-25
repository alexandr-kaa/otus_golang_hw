package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

type Worker struct {
	in    In
	stage Stage
}

func (w Worker) execStage() Out {
	if w.in == nil {
		return nil
	}
	outChan := w.stage(w.in)
	return outChan
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here

	innerChannel := in
	out := make(Bi)
	go func() {
		defer close(out)
		for _, stage := range stages {
			actual := stage
			worker := Worker{in: innerChannel,
				stage: actual}
			innerChannel = worker.execStage()
			if innerChannel == nil {
				return
			}
		}
		for {
			select {
			case <-done:
				return
			default:
				{
				}
			}
			select {
			case s, ok := <-innerChannel:
				if !ok {
					return
				}
				out <- s
			default:
				{
				}
			}
		}
	}()
	return out
}
