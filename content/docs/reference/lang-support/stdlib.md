
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Test results are for linux/amd64.

Package | Importable | Passes tests
--- | --- | --- |
archive/tar |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#archivetar)  | 
archive/zip |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
bufio |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bufio)  | 
bytes |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bytes)  | 
cmp |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
crypto/ed25519 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/fips140 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hkdf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptohmac)  | 
crypto/md5 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/mlkem |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/pbkdf2 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptorc4)  | 
crypto/rsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptorsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosha256)  | 
crypto/sha3 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosha512)  | 
crypto/subtle |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosubtle)  | 
crypto/tls |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/x509 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptox509)  | 
crypto/x509/pkix |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesql)  | 
database/sql/driver |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/buildinfo |  [<span style="color: red">✗</span> no](#debugbuildinfo)  |  <span style="color: gray">✗</span> no  | 
debug/dwarf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/elf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/gosym |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debuggosym)  | 
debug/macho |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/pe |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#debugpe)  | 
debug/plan9obj |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
embed |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
encoding/xml |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
errors |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#errors)  | 
expvar |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
flag |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#flag)  | 
fmt |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#fmt)  | 
go/ast |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/build |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gobuild)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/constant |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goconstant)  | 
go/doc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#godoc)  | 
go/doc/comment |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#godoccomment)  | 
go/format |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/importer |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goimporter)  | 
go/parser |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goparser)  | 
go/printer |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/scanner |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/token |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gotoken)  | 
go/types |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gotypes)  | 
go/version |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#hashcrc32)  | 
hash/crc64 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/fnv |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/maphash |  [<span style="color: red">✗</span> no](#hashmaphash)  |  <span style="color: gray">✗</span> no  | 
html |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
html/template |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#htmltemplate)  | 
image |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/color |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagecolor)  | 
image/color/palette |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/draw |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagedraw)  | 
image/gif |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/jpeg |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagejpeg)  | 
image/png |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
index/suffixarray |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#io)  | 
io/fs |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io/ioutil |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
iter |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#iter)  | 
log |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#log)  | 
log/slog |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#logslog)  | 
log/syslog |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#logsyslog)  | 
maps |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/big |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbig)  | 
math/bits |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathbits)  | 
math/cmplx |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathrand)  | 
math/rand/v2 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mathrandv2)  | 
mime |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
mime/multipart |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/cgi |  [<span style="color: red">✗</span> no](#nethttpcgi)  |  <span style="color: gray">✗</span> no  | 
net/http/cookiejar |  [<span style="color: red">✗</span> no](#nethttpcookiejar)  |  <span style="color: gray">✗</span> no  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#nethttpfcgi)  |  <span style="color: gray">✗</span> no  | 
net/http/httptest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/httptrace |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/httputil |  [<span style="color: red">✗</span> no](#nethttphttputil)  |  <span style="color: gray">✗</span> no  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#nethttppprof)  |  <span style="color: gray">✗</span> no  | 
net/mail |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/netip |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netnetip)  | 
net/rpc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netrpc)  | 
net/rpc/jsonrpc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netrpcjsonrpc)  | 
net/smtp |  [<span style="color: red">✗</span> no](#netsmtp)  |  <span style="color: gray">✗</span> no  | 
net/textproto |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/url |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#neturl)  | 
os |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
os/exec |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#osexec)  | 
os/signal |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#ossignal)  | 
os/user |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path/filepath |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#pathfilepath)  | 
plugin |  [<span style="color: red">✗</span> no](#plugin)  |  <span style="color: gray">✗</span> no  | 
reflect |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
regexp |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
regexp/syntax |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
slices |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#slices)  | 
sort |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#sort)  | 
strconv |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
strings |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strings)  | 
structs |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync/atomic |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syncatomic)  | 
syscall |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syscall)  | 
testing |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/fstest |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingfstest)  | 
testing/iotest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/quick |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingquick)  | 
testing/slogtest |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingslogtest)  | 
testing/synctest |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingsynctest)  | 
text/scanner |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
text/template |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#texttemplate)  | 
text/template/parse |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
time |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#time)  | 
time/tzdata |  [<span style="color: red">✗</span> no](#timetzdata)  |  <span style="color: gray">✗</span> no  | 
unicode |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unique |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
weak |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#weak)  | 



## archive/tar



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x000000000027f886: nil pointer dereference
    FAIL	archive/tar	1.413s








## bufio



The compiler gave the following error when running the tests for this package:


    FAIL	bufio	0.000s
    # bufio_test
    /home/ron/.gvm/gos/go1.25.5/src/bufio/net_test.go:31:21: undefined: net.ListenUnix
    /home/ron/.gvm/gos/go1.25.5/src/bufio/net_test.go:70:25: undefined: net.DialUnix






## bytes



The compiler gave the following error when running the tests for this package:


    FAIL	bytes	0.000s
    # bytes_test
    /home/ron/.gvm/gos/go1.25.5/src/math/rand/rng.go:239:6: interp: running for more than 3m0s, timing out (executed calls: 582422)
      %0 = icmp eq ptr %rng, null, !dbg !19942
    
    traceback:
    /home/ron/.gvm/gos/go1.25.5/src/math/rand/rng.go:239:6:
      %0 = icmp eq ptr %rng, null, !dbg !19942
    /home/ron/.gvm/gos/go1.25.5/src/math/rand/rng.go:234:25:
      %0 = call i64 @"(*math/rand.rngSource).Uint64"(ptr %rng, ptr undef), !dbg !19939
    /home/ron/.gvm/gos/go1.25.5/src/math/rand/rand.go:96:50:
      %4 = call i64 @"interface:{Int63:func:{}{basic:int64},Seed:func:{basic:int64}{}}.Int63$invoke"(ptr %invoke.func.value, ptr %invoke.func.typecode, ptr undef), !dbg !19933
    /home/ron/.gvm/gos/go1.25.5/src/math/rand/rand.go:110:52:
      %0 = call i64 @"(*math/rand.Rand).Int63"(ptr %r, ptr undef), !dbg !19932
    /home/ron/.gvm/gos/go1.25.5/src/math/rand/rand.go:145:14:
    [...more lines following...]












## compress/gzip



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestGZIPFilesHaveZeroMTimes (0.00s)
        skipping test on non-builder
        SkipNow is incomplete, requires runtime.Goexit()
        error evaluating GOROOT:  lstat /usr/local/go: file does not exist
        FailNow is incomplete, requires runtime.Goexit()
        skipping: GOROOT directory not found: /usr/local/go
        SkipNow is incomplete, requires runtime.Goexit()
        error collecting list of .gz files in GOROOT:  lstat : file does not exist
        FailNow is incomplete, requires runtime.Goexit()
        expected to find some .gz files under GOROOT
        FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	compress/gzip	0.551s
















## context



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCauseRace (0.00s)
        Cause returned context canceled, want TestCauseRace
    FAIL
    FAIL	context	0.024s






## crypto



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x00000000002662d7: nil pointer dereference
    FAIL	crypto	1.120s






## crypto/aes



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/aes	1.091s






## crypto/cipher



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/cipher	1.097s










## crypto/ecdh



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestLinker (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.25.5/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.25.5/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        [ build -o hello.exe hello.go]: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        [./hello.exe]: directory setting not implemented
        FailNow is incomplete, requires runtime.Goexit()
        unexpected output: 
        [ tool nm hello.exe]: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        no P256 symbols found in program using ecdh.P256, test is broken
    [...more lines following...]






## crypto/ecdsa



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/crypto/ecdsa/ecdsa_test.go:563:6]
    panic: runtime error at 0x00000000002bf5e1: index out of range
    FAIL	crypto/ecdsa	1.970s






## crypto/ed25519



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestEd25519Vectors (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.25.5/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
         env GOMODCACHE: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        fetching filippo.io/mostly-harmless/ed25519vectors@v0.0.0-20210322192420-30a2d7243a94
        failed to download filippo.io/mostly-harmless/ed25519vectors@v0.0.0-20210322192420-30a2d7243a94: exec: no command
            
        FailNow is incomplete, requires runtime.Goexit()
        failed to parse 'go mod download': unexpected end of JSON input
            
        FailNow is incomplete, requires runtime.Goexit()
        failed to read ed25519vectors.json: open ed25519vectors.json: file does not exist
    [...more lines following...]












## crypto/hmac



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/hmac	1.141s














## crypto/rc4



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/crypto/rc4/rc4.go]
    panic: runtime error at 0x00000000002607c7: index out of range
    FAIL	crypto/rc4	1.140s






## crypto/rsa



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	crypto/rsa	1.101s








## crypto/sha256



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/sha256	1.165s








## crypto/sha512



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/sha512	2.168s






## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	2.273s








## crypto/x509



The compiler gave the following error when running the tests for this package:


    FAIL	crypto/x509	0.000s
    # crypto/x509_test
    /home/ron/.gvm/gos/go1.25.5/src/crypto/x509/hybrid_pool_test.go:63:17: c.ConnectionState undefined (type *net.TLSConn has no field or method ConnectionState)








## database/sql



The compiler gave the following error when running the tests for this package:


    FAIL	database/sql	0.000s
    # database/sql
    /home/ron/.gvm/gos/go1.25.5/src/database/sql/sql_test.go:4840:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)








