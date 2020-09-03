
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

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
crypto/ed25519 |  [<span style="color: red">✗</span> no](#cryptoed25519)  |  <span style="color: gray">✗</span> no  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoelliptic)  | 
crypto/hmac |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptohmac)  | 
crypto/md5 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptorsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/subtle |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosubtle)  | 
crypto/tls |  [<span style="color: red">✗</span> no](#cryptotls)  |  <span style="color: gray">✗</span> no  | 
crypto/x509 |  [<span style="color: red">✗</span> no](#cryptox509)  |  <span style="color: gray">✗</span> no  | 
crypto/x509/pkix |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesql)  | 
database/sql/driver |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesqldriver)  | 
debug/dwarf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/elf |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debugelf)  | 
debug/gosym |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debuggosym)  | 
debug/macho |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/pe |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debugpe)  | 
debug/plan9obj |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
embed |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/ascii85 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/asn1 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingasn1)  | 
encoding/base32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/base64 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingbase64)  | 
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
image |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#image)  | 
image/color |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagecolor)  | 
image/color/palette |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/draw |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagedraw)  | 
image/gif |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagegif)  | 
image/jpeg |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagejpeg)  | 
image/png |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagepng)  | 
index/suffixarray |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#io)  | 
io/fs |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io/ioutil |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#ioioutil)  | 
log |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#log)  | 
log/syslog |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#logsyslog)  | 
math |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/big |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbig)  | 
math/bits |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbits)  | 
math/cmplx |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathrand)  | 
mime |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mime)  | 
mime/multipart |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mimemultipart)  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mimequotedprintable)  | 
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
strconv |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strconv)  | 
strings |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strings)  | 
sync |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync/atomic |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syncatomic)  | 
syscall |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syscall)  | 
testing |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/fstest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/iotest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/quick |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingquick)  | 
text/scanner |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#texttabwriter)  | 
text/template |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#texttemplate)  | 
text/template/parse |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#texttemplateparse)  | 
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


    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo1351446372/main.o:(bufio_test.TestZeroReader)
    error: failed to link /tmp/tinygo1351446372/main: exit status 1
    FAIL






## bytes



The compiler gave the following error when running the tests for this package:


    # bytes_test
    /home/ayke/src/github.com/tinygo-org/tinygo/src/sync/mutex.go:17:7: interp: running for more than 180 seconds, timing out (executed calls: 379544)
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !3046
    
    traceback:
    /home/ayke/src/github.com/tinygo-org/tinygo/src/sync/mutex.go:17:7:
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !3046
    /usr/local/go1.17/src/math/rand/rand.go:388:11:
      call void @"(*sync.Mutex).Lock"(%sync.Mutex* %1, i8* undef), !dbg !3058
    /usr/local/go1.17/src/math/rand/rand.go:84:50:
      %3 = call i64 @"interface:{Int63:func:{}{basic:int64},Seed:func:{basic:int64}{}}.Int63$invoke"(i8* %invoke.func.value, i64 %invoke.func.typecode, i8* undef), !dbg !3031
    /usr/local/go1.17/src/math/rand/rand.go:98:52:
      %0 = call i64 @"(*math/rand.Rand).Int63"(%"math/rand.Rand"* %r, i8* undef), !dbg !3030
    /usr/local/go1.17/src/math/rand/rand.go:133:14:
      %11 = call i32 @"(*math/rand.Rand).Int31"(%"math/rand.Rand"* %r, i8* undef), !dbg !3060
    [...more lines following...]










## compress/gzip



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo1110457428/main.o:(compress/gzip.TestIssue6550)
    error: failed to link /tmp/tinygo1110457428/main: exit status 1
    FAIL
















## context



The compiler gave the following error when running the tests for this package:


    # context_test
    /usr/local/go1.17/src/context/x_test.go:12:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestBackground: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:13:68: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTODO: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:14:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestWithCancel: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:15:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestParentFinishesChild: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:16:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestChildFinishesFirst: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:17:72: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestDeadline: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:18:71: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTimeout: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:19:79: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestCanceledTimeout: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:20:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestValues: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:21:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestAllocs: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:22:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestSimultaneousCancels: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:23:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestInterlockedCancels: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:24:76: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersCancel: missing method Deadline
    /usr/local/go1.17/src/context/x_test.go:25:77: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersTimeout: missing method Deadline
    [...more lines following...]






