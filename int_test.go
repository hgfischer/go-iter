package iter_test

import (
	"reflect"
	"testing"

	iter "github.com/hgfischer/go-iter"
)

var iterator *iter.Int

func failIfNotDeepEqual(t *testing.T, is, shouldBe interface{}) {
	if !reflect.DeepEqual(is, shouldBe) {
		t.Errorf("%#v != %#v", is, shouldBe)
	}
}

func TestSliceWithIncrestingSequence(t *testing.T) {
	iterator = iter.NewIntSeq(iter.Start(1), iter.Stop(10), iter.Step(2))
	failIfNotDeepEqual(t, iterator.All(), []int{1, 3, 5, 7, 9})
	failIfNotDeepEqual(t, iterator.Error(), nil)
	iterator = iter.NewIntSeq(iter.Stop(10))
	failIfNotDeepEqual(t, iterator.All(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	failIfNotDeepEqual(t, iterator.Error(), nil)
	iterator = iter.NewIntSeq(iter.Start(5), iter.Stop(10))
	failIfNotDeepEqual(t, iterator.All(), []int{5, 6, 7, 8, 9})
	failIfNotDeepEqual(t, iterator.Error(), nil)
	iterator = iter.NewIntSeq(iter.Start(5))
	failIfNotDeepEqual(t, iterator.All(), []int{})
	failIfNotDeepEqual(t, iterator.Error(), iter.ErrInvalidSequence)
	iterator = iter.NewIntSeq(iter.Stop(10), iter.Step(3))
	failIfNotDeepEqual(t, iterator.All(), []int{0, 3, 6, 9})
	failIfNotDeepEqual(t, iterator.Error(), nil)
	iterator = iter.NewIntSeq(iter.Start(5), iter.Stop(100), iter.Step(-1))
	failIfNotDeepEqual(t, iterator.All(), []int{})
	failIfNotDeepEqual(t, iterator.Error(), iter.ErrInvalidSequence)
}

func TestSliceWithDecreasingSequence(t *testing.T) {
	iterator = iter.NewIntSeq(iter.Start(10), iter.Stop(1), iter.Step(-2))
	failIfNotDeepEqual(t, iterator.All(), []int{10, 8, 6, 4, 2})
	failIfNotDeepEqual(t, iterator.Error(), nil)
	iterator = iter.NewIntSeq(iter.Start(5), iter.Step(-1))
	failIfNotDeepEqual(t, iterator.All(), []int{5, 4, 3, 2, 1})
	failIfNotDeepEqual(t, iterator.Error(), nil)
	iterator = iter.NewIntSeq(iter.Stop(10), iter.Step(-3))
	failIfNotDeepEqual(t, iterator.All(), []int{})
	failIfNotDeepEqual(t, iterator.Error(), iter.ErrInvalidSequence)
	iterator = iter.NewIntSeq(iter.Start(5), iter.Stop(100), iter.Step(-1))
	failIfNotDeepEqual(t, iterator.All(), []int{})
	failIfNotDeepEqual(t, iterator.Error(), iter.ErrInvalidSequence)
}

func TestSequenceLoopWithIncreasingSequence(t *testing.T) {
	accum := make([]int, 0)
	iterator, start := iter.NewIntSeqStart(iter.Stop(10), iter.Step(2))
	for seq, n := iterator, start; seq.Continue(); n = seq.Get() {
		failIfNotDeepEqual(t, seq.Error(), nil)
		accum = append(accum, n)
	}
	failIfNotDeepEqual(t, accum, []int{0, 2, 4, 6, 8})
	failIfNotDeepEqual(t, iterator.Error(), nil)
}

func TestSequenceLoopWithDecreasingSequence(t *testing.T) {
	accum := make([]int, 0)
	iterator, start := iter.NewIntSeqStart(iter.Start(10), iter.Stop(-11), iter.Step(-5))
	for seq, n := iterator, start; seq.Continue(); n = seq.Get() {
		failIfNotDeepEqual(t, seq.Error(), nil)
		accum = append(accum, n)
	}
	failIfNotDeepEqual(t, accum, []int{10, 5, 0, -5, -10})
	failIfNotDeepEqual(t, iterator.Error(), nil)
}

func TestChannelIteratorWithIncreasingSequence(t *testing.T) {
	accum := make([]int, 0)
	iterator = iter.NewIntSeq(iter.Stop(10), iter.Step(2))
	for n := range iterator.Iter() {
		accum = append(accum, n)
	}
	failIfNotDeepEqual(t, accum, []int{0, 2, 4, 6, 8})
	failIfNotDeepEqual(t, iterator.Error(), nil)
}

func TestChannelIteratorWithDecreasingSequence(t *testing.T) {
	accum := make([]int, 0)
	iterator = iter.NewIntSeq(iter.Start(10), iter.Stop(-11), iter.Step(-5))
	for n := range iterator.Iter() {
		accum = append(accum, n)
	}
	failIfNotDeepEqual(t, accum, []int{10, 5, 0, -5, -10})
	failIfNotDeepEqual(t, iterator.Error(), nil)
}

func TestInvalidChannelIterator(t *testing.T) {
	accum := make([]int, 0)
	iterator = iter.NewIntSeq(iter.Start(-10), iter.Step(-5))
	for n := range iterator.Iter() {
		accum = append(accum, n)
	}
	failIfNotDeepEqual(t, accum, []int{})
	failIfNotDeepEqual(t, iterator.Error(), iter.ErrInvalidSequence)
}

const (
	small  = 1000
	medium = 500000
	big    = 1000000
)

var (
	smaIncr = [3]int{0, small, 1}
	smaDecr = [3]int{small, 0, -1}
	medIncr = [3]int{0, medium, 1}
	medDecr = [3]int{medium, 0, -1}
	bigIncr = [3]int{0, big, 1}
	bigDecr = [3]int{big, 0, -1}

	smaIncrOps = []iter.IntSeqOptFunc{iter.Start(smaIncr[0]), iter.Stop(smaIncr[1]), iter.Step(smaIncr[2])}
	smaDecrOps = []iter.IntSeqOptFunc{iter.Start(smaDecr[0]), iter.Stop(smaDecr[1]), iter.Step(smaDecr[2])}
	medIncrOps = []iter.IntSeqOptFunc{iter.Start(medIncr[0]), iter.Stop(medIncr[1]), iter.Step(medIncr[2])}
	medDecrOps = []iter.IntSeqOptFunc{iter.Start(medDecr[0]), iter.Stop(medDecr[1]), iter.Step(medDecr[2])}
	bigIncrOps = []iter.IntSeqOptFunc{iter.Start(bigIncr[0]), iter.Stop(bigIncr[1]), iter.Step(bigIncr[2])}
	bigDecrOps = []iter.IntSeqOptFunc{iter.Start(bigDecr[0]), iter.Stop(bigDecr[1]), iter.Step(bigDecr[2])}

	smaOps = [][]iter.IntSeqOptFunc{smaIncrOps, smaDecrOps}
	medOps = [][]iter.IntSeqOptFunc{medIncrOps, medDecrOps}
	bigOps = [][]iter.IntSeqOptFunc{bigIncrOps, bigDecrOps}
)

func getSliceBenchmark(b *testing.B, allOps [][]iter.IntSeqOptFunc) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, ops := range allOps {
				for n := range iter.NewIntSeq(ops...).All() {
					_ = n
				}
			}
		}
	}
}

