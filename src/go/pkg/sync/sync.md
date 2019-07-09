# [Package sync](https://golang.org/pkg/sync/)

# Links

* [Golang学习 - sync 包 - GoLove - 博客园](https://www.cnblogs.com/golove/p/5918082.html)

# type WaitGroup struct

A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

A WaitGroup must not be copied after first use. 

```go
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
```

# type Once

```go
type Once struct {
        // contains filtered or unexported fields
}

func (o *Once) Do(f func())
```
```go
package main

import (
        "fmt"
        "sync"
)

func main() {
        var once sync.Once

        onceBody := func() {
                fmt.Println("Only once")
        }

        done := make(chan bool)

        N := 10
        for i := 0; i < N; i++ {
                go func() {
                        once.Do(onceBody)
                        done <- true
                }()
        }

        for i := 0; i < N; i++ {
                <-done
        }
}
```

# type Pool struct

Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists. 

```go
type Pool struct {

        // New optionally specifies a function to generate
        // a value when Get would otherwise return nil.
        // It may not be changed concurrently with calls to Get.
        New func() interface{}
        // contains filtered or unexported fields
}

func (p *Pool) Get() interface{}
func (p *Pool) Put(x interface{})
```

# type Locker interface

A Locker represents an object that can be locked and unlocked.

```go
type Locker interface {
        Lock()
        Unlock()
}
```

# type Mutex struct

**Mutex == mutual exclusion**

A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.

A Mutex must not be copied after first use. 

Lock locks m. If the lock is already in use, the calling goroutine blocks until the mutex is available. 

Unlock unlocks m. It is a run-time error if m is not locked on entry to Unlock.

A locked Mutex is not associated with a particular goroutine. It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it. 

```go
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
```

Example:

```go
func main() {
	var mu sync.Mutex
	var count int
	done := make(chan bool)
	for i := 0; i < 10; i++ {

		go func(i int) {
			mu.Lock()
			// 注意这里是 count += i, 不是count++
			count += i
			fmt.Print(count, " ")
			mu.Unlock()
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
	<-done
	}
	fmt.Println()
}
// result: (mutative)
// 1 10 12 15 19 24 30 37 45 45 or
// 9 16 20 20 21 23 26 32 37 45 or ...
```

# type RWMutex struct

Lock locks rw for writing.

RLock locks rw for reading. 

```go
func (rw *RWMutex) Lock()
func (rw *RWMutex) Unlock()
func (rw *RWMutex) RLock()
func (rw *RWMutex) RUnlock()
func (rw *RWMutex) RLocker() Locker
```

# type Cond struct

Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.

Each Cond has an associated Locker L (often a *Mutex or *RWMutex), which must be held when changing the condition and when calling the Wait method.

A Cond must not be copied after first use. 

```go
type Cond struct {
	// L is held while observing or changing the condition
	L Locker
	// contains filtered or unexported fields
}

func NewCond(l Locker) *Cond
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()
```

Because c.L is not locked when Wait first resumes, the caller typically cannot assume that the condition is true when Wait returns. Instead, the caller should Wait in a loop: 

```
c.L.Lock()
for !condition() {
    c.Wait()
}
... make use of condition ...
c.L.Unlock()
```

# type Map struct

Map is like a Go map[interface{}]interface{} but is safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.

The Map type is specialized. Most code should use a plain Go map instead, with separate locking or coordination, for better type safety and to make it easier to maintain other invariants along with the map content.

The Map type is optimized for two common use cases: (1) when the entry for a given key is only ever written once but read many times, as in caches that only grow, or (2) when multiple goroutines read, write, and overwrite entries for disjoint sets of keys. In these two cases, use of a Map may significantly reduce lock contention compared to a Go map paired with a separate Mutex or RWMutex.

The zero Map is empty and ready for use. A Map must not be copied after first use. 

```go
func (m *Map) Store(key, value interface{})
func (m *Map) Load(key interface{}) (value interface{}, ok bool)
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
func (m *Map) Delete(key interface{})
func (m *Map) Range(f func(key, value interface{}) bool)
```
