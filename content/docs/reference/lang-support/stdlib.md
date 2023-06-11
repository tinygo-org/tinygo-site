
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Test results are for linux/amd64.

Package | Importable | Passes tests
--- | --- | --- |
archive/tar |  [<span style="color: red">✗</span> no](#archivetar)  |  <span style="color: gray">✗</span> no  | 
archive/zip |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
bufio |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bufio)  | 
bytes |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
crypto/ecdh |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoecdh)  | 
crypto/ecdsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoecdsa)  | 
crypto/ed25519 |  [<span style="color: red">✗</span> no](#cryptoed25519)  |  <span style="color: gray">✗</span> no  | 
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
crypto/x509 |  [<span style="color: red">✗</span> no](#cryptox509)  |  <span style="color: gray">✗</span> no  | 
crypto/x509/pkix |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesql)  | 
database/sql/driver |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
encoding/asn1 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/base32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/base64 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/binary |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingbinary)  | 
encoding/csv |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/gob |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodinggob)  | 
encoding/hex |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
encoding/json |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingjson)  | 
encoding/pem |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingpem)  | 
encoding/xml |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingxml)  | 
errors |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#errors)  | 
expvar |  [<span style="color: red">✗</span> no](#expvar)  |  <span style="color: gray">✗</span> no  | 
flag |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#flag)  | 
fmt |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#fmt)  | 
go/ast |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/build |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gobuild)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/constant |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goconstant)  | 
go/doc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#godoc)  | 
go/doc/comment |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#godoccomment)  | 
go/format |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/importer |  [<span style="color: red">✗</span> no](#goimporter)  |  <span style="color: gray">✗</span> no  | 
go/parser |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goparser)  | 
go/printer |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
image/jpeg |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagejpeg)  | 
image/png |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagepng)  | 
index/suffixarray |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#io)  | 
io/fs |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io/ioutil |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
log |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#log)  | 
log/syslog |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#logsyslog)  | 
math |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/big |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbig)  | 
math/bits |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbits)  | 
math/cmplx |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathrand)  | 
mime |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mime)  | 
mime/multipart |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
net/textproto |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
regexp/syntax |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
    FAIL	bufio	0.833s












## compress/gzip



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-DC599358A5C8BD4AABB7CE20F999F63B34961798:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-DC599358A5C8BD4AABB7CE20F999F63B34961798:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-DC599358A5C8BD4AABB7CE20F999F63B34961798:(runtime.initAll)
    >>> referenced 1 more times
    
    ld.lld: error: undefined symbol: vendor/golang.org/x/sys/cpu.xgetbv
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-DC599358A5C8BD4AABB7CE20F999F63B34961798:(runtime.initAll)
    
    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    [...more lines following...]
















## context



The compiler gave the following error when running the tests for this package:


    FAIL	context	0.000s
    # context_test
    /usr/local/go1.20/src/context/x_test.go:12:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestBackground: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:13:68: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTODO: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:14:74: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestWithCancel: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:15:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestParentFinishesChild: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:16:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestChildFinishesFirst: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:17:72: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestDeadline: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:18:71: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:19:79: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestCanceledTimeout: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:20:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestValues: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:21:70: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestAllocs: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:22:83: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestSimultaneousCancels: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:23:82: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestInterlockedCancels: *testing.T does not implement context.testingT (missing method Deadline)
    /usr/local/go1.20/src/context/x_test.go:24:76: cannot use t (variable of type *testing.T) as context.testingT value in argument to XTestLayersCancel: *testing.T does not implement context.testingT (missing method Deadline)
    [...more lines following...]






## crypto



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-2488623C069C9DC12EF6E37F26080642CE5B006E:(crypto/subtle.XORBytes)
    
    ld.lld: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.20/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-2488623C069C9DC12EF6E37F26080642CE5B006E:((*crypto/aes.aesCipherAsm).Encrypt)
    FAIL	crypto	0.000s
    error: failed to link /tmp/tinygo1116462727/main: exit status 1