func BenchmarkSmallSlice(b *testing.B) {
	getSliceBenchmark(b, smaOps)(b)
}

func BenchmarkMediumSlice(b *testing.B) {
	getSliceBenchmark(b, medOps)(b)
}

func BenchmarkBigSlice(b *testing.B) {
	getSliceBenchmark(b, bigOps)(b)
}

func getSequenceLoopBenchmark(b *testing.B, allOps [][]iter.IntSeqOptFunc) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, ops := range allOps {
				for seq, n := iter.NewIntSeqStart(ops...); seq.Continue(); n = seq.Get() {
					_ = n
				}
			}
		}
	}
}

func BenchmarkSmallSequenceLoop(b *testing.B) {
	getSequenceLoopBenchmark(b, smaOps)(b)
}

func BenchmarkMediumSequenceLoop(b *testing.B) {
	getSequenceLoopBenchmark(b, medOps)(b)
}

func BenchmarkBigSequenceLoop(b *testing.B) {
	getSequenceLoopBenchmark(b, bigOps)(b)
}

func getChannelIteratorBenchmark(b *testing.B, allOps [][]iter.IntSeqOptFunc) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, ops := range allOps {
				for n := range iter.NewIntSeq(ops...).Iter() {
					_ = n
				}
			}
		}
	}
}

func BenchmarkSmallChannelIterator(b *testing.B) {
	getChannelIteratorBenchmark(b, smaOps)(b)
}

func BenchmarkMediumChannelIterator(b *testing.B) {
	getChannelIteratorBenchmark(b, medOps)(b)
}

func BenchmarkBigChannelIterator(b *testing.B) {
	getChannelIteratorBenchmark(b, bigOps)(b)
}

func getThreeClauseForLoopBenchmark(b *testing.B, incr, decr [3]int) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := incr[0]; n < incr[1]; n += incr[2] {
				_ = n
			}
			for n := decr[0]; n > decr[1]; n += decr[2] {
				_ = n
			}
		}
	}
}

func BenchmarkSmallThreeClauseForLoop(b *testing.B) {
	getThreeClauseForLoopBenchmark(b, smaIncr, smaDecr)(b)
}

func BenchmarkMediumThreeClauseForLoop(b *testing.B) {
	getThreeClauseForLoopBenchmark(b, medIncr, medDecr)(b)
}

func BenchmarkBigThreeClauseForLoop(b *testing.B) {
	getThreeClauseForLoopBenchmark(b, bigIncr, bigDecr)(b)
}