## crypto



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.17/src/crypto/cipher/xor_amd64.go:18)
    >>>               /tmp/tinygo525294958/main.o:(crypto/cipher.xorBytes)
    
    ld.lld-13: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:69 (/usr/local/go1.17/src/crypto/aes/cipher_asm.go:69)
    >>>               /tmp/tinygo525294958/main.o:((*crypto/aes.aesCipherAsm).Encrypt)
    error: failed to link /tmp/tinygo525294958/main: exit status 1
    FAIL






## crypto/aes



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:69 (/usr/local/go1.17/src/crypto/aes/cipher_asm.go:69)
    >>>               /tmp/tinygo3279067475/main.o:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld-13: error: undefined symbol: crypto/aes.decryptBlockAsm
    >>> referenced by cipher_asm.go:82 (/usr/local/go1.17/src/crypto/aes/cipher_asm.go:82)
    >>>               /tmp/tinygo3279067475/main.o:((*crypto/aes.aesCipherAsm).Decrypt)
    error: failed to link /tmp/tinygo3279067475/main: exit status 1
    FAIL






## crypto/cipher



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/aes.gcmAesData
    >>> referenced by aes_gcm.go:115 (/usr/local/go1.17/src/crypto/aes/aes_gcm.go:115)
    >>>               /tmp/tinygo4141715439/main.o:(interface:{NonceSize:func:{}{basic:int},Open:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8,named:error},Overhead:func:{}{basic:int},Seal:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8}}.Seal$invoke)
    >>> referenced by aes_gcm.go:122 (/usr/local/go1.17/src/crypto/aes/aes_gcm.go:122)
    >>>               /tmp/tinygo4141715439/main.o:(interface:{NonceSize:func:{}{basic:int},Open:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8,named:error},Overhead:func:{}{basic:int},Seal:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8}}.Seal$invoke)
    >>> referenced by aes_gcm.go:168 (/usr/local/go1.17/src/crypto/aes/aes_gcm.go:168)
    >>>               /tmp/tinygo4141715439/main.o:(interface:{NonceSize:func:{}{basic:int},Open:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8,named:error},Overhead:func:{}{basic:int},Seal:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8}}.Open$invoke)
    >>> referenced 1 more times
    
    ld.lld-13: error: undefined symbol: crypto/aes.gcmAesFinish
    >>> referenced by aes_gcm.go:116 (/usr/local/go1.17/src/crypto/aes/aes_gcm.go:116)
    >>>               /tmp/tinygo4141715439/main.o:(interface:{NonceSize:func:{}{basic:int},Open:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8,named:error},Overhead:func:{}{basic:int},Seal:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8}}.Seal$invoke)
    >>> referenced by aes_gcm.go:131 (/usr/local/go1.17/src/crypto/aes/aes_gcm.go:131)
    >>>               /tmp/tinygo4141715439/main.o:(interface:{NonceSize:func:{}{basic:int},Open:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8,named:error},Overhead:func:{}{basic:int},Seal:func:{slice:basic:uint8,slice:basic:uint8,slice:basic:uint8,slice:basic:uint8}{slice:basic:uint8}}.Seal$invoke)
    >>> referenced by aes_gcm.go:169 (/usr/local/go1.17/src/crypto/aes/aes_gcm.go:169)
    [...more lines following...]










## crypto/ecdsa



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.17/src/crypto/cipher/xor_amd64.go:18)
    >>>               /tmp/tinygo4260333813/main.o:((crypto/cipher.StreamReader).Read)
    
    ld.lld-13: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:69 (/usr/local/go1.17/src/crypto/aes/cipher_asm.go:69)
    >>>               /tmp/tinygo4260333813/main.o:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld-13: error: undefined symbol: crypto/elliptic.p256Mul
    >>> referenced by p256_asm.go:279 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:279)
    >>>               /tmp/tinygo4260333813/main.o:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:280 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:280)
    >>>               /tmp/tinygo4260333813/main.o:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:351 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:351)
    >>>               /tmp/tinygo4260333813/main.o:((*crypto/elliptic.p256Point).p256PointToAffine)
    [...more lines following...]






## crypto/ed25519



