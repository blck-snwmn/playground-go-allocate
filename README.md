# playground-go-allocate
## Benchmark
```
go test -bench=. -benchmem
```
### map
```
pkg: github.com/blck-snwmn/playground-go-allocate/map
Benchmark_setMap-16                     16622966                69.88 ns/op            0 B/op          0 allocs/op
Benchmark_returnGenMap-16                5555923               221.1 ns/op           256 B/op          2 allocs/op
Benchmark_returnGeneratedMap-16         698091499                1.789 ns/op           0 B/op          0 allocs/op
```


### slice
```
pkg: github.com/blck-snwmn/playground-go-allocate/slice
Benchmark_genSlice-16           458030658                2.586 ns/op           0 B/op          0 allocs/op
Benchmark_returnGenSlice-16     79088181                14.19 ns/op            3 B/op          1 allocs/op
Benchmark_returnSlice-16        686081095                1.799 ns/op           0 B/op          0 allocs/op
```

## Benchmark with gcflags
```
go test -gcflags "-m" -bench=. -benchmem
```


```
# github.com/blck-snwmn/playground-go-allocate/slice [github.com/blck-snwmn/playground-go-allocate/slice.test]
./slice_test.go:30:6: can inline none
./slice_test.go:11:7: inlining call to none
./slice_test.go:15:6: can inline Benchmark_returnGenSlice
./slice_test.go:18:7: inlining call to none
./slice_test.go:22:6: can inline Benchmark_returnSlice
./slice_test.go:25:7: inlining call to none
./slice.go:4:15: b does not escape
./slice.go:12:15: []byte{...} escapes to heap
./slice_test.go:7:25: b does not escape
./slice_test.go:9:15: make([]byte, 3) does not escape
./slice_test.go:15:31: b does not escape
./slice_test.go:22:28: b does not escape
# github.com/blck-snwmn/playground-go-allocate/map [github.com/blck-snwmn/playground-go-allocate/map.test]
./map_test.go:28:6: can inline none
./map_test.go:9:7: inlining call to none
./map_test.go:13:6: can inline Benchmark_returnGenMap
./map_test.go:16:7: inlining call to none
./map_test.go:20:6: can inline Benchmark_returnGeneratedMap
./map_test.go:23:7: inlining call to none
./map.go:4:13: m does not escape
./map.go:13:23: map[int]string{...} escapes to heap
./map_test.go:5:23: b does not escape
./map_test.go:7:15: make(map[int]string) does not escape
./map_test.go:13:29: b does not escape
./map_test.go:20:35: b does not escape
./map.go:21:23: map[int]string{...} escapes to heap
# github.com/blck-snwmn/playground-go-allocate/map.test
/tmp/go-build4285438691/b001/_testmain.go:41:6: can inline init.0
/tmp/go-build4285438691/b001/_testmain.go:49:24: inlining call to testing.MainStart
/tmp/go-build4285438691/b001/_testmain.go:49:42: testdeps.TestDeps{} escapes to heap
/tmp/go-build4285438691/b001/_testmain.go:49:24: &testing.M{...} escapes to heap
# github.com/blck-snwmn/playground-go-allocate/slice.test
/tmp/go-build755518354/b001/_testmain.go:41:6: can inline init.0
/tmp/go-build755518354/b001/_testmain.go:49:24: inlining call to testing.MainStart
/tmp/go-build755518354/b001/_testmain.go:49:42: testdeps.TestDeps{} escapes to heap
/tmp/go-build755518354/b001/_testmain.go:49:24: &testing.M{...} escapes to heap
```