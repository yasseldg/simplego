package sDate

import (
	"sync"
	"time"

	"github.com/yasseldg/simplego/sSlice"
)

type Waiter struct {
	Group   sync.WaitGroup
	process map[string]int
	poss    []int
	proc    chan int
	stop    chan bool
}

func NewWaiter(process map[string]int) *Waiter {
	l := len(process)
	w := &Waiter{
		proc:    make(chan int, l),
		stop:    make(chan bool, 1),
		process: process,
	}
	w.Group.Add(l)
	return w
}

// Wait waits for all the processes to finish, with a timeout if specified
func (w *Waiter) Wait(timeout time.Duration) {
	if timeout > 0 {
		go w.timeOut(timeout)
	}

	go func() {
		defer w.close()
		w.Group.Wait()
	}()
}

func (w *Waiter) Done(n int, stop bool) {
	select {
	case <-w.stop:
		return
	case w.proc <- n:
	default:
	}

	if stop {
		w.stop <- true
	}
}

func (w *Waiter) Proc() []string {
	for {
		select {
		case n, ok := <-w.proc:
			if !ok {
				return w.remaining()
			}
			if n == 0 {
				return w.remaining()
			}
			w.poss = append(w.poss, n)
			w.Group.Done()
		case <-w.stop:
			return w.remaining()
		}
	}
}

func (w *Waiter) timeOut(t time.Duration) {
	select {
	case <-w.stop:
		return
	case <-time.After(t):
		w.Done(0, true)
	}
}

func (w *Waiter) close() {
	close(w.proc)
	w.stop <- true
}

func (w *Waiter) remaining() []string {
	r := make([]string, 0)
	for name, p := range w.process {
		if !sSlice.InInts(p, w.poss) {
			r = append(r, name)
		}
	}
	return r
}
