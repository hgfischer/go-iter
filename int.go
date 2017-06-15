package iter

import (
	"context"
	"errors"
	"math"
)

var (
	// ErrInvalidSequence ...
	ErrInvalidSequence = errors.New("Invalid sequence")
)

// IntSeqOptFunc ...
type IntSeqOptFunc func(s *Int)

// Start sets the start of the iteration
func Start(n int) IntSeqOptFunc { return func(seq *Int) { seq.start = n } }

// Stop sets the stop of the iteration
func Stop(n int) IntSeqOptFunc { return func(seq *Int) { seq.stop = n } }

// Step sets the end of the iteration. It must be negative to decrement.
func Step(n int) IntSeqOptFunc { return func(seq *Int) { seq.step = n } }

// Ctx sets a context, necessary for using the channel-based Iterator
func Ctx(ctx context.Context) IntSeqOptFunc { return func(seq *Int) { seq.ctx = ctx } }

// Int represents a iterable sequence of integers
type Int struct {
	start int
	stop  int
	step  int
	curr  int
	ctx   context.Context
	err   error
}

// NewIntSeq creates a new Int Sequence generator
func NewIntSeq(opts ...IntSeqOptFunc) *Int {
	seq := &Int{step: 1}
	for _, opt := range opts {
		opt(seq)
	}
	seq.curr = seq.start
	if (seq.start > seq.stop && seq.step > 0) || (seq.start < seq.stop && seq.step < 0) {
		seq.err = ErrInvalidSequence
	}
	return seq
}

// NewIntSeqStart creates a new Int Sequence generator
func NewIntSeqStart(opts ...IntSeqOptFunc) (*Int, int) {
	seq := NewIntSeq(opts...)
	return seq, seq.start
}

func (s *Int) quantity() int {
	return int(math.Ceil(math.Abs(float64(s.stop-s.start) / float64(s.step))))
}

// All returns the entire integer slice at once
func (s *Int) All() []int {
	if s.err != nil {
		return []int{}
	}
	ints := make([]int, s.quantity())
	for i, n := 0, s.start; s.Continue(); i, n = i+1, s.Get() {
		ints[i] = n
	}
	return ints
}

// Continue returns true if the sequence hasn't ended
func (s *Int) Continue() bool {
	if s.err != nil {
		return false
	}
	if s.step >= 0 {
		return s.curr < s.stop
	}
	return s.curr > s.stop
}

// Get returns the current number in the sequence
func (s *Int) Get() int {
	s.curr += s.step
	return s.curr
}

// Iter returns a channel that can be iterated on
func (s *Int) Iter() <-chan int {
	ch := make(chan int, s.quantity())
	if s.err != nil {
		close(ch)
		return ch
	}
	go func() {
		for n := s.start; s.Continue(); n = s.Get() {
			if s.ctx == nil {
				ch <- n
			} else {
				select {
				case <-s.ctx.Done():
					break // avoid leaking of this goroutine when ctx is done.
				case ch <- n:
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (s *Int) Error() error {
	return s.err
}
