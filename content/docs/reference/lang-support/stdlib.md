
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Test results shown are for Linux.

Package | Importable | Passes tests
--- | --- | --- |
archive/tar |  [<span style="color: red">✗</span> no](#archivetar)  |  <span style="color: gray">✗</span> no  | 
archive/zip |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
bufio |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bufio)  | 
bytes |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bytes)  | 
compress/bzip2 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/flate |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/gzip |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#compressgzip)  | 
compress/lzw |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/zlib |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
container/heap |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
container/list |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
container/ring |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
context |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#context)  | 
crypto |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#crypto)  | 
crypto/aes |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoaes)  | 
crypto/cipher |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptocipher)  | 
crypto/des |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/dsa |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ecdsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoecdsa)  | 
crypto/ed25519 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoelliptic)  | 
crypto/hmac |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptorsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/subtle |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosubtle)  | 
crypto/tls |  [<span style="color: red">✗</span> no](#cryptotls)  |  <span style="color: gray">✗</span> no  | 
crypto/x509 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptox509)  | 
crypto/x509/pkix |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesql)  | 
database/sql/driver |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesqldriver)  | 
debug/buildinfo |  [<span style="color: red">✗</span> no](#debugbuildinfo)  |  <span style="color: gray">✗</span> no  | 
debug/dwarf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/elf |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debugelf)  | 
debug/gosym |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debuggosym)  | 
debug/macho |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/pe |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debugpe)  | 
debug/plan9obj |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
embed |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#embed)  | 
encoding |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/ascii85 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/asn1 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingasn1)  | 
encoding/base32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/base64 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/binary |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingbinary)  | 
encoding/csv |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/gob |  [<span style="color: red">✗</span> no](#encodinggob)  |  <span style="color: gray">✗</span> no  | 
encoding/hex |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/json |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingjson)  | 
encoding/pem |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingpem)  | 
encoding/xml |  [<span style="color: red">✗</span> no](#encodingxml)  |  <span style="color: gray">✗</span> no  | 
errors |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#errors)  | 
expvar |  [<span style="color: red">✗</span> no](#expvar)  |  <span style="color: gray">✗</span> no  | 
flag |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#flag)  | 
fmt |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#fmt)  | 
go/ast |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goast)  | 
go/build |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gobuild)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gobuildconstraint)  | 
go/constant |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goconstant)  | 
go/doc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#godoc)  | 
go/format |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goformat)  | 
go/importer |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goimporter)  | 
go/parser |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goparser)  | 
go/printer |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goprinter)  | 
go/scanner |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/token |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gotoken)  | 
go/types |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gotypes)  | 
hash |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#hashcrc32)  | 
hash/crc64 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/fnv |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/maphash |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#hashmaphash)  | 
html |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
html/template |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#htmltemplate)  | 
image |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/color |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagecolor)  | 
image/color/palette |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/draw |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagedraw)  | 
image/gif |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagegif)  | 
image/jpeg |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/png |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagepng)  | 
index/suffixarray |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#io)  | 
io/fs |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#iofs)  | 
io/ioutil |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
log |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#log)  | 
log/syslog |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#logsyslog)  | 
math |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/big |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbig)  | 
math/bits |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbits)  | 
math/cmplx |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathrand)  | 
mime |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mime)  | 
mime/multipart |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mimemultipart)  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http |  [<span style="color: red">✗</span> no](#nethttp)  |  <span style="color: gray">✗</span> no  | 
net/http/cgi |  [<span style="color: red">✗</span> no](#nethttpcgi)  |  <span style="color: gray">✗</span> no  | 
net/http/cookiejar |  [<span style="color: red">✗</span> no](#nethttpcookiejar)  |  <span style="color: gray">✗</span> no  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#nethttpfcgi)  |  <span style="color: gray">✗</span> no  | 
net/http/httptest |  [<span style="color: red">✗</span> no](#nethttphttptest)  |  <span style="color: gray">✗</span> no  | 
net/http/httptrace |  [<span style="color: red">✗</span> no](#nethttphttptrace)  |  <span style="color: gray">✗</span> no  | 
net/http/httputil |  [<span style="color: red">✗</span> no](#nethttphttputil)  |  <span style="color: gray">✗</span> no  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#nethttppprof)  |  <span style="color: gray">✗</span> no  | 
net/mail |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/netip |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netnetip)  | 
net/rpc |  [<span style="color: red">✗</span> no](#netrpc)  |  <span style="color: gray">✗</span> no  | 
net/rpc/jsonrpc |  [<span style="color: red">✗</span> no](#netrpcjsonrpc)  |  <span style="color: gray">✗</span> no  | 
net/smtp |  [<span style="color: red">✗</span> no](#netsmtp)  |  <span style="color: gray">✗</span> no  | 
net/textproto |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#nettextproto)  | 
net/url |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#neturl)  | 
os |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
os/exec |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#osexec)  | 
os/signal |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#ossignal)  | 
os/user |  [<span style="color: red">✗</span> no](#osuser)  |  <span style="color: gray">✗</span> no  | 
path |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path/filepath |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#pathfilepath)  | 
plugin |  [<span style="color: red">✗</span> no](#plugin)  |  <span style="color: gray">✗</span> no  | 
reflect |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
regexp |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#regexp)  | 
regexp/syntax |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#regexpsyntax)  | 
sort |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#sort)  | 
strconv |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
strings |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strings)  | 
sync |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync/atomic |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syncatomic)  | 
syscall |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syscall)  | 
testing |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/fstest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/iotest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/quick |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingquick)  | 
text/scanner |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
text/template |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#texttemplate)  | 
text/template/parse |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
time |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#time)  | 
time/tzdata |  [<span style="color: red">✗</span> no](#timetzdata)  |  <span style="color: gray">✗</span> no  | 
unicode |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 



## archive/tar


This package cannot be imported because the following dependencies cannot be compiled:

  * [os/user](#osuser)







## bufio



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestReadStringAllocs (0.00s)
        Unexpected number of allocations, got 0.000000, want 1
    FAIL
    FAIL	bufio	0.315s
    FAIL






## bytes



The compiler gave the following error when running the tests for this package:


    # bytes_test
    /home/ayke/src/github.com/tinygo-org/tinygo/src/sync/mutex.go:17:7: interp: running for more than 3m0s, timing out (executed calls: 304674)
      <badref> = icmp eqfatal error: unexpected signal during runtime execution
    [signal SIGSEGV: segmentation violation code=0x1 addr=0x30 pc=0x7f60eff95d44]
    
    runtime stack:
    runtime.throw({0x2c7b65?, 0x2?})
    	/usr/local/go/src/runtime/panic.go:992 +0x71
    runtime.sigpanic()
    	/usr/local/go/src/runtime/signal_unix.go:802 +0x389
    
    goroutine 21 [syscall]:
    runtime.cgocall(0x85f840, 0xc0069f5ba0)
    	/usr/local/go/src/runtime/cgocall.go:157 +0x5c fp=0xc0069f5b78 sp=0xc0069f5b40 pc=0x50e8dc
    tinygo.org/x/go-llvm._Cfunc_LLVMDumpValue(0x7f60bc3ba6c0)
    [...more lines following...]










## compress/gzip



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestGZIPFilesHaveZeroMTimes (0.67s)
        skipping test on non-builder
        SkipNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	compress/gzip	72.849s
    FAIL
















## context



The compiler gave the following error when running the tests for this package:


    # context_test
    /usr/local/go1.18.6/src/context/x_test.go:12:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestBackground: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:13:68: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTODO: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:14:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestWithCancel: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:15:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestParentFinishesChild: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:16:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestChildFinishesFirst: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:17:72: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestDeadline: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:18:71: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:19:79: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestCanceledTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:20:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestValues: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:21:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestAllocs: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:22:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestSimultaneousCancels: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:23:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestInterlockedCancels: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:24:76: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersCancel: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18.6/src/context/x_test.go:25:77: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    [...more lines following...]






## crypto



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.18.6/src/crypto/cipher/xor_amd64.go:18)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-949791DEA8E6E45F71C8B6AE668E125F19F36856:(crypto/cipher.xorBytes)
    
    ld.lld-14: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:68 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:68)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-949791DEA8E6E45F71C8B6AE668E125F19F36856:((*crypto/aes.aesCipherAsm).Encrypt)
    error: failed to link /tmp/tinygo2862753490/main: exit status 1
    FAIL






## crypto/aes



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/aes.decryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-CA4753056032B6CD353BF1476BC8AE01682D11C8:((*crypto/aes.aesCipherAsm).Decrypt)
    
    ld.lld-14: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:68 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:68)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-CA4753056032B6CD353BF1476BC8AE01682D11C8:((*crypto/aes.aesCipherAsm).Encrypt)
    error: failed to link /tmp/tinygo2032559460/main: exit status 1
    FAIL






## crypto/cipher



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.18.6/src/crypto/cipher/xor_amd64.go:18)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-25FB53323ADDEF3CCF37F77FC23E8FFFE1088DD9:(crypto/cipher.xorBytes)
    
    ld.lld-14: error: undefined symbol: crypto/aes.gcmAesInit
    >>> referenced by aes_gcm.go:56 (/usr/local/go1.18.6/src/crypto/aes/aes_gcm.go:56)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-25FB53323ADDEF3CCF37F77FC23E8FFFE1088DD9:(crypto/cipher.newGCMWithNonceAndTagSize)
    
    ld.lld-14: error: undefined symbol: crypto/aes.decryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-25FB53323ADDEF3CCF37F77FC23E8FFFE1088DD9:((*crypto/aes.aesCipherAsm).Decrypt)
    
    ld.lld-14: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:68 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:68)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-25FB53323ADDEF3CCF37F77FC23E8FFFE1088DD9:((*crypto/aes.aesCipherAsm).Encrypt)
    [...more lines following...]










## crypto/ecdsa



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.18.6/src/crypto/cipher/xor_amd64.go:18)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-96614A4D104165B16CF0826152A1C9CCA84DA924:((crypto/cipher.StreamReader).Read)
    
    ld.lld-14: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:68 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:68)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-96614A4D104165B16CF0826152A1C9CCA84DA924:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld-14: error: undefined symbol: crypto/elliptic.p256Mul
    >>> referenced by p256_asm.go:284 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:284)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-96614A4D104165B16CF0826152A1C9CCA84DA924:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:285 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:285)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-96614A4D104165B16CF0826152A1C9CCA84DA924:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:356 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:356)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-96614A4D104165B16CF0826152A1C9CCA84DA924:((*crypto/elliptic.p256Point).p256PointToAffine)
    [...more lines following...]






