# GOGC

The `GOGC` variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is `GOGC=100`. Setting `GOGC=off` disables the garbage collector entirely. [runtime/debug.SetGCPercent](https://pkg.go.dev/runtime/debug#SetGCPercent) allows changing this percentage at run time.

`GOGC` is one of the oldest environment variable supported by the Go runtime. It’s possibly older than `GOROOT`, but nowhere near as well known.

`GOGC` controls the aggressiveness of the garbage collector. By default this value is assumed to be 100, which means garbage collection will not be triggered until the heap has grown by 100% since the previous collection. Effectively `GOGC=100` (the default) means the garbage collector will run each time the live heap doubles.

Setting this value higher, say `GOGC=200`, will delay the start of a garbage collection cycle until the live heap has grown to 200% of the previous size. Setting the value lower, say `GOGC=20` will cause the garbage collector to be triggered more often as less new data can be allocated on the heap before triggering a collection.

Setting `GOGC=off` will disable garbage collection entirely.

With the introduction of the low latency collector in Go 1.5, phrases like “trigger a garbage collection cycle” become more fluid, but the underlying message that values of `GOGC` greater than 100 mean the garbage collector will run less often, and for values of GOGC less than 100, more often, remains the same.