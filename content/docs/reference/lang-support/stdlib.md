
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
strconv |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strconv)  | 
strings |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strings)  | 
sync |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync/atomic |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syncatomic)  | 
syscall |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo189667552/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo189667552/main: exit status 1
    FAIL






## bytes



The compiler gave the following error when running the tests for this package:


    # bytes_test
    /Users/dkegel/src/tinygo/src/sync/mutex.go:17:7: interp: running for more than 180 seconds, timing out (executed calls: 532318)
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !2971
    
    traceback:
    /Users/dkegel/src/tinygo/src/sync/mutex.go:17:7:
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !2971
    /usr/local/go1.18/src/math/rand/rand.go:388:11:
      call void @"(*sync.Mutex).Lock"(%sync.Mutex* %1, i8* undef), !dbg !2983
    /usr/local/go1.18/src/math/rand/rand.go:84:50:
      %3 = call i64 @"interface:{Int63:func:{}{basic:int64},Seed:func:{basic:int64}{}}.Int63$invoke"(i8* %invoke.func.value, i64 %invoke.func.typecode, i8* undef), !dbg !2956
    /usr/local/go1.18/src/math/rand/rand.go:98:52:
      %0 = call i64 @"(*math/rand.Rand).Int63"(%"math/rand.Rand"* %r, i8* undef), !dbg !2955
    /usr/local/go1.18/src/math/rand/rand.go:133:14:
      %11 = call i32 @"(*math/rand.Rand).Int31"(%"math/rand.Rand"* %r, i8* undef), !dbg !2985
    [...more lines following...]










## compress/gzip



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo345307923/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo345307923/main: exit status 1
    FAIL
















## context



The compiler gave the following error when running the tests for this package:


    # context_test
    /usr/local/go1.18/src/context/x_test.go:12:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestBackground: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:13:68: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTODO: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:14:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestWithCancel: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:15:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestParentFinishesChild: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:16:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestChildFinishesFirst: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:17:72: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestDeadline: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:18:71: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:19:79: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestCanceledTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:20:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestValues: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:21:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestAllocs: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:22:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestSimultaneousCancels: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:23:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestInterlockedCancels: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:24:76: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersCancel: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.18/src/context/x_test.go:25:77: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    [...more lines following...]






## crypto



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _crypto/cipher.xorBytesSSE2
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo209964351/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.encryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo209964351/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo209964351/main: exit status 1
    FAIL






## crypto/aes



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _crypto/aes.encryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3086193382/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.decryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3086193382/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3086193382/main: exit status 1
    FAIL






## crypto/cipher



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _crypto/cipher.xorBytesSSE2
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1517805482/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.gcmAesInit
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1517805482/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.encryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1517805482/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.decryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1517805482/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.gcmAesFinish
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1517805482/main.o
    
    [...more lines following...]










## crypto/ecdsa



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _crypto/cipher.xorBytesSSE2
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1473379088/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.encryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1473379088/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256LittleToBig
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1473379088/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256LittleToBig
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1473379088/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256FromMont
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1473379088/main.o
    
    [...more lines following...]






## crypto/ed25519



The compiler gave the following error when this package was imported:


    tinygo:ld.lld: error: undefined symbol: _crypto/ed25519/internal/edwards25519/field.feSquare
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1818820213/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/ed25519/internal/edwards25519/field.feMul
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1818820213/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1818820213/main: exit status 1






## crypto/elliptic



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2238568371/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256LittleToBig
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2238568371/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256LittleToBig
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2238568371/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256FromMont
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2238568371/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256FromMont
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2238568371/main.o
    
    [...more lines following...]






## crypto/hmac



The compiler gave the following error when running the tests for this package:


    panic: crypto/hmac: hash generation function does not produce unique values
    FAIL	crypto/hmac	1.451s
    FAIL












## crypto/rsa



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _crypto/cipher.xorBytesSSE2
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2829435925/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/aes.encryptBlockAsm
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2829435925/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256LittleToBig
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2829435925/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256LittleToBig
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2829435925/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/elliptic.p256FromMont
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2829435925/main.o
    
    [...more lines following...]












## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	1.467s
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
    /usr/local/go1.18/src/database/sql/sql_test.go:4556:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)
    FAIL






## database/sql/driver



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Bytes()
    FAIL	database/sql/driver	7.808s
    FAIL






## debug/buildinfo



The compiler gave the following error when this package was imported:


    # debug/buildinfo
    /usr/local/go1.18/src/debug/buildinfo/buildinfo.go:78:19: ParseBuildInfo not declared by package debug








## debug/elf



The compiler gave the following error when running the tests for this package:


    # debug/elf
    /usr/local/go1.18/src/debug/elf/file_test.go:905:10: ResolveIPAddr not declared by package net
    FAIL






