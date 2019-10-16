package gorun

import (
	"testing"
	"time"
)

var TestA map[string]string

func testFunc(w *Worker) {
	defer w.Stop()
	for {
		if w.IsRun() {
			if TestA == nil {
				TestA = make(map[string]string)
			}
			TestA[w.Name] = "1"
			time.Sleep(5 * time.Second)
		}
	}
}

func TestWorkerManagerCreate(t *testing.T) {
	t.Log(`Create new worker manager`)
	var tm *WorkerManager
	tm = New()
	if tm == nil {
		t.Errorf(`Error create new worker manager`)
	}
}

func TestAddWorker(t *testing.T) {
	t.Log(`Add new worker`)
	_, e := WM().AddWorker(`test_worker`, nil, testFunc)
	if e != nil {
		t.Errorf(`Error create worker: %s`, e)
	} else {
		t.Log(`Worker created`)
	}
}

func TestTwoWorkersWithSameName(t *testing.T) {
	t.Log(`Add two workers with the same name`)
	_, e := WM().AddWorker(`test_worker_1`, nil, testFunc)
	if e != nil {
		t.Errorf(`Error create worker: %s`, e)
	} else {
		t.Log(`Worker test_worker_1 created`)
	}

	_, e = WM().AddWorker(`test_worker_1`, nil, testFunc)
	if e == nil {
		t.Errorf(`Worker with the same name exists`)
	} else {
		t.Log(`Worker test_worker_1 not created`)
	}
}

func checkWorkers(t *testing.T, names ...string) {
	for _, n := range names {
		if _, check := TestA[n]; !check {
			t.Errorf(`Worker with name %s not working`, n)
		} else {
			t.Log(`Worker with name`, n, `work`)
		}
	}
}

func TestRunWorkers(t *testing.T) {
	WM().RunAll()
	time.Sleep(100 * time.Millisecond)
	checkWorkers(t, "test_worker", "test_worker_1")
}

func TestStopAllWorkers(t *testing.T) {
	WM().StopAll()
	time.Sleep(100 * time.Millisecond)
	if len(WM().Workers) > 0 {
		t.Errorf(`workers not finished`)
	} else {
		t.Log(`all workers finished`)
	}
}
