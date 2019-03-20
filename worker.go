package gorun

// WorkerCreate - constant for create status
const WorkerCreate = 1

// WorkerRun - constant for run status
const WorkerRun = 2

// WorkerStop - constant for stop status
const WorkerStop = 3

// WorkerPause - constant for stop status
const WorkerPause = 4

// Worker - basic worker struct
// Status - current worker status
// Name - worker unique name
// Params - worker outer params
// Fn - worker function
// tm - pointer to worker manager
// isRun - flag for workers, that gorutine is start
type Worker struct {
	Status int
	Name   string
	Params interface{}
	Fn     func(worker *Worker)
	tm     *WorkerManager
	run    bool
}

// Run - run worker gorutine
func (w *Worker) Run() {
	if w.Status == WorkerCreate || w.Status == WorkerPause {
		w.Status = WorkerRun
		if !w.run {
			w.run = true
			go w.Fn(w)
		}
	}
}

// Stop - stop and remove worker
func (w *Worker) Stop() {
	w.Status = WorkerStop
	w.tm.removeWorker(w.Name)
}

// Pause - pause worker
func (w *Worker) Pause() {
	w.Status = WorkerPause
}

// IsRun - check, that worker can run
func (w *Worker) IsRun() bool {
	if w.Status == WorkerRun {
		return true
	}
	return false
}

// IsStop - check, that worker must stop worker and finish gorutine
func (w *Worker) IsStop() bool {
	if w.Status == WorkerStop {
		return true
	}
	return false
}

// IsPause - check, that worker is pause
func (w *Worker) IsPause() bool {
	if w.Status == WorkerPause {
		return true
	}
	return false
}
