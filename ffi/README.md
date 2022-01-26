You cannot use a cgo shared library in a Go program, because you cannot have multiple Go runtimes in the same process.

Trying to do so will give the error:
```
# command-line-arguments
cgo-gcc-prolog:67:33: warning: unused variable '_cgo_a' [-Wunused-variable]
fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x43fd5e2]

goroutine 1 [running, locked to thread]:
runtime.throw({0x40a875b?, 0x1c00011b800?})
        /usr/local/go/src/runtime/panic.go:992 +0x71 fp=0x1c00004a960 sp=0x1c00004a930 pc=0x402f6d1
runtime: unexpected return pc for runtime.sigpanic called from 0x43fd5e2
stack: frame={sp:0x1c00004a960, fp:0x1c00004a9b0} stack=[0x1c00004a000,0x1c00004b000)
....
0x000001c00004aaa0:  0x0000000000000000  0x0000000000000000 
runtime.sigpanic()
        /usr/local/go/src/runtime/signal_unix.go:781 +0x3a9 fp=0x1c00004a9b0 sp=0x1c00004a960 pc=0x4043449
exit status 2
```