The compiler gave the following error when this package was imported:


    ld.lld-13: error: undefined symbol: crypto/ed25519/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:305 (/usr/local/go1.17/src/crypto/ed25519/internal/edwards25519/field/fe.go:305)
    >>>               /tmp/tinygo670425953/main.o:((*crypto/ed25519/internal/edwards25519/field.Element).Square)
    
    ld.lld-13: error: undefined symbol: crypto/ed25519/internal/edwards25519/field.feMul
    >>> referenced by fe.go:299 (/usr/local/go1.17/src/crypto/ed25519/internal/edwards25519/field/fe.go:299)
    >>>               /tmp/tinygo670425953/main.o:((*crypto/ed25519/internal/edwards25519/field.Element).Multiply)
    error: failed to link /tmp/tinygo670425953/main: exit status 1






## crypto/elliptic



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/elliptic.p256Mul
    >>> referenced by p256_asm.go:242 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:242)
    >>>               /tmp/tinygo2465557329/main.o:(interface:{Add:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,pointer:named:math/big.Int,pointer:named:math/big.Int}{pointer:named:math/big.Int,pointer:named:math/big.Int},CombinedMult:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,slice:basic:uint8,slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int},Double:func:{pointer:named:math/big.Int,pointer:named:math/big.Int}{pointer:named:math/big.Int,pointer:named:math/big.Int},IsOnCurve:func:{pointer:named:math/big.Int,pointer:named:math/big.Int}{basic:bool},Params:func:{}{pointer:named:crypto/elliptic.CurveParams},ScalarBaseMult:func:{slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int},ScalarMult:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int}}.CombinedMult$invoke)
    >>> referenced by p256_asm.go:243 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:243)
    >>>               /tmp/tinygo2465557329/main.o:(interface:{Add:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,pointer:named:math/big.Int,pointer:named:math/big.Int}{pointer:named:math/big.Int,pointer:named:math/big.Int},CombinedMult:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,slice:basic:uint8,slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int},Double:func:{pointer:named:math/big.Int,pointer:named:math/big.Int}{pointer:named:math/big.Int,pointer:named:math/big.Int},IsOnCurve:func:{pointer:named:math/big.Int,pointer:named:math/big.Int}{basic:bool},Params:func:{}{pointer:named:crypto/elliptic.CurveParams},ScalarBaseMult:func:{slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int},ScalarMult:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int}}.CombinedMult$invoke)
    >>> referenced by p256_asm.go:279 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:279)
    >>>               /tmp/tinygo2465557329/main.o:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced 20 more times
    
    ld.lld-13: error: undefined symbol: crypto/elliptic.p256PointAddAsm
    >>> referenced by p256_asm.go:254 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:254)
    >>>               /tmp/tinygo2465557329/main.o:(interface:{Add:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,pointer:named:math/big.Int,pointer:named:math/big.Int}{pointer:named:math/big.Int,pointer:named:math/big.Int},CombinedMult:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,slice:basic:uint8,slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int},Double:func:{pointer:named:math/big.Int,pointer:named:math/big.Int}{pointer:named:math/big.Int,pointer:named:math/big.Int},IsOnCurve:func:{pointer:named:math/big.Int,pointer:named:math/big.Int}{basic:bool},Params:func:{}{pointer:named:crypto/elliptic.CurveParams},ScalarBaseMult:func:{slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int},ScalarMult:func:{pointer:named:math/big.Int,pointer:named:math/big.Int,slice:basic:uint8}{pointer:named:math/big.Int,pointer:named:math/big.Int}}.CombinedMult$invoke)
    >>> referenced by p256_asm.go:464 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:464)
    >>>               /tmp/tinygo2465557329/main.o:((*crypto/elliptic.p256Point).p256ScalarMult)
    >>> referenced by p256_asm.go:465 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:465)
    [...more lines following...]






## crypto/hmac



The compiler gave the following error when running the tests for this package:


    panic: crypto/hmac: hash generation function does not produce unique values
    FAIL	crypto/hmac	0.038s
    FAIL












