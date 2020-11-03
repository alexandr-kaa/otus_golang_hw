package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here
	innerChannel := in
	for _, stage := range stages {
		actual := stage
		out := make(Bi)
		go func(stage Stage, out Bi, in In) {
			defer close(out)
			for s := range stage(in) {
				select {
				case <-done:
					return
				default:
				}
				select {
				case out <- s:
				case <-done:
					return
				}
			}
		}(actual, out, innerChannel)
		innerChannel = out
	}
	return innerChannel
}