## debug/buildinfo



The compiler gave the following error when this package was imported:


    # debug/buildinfo
    /home/ron/.gvm/gos/go1.25.5/src/debug/buildinfo/buildinfo.go:93:19: undefined: debug.ParseBuildInfo










## debug/gosym



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/debug/elf/file.go:129:5]
    panic: runtime error at 0x000000000026cff6: nil pointer dereference
    FAIL	debug/gosym	1.167s








## debug/pe



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/debug/pe/file_test.go]
    panic: runtime error at 0x000000000028a7f4: index out of range
    FAIL	debug/pe	1.122s




















## encoding/binary



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/encoding/binary/binary.go:116:7]
    panic: runtime error at 0x000000000025c723: index out of range
    FAIL	encoding/binary	1.091s








## encoding/gob



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	encoding/gob	1.110s








## encoding/json



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/json	0.000s
    # encoding/json
    /home/ron/.gvm/gos/go1.25.5/src/encoding/json/encode_test.go:1415:6: wg.Go undefined (type sync.WaitGroup has no field or method Go)






## encoding/pem



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	encoding/pem	1.567s








## errors



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAs (0.00s)
        --- FAIL: TestAs/8:As(Errorf(...,_err),_0x73ef014b8840) (0.00s)
            match: got true; want false
            FailNow is incomplete, requires runtime.Goexit()
            got &errors.errorString{s:"err"}, want <nil>
            FailNow is incomplete, requires runtime.Goexit()
        --- FAIL: TestAs/10:As(Errorf(...,_path_error),_0x73ef014b8840) (0.00s)
            got errors_test.wrapped{msg:"path error", err:(*fs.PathError)(0x73ef014c1d80)}, want &fs.PathError{Op:"open", Path:"non-existing", Err:(*errors.errorString)(0x267b88)}
            FailNow is incomplete, requires runtime.Goexit()
        --- FAIL: TestAs/16:As(Errorf(...,_multiError),_0x73ef014b8840) (0.00s)
            got errors_test.multiErr{errors_test.wrapped{msg:"path error", err:(*fs.PathError)(0x73ef014c1d80)}}, want &fs.PathError{Op:"open", Path:"non-existing", Err:(*errors.errorString)(0x267b88)}
            FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	errors	0.001s