## crypto/ed25519



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestEd25519Vectors (0.00s)
        failed to run `go mod download -json filippo.io/mostly-harmless/ed25519vectors@v0.0.0-20210322192420-30a2d7243a94`, output: 
        FailNow is incomplete, requires runtime.Goexit()
        unexpected end of JSON input
        FailNow is incomplete, requires runtime.Goexit()
        failed to read ed25519vectors.json: open ed25519vectors.json: file does not exist
        FailNow is incomplete, requires runtime.Goexit()
        unexpected end of JSON input
        FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	crypto/ed25519	0.861s
    FAIL






## crypto/elliptic



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/elliptic.p256Mul
    >>> referenced by p256_asm.go:284 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:284)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8693EC8B52B2B3F99B7E1A375582EB30DEBD4532:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:285 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:285)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8693EC8B52B2B3F99B7E1A375582EB30DEBD4532:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:319 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:319)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8693EC8B52B2B3F99B7E1A375582EB30DEBD4532:((*crypto/elliptic.p256Point).p256PointToAffine)
    >>> referenced 20 more times
    
    ld.lld-14: error: undefined symbol: crypto/elliptic.p256PointDoubleAsm
    >>> referenced by p256_asm.go:460 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:460)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8693EC8B52B2B3F99B7E1A375582EB30DEBD4532:((*crypto/elliptic.p256Point).p256ScalarMult)
    >>> referenced by p256_asm.go:461 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:461)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8693EC8B52B2B3F99B7E1A375582EB30DEBD4532:((*crypto/elliptic.p256Point).p256ScalarMult)
    >>> referenced by p256_asm.go:462 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:462)
    [...more lines following...]