## crypto/aes



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/aes.decryptBlockAsm
    >>> referenced by cipher_asm.go:95 (/usr/local/go1.20/src/crypto/aes/cipher_asm.go:95)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-2D00ED78EE01FF7528848C59DF182739099E48B0:((*crypto/aes.aesCipherAsm).Decrypt)
    
    ld.lld: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.20/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-2D00ED78EE01FF7528848C59DF182739099E48B0:((*crypto/aes.aesCipherAsm).Encrypt)
    FAIL	crypto/aes	0.000s
    error: failed to link /tmp/tinygo3158500475/main: exit status 1






## crypto/cipher



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-85892024269DA2D4324A4B2FCF26D2931ACFD6C2:(crypto/subtle.XORBytes)
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-85892024269DA2D4324A4B2FCF26D2931ACFD6C2:((*crypto/cipher.gcm).auth)
    
    ld.lld: error: undefined symbol: crypto/aes.gcmAesInit
    >>> referenced by aes_gcm.go:49 (/usr/local/go1.20/src/crypto/aes/aes_gcm.go:49)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-85892024269DA2D4324A4B2FCF26D2931ACFD6C2:(crypto/cipher.newGCMWithNonceAndTagSize)
    
    ld.lld: error: undefined symbol: crypto/aes.decryptBlockAsm
    >>> referenced by cipher_asm.go:95 (/usr/local/go1.20/src/crypto/aes/cipher_asm.go:95)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-85892024269DA2D4324A4B2FCF26D2931ACFD6C2:((*crypto/aes.aesCipherAsm).Decrypt)
    
    ld.lld: error: undefined symbol: crypto/aes.encryptBlockAsm
    [...more lines following...]










## crypto/ecdh



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:309 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:309)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8B5337492012101A96AFD7F823C76B7EE0629FAF:((*crypto/internal/edwards25519/field.Element).Square)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feMul
    >>> referenced by fe.go:303 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:303)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8B5337492012101A96AFD7F823C76B7EE0629FAF:((*crypto/internal/edwards25519/field.Element).Multiply)
    
    ld.lld: error: undefined symbol: crypto/internal/nistec.p256Sqr
    >>> referenced by p256_asm.go:563 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:563)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8B5337492012101A96AFD7F823C76B7EE0629FAF:(crypto/internal/nistec.p256Inverse)
    >>> referenced by p256_asm.go:565 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:565)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8B5337492012101A96AFD7F823C76B7EE0629FAF:(crypto/internal/nistec.p256Inverse)
    >>> referenced by p256_asm.go:567 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:567)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-8B5337492012101A96AFD7F823C76B7EE0629FAF:(crypto/internal/nistec.p256Inverse)
    [...more lines following...]






## crypto/ecdsa



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D2CF4BBB69E0C8F4FB42E53CB99819414C667E08:((crypto/cipher.StreamReader).Read)
    
    ld.lld: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.20/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D2CF4BBB69E0C8F4FB42E53CB99819414C667E08:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:309 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:309)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D2CF4BBB69E0C8F4FB42E53CB99819414C667E08:((*crypto/internal/edwards25519/field.Element).Square)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feMul
    >>> referenced by fe.go:303 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:303)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-D2CF4BBB69E0C8F4FB42E53CB99819414C667E08:((*crypto/internal/edwards25519/field.Element).Multiply)
    [...more lines following...]






## crypto/ed25519



The compiler gave the following error when this package was imported:


    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:309 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:309)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-0789AEB7EFB32DDEE9CB09DBEDFBFD637F9C1A48:((*crypto/internal/edwards25519/field.Element).Square)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feMul
    >>> referenced by fe.go:303 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:303)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-0789AEB7EFB32DDEE9CB09DBEDFBFD637F9C1A48:((*crypto/internal/edwards25519/field.Element).Multiply)
    error: failed to link /tmp/tinygo3040995961/main: exit status 1






