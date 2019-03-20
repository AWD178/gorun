# gorun
Simple goroutine manager 

# example
```go
package main

import (
	"gorun"
	"fmt"
	"time"
)

func main() {
	gorun.WM().AddWorker("testWorker", map[string]interface{}{"data": 123}, func(w *gorun.Worker) {
		for {
			if w.IsRun() {
				fmt.Println(`Worker `, w.Name, ` is working`)
				time.Sleep(1 * time.Second)
			}

			if w.IsPause() {
				fmt.Println(`Worker `, w.Name, ` is pause`)
				time.Sleep(1 * time.Second)
			}
		}
	})



	if w, err := gorun.WM().Get(`testWorker`); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(`Worker:`, w.Name, `exists`)
	}


	if w, err := gorun.WM().Get(`testWorker2`); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(`Worker:`, w.Name, `exists`)
	}

	if w, err := gorun.WM().Get(`testWorker32`); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(`Worker:`, w.Name, `exists`)
	}




	gorun.WM().RunAll()

	time.Sleep(5 * time.Second)
	gorun.WM().PauseAll()

	time.Sleep(5 * time.Second)
	gorun.WM().RunAll()

	time.Sleep(5 * time.Second)


	gorun.WM().RemoveWorker(`testWorker`)
	if w, err := gorun.WM().Get(`testWorker`); err == nil {
		w.Pause()
		time.Sleep(5  * time.Second)
	}

	gorun.WM().RemoveWorker(`testWorker`)
	gorun.WM().RemoveWorker(`testWorker`)

	fmt.Println(gorun.WM().Workers)

}

```

##API
``` gorun.WM() ``` - return worker manager singleton instance

``` gorun.New() ``` - create new worker manager

``` gorun.WM().AddWorker(worker_name string, worker_data interface{}, gorutine_fn)``` - add new worker

params:
name - ```string``` - worker name

params - ```interface{}``` - worker data

gorutine_fn -```func (w *gorun.Worker)``` - worker gorutine function 

#
``` gorun.WM().Get(string) ``` - return worker by name or error

``` gorun.WM().RunAll() ``` - run all workers

``` gorun.WM().StopAll()``` - stop all workers and remove

``` gorun.WM().PauseAll()``` - pause all workers without remove

``` gorun.WM().GetWorkers() ``` - return list of workers ```map[string]*Worker```

``` gorun.WM().RemoveWorker(worker_name string)``` - stop and remove worker, return error if worker not found

``` gorun.WM().UpdateWorker(worker_name string, params interface)``` - update worker data, return error if worker not found (stop or remove)

``` gorun.WM().SetStatus(worker_name string, status int)``` - change worker status

available statuses:

* ```gorun.WorkerCreate``` - worker created, not run

* ```gorun.WorkerRun``` - worker run

* ```gorun.WorkerStop``` - worker removed

* ```gorun.WorkerPause``` - worker paused, 


## Worker API
```worker.Run()``` - run worker

```worker.Stop()``` - stop worker and remove from worker manager

```worker.Pause()``` - pause worker

```worker.IsRun()``` - check, that worker status is run

```worker.IsStop()``` - check, that worker status is stop

```worker.IsPause()``` - check, that worker status is pause



#Notice
in gorutine use ```defer worker.Stop()``` - for remove worker from worker manager