## crypto/rsa



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.18.6/src/crypto/cipher/xor_amd64.go:18)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B6384F61356C78646301C371098314DAECF7356F:((crypto/cipher.StreamReader).Read)
    
    ld.lld-14: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:68 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:68)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B6384F61356C78646301C371098314DAECF7356F:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld-14: error: undefined symbol: crypto/elliptic.p256Mul
    >>> referenced by p256_asm.go:284 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:284)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B6384F61356C78646301C371098314DAECF7356F:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:285 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:285)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B6384F61356C78646301C371098314DAECF7356F:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:356 (/usr/local/go1.18.6/src/crypto/elliptic/p256_asm.go:356)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B6384F61356C78646301C371098314DAECF7356F:((*crypto/elliptic.p256Point).p256PointToAffine)
    [...more lines following...]












## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	0.001s
    FAIL






## crypto/tls



The compiler gave the following error when this package was imported:


    ld.lld-14: error: undefined symbol: vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by cpu_x86.go:64 (/usr/local/go1.18.6/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:64)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D1D261DAB7C6E5F4D5E1B9FA8DF98FB03DF78C31:(runtime.run$1$gowrapper)
    >>> referenced by cpu_x86.go:70 (/usr/local/go1.18.6/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:70)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D1D261DAB7C6E5F4D5E1B9FA8DF98FB03DF78C31:(runtime.run$1$gowrapper)
    >>> referenced by cpu_x86.go:109 (/usr/local/go1.18.6/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:109)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D1D261DAB7C6E5F4D5E1B9FA8DF98FB03DF78C31:(runtime.run$1$gowrapper)
    >>> referenced 1 more times
    
    ld.lld-14: error: undefined symbol: vendor/golang.org/x/sys/cpu.xgetbv
    >>> referenced by cpu_x86.go:88 (/usr/local/go1.18.6/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:88)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D1D261DAB7C6E5F4D5E1B9FA8DF98FB03DF78C31:(runtime.run$1$gowrapper)
    error: failed to link /tmp/tinygo3417349174/main: exit status 1