## crypto/elliptic



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/internal/nistec.p256PointAddAsm
    >>> referenced by p256_asm.go:372 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:372)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-21418DD625739E361D10381CEF2116024FD473AA:((*crypto/internal/nistec.P256Point).Add)
    >>> referenced by p256_asm.go:652 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:652)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-21418DD625739E361D10381CEF2116024FD473AA:((*crypto/internal/nistec.P256Point).ScalarMult)
    >>> referenced by p256_asm.go:653 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:653)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-21418DD625739E361D10381CEF2116024FD473AA:((*crypto/internal/nistec.P256Point).ScalarMult)
    >>> referenced 7 more times
    
    ld.lld: error: undefined symbol: crypto/internal/nistec.p256PointDoubleAsm
    >>> referenced by p256_asm.go:373 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:373)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-21418DD625739E361D10381CEF2116024FD473AA:((*crypto/internal/nistec.P256Point).Add)
    >>> referenced by p256_asm.go:643 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:643)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-21418DD625739E361D10381CEF2116024FD473AA:((*crypto/internal/nistec.P256Point).ScalarMult)
    >>> referenced by p256_asm.go:644 (/usr/local/go1.20/src/crypto/internal/nistec/p256_asm.go:644)
    [...more lines following...]














## crypto/rsa



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/internal/bigmod.montgomeryLoop
    >>> referenced by nat.go:599 (/usr/local/go1.20/src/crypto/internal/bigmod/nat.go:599)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EBC3D31FF2668FC7A7C761DEF96D971602D6951:((*crypto/internal/bigmod.Nat).montgomeryMul)
    
    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EBC3D31FF2668FC7A7C761DEF96D971602D6951:((crypto/cipher.StreamReader).Read)
    
    ld.lld: error: undefined symbol: crypto/aes.encryptBlockAsm
    >>> referenced by cipher_asm.go:81 (/usr/local/go1.20/src/crypto/aes/cipher_asm.go:81)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EBC3D31FF2668FC7A7C761DEF96D971602D6951:((*crypto/aes.aesCipherAsm).Encrypt)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:309 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:309)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3EBC3D31FF2668FC7A7C761DEF96D971602D6951:((*crypto/internal/edwards25519/field.Element).Square)
    [...more lines following...]












## crypto/subtle



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-414ED2EC413D93DD8C112FA3A0B6257CABEB7F13:(crypto/subtle.XORBytes)
    FAIL	crypto/subtle	0.000s
    error: failed to link /tmp/tinygo3362733111/main: exit status 1






## crypto/tls


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#cryptoed25519)
  * [crypto/x509](#cryptox509)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#cryptoed25519)







## database/sql



The compiler gave the following error when running the tests for this package:


    FAIL	database/sql	0.000s
    # database/sql
    /usr/local/go1.20/src/database/sql/sql_test.go:4574:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)








## debug/buildinfo



The compiler gave the following error when this package was imported:


    # debug/buildinfo
    /usr/local/go1.20/src/debug/buildinfo/buildinfo.go:79:19: undefined: debug.ParseBuildInfo








## debug/elf



The compiler gave the following error when running the tests for this package:


    FAIL	debug/elf	0.000s
    # debug/elf
    /usr/local/go1.20/src/debug/elf/file_test.go:906:10: undefined: net.ResolveIPAddr






## debug/gosym



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x00000000002549ab: nil pointer dereference
    FAIL	debug/gosym	1.671s








## debug/pe



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x00000000002737d1: index out of range
    FAIL	debug/pe	1.148s








## embed



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3B6D6B682F19B044232E3E4C73B8A78F898E93C6:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3B6D6B682F19B044232E3E4C73B8A78F898E93C6:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3B6D6B682F19B044232E3E4C73B8A78F898E93C6:(runtime.initAll)
    >>> referenced 1 more times
    
    ld.lld: error: undefined symbol: vendor/golang.org/x/sys/cpu.xgetbv
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-3B6D6B682F19B044232E3E4C73B8A78F898E93C6:(runtime.initAll)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:309 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:309)
    [...more lines following...]
















## encoding/binary



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/binary	1.025s








## encoding/gob



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	encoding/gob	1.418s








## encoding/json



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: vendor/golang.org/x/sys/cpu.cpuid
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B546F5E015501B15C784EB346D71B542493E6867:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B546F5E015501B15C784EB346D71B542493E6867:(runtime.initAll)
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B546F5E015501B15C784EB346D71B542493E6867:(runtime.initAll)
    >>> referenced 1 more times
    
    ld.lld: error: undefined symbol: vendor/golang.org/x/sys/cpu.xgetbv
    >>> referenced by main
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-B546F5E015501B15C784EB346D71B542493E6867:(runtime.initAll)
    
    ld.lld: error: undefined symbol: crypto/subtle.xorBytes
    >>> referenced by xor.go:22 (/usr/local/go1.20/src/crypto/subtle/xor.go:22)
    [...more lines following...]






