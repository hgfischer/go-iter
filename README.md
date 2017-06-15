# go-iter

[![Build Status](https://travis-ci.org/hgfischer/go-iter.svg?branch=master)](https://travis-ci.org/hgfischer/go-iter)

This is a small research on how to build number sequence generators for iteration.

I got inspired and puzzled by the article at [https://dev.to/loderunner/custom-range-iterators-in-go] so I decided to try some ideas myself, like adding support for Context in the channel based iterator, so it could be a safe option, and adding support for an interesting way to build structs by using Optional named arguments.

## Idioms

The `iter.Int` has support for 3 different idioms

### Slice Idiom

This is pretty similar to how the Python's `range` function works, which is generating the sequence and returning a list. The Go version in this package returns a slice, that can be iterated:

```code:go
    for n := iter.NewIntSeq(iter.Start(1), iter.Stop(100), iter.Step(2)).All() {
        fmt.Println(n)
    }
```

### Sequence Loop Idiom

This basically encapsulates a 3-clause for loop in a struct. There is not a single advantage here. But let's just test anyway:

```code:go
    for seq, n := iter.NewIntSeqStart(iter.Stop(10), iter.Step(2)); seq.Continue(); seq.Get() {
        fmt.Println(n)
    }
```

### Channel Iterator Idiom

This is, by far, the worst option, for many reasons [1][2]. Adding support for `context.Context` helps a bit, but this is still a slow option, just for the sake of using `range`:

```code:go
    for n := range iter.NewIntSeq(iter.Stop(100)).Iter() {
        fmt.Println(n)
    }
```

## Benchmarks

To run the benchmarks, you can either use `make bench` or run `go test -v -bench=. -benchmem ./...`. The following are the results I got in my desktop:

```code:go
BenchmarkSmallSlice-8                     200000              8025 ns/op
BenchmarkMediumSlice-8                       500           2821103 ns/op
BenchmarkBigSlice-8                          300           5726493 ns/op
BenchmarkSmallSequenceLoop-8              500000              3529 ns/op
BenchmarkMediumSequenceLoop-8               1000           1708549 ns/op
BenchmarkBigSequenceLoop-8                   500           3411812 ns/op
BenchmarkSmallChannelIterator-8            10000            107744 ns/op
BenchmarkMediumChannelIterator-8              20          56516534 ns/op
BenchmarkBigChannelIterator-8                 10         113113557 ns/op
BenchmarkSmallThreeClauseForLoop-8       3000000               528 ns/op
BenchmarkMediumThreeClauseForLoop-8         5000            259942 ns/op
BenchmarkBigThreeClauseForLoop-8            3000            526436 ns/op
PASS
ok      github.com/hgfischer/go-iter    20.067s
```

### Conclusion

This was a fun coding practice, but if you can, stick with the native 3-clause for loops. They are much simpler, and also the fastest option!