## crypto/x509



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.18.6/src/crypto/cipher/xor_amd64.go:18)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EDAFF76F77DB8AE43A27641CDCE98BBE04FAD89:(crypto/cipher.xorBytes)
    
    ld.lld-14: error: undefined symbol: crypto/aes.decryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EDAFF76F77DB8AE43A27641CDCE98BBE04FAD89:((*crypto/aes.aesCipherAsm).Decrypt)
    
    ld.lld-14: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:68 (/usr/local/go1.18.6/src/crypto/aes/cipher_asm.go:68)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EDAFF76F77DB8AE43A27641CDCE98BBE04FAD89:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld-14: error: undefined symbol: vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by cpu_x86.go:64 (/usr/local/go1.18.6/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:64)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EDAFF76F77DB8AE43A27641CDCE98BBE04FAD89:(runtime.run$1$gowrapper)
    [...more lines following...]








## database/sql



The compiler gave the following error when running the tests for this package:


    # database/sql
    /usr/local/go1.18.6/src/database/sql/sql_test.go:4556:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)
    FAIL






## database/sql/driver



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Bytes()
    FAIL	database/sql/driver	0.002s
    FAIL






## debug/buildinfo



The compiler gave the following error when this package was imported:


    # debug/buildinfo
    /usr/local/go1.18.6/src/debug/buildinfo/buildinfo.go:78:19: ParseBuildInfo not declared by package debug








## debug/elf



The compiler gave the following error when running the tests for this package:


    # debug/elf
    /usr/local/go1.18.6/src/debug/elf/file_test.go:905:10: ResolveIPAddr not declared by package net
    FAIL






## debug/gosym



The compiler gave the following error when running the tests for this package:


    panic: runtime error: nil pointer dereference
    FAIL	debug/gosym	0.004s
    FAIL








## debug/pe



The compiler gave the following error when running the tests for this package:


    panic: runtime error: index out of range
    FAIL	debug/pe	0.004s
    FAIL








## embed