## encoding/pem



The compiler gave the following error when running the tests for this package:


    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feSquare
    >>> referenced by fe.go:309 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:309)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-FD0540891799D0F9AF48E24892644EB0DE3691B1:((*crypto/internal/edwards25519/field.Element).Square)
    
    ld.lld: error: undefined symbol: crypto/internal/edwards25519/field.feMul
    >>> referenced by fe.go:303 (/usr/local/go1.20/src/crypto/internal/edwards25519/field/fe.go:303)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-FD0540891799D0F9AF48E24892644EB0DE3691B1:((*crypto/internal/edwards25519/field.Element).Multiply)
    FAIL	encoding/pem	0.000s
    error: failed to link /tmp/tinygo1826234130/main: exit status 1






## encoding/xml



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/xml	1.169s






## errors



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAs (0.01s)
        --- FAIL: TestAs/8:As(Errorf(...,_err),_0x400081ee90) (0.00s)
            match: got true; want false
            FailNow is incomplete, requires runtime.Goexit()
            got &errors.errorString{s:"err"}, want <nil>
            FailNow is incomplete, requires runtime.Goexit()
        --- FAIL: TestAs/10:As(Errorf(...,_path_error),_0x400081ee90) (0.00s)
            got errors_test.wrapped{msg:"path error", err:(*fs.PathError)(0x400081eef0)}, want &fs.PathError{Op:"open", Path:"non-existing", Err:(*errors.errorString)(0x2520d8)}
            FailNow is incomplete, requires runtime.Goexit()
        --- FAIL: TestAs/16:As(Errorf(...,_multiError),_0x400081ee90) (0.00s)
            got errors_test.multiErr{errors_test.wrapped{msg:"path error", err:(*fs.PathError)(0x400081eef0)}}, want &fs.PathError{Op:"open", Path:"non-existing", Err:(*errors.errorString)(0x2520d8)}
            FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	errors	0.038s






## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## flag



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestExitCode (0.01s)
        unexpected exit code for test case {flag:-h flagHandle: expectExit:0} 
            : got -1, expect 0
        unexpected exit code for test case {flag:-help flagHandle: expectExit:0} 
            : got -1, expect 0
        unexpected exit code for test case {flag:-undefined flagHandle: expectExit:2} 
            : got -1, expect 2
        unexpected exit code for test case {flag:-h flagHandle:h expectExit:123} 
            : got -1, expect 123
        unexpected exit code for test case {flag:-help flagHandle:help expectExit:123} 
            : got -1, expect 123
    FAIL
    FAIL	flag	0.047s






## fmt



The compiler gave the following error when running the tests for this package:


    panic: reflect: call of Slice on array Value
    FAIL	fmt	1.154s








## go/build



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestDotSlashImport (0.01s)
        import ".": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
        import "./file": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
    --- FAIL: TestLocalDirectory (0.01s)
        import ".": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
    --- FAIL: TestImportCmd (0.00s)
        import "cmd/internal/objfile": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
    --- FAIL: TestImportDirNotExist (0.01s)
        --- FAIL: TestImportDirNotExist/GO111MODULE=on (0.00s)
            Import(full, 0) got error: "go/build: go list go/build/doesnotexist: fork/exec /usr/local/go/bin/go: operation not implemented\n\n", want "cannot find package" or "is not in GOROOT" error
            Import(full, FindOnly) got error: "go/build: go list go/build/doesnotexist: fork/exec /usr/local/go/bin/go: operation not implemented\n\n", want "cannot find package" or "is not in GOROOT" error
    [...more lines following...]








## go/constant



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x00000000002549ec: divide by zero
    FAIL	go/constant	1.037s






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	1.349s