## flag



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestExitCode (0.00s)
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
    --- FAIL: TestDefineAfterSet (0.00s)
        DefineAfterSet
            : expected panic("flag myFlag set at .*/flag_test.go:.* before being defined"), but got panic("flag myFlag set at ?:0 before being defined")
    FAIL
    [...more lines following...]






## fmt



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Value).Method()
    FAIL	fmt	1.093s








## go/build



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestDotSlashImport (0.00s)
        import ".": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
        import "./file": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
    --- FAIL: TestLocalDirectory (0.00s)
        import ".": unknown compiler "tinygo"
        FailNow is incomplete, requires runtime.Goexit()
        ImportPath=".", want "go/build"
        FailNow is incomplete, requires runtime.Goexit()
    --- FAIL: TestImportCmd (0.00s)
        go/build: go list cmd/internal/objfile: files setting not implemented
            
            
        FailNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## go/constant



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/go/constant/value.go]
    panic: runtime error at 0x000000000026b79c: divide by zero
    FAIL	go/constant	1.129s






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	1.106s






## go/doc/comment



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestStd (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.25.5/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        stdPkgs is out of date: regenerate with 'go generate'
            diff stdPkgs want
            --- stdPkgs
            +++ want
            @@ -1,40 +1,1 @@
            -bufio
            -bytes
            -cmp
    [...more lines following...]








## go/importer



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x000000000033f069: nil pointer dereference
    FAIL	go/importer	1.087s






## go/parser



The compiler gave the following error when running the tests for this package:


    FAIL	go/parser	1.171s










## go/token



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	go/token	1.135s






## go/types



The compiler gave the following error when running the tests for this package:


    FAIL	go/types	0.000s
    # go/types_test
    /home/ron/.gvm/gos/go1.25.5/src/go/types/self_test.go:102:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)












## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	1.152s










## hash/maphash



The compiler gave the following error when this package was imported:


    # hash/maphash
    /home/ron/.gvm/gos/go1.25.5/src/hash/maphash/maphash.go:295:6: undefined: abi.EscapeNonString
    /home/ron/.gvm/gos/go1.25.5/src/hash/maphash/maphash.go:301:6: undefined: abi.EscapeNonString








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	1.105s








## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	1.095s








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	1.158s








## image/jpeg



The compiler gave the following error when running the tests for this package:


    FAIL	image/jpeg	0.000s
    /home/ron/.gvm/gos/go1.25.5/src/image/jpeg/reader_test.go:253: linker could not find symbol runtime/debug.SetTraceback










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
    FAIL	io	5.028s










## iter



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestPull (0.00s)
        --- FAIL: TestPull/0 (0.00s)
            have 0 extra goroutines, want 1
            have 0 extra goroutines, want 1
        --- FAIL: TestPull/1 (0.00s)
            have 0 extra goroutines, want 1
            have 0 extra goroutines, want 1
            have 0 extra goroutines, want 1
        --- FAIL: TestPull/2 (0.00s)
            have 0 extra goroutines, want 1
            have 0 extra goroutines, want 1
            have 0 extra goroutines, want 1
            have 0 extra goroutines, want 1
        --- FAIL: TestPull/3 (0.00s)
            have 0 extra goroutines, want 1
    [...more lines following...]






## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.00s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2025/12/15 13:33:52.466132 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2025/12/15 13:33:52.466247 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2025/12/15 13:33:52.466279 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2025/12/15 13:33:52.466305 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2025/12/15 13:33:52.466329 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2025/12/15 13:33:52.466443 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2025/12/15 13:33:52.466548 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2025/12/15 13:33:52.466578 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/slog



The compiler gave the following error when running the tests for this package:


    panic: runtime.Caller failed
    FAIL	log/slog	1.114s






## log/syslog



The compiler gave the following error when running the tests for this package:


    FAIL	log/syslog	0.000s
    # log/syslog
    /home/ron/.gvm/gos/go1.25.5/src/log/syslog/syslog_test.go:35:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /home/ron/.gvm/gos/go1.25.5/src/log/syslog/syslog_test.go:110:15: undefined: net.ListenPacket










## math/big