The compiler gave the following error when running the tests for this package:


    ld.lld-14: error: undefined symbol: vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-C43B0A7BA1F52489FB49202AC232B5E3D0CEDFF4:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-C43B0A7BA1F52489FB49202AC232B5E3D0CEDFF4:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-C43B0A7BA1F52489FB49202AC232B5E3D0CEDFF4:(runtime.initAll)
    >>> referenced 1 more times
    
    ld.lld-14: error: undefined symbol: vendor/golang.org/x/sys/cpu.xgetbv
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-C43B0A7BA1F52489FB49202AC232B5E3D0CEDFF4:(runtime.initAll)
    error: failed to link /tmp/tinygo1088601438/main: exit status 1
    FAIL










## encoding/asn1



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Addr()
    FAIL	encoding/asn1	0.003s
    FAIL










## encoding/binary



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Slice()
    FAIL	encoding/binary	0.002s
    FAIL








## encoding/gob



The compiler gave the following error when this package was imported:


    # encoding/gob
    /usr/local/go1.18.6/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect








## encoding/json



The compiler gave the following error when running the tests for this package:


    # encoding/json
    /usr/local/go1.18.6/src/encoding/json/bench_test.go:344:22: StructOf not declared by package reflect
    /usr/local/go1.18.6/src/encoding/json/stream_test.go:298:15: Pipe not declared by package net
    FAIL






## encoding/pem



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	encoding/pem	0.007s
    FAIL






## encoding/xml



The compiler gave the following error when this package was imported:


    # encoding/xml
    /usr/local/go1.18.6/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
    /usr/local/go1.18.6/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)






## errors



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: reflect.Zero()
    FAIL	errors	0.006s
    FAIL






## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## flag



The compiler gave the following error when running the tests for this package:


    flag provided but not defined: -x
    panic: unimplemented: reflect.Zero()
    FAIL	flag	0.004s
    FAIL






## fmt



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Slice()
    FAIL	fmt	0.003s
    FAIL






## go/ast



The compiler gave the following error when running the tests for this package:


    panic: runtime error: type assert failed
    FAIL	go/ast	0.014s
    FAIL






## go/build



The compiler gave the following error when running the tests for this package:


    # go/build
    /usr/local/go1.18.6/src/go/build/build_test.go:556:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:577:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:599:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:629:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:641:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:642:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:695:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:696:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18.6/src/go/build/build_test.go:697:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    FAIL






## go/build/constraint



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	go/build/constraint	0.007s
    FAIL






## go/constant



The compiler gave the following error when running the tests for this package:


    panic: runtime error: divide by zero
    FAIL	go/constant	0.013s
    FAIL






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	0.011s
    FAIL






## go/format



The compiler gave the following error when running the tests for this package:


    FAIL	go/format	374.574s
    FAIL






## go/importer



The compiler gave the following error when running the tests for this package:


    panic: runtime error: nil pointer dereference
    FAIL	go/importer	0.001s
    FAIL






## go/parser



The compiler gave the following error when running the tests for this package:


    FAIL	go/parser	750.695s
    FAIL






## go/printer



The compiler gave the following error when running the tests for this package:


    FAIL	go/printer	740.362s
    FAIL








## go/token



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18.6/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## go/types



The compiler gave the following error when running the tests for this package:


    # go/types_test
    /usr/local/go1.18.6/src/go/types/self_test.go:98:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    FAIL










## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	0.007s
    FAIL










## hash/maphash



The compiler gave the following error when running the tests for this package:


    panic: runtime error: out of memory
    FAIL	hash/maphash	100.856s
    FAIL








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	0.006s
    FAIL








## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	0.007s
    FAIL








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	0.094s
    FAIL






## image/gif



The compiler gave the following error when running the tests for this package:


    # image/gif
    /usr/local/go1.18.6/src/image/gif/reader_test.go:415:14: SetGCPercent not declared by package debug
    /usr/local/go1.18.6/src/image/gif/reader_test.go:415:33: SetGCPercent not declared by package debug
    /usr/local/go1.18.6/src/image/gif/reader_test.go:421:26: s1.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    /usr/local/go1.18.6/src/image/gif/reader_test.go:421:41: s0.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    FAIL








## image/png



The compiler gave the following error when running the tests for this package:


    panic: runtime error: out of memory
    FAIL	image/png	2.921s
    FAIL