## debug/gosym



The compiler gave the following error when running the tests for this package:


    panic: runtime error: nil pointer dereference
    FAIL	debug/gosym	0.714s
    FAIL








## debug/pe



The compiler gave the following error when running the tests for this package:


    panic: runtime error: index out of range
    FAIL	debug/pe	0.561s
    FAIL








## embed



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1909912716/main.o
    
    tinygo:ld.lld: error: undefined symbol: _vendor/golang.org/x/sys/cpu.xgetbv
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1909912716/main.o
    
    tinygo:ld.lld: error: undefined symbol: _vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1909912716/main.o
    
    tinygo:ld.lld: error: undefined symbol: _vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1909912716/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/ed25519/internal/edwards25519/field.feSquare
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1909912716/main.o
    
    [...more lines following...]










## encoding/asn1



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Addr()
    FAIL	encoding/asn1	0.839s
    FAIL








## encoding/base64



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3292315030/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3292315030/main: exit status 1
    FAIL






## encoding/binary



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Slice()
    FAIL	encoding/binary	3.803s
    FAIL








## encoding/gob



The compiler gave the following error when this package was imported:


    # encoding/gob
    /usr/local/go1.18/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect








## encoding/json



The compiler gave the following error when running the tests for this package:


    # encoding/json
    /usr/local/go1.18/src/encoding/json/bench_test.go:344:22: StructOf not declared by package reflect
    /usr/local/go1.18/src/encoding/json/stream_test.go:298:15: Pipe not declared by package net
    FAIL






## encoding/pem



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _crypto/ed25519/internal/edwards25519/field.feSquare
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo20248966/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/ed25519/internal/edwards25519/field.feMul
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo20248966/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/x509/internal/macos.syscall
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo20248966/main.o
    
    tinygo:ld.lld: error: undefined symbol: _internal/abi.FuncPCABI0
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo20248966/main.o
    
    tinygo:ld.lld: error: undefined symbol: _crypto/x509/internal/macos.x509_CFStringCreateWithBytes_trampoline
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo20248966/main.o
    failed to run tool: ld.lld
    [...more lines following...]






## encoding/xml



The compiler gave the following error when this package was imported:


    # encoding/xml
    /usr/local/go1.18/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
    /usr/local/go1.18/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)






## errors



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: reflect.Zero()
    FAIL	errors	0.762s
    FAIL






## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## flag



The compiler gave the following error when running the tests for this package:


    flag provided but not defined: -x
    panic: unimplemented: reflect.Zero()
    FAIL	flag	1.062s
    FAIL






## fmt



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Slice()
    FAIL	fmt	0.857s
    FAIL






## go/ast



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	go/ast	0.825s
    FAIL






## go/build



The compiler gave the following error when running the tests for this package:


    # go/build
    /usr/local/go1.18/src/go/build/build_test.go:556:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:577:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:599:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:629:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:641:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:642:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:695:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:696:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    /usr/local/go1.18/src/go/build/build_test.go:697:4: t.Setenv undefined (type *testing.T has no field or method Setenv)
    FAIL






## go/build/constraint



The compiler gave the following error when running the tests for this package:


    panic: invalid syntax at ²
    FAIL	go/build/constraint	0.645s
    FAIL






## go/constant



The compiler gave the following error when running the tests for this package:


    panic: division by zero
    FAIL	go/constant	0.769s
    FAIL






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unknown type: E
    
    goroutine 30 [running]:
    github.com/tinygo-org/tinygo/compiler.(*compilerContext).makeLLVMType(0xc0017fc5b0?, {0x839de00?, 0xc0050a7b30?})
    	/Users/dkegel/src/tinygo/compiler/compiler.go:415 +0x82c
    github.com/tinygo-org/tinygo/compiler.(*compilerContext).getLLVMType(0xc0017fc5b0, {0x839de00, 0xc0050a7b30})
    	/Users/dkegel/src/tinygo/compiler/compiler.go:326 +0x6e
    github.com/tinygo-org/tinygo/compiler.(*compilerContext).makeLLVMType(0xc0017fc5b0, {0x839dd88?, 0xc001a09650?})
    	/Users/dkegel/src/tinygo/compiler/compiler.go:395 +0x185
    github.com/tinygo-org/tinygo/compiler.(*compilerContext).getLLVMType(0xc0017fc5b0, {0x839dd88, 0xc001a09650})
    	/Users/dkegel/src/tinygo/compiler/compiler.go:326 +0x6e
    github.com/tinygo-org/tinygo/compiler.(*compilerContext).getFunction(0xc0017fc5b0, 0xc0056df040)
    	/Users/dkegel/src/tinygo/compiler/symbol.go:78 +0x385
    github.com/tinygo-org/tinygo/compiler.(*builder).createFunctionCall(0xc0033cf320, 0xc000833d40)
    	/Users/dkegel/src/tinygo/compiler/compiler.go:1463 +0xab3
    [...more lines following...]






