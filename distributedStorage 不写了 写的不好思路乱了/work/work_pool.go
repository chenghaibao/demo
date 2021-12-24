package work

type WorkPool struct {
	Pool chan string
}

var Pool *WorkPool

func NewPool() {
	work := make(chan string, 100)
	Pool = &WorkPool{
		work,
	}
}