## io



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestMultiWriter_WriteStringSingleAlloc (0.00s)
        num allocations = 0; want 1
    --- FAIL: TestMultiWriterSingleChainFlatten (0.00s)
        multiWriter did not flatten chained multiWriters: expected writeDepth 12, got 4
    --- FAIL: TestMultiReaderFlatten (0.00s)
        multiReader did not flatten chained multiReaders: expected readDepth 3, got 1
    --- FAIL: TestMultiReaderFreesExhaustedReaders (5.00s)
        timeout waiting for collection of buf1
        FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	io	5.083s
    FAIL






## io/fs



The compiler gave the following error when running the tests for this package:


    FAIL	io/fs	0.010s
    FAIL








## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.01s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/09/29 18:41:45.180641 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/09/29 18:41:45.180706 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/09/29 18:41:45.181471 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/09/29 18:41:45.181544 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/09/29 18:41:45.182159 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/09/29 18:41:45.182420 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/09/29 18:41:45.182505 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/09/29 18:41:45.182696 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/syslog



The compiler gave the following error when running the tests for this package:


    # log/syslog
    /usr/local/go1.18.6/src/log/syslog/syslog_test.go:22:25: PacketConn not declared by package net
    /usr/local/go1.18.6/src/log/syslog/syslog_test.go:35:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /usr/local/go1.18.6/src/log/syslog/syslog_test.go:110:15: ListenPacket not declared by package net
    FAIL








## math/big



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18.6/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## math/bits



The compiler gave the following error when running the tests for this package:


    panic: runtime error: divide by zero
    FAIL	math/bits	0.003s
    FAIL








## math/rand



The compiler gave the following error when running the tests for this package:


    # math/rand_test
    /usr/local/go1.18.6/src/math/rand/regress_test.go:38:18: rv.Type().Method undefined (type reflect.Type has no field or method Method)
    /usr/local/go1.18.6/src/math/rand/regress_test.go:39:12: rv.Method undefined (type reflect.Value has no field or method Method)
    FAIL






## mime



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	mime	0.007s
    FAIL






## mime/multipart



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapRange()
    FAIL	mime/multipart	0.007s
    FAIL










## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [net/http/httptrace](#nethttphttptrace)





## net/http/cgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [net/http](#nethttp)