## crypto/rsa



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/cipher.xorBytesSSE2
    >>> referenced by xor_amd64.go:18 (/usr/local/go1.17/src/crypto/cipher/xor_amd64.go:18)
    >>>               /tmp/tinygo2980328936/main.o:((crypto/cipher.StreamReader).Read)
    
    ld.lld-13: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:69 (/usr/local/go1.17/src/crypto/aes/cipher_asm.go:69)
    >>>               /tmp/tinygo2980328936/main.o:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld-13: error: undefined symbol: crypto/elliptic.p256Mul
    >>> referenced by p256_asm.go:279 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:279)
    >>>               /tmp/tinygo2980328936/main.o:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:280 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:280)
    >>>               /tmp/tinygo2980328936/main.o:((crypto/elliptic.p256Curve).ScalarMult)
    >>> referenced by p256_asm.go:351 (/usr/local/go1.17/src/crypto/elliptic/p256_asm.go:351)
    >>>               /tmp/tinygo2980328936/main.o:((*crypto/elliptic.p256Point).p256PointToAffine)
    [...more lines following...]












## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	0.016s
    FAIL






## crypto/tls


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#cryptoed25519)
  * [crypto/x509](#cryptox509)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#cryptoed25519)







## database/sql



The compiler gave the following error when running the tests for this package:


    # database/sql
    /usr/local/go1.17/src/database/sql/sql_test.go:4446:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)
    FAIL






## database/sql/driver



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Bytes()
    FAIL	database/sql/driver	0.019s
    FAIL








## debug/elf



The compiler gave the following error when running the tests for this package:


    # debug/elf
    /usr/local/go1.17/src/debug/elf/file_test.go:905:10: ResolveIPAddr not declared by package net
    FAIL






## debug/gosym



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: syscall.RawSyscall
    >>> referenced by zsyscall_linux_amd64.go:1661 (/usr/local/go1.17/src/syscall/zsyscall_linux_amd64.go:1661)
    >>>               /tmp/tinygo435389903/main.o:(os.Pipe)
    error: failed to link /tmp/tinygo435389903/main: exit status 1
    FAIL








## debug/pe



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: syscall.RawSyscall
    >>> referenced by zsyscall_linux_amd64.go:1661 (/usr/local/go1.17/src/syscall/zsyscall_linux_amd64.go:1661)
    >>>               /tmp/tinygo3484213178/main.o:(os.Pipe)
    error: failed to link /tmp/tinygo3484213178/main: exit status 1
    FAIL














## encoding/asn1



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Addr()
    FAIL	encoding/asn1	0.005s
    FAIL








## encoding/base64



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo3737270656/main.o:(encoding/base64.TestDecoderIssue3577)
    error: failed to link /tmp/tinygo3737270656/main: exit status 1
    FAIL






## encoding/binary



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Slice()
    FAIL	encoding/binary	0.010s
    FAIL








## encoding/gob



The compiler gave the following error when this package was imported:


    # encoding/gob
    /usr/local/go1.17/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect








## encoding/json



The compiler gave the following error when running the tests for this package:


    # encoding/json
    /usr/local/go1.17/src/encoding/json/bench_test.go:344:22: StructOf not declared by package reflect
    /usr/local/go1.17/src/encoding/json/stream_test.go:298:15: Pipe not declared by package net
    FAIL






## encoding/pem



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: crypto/ed25519/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:305 (/usr/local/go1.17/src/crypto/ed25519/internal/edwards25519/field/fe.go:305)
    >>>               /tmp/tinygo3851192582/main.o:((*crypto/ed25519/internal/edwards25519/field.Element).Square)
    
    ld.lld-13: error: undefined symbol: crypto/ed25519/internal/edwards25519/field.feMul
    >>> referenced by fe.go:299 (/usr/local/go1.17/src/crypto/ed25519/internal/edwards25519/field/fe.go:299)
    >>>               /tmp/tinygo3851192582/main.o:((*crypto/ed25519/internal/edwards25519/field.Element).Multiply)
    error: failed to link /tmp/tinygo3851192582/main: exit status 1
    FAIL






## encoding/xml



The compiler gave the following error when this package was imported:


    # encoding/xml
    /usr/local/go1.17/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
    /usr/local/go1.17/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)






## errors



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: reflect.Zero()
    FAIL	errors	0.012s
    FAIL






## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## flag



The compiler gave the following error when running the tests for this package:


    # flag_test
    /usr/local/go1.17/src/flag/flag_test.go:647:27: cmd.ProcessState.ExitCode undefined (type *os.ProcessState has no field or method ExitCode)
    FAIL






## fmt



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Slice()
    FAIL	fmt	0.002s
    FAIL






