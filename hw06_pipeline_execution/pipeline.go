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

func (w Worker) execStage(done In) Out {
	if w.in == nil {
		return nil
	}

	out := make(Bi)

	outChan := w.stage(w.in)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			default:
				{
				}
			}
			select {
			case s, ok := <-outChan:
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

	return Out(out)
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here

	outChannel := in
	for _, stage := range stages {
		actual := stage
		worker := Worker{in: outChannel,
			stage: actual}
		outChannel = worker.execStage(done)
		if outChannel == nil {
			return nil
		}
	}
	return outChannel
}
