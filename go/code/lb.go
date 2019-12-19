package main

import "fmt"

type Request struct {
	fn func() int // The operation to perform.
	c  chan int   // The channel to return the result.
}

func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		Sleep(rand.Int63n(nWorker * 2 * Second))
		work <- Request{workFn, c} // send request
		result := <-c              // wait for answer
		furtherProcess(result)
	}
}

// A channel of requests, plus some load tracking data.
type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

// The channel of requests (w.requests) delivers requests to each worker. The balancer tracks the number of pending requests as a measure of load.
// Each response goes directly to its requester.
// Could run the loop body as a goroutine for parallelism.
func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}

// Balancer definition

// The load balancer needs a pool of workers and a single channel to which requesters can report task completion.
type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

// Now we balance by making the Pool a heap tracked by load.
func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

func main() {
}