## go/ast



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	go/ast	0.003s
    FAIL






## go/build



The compiler gave the following error when running the tests for this package:


    # go/build
    /usr/local/go1.17/src/go/build/build_test.go:556:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:577:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:599:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:629:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:641:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:642:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:695:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:696:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.17/src/go/build/build_test.go:697:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    FAIL






## go/build/constraint



The compiler gave the following error when running the tests for this package:


    panic: invalid syntax at ²
    FAIL	go/build/constraint	0.010s
    FAIL






## go/constant



The compiler gave the following error when running the tests for this package:


    panic: division by zero
    FAIL	go/constant	0.003s
    FAIL






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	0.001s
    FAIL






## go/format



The compiler gave the following error when running the tests for this package:


    panic: go/parser internal error: unbalanced scopes
    FAIL	go/format	0.025s
    FAIL






## go/importer



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: syscall.RawSyscall
    >>> referenced by zsyscall_linux_amd64.go:1661 (/usr/local/go1.17/src/syscall/zsyscall_linux_amd64.go:1661)
    >>>               /tmp/tinygo340748001/main.o:(os.Pipe)
    error: failed to link /tmp/tinygo340748001/main: exit status 1
    FAIL






## go/parser



The compiler gave the following error when running the tests for this package:


    panic: go/parser internal error: unbalanced scopes
    FAIL	go/parser	0.020s
    FAIL






## go/printer



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo2267952612/main.o:(go/printer.TestFiles$1)
    error: failed to link /tmp/tinygo2267952612/main: exit status 1
    FAIL








## go/token



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.17/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## go/types



The compiler gave the following error when running the tests for this package:


    # go/types_test
    /usr/local/go1.17/src/go/types/self_test.go:102:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    FAIL










## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	0.017s
    FAIL










## hash/maphash



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: runtime.fastrand
    >>> referenced by maphash.go:210 (/usr/local/go1.17/src/hash/maphash/maphash.go:210)
    >>>               /tmp/tinygo2719359244/main.o:(hash/maphash.MakeSeed)
    >>> referenced by maphash.go:211 (/usr/local/go1.17/src/hash/maphash/maphash.go:211)
    >>>               /tmp/tinygo2719359244/main.o:(hash/maphash.MakeSeed)
    
    ld.lld-13: error: undefined symbol: runtime.memhash
    >>> referenced by maphash.go:232 (/usr/local/go1.17/src/hash/maphash/maphash.go:232)
    >>>               /tmp/tinygo2719359244/main.o:(hash/maphash.rthash)
    error: failed to link /tmp/tinygo2719359244/main: exit status 1
    FAIL








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	0.007s
    FAIL






## image



The compiler gave the following error when running the tests for this package:


    panic: image: NewRGBA Rectangle has huge or negative dimensions
    FAIL	image	0.008s
    FAIL






## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	0.015s
    FAIL








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	0.113s
    FAIL






## image/gif



The compiler gave the following error when running the tests for this package:


    # image/gif
    /usr/local/go1.17/src/image/gif/reader_test.go:415:14: SetGCPercent not declared by package debug
    /usr/local/go1.17/src/image/gif/reader_test.go:415:33: SetGCPercent not declared by package debug
    /usr/local/go1.17/src/image/gif/reader_test.go:421:26: s1.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    /usr/local/go1.17/src/image/gif/reader_test.go:421:41: s0.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    FAIL






## image/jpeg



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo3875756092/main.o:(image/jpeg.TestLargeImageWithShortData)
    error: failed to link /tmp/tinygo3875756092/main: exit status 1
    FAIL






## image/png



The compiler gave the following error when running the tests for this package:


    panic: runtime error: out of memory
    FAIL	image/png	4.196s
    FAIL








## io



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo2342414757/main.o:(io_test.TestMultiReaderFreesExhaustedReaders)
    error: failed to link /tmp/tinygo2342414757/main: exit status 1
    FAIL








## io/ioutil



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestTempDir_BadDir (0.00s)
        TempDir error = &T{Op:"stat", Path:"/tmp/TestTempDir_BadDir83397427/not-exist", Err:(T)(0x2479d0)}; want PathError for path "/tmp/TestTempDir_BadDir83397427/not-exist" satisifying os.IsNotExist
    FAIL
    FAIL	io/ioutil	0.012s
    FAIL






## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.01s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "XXX2022/03/02 19:31:49.648250 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "XXX2022/03/02 19:31:49.648682 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "XXX2022/03/02 19:31:49.649183 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(60|62): hello 23 world$" is "XXX2022/03/02 19:31:49.650102 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(60|62): XXXhello 23 world$" is "2022/03/02 19:31:49.650180 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(60|62): XXXhello 23 world$" is "2022/03/02 19:31:49.650655 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(60|62): XXXhello 23 world$" is "2022/03/02 19:31:49.651629 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(60|62): XXXhello 23 world$" is "2022/03/02 19:31:49.651703 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/syslog



The compiler gave the following error when running the tests for this package:


    # log/syslog
    /usr/local/go1.17/src/log/syslog/syslog_test.go:23:25: PacketConn not declared by package net
    /usr/local/go1.17/src/log/syslog/syslog_test.go:36:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /usr/local/go1.17/src/log/syslog/syslog_test.go:104:15: ListenPacket not declared by package net
    FAIL








## math/big



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.17/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## math/bits



The compiler gave the following error when running the tests for this package:


    panic: overflow
    FAIL	math/bits	0.021s
    FAIL








## math/rand



The compiler gave the following error when running the tests for this package:


    # math/rand_test
    /usr/local/go1.17/src/math/rand/regress_test.go:38:18: rv.Type().Method undefined (type reflect.Type has no field or method Method)
    /usr/local/go1.17/src/math/rand/regress_test.go:39:12: rv.Method undefined (type reflect.Value has no field or method Method)
    FAIL






## mime



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	mime	0.003s
    FAIL






## mime/multipart



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapRange()
    FAIL	mime/multipart	0.004s
    FAIL






