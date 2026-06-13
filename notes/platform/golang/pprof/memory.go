package pprof

// [Flag] --memprofile=<profile name>
// --memprofilerate // The rate of 1 will lead to collection of information about all allocations, but it can slow down execution. The default sampling rate is 1 sample per 512KB of allocated memory.
// --inuse/alloc_space // collected with net/http/pprof on live applications,
// --inuse/alloc_objects // collected at program end
// [Flag for reporting granularity]
// --functions
// --lines
// --addresses
// --files

// [Sample]
// heap profile: 4: 266528 [123: 11284472] @ heap/1048576
// 1: 262144 [4: 376832] @ 0x28d9f 0x2a201 0x2a28a 0x2624d 0x26188 0x94ca3 0x94a0b 0x17add6 0x17ae9f 0x1069d3 0xfe911 0xf0a3e 0xf0d22 0x21a70
// #   0x2a201 cnew+0xc1   runtime/malloc.goc:718
// #   0x2a28a runtime.cnewarray+0x3a          runtime/malloc.goc:731
// #   0x2624d makeslice1+0x4d             runtime/slice.c:57
// #   0x26188 runtime.makeslice+0x98          runtime/slice.c:38
// #   0x94ca3 bytes.makeSlice+0x63            bytes/buffer.go:191
// #   0x94a0b bytes.(*Buffer).ReadFrom+0xcb       bytes/buffer.go:163
// #   0x17add6    io/ioutil.readAll+0x156         io/ioutil/ioutil.go:32
// #   0x17ae9f    io/ioutil.ReadAll+0x3f          io/ioutil/ioutil.go:41
// #   0x1069d3    godoc/vfs.ReadFile+0x133            godoc/vfs/vfs.go:44
// #   0xfe911 godoc.func·023+0x471            godoc/meta.go:80
// #   0xf0a3e godoc.(*Corpus).updateMetadata+0x9e     godoc/meta.go:101
// #   0xf0d22 godoc.(*Corpus).refreshMetadataLoop+0x42    godoc/meta.go:141

// 2: 4096 [2: 4096] @ 0x28d9f 0x29059 0x1d252 0x1d450 0x106993 0xf1225 0xe1489 0xfbcad 0x21a70
// #   0x1d252 newdefer+0x112              runtime/panic.c:49
// #   0x1d450 runtime.deferproc+0x10          runtime/panic.c:132
// #   0x106993    godoc/vfs.ReadFile+0xf3         godoc/vfs/vfs.go:43
// #   0xf1225 godoc.(*Corpus).parseFile+0x75      godoc/parser.go:20

// 1: 262144 [4: 376832]
// number of currently live objects,
// amount of memory occupied by live objects,
// total number of allocations
// amount of memory occupied by all allocations

// [Solution]
// 1. [combine objects to large objects] to reduce the number of memory allocations
// 2. [Local variables that escape]
// for k, v := range m {
//     k, v := k, v   // copy for capturing by the goroutine
//     go func() {/* use k and v */}()
// }
// =>
// for k, v := range m {
//     x := struct{ k, v string }{k, v}   // copy for capturing by the goroutine
//     go func() {/* use x.k and x.v */}()
// }
// 3.  If you know a typical size of the slice, you can [preallocate a backing array]
// type X struct {
//     buf      []byte
//     bufArray [16]byte // Buf usually does not grow beyond 16 bytes.
// }

// func MakeX() *X {
//     x := &X{}
//     // Preinitialize buf with the backing array.
//     x.buf = x.bufArray[:0]
//     return x
// }
// 4. Use [smaller data types]. For example, use int8 instead of int
// 5. Objects that [do not contain any pointers] (note that strings, slices, maps and chans contain implicit pointers), are not scanned by garbage collector. a 1GB byte slice virtually does not affect garbage collection time. Some possibilities are: replace pointers with indices, split object into two parts one of which does not contain pointers.
// 6. [sync.Pool] type that allows to reuse the same object several times in between garbage collections. However, be aware that, as any manual memory management scheme, incorrect use of sync.Pool can lead to <use-after-free> bugs.
