package sDate

import (
	"sort"
	"sync"
	"time"

	"github.com/yasseldg/simplego/unique"
)

type Waiter struct {
	Group sync.WaitGroup
	names []string
	poss  []int
	proc  chan int
	stop  chan bool
}

func NewWaiter(names ...string) *Waiter {
	w := &Waiter{
		proc:  make(chan int),
		stop:  make(chan bool, 1),
		names: names,
	}
	w.Group.Add(len(names))
	return w
}

func (w *Waiter) Done(n int, stop bool) {
	select {
	case <-w.stop:
		return
	default:
		w.proc <- n
	}

	if stop {
		w.stop <- true
	}
}

func (w *Waiter) timeOut(t time.Duration) {
	go func() {
		time.Sleep(t)

		w.Done(0, true)
	}()
}

func (w *Waiter) close() {
	close(w.proc)
	w.stop <- true
}

func (w *Waiter) remaining() []string {

	sort.Ints(w.poss)
	unique.Ints(&w.poss)
	sort.Sort(sort.Reverse(sort.IntSlice(w.poss)))

	for _, p := range w.poss {
		w.names = append(w.names[:p], w.names[p+1:]...)
	}

	return w.names
}

func (w *Waiter) Proc() []string {
	for n := range w.proc {
		if n == 0 {
			return w.remaining()
		}

		w.poss = append(w.poss, n-1)
		w.Group.Done()
	}

	return []string{}
}

// wg.Wait and TimeOut, timeout = 0 = disable
func (w *Waiter) Wait(timeout time.Duration) {

	if timeout > 0 {
		w.timeOut(timeout)
	}

	go func() {
		defer w.close()
		w.Group.Wait()
	}()
}