## mime/quotedprintable



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: syscall.RawSyscall
    >>> referenced by zsyscall_linux_amd64.go:1661 (/usr/local/go1.17/src/syscall/zsyscall_linux_amd64.go:1661)
    >>>               /tmp/tinygo3616146811/main.o:(os.Pipe)
    
    ld.lld-13: error: undefined symbol: time.startTimer
    >>> referenced by sleep.go:96 (/usr/local/go1.17/src/time/sleep.go:96)
    >>>               /tmp/tinygo3616146811/main.o:(mime/quotedprintable.TestExhaustive$1)
    error: failed to link /tmp/tinygo3616146811/main: exit status 1
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
  * [crypto/x509](#cryptox509)
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
    FAIL	net/textproto	0.004s
    FAIL






## net/url



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.17/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL








## os/exec



The compiler gave the following error when running the tests for this package:


    # os/user
    ../../../../../usr/local/go1.17/src/os/user/cgo_lookup_unix.go:21:6: not implemented: build constraints in #cgo line
    FAIL






## os/signal



The compiler gave the following error when running the tests for this package:


    can't load test package: ../../../.cache/tinygo/goroot-7a0feacad3b2376d3ea46a57eaccf73bf6f44920673e97de4ffa85431952c656/src/os/signal/signal_test.go:19:2: package runtime/trace is not in GOROOT (/home/ayke/.cache/tinygo/goroot-7a0feacad3b2376d3ea46a57eaccf73bf6f44920673e97de4ffa85431952c656/src/runtime/trace)






## os/user



The compiler gave the following error when this package was imported:


    # os/user
    ../../../../../usr/local/go1.17/src/os/user/cgo_lookup_unix.go:21:6: not implemented: build constraints in #cgo line








## path/filepath



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	path/filepath	0.031s
    FAIL






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    ../../../../../usr/local/go1.17/src/plugin/plugin_dlopen.go:11:6: not implemented: build constraints in #cgo line








## regexp



The compiler gave the following error when running the tests for this package:


    FAIL	regexp	0.124s
    FAIL






## regexp/syntax



The compiler gave the following error when running the tests for this package:


    # regexp/syntax
    /usr/local/go1.17/src/unicode/letter.go:122:6: interp: running for more than 180 seconds, timing out (executed calls: 3139442)
      %range_ = alloca %unicode.Range32, align 8, !dbg !2558
    
    traceback:
    /usr/local/go1.17/src/unicode/letter.go:122:6:
      %range_ = alloca %unicode.Range32, align 8, !dbg !2558
    /usr/local/go1.17/src/unicode/letter.go:176:14:
      %5 = call i1 @unicode.is32(%unicode.Range32* %indexaddr.ptr16, i64 %len15, i64 %4, i32 %r, i8* undef), !dbg !2574
    /usr/local/go1.17/src/unicode/letter.go:187:25:
      %4 = call i1 @unicode.isExcludingLatin(%unicode.RangeTable* %3, i32 %r, i8* undef), !dbg !2561
    /usr/local/go1.17/src/regexp/syntax/parse_test.go:389:7:
      %17 = call i1 %7(i32 %5, i8* %f.context), !dbg !2589
    /usr/local/go1.17/src/regexp/syntax:
      %3 = call %runtime._string @"regexp/syntax.mkCharClass"(i8* undef, void ()* bitcast (i1 (i32, i8*)* @unicode.IsUpper to void ()*), i8* undef), !dbg !2566
    [...more lines following...]






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (3.58s)
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
    --- FAIL: TestCountSortOps (0.79s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]






## strconv



The compiler gave the following error when running the tests for this package:


    panic: invalid bitSize
    FAIL	strconv	6.017s
    FAIL






## strings



The compiler gave the following error when running the tests for this package:


    # strings_test
    /home/ayke/src/github.com/tinygo-org/tinygo/src/sync/mutex.go:17:7: interp: running for more than 180 seconds, timing out (executed calls: 391871)
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !2981
    
    traceback:
    /home/ayke/src/github.com/tinygo-org/tinygo/src/sync/mutex.go:17:7:
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !2981
    /usr/local/go1.17/src/math/rand/rand.go:388:11:
      call void @"(*sync.Mutex).Lock"(%sync.Mutex* %1, i8* undef), !dbg !2993
    /usr/local/go1.17/src/math/rand/rand.go:84:50:
      %3 = call i64 @"interface:{Int63:func:{}{basic:int64},Seed:func:{basic:int64}{}}.Int63$invoke"(i8* %invoke.func.value, i64 %invoke.func.typecode, i8* undef), !dbg !2966
    /usr/local/go1.17/src/math/rand/rand.go:98:52:
      %0 = call i64 @"(*math/rand.Rand).Int63"(%"math/rand.Rand"* %r, i8* undef), !dbg !2965
    /usr/local/go1.17/src/math/rand/rand.go:133:14:
      %11 = call i32 @"(*math/rand.Rand).Int31"(%"math/rand.Rand"* %r, i8* undef), !dbg !2995
    [...more lines following...]








## sync/atomic



The compiler gave the following error when running the tests for this package:


    # sync/atomic_test
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1217:13: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1218:60: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1259:13: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1260:60: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1304:13: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1305:60: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1350:13: NumCPU not declared by package runtime
    /usr/local/go1.17/src/sync/atomic/atomic_test.go:1351:60: NumCPU not declared by package runtime
    FAIL






## syscall



The compiler gave the following error when running the tests for this package:


    # os/user
    ../../../../../usr/local/go1.17/src/os/user/cgo_lookup_unix.go:21:6: not implemented: build constraints in #cgo line
    FAIL












## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	0.001s
    FAIL








## text/tabwriter



The compiler gave the following error when running the tests for this package:


    panic: cannot write
    FAIL	text/tabwriter	0.002s
    FAIL






## text/template



The compiler gave the following error when running the tests for this package:


    ld.lld-13: error: undefined symbol: syscall.RawSyscall
    >>> referenced by zsyscall_linux_amd64.go:1661 (/usr/local/go1.17/src/syscall/zsyscall_linux_amd64.go:1661)
    >>>               /tmp/tinygo2777764320/main.o:(os.Pipe)
    error: failed to link /tmp/tinygo2777764320/main: exit status 1
    FAIL






## text/template/parse



The compiler gave the following error when running the tests for this package:


    panic: template: root:1: unexpected "}}" in define clause
    FAIL	text/template/parse	0.010s
    FAIL






## time



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.17/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## time/tzdata



The compiler gave the following error when this package was imported:


    ld.lld-13: error: undefined symbol: time.registerLoadFromEmbeddedTZData
    >>> referenced by scheduler_any.go:24 (/home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/scheduler_any.go:24)
    >>>               /tmp/tinygo163579754/main.o:(runtime.run$1$gowrapper)
    error: failed to link /tmp/tinygo163579754/main: exit status 1













