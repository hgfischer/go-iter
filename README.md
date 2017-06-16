# go-iter

[![Build Status](https://travis-ci.org/hgfischer/go-iter.svg?branch=master)](https://travis-ci.org/hgfischer/go-iter)

This is a small research on how to build number sequence generators for iteration.

I got inspired and puzzled by the article at [https://dev.to/loderunner/custom-range-iterators-in-go] so I decided to try some ideas myself, like adding support for Context in the channel based iterator, so it could be a safe option, and adding support for an interesting way to build structs by using Optional named arguments.

## Idioms

The `iter.Int` has support for 3 different idioms

### Slice Idiom

This is pretty similar to how the Python's `range` function works, which is generating the sequence and returning a list. The Go version in this package returns a slice, that can be iterated:

```code:go
    for n := range iter.NewIntSeq(iter.Start(1), iter.Stop(100), iter.Step(2)).All() {
        fmt.Println(n)
    }
```

### Sequence Loop Idiom

This is close to how the Python's `xrange` function works. It basically encapsulates a 3-clause for loop in a struct. There is not a single advantage here. But let's just test anyway:

```code:go
    for seq, n := iter.NewIntSeqStart(iter.Stop(10), iter.Step(2)); seq.Continue(); seq.Get() {
        fmt.Println(n)
    }
```

### Channel Iterator Idiom

This is, by far, the worst option, for many reasons including memory and resource leaks [1], and worse performance because of locking & context switching. Adding support for `context.Context` helps a bit, but this is still a slow option, just for the sake of using `range`:

```code:go
    for n := range iter.NewIntSeq(iter.Stop(100)).Iter() {
        fmt.Println(n)
    }
```

## Benchmarks

To run the benchmarks, you can either use `make bench` or run `go test -v -bench=. -benchmem ./...`. The following are the results I got in my desktop:

```code:go
BenchmarkSmallSlice-8                     200000              6748 ns/op           16512 B/op          4 allocs/op
BenchmarkMediumSlice-8                       500           3034537 ns/op         8011908 B/op          4 allocs/op
BenchmarkBigSlice-8                          200           5960648 ns/op        16007296 B/op          4 allocs/op
BenchmarkSmallSequenceLoop-8              500000              3531 ns/op             128 B/op          2 allocs/op
BenchmarkMediumSequenceLoop-8               1000           1717939 ns/op             128 B/op          2 allocs/op
BenchmarkBigSequenceLoop-8                   500           3447407 ns/op             128 B/op          2 allocs/op
BenchmarkSmallChannelIterator-8            10000            108716 ns/op           16512 B/op          4 allocs/op
BenchmarkMediumChannelIterator-8              20          56450098 ns/op         8012062 B/op          4 allocs/op
BenchmarkBigChannelIterator-8                 10         112913903 ns/op        16007483 B/op          4 allocs/op
BenchmarkSmallThreeClauseForLoop-8       3000000               531 ns/op               0 B/op          0 allocs/op
BenchmarkMediumThreeClauseForLoop-8         5000            261893 ns/op               0 B/op          0 allocs/op
BenchmarkBigThreeClauseForLoop-8            3000            525303 ns/op               0 B/op          0 allocs/op
```

### Conclusion

This was a fun coding practice, but if you can, stick with the native 3-clause for loops. They are much simpler, and also the fastest option!

## References

[1] https://github.com/golang/go/issues/19702