## go/format



The compiler gave the following error when running the tests for this package:


    panic: go/parser internal error: unbalanced scopes
    FAIL	go/format	0.565s
    FAIL






## go/importer



The compiler gave the following error when running the tests for this package:


    panic: runtime error: nil pointer dereference
    FAIL	go/importer	0.799s
    FAIL






## go/parser



The compiler gave the following error when running the tests for this package:


    panic: go/parser internal error: unbalanced scopes
    FAIL	go/parser	0.680s
    FAIL






## go/printer



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1882669185/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1882669185/main: exit status 1
    FAIL








## go/token



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## go/types



The compiler gave the following error when running the tests for this package:


    # go/types_test
    /usr/local/go1.18/src/go/types/self_test.go:98:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    FAIL










## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	1.495s
    FAIL










## hash/maphash



The compiler gave the following error when running the tests for this package:


    panic: runtime error: out of memory
    FAIL	hash/maphash	168.656s
    FAIL








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	0.464s
    FAIL






## image



The compiler gave the following error when running the tests for this package:


    panic: image: NewRGBA Rectangle has huge or negative dimensions
    FAIL	image	0.763s
    FAIL






## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	0.414s
    FAIL








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	0.809s
    FAIL






## image/gif



The compiler gave the following error when running the tests for this package:


    # image/gif
    /usr/local/go1.18/src/image/gif/reader_test.go:415:14: SetGCPercent not declared by package debug
    /usr/local/go1.18/src/image/gif/reader_test.go:415:33: SetGCPercent not declared by package debug
    /usr/local/go1.18/src/image/gif/reader_test.go:421:26: s1.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    /usr/local/go1.18/src/image/gif/reader_test.go:421:41: s0.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    FAIL






## image/jpeg



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2239879435/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo2239879435/main: exit status 1
    FAIL






## image/png



The compiler gave the following error when running the tests for this package:


    panic: runtime error: out of memory
    FAIL	image/png	2.466s
    FAIL








## io



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1333954222/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo1333954222/main: exit status 1
    FAIL








## io/ioutil



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestTempDir_BadDir (0.00s)
        TempDir error = &T{Op:"stat", Path:"/var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/TestTempDir_BadDir99598472/not-exist", Err:(T)(0x10180c878)}; want PathError for path "/var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/TestTempDir_BadDir99598472/not-exist" satisifying os.IsNotExist
    FAIL
    FAIL	io/ioutil	0.959s
    FAIL






## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.00s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/04/28 16:01:17.295555 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/04/28 16:01:17.295970 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/04/28 16:01:17.296884 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2022/04/28 16:01:17.297313 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/04/28 16:01:17.297355 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/04/28 16:01:17.297393 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/04/28 16:01:17.297772 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2022/04/28 16:01:17.297802 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/syslog



The compiler gave the following error when running the tests for this package:


    # log/syslog
    /usr/local/go1.18/src/log/syslog/syslog_test.go:22:25: PacketConn not declared by package net
    /usr/local/go1.18/src/log/syslog/syslog_test.go:35:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /usr/local/go1.18/src/log/syslog/syslog_test.go:110:15: ListenPacket not declared by package net
    FAIL








## math/big



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## math/bits



The compiler gave the following error when running the tests for this package:


    panic: overflow
    FAIL	math/bits	2.504s
    FAIL








## math/rand



The compiler gave the following error when running the tests for this package:


    # math/rand_test
    /usr/local/go1.18/src/math/rand/regress_test.go:38:18: rv.Type().Method undefined (type reflect.Type has no field or method Method)
    /usr/local/go1.18/src/math/rand/regress_test.go:39:12: rv.Method undefined (type reflect.Value has no field or method Method)
    FAIL






## mime



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	mime	1.294s
    FAIL






## mime/multipart



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapRange()
    FAIL	mime/multipart	0.512s
    FAIL






## mime/quotedprintable



The compiler gave the following error when running the tests for this package:


    tinygo:ld.lld: error: undefined symbol: _time.startTimer
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3869451097/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3869451097/main: exit status 1
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







## net/netip



The compiler gave the following error when running the tests for this package:


    # net/netip_test
    /usr/local/go1.18/src/net/netip/netip_test.go:1815:25: UDPAddr not declared by package net
    /usr/local/go1.18/src/net/netip/fuzz_test.go:177:32: stdip.IsPrivate undefined (type net.IP has no field or method IsPrivate)
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
    FAIL	net/textproto	0.808s
    FAIL






## net/url



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL








## os/exec



The compiler gave the following error when running the tests for this package:


    # os/user
    ../../../../../usr/local/go1.18/src/os/user/cgo_lookup_unix.go:18:6: not implemented: build constraints in #cgo line
    ../../../../../usr/local/go1.18/src/os/user/cgo_lookup_unix.go:21:10: fatal: 'pwd.h' file not found
    FAIL






## os/signal



The compiler gave the following error when running the tests for this package:


    can't load test package: ../../../Library/Caches/tinygo/goroot-6cdeec8e8f6bb3246a475995a5c5424095001b1df68f35e053b64525f5faf840/src/os/signal/signal_test.go:18:2: package runtime/trace is not in GOROOT (/Users/dkegel/Library/Caches/tinygo/goroot-6cdeec8e8f6bb3246a475995a5c5424095001b1df68f35e053b64525f5faf840/src/runtime/trace)






## os/user



The compiler gave the following error when this package was imported:


    # os/user
    ../../../../../usr/local/go1.18/src/os/user/cgo_lookup_unix.go:18:6: not implemented: build constraints in #cgo line
    ../../../../../usr/local/go1.18/src/os/user/cgo_lookup_unix.go:21:10: fatal: 'pwd.h' file not found








## path/filepath



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).MapKeys()
    FAIL	path/filepath	1.553s
    FAIL






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    ../../../../../usr/local/go1.18/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line
    ../../../../../usr/local/go1.18/src/plugin/plugin_dlopen.go:11:10: fatal: 'dlfcn.h' file not found








## regexp



The compiler gave the following error when running the tests for this package:


    FAIL	regexp	0.584s
    FAIL






## regexp/syntax



The compiler gave the following error when running the tests for this package:


    panic: regexp/syntax: internal error
    FAIL	regexp/syntax	1.367s
    FAIL






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (2.34s)
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
    --- FAIL: TestCountSortOps (0.47s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]






## strconv



The compiler gave the following error when running the tests for this package:


    panic: invalid bitSize
    FAIL	strconv	7.906s
    FAIL






## strings



The compiler gave the following error when running the tests for this package:


    # strings_test
    /Users/dkegel/src/tinygo/src/sync/mutex.go:17:7: interp: running for more than 180 seconds, timing out (executed calls: 546895)
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !2903
    
    traceback:
    /Users/dkegel/src/tinygo/src/sync/mutex.go:17:7:
      %0 = icmp eq %sync.Mutex* %m, null, !dbg !2903
    /usr/local/go1.18/src/math/rand/rand.go:388:11:
      call void @"(*sync.Mutex).Lock"(%sync.Mutex* %1, i8* undef), !dbg !2915
    /usr/local/go1.18/src/math/rand/rand.go:84:50:
      %3 = call i64 @"interface:{Int63:func:{}{basic:int64},Seed:func:{basic:int64}{}}.Int63$invoke"(i8* %invoke.func.value, i64 %invoke.func.typecode, i8* undef), !dbg !2888
    /usr/local/go1.18/src/math/rand/rand.go:98:52:
      %0 = call i64 @"(*math/rand.Rand).Int63"(%"math/rand.Rand"* %r, i8* undef), !dbg !2887
    /usr/local/go1.18/src/math/rand/rand.go:133:14:
      %11 = call i32 @"(*math/rand.Rand).Int31"(%"math/rand.Rand"* %r, i8* undef), !dbg !2917
    [...more lines following...]








## sync/atomic



The compiler gave the following error when running the tests for this package:


    # sync/atomic_test
    /usr/local/go1.18/src/sync/atomic/atomic_test.go:1202:14: SetGCPercent not declared by package debug
    /usr/local/go1.18/src/sync/atomic/atomic_test.go:1202:33: SetGCPercent not declared by package debug
    FAIL














## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	0.774s
    FAIL








## text/tabwriter



The compiler gave the following error when running the tests for this package:


    panic: cannot write
    FAIL	text/tabwriter	1.302s
    FAIL






## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	0.841s
    FAIL






## text/template/parse



The compiler gave the following error when running the tests for this package:


    panic: template: root:1: unexpected "}}" in define clause
    FAIL	text/template/parse	0.510s
    FAIL






## time



The compiler gave the following error when running the tests for this package:


    # encoding/gob
    /usr/local/go1.18/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
    FAIL






## time/tzdata



The compiler gave the following error when this package was imported:


    tinygo:ld.lld: error: undefined symbol: _time.registerLoadFromEmbeddedTZData
    >>> referenced by /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3893541617/main.o
    failed to run tool: ld.lld
    error: failed to link /var/folders/v0/0k9hwftd29b18z9z7s2g2sfc0000gn/T/tinygo3893541617/main: exit status 1