## go/doc/comment



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestStd (0.01s)
        fork/exec /usr/local/go/bin/go: operation not implemented
        FailNow is incomplete, requires runtime.Goexit()
        stdPkgs is out of date: regenerate with 'go generate'
            diff stdPkgs want
            --- stdPkgs
            +++ want
            @@ -1,33 +1,1 @@
            -bufio
            -bytes
            -context
            -crypto
            -embed
            -encoding
            -errors
    [...more lines following...]








## go/importer



The compiler gave the following error when this package was imported:


    # internal/pkgbits
    /usr/local/go1.20/src/internal/pkgbits/sync.go:38:64: frame.Entry undefined (type runtime.Frame has no field or method Entry)






## go/parser



The compiler gave the following error when running the tests for this package:


    FAIL	go/parser	1.963s










## go/token



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	go/token	1.369s






## go/types



The compiler gave the following error when running the tests for this package:


    FAIL	go/types	0.000s
    # internal/pkgbits
    /usr/local/go1.20/src/internal/pkgbits/sync.go:38:64: frame.Entry undefined (type runtime.Frame has no field or method Entry)










## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	1.532s










## hash/maphash



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x0000000000223ba1: out of memory
    FAIL	hash/maphash	144.258s








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	1.105s








## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	1.340s








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	1.669s






## image/gif



The compiler gave the following error when running the tests for this package:


    FAIL	image/gif	0.000s
    # image/gif
    /usr/local/go1.20/src/image/gif/reader_test.go:421:26: s1.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)
    /usr/local/go1.20/src/image/gif/reader_test.go:421:41: s0.HeapAlloc undefined (type *runtime.MemStats has no field or method HeapAlloc)






## image/jpeg



The compiler gave the following error when running the tests for this package:


    FAIL	image/jpeg	0.000s
    # image/jpeg
    /usr/local/go1.20/src/image/jpeg/reader_test.go:253:9: undefined: debug.SetTraceback






## image/png



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x00000000002593d3: out of memory
    FAIL	image/png	3.639s








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
    FAIL	io	5.193s










## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.02s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2023/06/11 18:35:24.852021 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2023/06/11 18:35:24.852221 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2023/06/11 18:35:24.853505 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): hello 23 world$" is "XXX2023/06/11 18:35:24.853693 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2023/06/11 18:35:24.854965 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2023/06/11 18:35:24.855138 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2023/06/11 18:35:24.856385 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(61|63): XXXhello 23 world$" is "2023/06/11 18:35:24.856563 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/syslog



The compiler gave the following error when running the tests for this package:


    FAIL	log/syslog	0.000s
    # log/syslog
    /usr/local/go1.20/src/log/syslog/syslog_test.go:22:25: undefined: net.PacketConn
    /usr/local/go1.20/src/log/syslog/syslog_test.go:35:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /usr/local/go1.20/src/log/syslog/syslog_test.go:110:15: undefined: net.ListenPacket








## math/big



The compiler gave the following error when running the tests for this package:


    FAIL	math/big	0.000s
    # math/big
    /usr/local/go1.20/src/math/bits/bits.go:472:10: interp: running for more than 3m0s, timing out (executed calls: 28656473)
      <badref> = and half <badref>, <null operand!>fatal error: unexpected signal during runtime execution
    [signal SIGSEGV: segmentation violation code=0x1 addr=0xfff12bed61e4 pc=0xffff100db604]
    
    runtime stack:
    runtime.throw({0x7e48c2?, 0xfffee2c86230?})
    	/usr/local/go/src/runtime/panic.go:1047 +0x40 fp=0xfffee2c85ed0 sp=0xfffee2c85ea0 pc=0x44fa50
    runtime.sigpanic()
    	/usr/local/go/src/runtime/signal_unix.go:821 +0x240 fp=0xfffee2c85f10 sp=0xfffee2c85ed0 pc=0x466a90
    
    goroutine 8 [syscall]:
    runtime.cgocall(0x72cb60, 0x400bb79bf8)
    	/usr/local/go/src/runtime/cgocall.go:157 +0x50 fp=0x400bb79bc0 sp=0x400bb79b80 pc=0x4203e0
    [...more lines following...]






## math/bits



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x0000000000246ace: divide by zero
    FAIL	math/bits	1.337s








## math/rand



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand	26.906s






