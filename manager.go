package gorun

import (
	"fmt"
)

// WorkerManager - manage workers, update worker status, reg new workers
// Workers - slice of Worker pointers
type WorkerManager struct {
	Workers map[string]*Worker
}

var workerManager *WorkerManager

// New - create new worker manager
// return
// WM - new WorkerManager instance
func New() *WorkerManager {
	return &WorkerManager{
		Workers: make(map[string]*Worker),
	}
}

// WM - singleton WorkerManager instance
// return
// * WM - pointer to WorkerManager instance
func WM() *WorkerManager {
	if workerManager == nil {
		workerManager = New()
	}
	return workerManager
}

func (tm *WorkerManager) workerExists(name string) bool {
	_, check := tm.Workers[name]
	return check
}

func (tm *WorkerManager) removeWorker(name string) {
	delete(tm.Workers, name)
}

// AddWorker - add new worker to worker manager
// return
// err - error, if worker with the same name exists
func (tm *WorkerManager) AddWorker(name string, params interface{}, cb func(*Worker)) (*Worker, error) {
	if tm.workerExists(name) {
		return nil, fmt.Errorf(`worker with name [%s] exists`, name)
	}
	tm.Workers[name] = &Worker{
		Name:   name,
		Status: WorkerCreate,
		Params: params,
		Fn:     cb,
		tm:     tm,
	}

	return tm.Workers[name], nil
}

// GetWorkers - return workers list in queue
// return
// list - map[string]*Workers - list of available workers
func (tm *WorkerManager) GetWorkers() map[string]*Worker {
	return tm.Workers
}

// RunAll - run all workers
func (tm *WorkerManager) RunAll() {
	for _, worker := range tm.Workers {
		worker.Run()
	}
}

// StopAll - stop all workers
func (tm *WorkerManager) StopAll() {
	for _, worker := range tm.Workers {
		worker.Stop()
	}
}

// PauseAll - pause all workers
func (tm *WorkerManager) PauseAll() {
	for _, worker := range tm.Workers {
		worker.Pause()
	}
}

// UpdateWorker - update worker data
func (tm *WorkerManager) UpdateWorker(name string, params interface{}) error {
	if !tm.workerExists(name) {
		return fmt.Errorf(`worker %s not found`, name)
	}
	tm.Workers[name].Params = params
	return nil
}

// SetStatus - update worker status
func (tm *WorkerManager) SetStatus(name string, status int) error {
	if !tm.workerExists(name) {
		return fmt.Errorf(`worker %s not found`, name)
	}
	if inArray([]int{WorkerCreate, WorkerRun, WorkerPause, WorkerStop}, status) {
		tm.Workers[name].Status = status
		return nil
	}

	return fmt.Errorf(`worker status not found`)
}

// RemoveWorker - remove worker by worker name
func (tm *WorkerManager) RemoveWorker(name string) error {
	if !tm.workerExists(name) {
		return fmt.Errorf(`worker %s not found`, name)
	}
	tm.Workers[name].Stop()
	return nil
}

// Get - return worker by name
func (tm *WorkerManager) Get(name string) (*Worker, error) {
	if !tm.workerExists(name) {
		return nil, fmt.Errorf(`worker %s not found`, name)
	}
	return tm.Workers[name], nil
}

// Run - run worker or return error
func (tm *WorkerManager) Run(name string) (*Worker, error) {
	if !tm.workerExists(name) {
		return nil, fmt.Errorf(`worker %s not found`, name)
	}
	tm.Workers[name].Run()
	return tm.Workers[name], nil
}

func inArray(arr []int, need int) bool {
	for _, i := range arr {
		if arr[i] == need {
			return true
		}
	}
	return false
}
