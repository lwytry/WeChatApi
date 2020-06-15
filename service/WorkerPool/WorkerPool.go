package WorkerPool

type Job interface {
	Do()
}

type WorkerPool struct {
	workerLen int
	JobQueue chan  Job
	WorkerQueue chan chan Job
}


func NewWorkerPool(workerLen, jobLen int) *WorkerPool {
	return &WorkerPool{
		workerLen:   workerLen,
		JobQueue:    make(chan Job, jobLen),
		WorkerQueue: make(chan chan Job, workerLen),
	}
}


var (
	workerNum = 100 * 100
	jobNum = 100 * 100
	WP *WorkerPool
)

func init () {
	WP = NewWorkerPool(workerNum, jobNum)
	WP.Start()
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerLen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}

	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}