## mime



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x000000000022dc26: deadlocked: no event source
    FAIL	mime	1.292s












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
  * [net/http/httptrace](#nethttphttptrace)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)







## net/netip



The compiler gave the following error when running the tests for this package:


    FAIL	net/netip	0.000s
    # net/netip_test
    /usr/local/go1.20/src/net/netip/fuzz_test.go:177:32: stdip.IsPrivate undefined (type net.IP has no field or method IsPrivate)






## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/rpc](#netrpc)





## net/smtp


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)







## net/url



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	net/url	1.292s








## os/exec



The compiler gave the following error when running the tests for this package:


    FAIL	os/exec	0.000s
    # os/user
    ../../../../../../usr/local/go1.20/src/os/user/cgo_lookup_cgo.go:14:6: not implemented: build constraints in #cgo line






## os/signal



The compiler gave the following error when running the tests for this package:


    FAIL	os/signal	0.000s
    # internal/testpty
    ../../tinygo/lib/musl/include/fcntl.h:22:10: fatal: 'bits/fcntl.h' file not found
    ../../../../../../usr/local/go1.20/src/internal/testpty/pty_cgo.go:11:10: note: in file included from pty_cgo.go!cgo.c:4:






## os/user



The compiler gave the following error when this package was imported:


    # os/user
    ../../../../../../usr/local/go1.20/src/os/user/cgo_lookup_cgo.go:14:6: not implemented: build constraints in #cgo line








## path/filepath



The compiler gave the following error when running the tests for this package:


    FAIL	path/filepath	1.372s






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    ../../../../../../usr/local/go1.20/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line








## regexp



The compiler gave the following error when running the tests for this package:


    FAIL	regexp	10.812s








## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (9.51s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
        Stable      100 elements:         930 Swap,        820 Less
        Stable      300 elements:        3937 Swap,       2993 Less
        Stable     1000 elements:       19510 Swap,      12224 Less
        Stable     3000 elements:       82341 Swap,      43082 Less
        Stable    10000 elements:      349590 Swap,     165816 Less
        Stable    30000 elements:     1341046 Swap,     558467 Less
        Stable   100000 elements:     5787139 Swap,    2084649 Less
        Stable   300000 elements:    20447012 Swap,    6868979 Less
        Stable  1000000 elements:    84481603 Swap,   25116491 Less
    --- FAIL: TestCountSortOps (1.47s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## strings



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestBuilderGrow (0.18s)
        growLen=100: got 0 allocs during Write; want 1
        growLen=1000: got 0 allocs during Write; want 1
        growLen=10000: got 0 allocs during Write; want 1
        growLen=100000: got 0 allocs during Write; want 1
    --- FAIL: TestBuilderAllocs (0.02s)
        Builder allocs = 0; want 1
    FAIL
    FAIL	strings	40.097s








## sync/atomic



The compiler gave the following error when running the tests for this package:


    FAIL	sync/atomic	5.916s






## syscall



The compiler gave the following error when running the tests for this package:


    FAIL	syscall	0.000s
    # os/user
    ../../../../../../usr/local/go1.20/src/os/user/cgo_lookup_cgo.go:14:6: not implemented: build constraints in #cgo line












## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	0.972s










## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	1.088s








## time



The compiler gave the following error when running the tests for this package:


    FAIL	time	0.000s
    # time_test
    /usr/local/go1.20/src/time/sleep_test.go:714:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /usr/local/go1.20/src/time/sleep_test.go:715:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /usr/local/go1.20/src/time/sleep_test.go:716:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /usr/local/go1.20/src/time/sleep_test.go:787:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /usr/local/go1.20/src/time/sleep_test.go:788:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /usr/local/go1.20/src/time/sleep_test.go:789:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)






## time/tzdata



The compiler gave the following error when this package was imported:


    ld.lld: error: undefined symbol: time.registerLoadFromEmbeddedTZData
    >>> referenced by scheduler_any.go:24 (/home/ayke/src/tinygo/tinygo/src/runtime/scheduler_any.go:24)
    >>>               /home/ayke/.cache/tinygo/thinlto/llvmcache-809755382705B4E52ACE223C1D98D60EC110E4FB:(runtime.run$1$gowrapper)
    error: failed to link /tmp/tinygo4244802663/main: exit status 1