The compiler gave the following error when running the tests for this package:


    FAIL	math/big	0.000s
    # math/big
    /home/ron/.gvm/gos/go1.25.5/src/math/bits/bits.go:472:10: interp: running for more than 3m0s, timing out (executed calls: 31156621)
      %0 = and i64 %x, 4294967295, !dbg !31805
    
    traceback:
    /home/ron/.gvm/gos/go1.25.5/src/math/bits/bits.go:472:10:
      %0 = and i64 %x, 4294967295, !dbg !31805
    /home/ron/.gvm/gos/go1.25.5/src/math/bits/bits.go:450:15:
      %9 = call { i64, i64 } @"math/bits.Mul64"(i64 %x, i64 %y, ptr undef), !dbg !31813
    /home/ron/.gvm/gos/go1.25.5/src/math/big/arith.go:50:20:
      %0 = call { i64, i64 } @"math/bits.Mul"(i64 %x, i64 %y, ptr undef), !dbg !31800
    /home/ron/.gvm/gos/go1.25.5/src/math/big/arith.go:250:24:
      %14 = call { i64, i64 } @"math/big.mulAddWWW_g"(i64 %13, i64 %y, i64 %7, ptr undef), !dbg !31817
    /home/ron/.gvm/gos/go1.25.5/src/math/big/arith_decl_pure.go:26:20:
    [...more lines following...]






## math/bits



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/math/bits/bits.go:510:21]
    panic: runtime error at 0x000000000022d830: divide by zero
    FAIL	math/bits	1.623s








## math/rand



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand	3.081s






## math/rand/v2



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand/v2	3.623s
















## net/http/cgi



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-1980316787/main.go:2:8: package net/http/cgi is not in std (/home/ron/.cache/tinygo/goroot-ab68f8ed5d891c12323b7956a84922b36d555b19c5ca444b3d1f322300c5d23f/src/net/http/cgi)






## net/http/cookiejar



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-3073495770/main.go:2:8: package net/http/cookiejar is not in std (/home/ron/.cache/tinygo/goroot-ab68f8ed5d891c12323b7956a84922b36d555b19c5ca444b3d1f322300c5d23f/src/net/http/cookiejar)






## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http/cgi](#nethttpcgi)









## net/http/httputil



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-1974920281/main.go:2:8: package net/http/httputil is not in std (/home/ron/.cache/tinygo/goroot-ab68f8ed5d891c12323b7956a84922b36d555b19c5ca444b3d1f322300c5d23f/src/net/http/httputil)






## net/http/pprof



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-530727302/main.go:2:8: package net/http/pprof is not in std (/home/ron/.cache/tinygo/goroot-ab68f8ed5d891c12323b7956a84922b36d555b19c5ca444b3d1f322300c5d23f/src/net/http/pprof)