## net/http/cookiejar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)
  * [net/http/cgi](#nethttpcgi)





## net/http/httptest


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [net/http](#nethttp)





## net/http/httptrace


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)





## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)







## net/netip



The compiler gave the following error when running the tests for this package:


    # net/netip_test
    /usr/local/go1.18.6/src/net/netip/netip_test.go:1815:25: UDPAddr not declared by package net
    /usr/local/go1.18.6/src/net/netip/fuzz_test.go:177:32: stdip.IsPrivate undefined (type net.IP has no field or method IsPrivate)
    FAIL






## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/gob](#encodinggob)
  * [net/http](#nethttp)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/rpc](#netrpc)





## net/smtp


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)





## net/textproto



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	net/textproto	0.006s
    FAIL






## net/url



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18.6/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL








## os/exec



The compiler gave the following error when running the tests for this package:


    # os/user
    ../../../../../usr/local/go1.18.6/src/os/user/cgo_lookup_unix.go:18:6: not implemented: build constraints in #cgo line
    FAIL






## os/signal



The compiler gave the following error when running the tests for this package:


    # os/signal
    /usr/local/go1.18.6/src/os/signal/signal_test.go:435:23: t.Deadline undefined (type *testing.T has no field or method Deadline)
    /usr/local/go1.18.6/src/os/signal/signal_test.go:479:23: t.Deadline undefined (type *testing.T has no field or method Deadline)
    /usr/local/go1.18.6/src/os/signal/signal_test.go:549:24: t.Deadline undefined (type *testing.T has no field or method Deadline)
    /usr/local/go1.18.6/src/os/signal/signal_test.go:599:23: t.Deadline undefined (type *testing.T has no field or method Deadline)
    /usr/local/go1.18.6/src/os/signal/signal_test.go:735:25: t.Deadline undefined (type *testing.T has no field or method Deadline)
    FAIL






## os/user



The compiler gave the following error when this package was imported:


    # os/user
    ../../../../../usr/local/go1.18.6/src/os/user/cgo_lookup_unix.go:18:6: not implemented: build constraints in #cgo line








## path/filepath



The compiler gave the following error when running the tests for this package:


    FAIL	path/filepath	0.011s
    FAIL






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    ../../../../../usr/local/go1.18.6/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line








## regexp



The compiler gave the following error when running the tests for this package:


    FAIL	regexp	0.266s
    FAIL






## regexp/syntax



The compiler gave the following error when running the tests for this package:


    # regexp/syntax
    /usr/local/go1.18.6/src/unicode/letter.go:122:6: interp: running for more than 3m0s, timing out (executed calls: 2918928)
      <badref> = alloca fatal error: unexpected signal during runtime execution
    [signal SIGSEGV: segmentation violation code=0x1 addr=0x108 pc=0x7f83239ef3b7]
    
    runtime stack:
    runtime.throw({0x2c7b65?, 0x7f83239c054f?})
    	/usr/local/go/src/runtime/panic.go:992 +0x71
    runtime.sigpanic()
    	/usr/local/go/src/runtime/signal_unix.go:802 +0x389
    
    goroutine 8 [syscall]:
    runtime.cgocall(0x85f840, 0xc004fadba0)
    	/usr/local/go/src/runtime/cgocall.go:157 +0x5c fp=0xc004fadb78 sp=0xc004fadb40 pc=0x50e8dc
    tinygo.org/x/go-llvm._Cfunc_LLVMDumpValue(0x2f79250)
    [...more lines following...]






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (4.91s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
        Stable      100 elements:         900 Swap,        774 Less
        Stable      300 elements:        3896 Swap,       3023 Less
        Stable     1000 elements:       19935 Swap,      12480 Less
        Stable     3000 elements:       82471 Swap,      42842 Less
        Stable    10000 elements:      349396 Swap,     165950 Less
        Stable    30000 elements:     1339352 Swap,     558151 Less
        Stable   100000 elements:     5787542 Swap,    2084840 Less
        Stable   300000 elements:    20444960 Swap,    6868106 Less
        Stable  1000000 elements:    84504188 Swap,   25119300 Less
    --- FAIL: TestCountSortOps (0.99s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## strings



The compiler gave the following error when running the tests for this package:


    # strings_test
    /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/runtime.go:37:6: interp: running for more than 3m0s, timing out (executed calls: 303856)
      callfatal error: unexpected signal during runtime execution
    [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x7ff97639c5d8]
    
    runtime stack:
    runtime.throw({0x2c7b65?, 0x7ff97923d1f6?})
    	/usr/local/go/src/runtime/panic.go:992 +0x71
    runtime.sigpanic()
    	/usr/local/go/src/runtime/signal_unix.go:802 +0x389
    
    goroutine 35 [syscall]:
    runtime.cgocall(0x85f840, 0xc004269ba0)
    	/usr/local/go/src/runtime/cgocall.go:157 +0x5c fp=0xc004269b78 sp=0xc004269b40 pc=0x50e8dc
    tinygo.org/x/go-llvm._Cfunc_LLVMDumpValue(0x7ff92cafd6b0)
    [...more lines following...]








## sync/atomic



The compiler gave the following error when running the tests for this package:


    # sync/atomic_test
    /usr/local/go1.18.6/src/sync/atomic/atomic_test.go:1202:14: SetGCPercent not declared by package debug
    /usr/local/go1.18.6/src/sync/atomic/atomic_test.go:1202:33: SetGCPercent not declared by package debug
    FAIL






## syscall



The compiler gave the following error when running the tests for this package:


    # os/user
    ../../../../../usr/local/go1.18.6/src/os/user/cgo_lookup_unix.go:18:6: not implemented: build constraints in #cgo line
    FAIL












## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	0.004s
    FAIL










## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	0.003s
    FAIL








## time



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18.6/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## time/tzdata



The compiler gave the following error when this package was imported:


    ld.lld-14: error: undefined symbol: time.registerLoadFromEmbeddedTZData
    >>> referenced by scheduler_any.go:25 (/home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/scheduler_any.go:25)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-191D9438BAF043FD1177ECD11661EC8F3C4DE9FE:(runtime.run$1$gowrapper)
    error: failed to link /tmp/tinygo831970216/main: exit status 1