## net/netip



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestInlining (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.25.5/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.25.5/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        go build: exec: no command, 
        FailNow is incomplete, requires runtime.Goexit()
        "(*uint128).halves" is no longer inlinable
        "Addr.BitLen" is no longer inlinable
        "Addr.hasZone" is no longer inlinable
        "Addr.Is4" is no longer inlinable
        "Addr.Is4In6" is no longer inlinable
        "Addr.Is6" is no longer inlinable
    [...more lines following...]






## net/rpc



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	net/rpc	1.101s






## net/rpc/jsonrpc



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	net/rpc/jsonrpc	1.089s






## net/smtp



The compiler gave the following error when this package was imported:


    # net/smtp
    /home/ron/.gvm/gos/go1.25.5/src/net/smtp/smtp.go:72:24: undefined: tls.Conn
    /home/ron/.gvm/gos/go1.25.5/src/net/smtp/smtp.go:172:25: undefined: tls.Conn








## net/url



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	net/url	1.112s








## os/exec



The compiler gave the following error when running the tests for this package:


    FAIL	os/exec	0.000s
    # os/exec_test
    /home/ron/.gvm/gos/go1.25.5/src/os/exec/exec_test.go:260:17: undefined: net.FileListener
    /home/ron/.gvm/gos/go1.25.5/src/os/exec/exec_test.go:405:92: cmd.ProcessState.Pid undefined (type *os.ProcessState has no field or method Pid)
    /home/ron/.gvm/gos/go1.25.5/src/os/exec/exec_test.go:704:34: ln.(*net.TCPListener).File undefined (type *net.TCPListener has no field or method File)
    /home/ron/.gvm/gos/go1.25.5/src/os/exec/exec_test.go:709:18: undefined: net.FileListener
    /home/ron/.gvm/gos/go1.25.5/src/os/exec/exec_test.go:720:5: ts.StartTLS undefined (type *httptest.Server has no field or method StartTLS)
    /home/ron/.gvm/gos/go1.25.5/src/os/exec/exec_test.go:809:35: ln.(*net.TCPListener).File undefined (type *net.TCPListener has no field or method File)






## os/signal



The compiler gave the following error when running the tests for this package:


    FAIL	os/signal	0.000s
    /home/ron/.gvm/gos/go1.25.5/src/os/signal/signal_unix.go:61: linker could not find symbol os/signal.signal_ignored
    /home/ron/Development/tinygo/tinygo-122/src/syscall/syscall_linux.go:1127: linker could not find symbol syscall.runtime_doAllThreadsSyscall










## path/filepath



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestWindowsGlob (0.00s)
        skipping windows specific test
        SkipNow is incomplete, requires runtime.Goexit()
        tmpDir path "/tmp/TestWindowsGlob227586258/000" must have drive letter in it
        FailNow is incomplete, requires runtime.Goexit()
        Glob("/tmp/TestWindowsGlob227586258/000\\a") returns [], but ["/tmp/TestWindowsGlob227586258/000\\a"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\a") returns [], but ["/tmp/TestWindowsGlob227586258/000\\a"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\a") returns [], but ["/tmp/TestWindowsGlob227586258/000\\a"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\b") returns [], but ["/tmp/TestWindowsGlob227586258/000\\b"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\b") returns [], but ["/tmp/TestWindowsGlob227586258/000\\b"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\b") returns [], but ["/tmp/TestWindowsGlob227586258/000\\b"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\*") returns [], but ["/tmp/TestWindowsGlob227586258/000\\a" "/tmp/TestWindowsGlob227586258/000\\b" "/tmp/TestWindowsGlob227586258/000\\dir"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\*") returns [], but ["/tmp/TestWindowsGlob227586258/000\\a" "/tmp/TestWindowsGlob227586258/000\\b" "/tmp/TestWindowsGlob227586258/000\\dir"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\*") returns [], but ["/tmp/TestWindowsGlob227586258/000\\a" "/tmp/TestWindowsGlob227586258/000\\b" "/tmp/TestWindowsGlob227586258/000\\dir"] expected
        Glob("/tmp/TestWindowsGlob227586258/000\\d*") returns [], but ["/tmp/TestWindowsGlob227586258/000\\dir"] expected
    [...more lines following...]






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    /home/ron/.gvm/gos/go1.25.5/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line












## slices



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/slices/slices.go]
    panic: runtime error at 0x000000000028f2a0: slice out of range
    FAIL	slices	1.099s






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (1.02s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
        Stable      100 elements:         939 Swap,        809 Less
        Stable      300 elements:        3892 Swap,       2931 Less
        Stable     1000 elements:       19580 Swap,      12290 Less
        Stable     3000 elements:       81917 Swap,      42714 Less
        Stable    10000 elements:      347895 Swap,     165662 Less
        Stable    30000 elements:     1345166 Swap,     559451 Less
        Stable   100000 elements:     5785163 Swap,    2086453 Less
        Stable   300000 elements:    20448330 Swap,    6867550 Less
        Stable  1000000 elements:    84484245 Swap,   25118775 Less
    --- FAIL: TestCountSortOps (0.59s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## strings



The compiler gave the following error when running the tests for this package:


    FAIL	strings	0.000s
    # strings_test
    /home/ron/Development/tinygo/tinygo-122/src/runtime/runtime.go:48:6: interp: running for more than 3m0s, timing out (executed calls: 842273)
      call void @llvm.memmove.p0.p0.i64(ptr %dst, ptr %src, i64 %size, i1 false), !dbg !20241
    
    traceback:
    /home/ron/Development/tinygo/tinygo-122/src/runtime/runtime.go:48:6:
      call void @llvm.memmove.p0.p0.i64(ptr %dst, ptr %src, i64 %size, i1 false), !dbg !20241
    /home/ron/Development/tinygo/tinygo-122/src/runtime/slice.go:21:10:
      call void @runtime.memmove(ptr %7, ptr %elemsBuf, i64 %8, ptr undef), !dbg !20271
    /home/ron/.gvm/gos/go1.25.5/src/strings/strings_test.go:1884:13:
      %append.new = call { ptr, i64, i64 } @runtime.sliceAppend(ptr %append.srcBuf, ptr %append.elemsBuf, i64 %append.srcLen, i64 %append.srcCap, i64 %append.elemsLen, i64 1, ptr undef), !dbg !20273
    /home/ron/.gvm/gos/go1.25.5/src/strings:
      %56 = call %runtime._string @strings_test.makeBenchInputHard(ptr undef), !dbg !20252
    
    [...more lines following...]










## sync/atomic



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.25.5/src/sync/atomic/doc.go:93:6]
    panic: runtime error at 0x000000000026029b: caught signal SIGSEGV
    FAIL	sync/atomic	3.639s






## syscall



The compiler gave the following error when running the tests for this package:


    FAIL	syscall	0.000s
    # syscall_test
    /home/ron/.gvm/gos/go1.25.5/src/syscall/creds_test.go:53:19: undefined: net.FileConn
    /home/ron/.gvm/gos/go1.25.5/src/syscall/creds_test.go:60:19: undefined: net.FileConn
    /home/ron/.gvm/gos/go1.25.5/src/syscall/syscall_unix_test.go:183:16: undefined: net.FileConn
    /home/ron/.gvm/gos/go1.25.5/src/syscall/syscall_unix_test.go:238:14: undefined: net.UnixConn
    /home/ron/.gvm/gos/go1.25.5/src/syscall/syscall_unix_test.go:242:18: undefined: net.FileConn








## testing/fstest



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x0000000000232228: nil pointer dereference
    FAIL	testing/fstest	1.108s








## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	1.102s






## testing/slogtest



The compiler gave the following error when running the tests for this package:


    panic: runtime.Caller failed
    FAIL	testing/slogtest	1.115s






## testing/synctest



The compiler gave the following error when running the tests for this package:


    FAIL	testing/synctest	0.000s
    # testing/synctest_test
    /home/ron/.gvm/gos/go1.25.5/src/testing/synctest/example_test.go:109:4: unknown field DialContext in struct literal of type http.Transport
    /home/ron/.gvm/gos/go1.25.5/src/testing/synctest/example_test.go:113:4: unknown field ExpectContinueTimeout in struct literal of type http.Transport










## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	1.105s








## time



The compiler gave the following error when running the tests for this package:


    FAIL	time	0.000s
    # time_test
    /home/ron/.gvm/gos/go1.25.5/src/time/sleep_test.go:977:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.25.5/src/time/sleep_test.go:978:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.25.5/src/time/sleep_test.go:979:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.25.5/src/time/sleep_test.go:1050:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.25.5/src/time/sleep_test.go:1051:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.25.5/src/time/sleep_test.go:1052:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)






## time/tzdata



The compiler gave the following error when this package was imported:


    /home/ron/Development/tinygo/tinygo-122/src/runtime/scheduler_threads.go:26: linker could not find symbol time.registerLoadFromEmbeddedTZData
















## weak



The compiler gave the following error when running the tests for this package:


    FAIL	weak	0.000s
    /home/ron/.gvm/gos/go1.25.5/src/weak/pointer.go: linker could not find symbol weak.runtime_makeStrongFromWeak
    /home/ron/.gvm/gos/go1.25.5/src/weak/pointer.go: linker could not find symbol weak.runtime_makeStrongFromWeak





