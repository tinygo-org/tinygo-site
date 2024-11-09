
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
context |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#crypto)  | 
crypto/aes |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/cipher |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptocipher)  | 
crypto/des |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/dsa |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ecdh |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoecdh)  | 
crypto/ecdsa |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ed25519 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptorsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/subtle |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosubtle)  | 
crypto/tls |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/x509 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptox509)  | 
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
hash/maphash |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#hashmaphash)  | 
html |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
html/template |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#htmltemplate)  | 
image |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/color |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagecolor)  | 
image/color/palette |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/draw |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagedraw)  | 
image/gif |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/jpeg |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagejpeg)  | 
image/png |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#imagepng)  | 
index/suffixarray |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#io)  | 
io/fs |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#iofs)  | 
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
os/user |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path/filepath |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#pathfilepath)  | 
plugin |  [<span style="color: red">✗</span> no](#plugin)  |  <span style="color: gray">✗</span> no  | 
reflect |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
regexp |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#regexp)  | 
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



## archive/tar



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x000000000026eb43: nil pointer dereference
    FAIL	archive/tar	0.169s








## bufio



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestReadStringAllocs (0.00s)
        Unexpected number of allocations, got 0.000000, want 1
    FAIL
    FAIL	bufio	0.084s






## bytes



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/src/internal/bytealg/bytealg.go:304:13]
    panic: runtime error at 0x000000000022bb32: out of memory
    FAIL	bytes	5.594s












## compress/gzip



The compiler gave the following error when running the tests for this package:


    FAIL		0.000s
    /home/ron/.gvm/gos/go1.23/src/compress/gzip/example_test.go:14:2: package net/http/httptest is not in std (/home/ron/.cache/tinygo/goroot-59c0c8df7700d2980b426f170e6754c2d2901b0680adb52840eeef4402e8c177/src/net/http/httptest)


















## crypto



The compiler gave the following error when running the tests for this package:


    panic: runtime error at 0x000000000028e227: index out of range
    FAIL	crypto	0.004s








## crypto/cipher



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestGCMAsm (0.33s)
        no assembly implementation of GCM
        SkipNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	crypto/cipher	0.334s










## crypto/ecdh



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestLinker (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.23/bin/go tool -n compile: fork/exec /home/ron/.gvm/gos/go1.23/bin/go: operation not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.23/bin/go tool -n compile: fork/exec /home/ron/.gvm/gos/go1.23/bin/go: operation not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        [ build -o hello.exe hello.go]: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        [./hello.exe]: fork/exec ./hello.exe: operation not implemented
        FailNow is incomplete, requires runtime.Goexit()
        unexpected output: 
        [ tool nm hello.exe]: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        no P384 symbols found in program using ecdh.P384, test is broken
    [...more lines following...]








## crypto/ed25519



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestEd25519Vectors (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.23/bin/go tool -n compile: fork/exec /home/ron/.gvm/gos/go1.23/bin/go: operation not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        failed to run `go mod download -json filippo.io/mostly-harmless/ed25519vectors@v0.0.0-20210322192420-30a2d7243a94`, output: 
        FailNow is incomplete, requires runtime.Goexit()
        unexpected end of JSON input
        FailNow is incomplete, requires runtime.Goexit()
        failed to read ed25519vectors.json: open ed25519vectors.json: file does not exist
        FailNow is incomplete, requires runtime.Goexit()
        unexpected end of JSON input
        FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	crypto/ed25519	0.152s
    [...more lines following...]
















## crypto/rsa



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	crypto/rsa	0.037s












## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	0.001s








## crypto/x509



The compiler gave the following error when running the tests for this package:


    FAIL	crypto/x509	0.000s
    # crypto/x509_test
    /home/ron/.gvm/gos/go1.23/src/crypto/x509/hybrid_pool_test.go:63:17: c.ConnectionState undefined (type *net.TLSConn has no field or method ConnectionState)








## database/sql



The compiler gave the following error when running the tests for this package:


    FAIL	database/sql	0.000s
    # database/sql
    /home/ron/.gvm/gos/go1.23/src/database/sql/sql_test.go:4804:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)








## debug/buildinfo



The compiler gave the following error when this package was imported:


    # debug/buildinfo
    /home/ron/.gvm/gos/go1.23/src/debug/buildinfo/buildinfo.go:89:19: undefined: debug.ParseBuildInfo








## debug/elf



The compiler gave the following error when running the tests for this package:


    FAIL	debug/elf	0.000s
    # debug/elf
    /home/ron/.gvm/gos/go1.23/src/debug/elf/file_test.go:1159:10: undefined: net.ResolveIPAddr






## debug/gosym



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/debug/elf/file.go]
    panic: runtime error at 0x000000000025c02b: nil pointer dereference
    FAIL	debug/gosym	0.004s








## debug/pe



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/debug/pe/file_test.go]
    panic: runtime error at 0x0000000000275fa9: index out of range
    FAIL	debug/pe	0.004s




















## encoding/binary



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/encoding/binary/binary.go:109:7]
    panic: runtime error at 0x000000000024c031: index out of range
    FAIL	encoding/binary	0.006s








## encoding/gob



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	encoding/gob	0.001s








## encoding/json



The compiler gave the following error when running the tests for this package:


    FAIL		0.000s
    /home/ron/.gvm/gos/go1.23/src/encoding/json/stream_test.go:14:2: package net/http/httptest is not in std (/home/ron/.cache/tinygo/goroot-59c0c8df7700d2980b426f170e6754c2d2901b0680adb52840eeef4402e8c177/src/net/http/httptest)






## encoding/pem



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	encoding/pem	1.136s






## encoding/xml



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/xml	0.022s






## errors



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAs (0.00s)
        --- FAIL: TestAs/8:As(Errorf(...,_err),_0x78d6cd613c10) (0.00s)
            match: got true; want false
            FailNow is incomplete, requires runtime.Goexit()
            got &errors.errorString{s:"err"}, want <nil>
            FailNow is incomplete, requires runtime.Goexit()
        --- FAIL: TestAs/10:As(Errorf(...,_path_error),_0x78d6cd613c10) (0.00s)
            got errors_test.wrapped{msg:"path error", err:(*fs.PathError)(0x78d6cd613c70)}, want &fs.PathError{Op:"open", Path:"non-existing", Err:(*errors.errorString)(0x253a08)}
            FailNow is incomplete, requires runtime.Goexit()
        --- FAIL: TestAs/16:As(Errorf(...,_multiError),_0x78d6cd613c10) (0.00s)
            got errors_test.multiErr{errors_test.wrapped{msg:"path error", err:(*fs.PathError)(0x78d6cd613c70)}}, want &fs.PathError{Op:"open", Path:"non-existing", Err:(*errors.errorString)(0x253a08)}
            FailNow is incomplete, requires runtime.Goexit()
    FAIL
    FAIL	errors	0.001s






## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





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


    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/src/reflect/type.go:508:11]
    panic: runtime error at 0x0000000000228f19: nil pointer dereference
    FAIL	fmt	0.004s








## go/build



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestDotSlashImport (0.01s)
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
        go/build: go list cmd/internal/objfile: fork/exec /usr/local/go/bin/go: operation not implemented
            
            
        FailNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## go/constant



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/go/constant/value.go]
    panic: runtime error at 0x0000000000254be1: divide by zero
    FAIL	go/constant	0.004s






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	0.001s






## go/doc/comment



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestStd (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.23/bin/go tool -n compile: fork/exec /home/ron/.gvm/gos/go1.23/bin/go: operation not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        stdPkgs is out of date: regenerate with 'go generate'
            diff stdPkgs want
            --- stdPkgs
            +++ want
            @@ -1,39 +1,1 @@
            -bufio
            -bytes
            -cmp
    [...more lines following...]








## go/importer



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x000000000032148c: nil pointer dereference
    FAIL	go/importer	0.006s






## go/parser



The compiler gave the following error when running the tests for this package:


    FAIL	go/parser	0.075s










## go/token



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	go/token	0.001s






## go/types



The compiler gave the following error when running the tests for this package:


    FAIL	go/types	0.000s
    # go/types_test
    /home/ron/.gvm/gos/go1.23/src/go/types/self_test.go:103:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)












## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	0.014s










## hash/maphash



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/src/runtime/slice.go:57:14]
    panic: runtime error at 0x0000000000224cde: out of memory
    FAIL	hash/maphash	24.756s








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	0.003s








## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	0.001s








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	0.040s








## image/jpeg



The compiler gave the following error when running the tests for this package:


    FAIL	image/jpeg	0.000s
    # image/jpeg
    /home/ron/.gvm/gos/go1.23/src/image/jpeg/reader_test.go:253:9: undefined: debug.SetTraceback






## image/png



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/image/png/reader.go:506:12]
    panic: runtime error at 0x00000000002593c3: out of memory
    FAIL	image/png	0.864s








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
    FAIL	io	5.033s






## io/fs



The compiler gave the following error when running the tests for this package:


    FAIL	io/fs	0.004s








## iter



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/iter/pull_test.go]
    panic: runtime error at 0x0000000000242613: nil pointer dereference
    FAIL	iter	0.005s






## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.00s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "XXX2024/10/26 11:27:22.813528 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "XXX2024/10/26 11:27:22.813562 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "XXX2024/10/26 11:27:22.813600 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(63|65): hello 23 world$" is "XXX2024/10/26 11:27:22.813798 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(63|65): XXXhello 23 world$" is "2024/10/26 11:27:22.813838 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(63|65): XXXhello 23 world$" is "2024/10/26 11:27:22.814041 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(63|65): XXXhello 23 world$" is "2024/10/26 11:27:22.814243 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(63|65): XXXhello 23 world$" is "2024/10/26 11:27:22.814269 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/slog



The compiler gave the following error when running the tests for this package:


    panic: runtime.Caller failed
    FAIL	log/slog	0.001s






## log/syslog



The compiler gave the following error when running the tests for this package:


    FAIL	log/syslog	0.000s
    # log/syslog
    /home/ron/.gvm/gos/go1.23/src/log/syslog/syslog_test.go:35:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /home/ron/.gvm/gos/go1.23/src/log/syslog/syslog_test.go:110:15: undefined: net.ListenPacket










## math/big



The compiler gave the following error when running the tests for this package:


    FAIL	math/big	0.000s
    # math/big
    /home/ron/.gvm/gos/go1.23/src/math/bits/bits.go:472:10: interp: running for more than 3m0s, timing out (executed calls: 30224089)
      %0 = and i64 %x, 4294967295, !dbg !30956
    
    traceback:
    /home/ron/.gvm/gos/go1.23/src/math/bits/bits.go:472:10:
      %0 = and i64 %x, 4294967295, !dbg !30956
    /home/ron/.gvm/gos/go1.23/src/math/bits/bits.go:450:15:
      %9 = call { i64, i64 } @"math/bits.Mul64"(i64 %x, i64 %y, ptr undef), !dbg !30964
    /home/ron/.gvm/gos/go1.23/src/math/big/arith.go:51:20:
      %0 = call { i64, i64 } @"math/bits.Mul"(i64 %x, i64 %y, ptr undef), !dbg !30951
    /home/ron/.gvm/gos/go1.23/src/math/big/arith.go:192:24:
      %13 = call { i64, i64 } @"math/big.mulAddWWW_g"(i64 %12, i64 %y, i64 %6, ptr undef), !dbg !30969
    /home/ron/.gvm/gos/go1.23/src/math/big/arith_decl_pure.go:44:20:
    [...more lines following...]






## math/bits



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/math/bits/bits.go:510:21]
    panic: runtime error at 0x000000000021de1a: divide by zero
    FAIL	math/bits	0.002s








## math/rand



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand	1.918s






## math/rand/v2



The compiler gave the following error when running the tests for this package:


    FAIL	math/rand/v2	0.000s
    /home/ron/.gvm/gos/go1.23/src/internal/chacha8rand/chacha8.go:66: linker could not find symbol internal/chacha8rand.block
    /home/ron/.gvm/gos/go1.23/src/internal/chacha8rand/chacha8.go:91: linker could not find symbol internal/chacha8rand.block
    /home/ron/.gvm/gos/go1.23/src/internal/chacha8rand/chacha8.go:153: linker could not find symbol internal/chacha8rand.block














## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http/httptrace](#nethttphttptrace)





## net/http/cgi


This package cannot be imported because the following dependencies cannot be compiled:

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

  * [net/http](#nethttp)





## net/http/httptrace



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-340311014/main.go:2:8: package net/http/httptrace is not in std (/home/ron/.cache/tinygo/goroot-59c0c8df7700d2980b426f170e6754c2d2901b0680adb52840eeef4402e8c177/src/net/http/httptrace)






## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)
  * [net/http/httptrace](#nethttphttptrace)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)







## net/netip



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestInlining (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.23/bin/go tool -n compile: fork/exec /home/ron/.gvm/gos/go1.23/bin/go: operation not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        skipping test: 'go build' unavailable: /home/ron/.gvm/gos/go1.23/bin/go tool -n compile: fork/exec /home/ron/.gvm/gos/go1.23/bin/go: operation not implemented
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


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/rpc](#netrpc)





## net/smtp



The compiler gave the following error when this package was imported:


    # net/smtp
    /home/ron/.gvm/gos/go1.23/src/net/smtp/smtp.go:72:24: undefined: tls.Conn
    /home/ron/.gvm/gos/go1.23/src/net/smtp/smtp.go:172:25: undefined: tls.Conn








## net/url



The compiler gave the following error when running the tests for this package:


    panic: reflect: unimplemented: AssignableTo with interface
    FAIL	net/url	0.001s








## os/exec



The compiler gave the following error when running the tests for this package:


    FAIL		0.000s
    /home/ron/.gvm/gos/go1.23/src/os/exec/exec_test.go:23:2: package net/http/httptest is not in std (/home/ron/.cache/tinygo/goroot-59c0c8df7700d2980b426f170e6754c2d2901b0680adb52840eeef4402e8c177/src/net/http/httptest)






## os/signal



The compiler gave the following error when running the tests for this package:


    FAIL	os/signal	0.000s
    # internal/testpty
    /home/ron/.gvm/gos/go1.23/src/internal/testpty/pty_cgo.go:11:10: note: in file included from pty_cgo.go!cgo.c:4:
    ../../tinygo/tinygo-122/lib/musl/include/fcntl.h:22:10: fatal: 'bits/fcntl.h' file not found










## path/filepath



The compiler gave the following error when running the tests for this package:


    FAIL	path/filepath	0.002s






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    /home/ron/.gvm/gos/go1.23/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line








## regexp



The compiler gave the following error when running the tests for this package:


    FAIL	regexp	0.860s








## slices



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/slices/slices.go]
    panic: runtime error at 0x000000000027fdad: slice out of range
    FAIL	slices	0.005s






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (1.17s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
        Stable      100 elements:         934 Swap,        840 Less
        Stable      300 elements:        3965 Swap,       3007 Less
        Stable     1000 elements:       19722 Swap,      12157 Less
        Stable     3000 elements:       82381 Swap,      43169 Less
        Stable    10000 elements:      347890 Swap,     165488 Less
        Stable    30000 elements:     1340159 Swap,     559016 Less
        Stable   100000 elements:     5788412 Swap,    2086988 Less
        Stable   300000 elements:    20452076 Swap,    6864623 Less
        Stable  1000000 elements:    84469132 Swap,   25114223 Less
    --- FAIL: TestCountSortOps (0.76s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## strings



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/src/internal/bytealg/bytealg.go:304:13]
    panic: runtime error at 0x000000000022a7f3: out of memory
    FAIL	strings	8.484s










## sync/atomic



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/gos/go1.23/src/sync/atomic/doc.go:94:6]
    panic: runtime error at 0x000000000025039d: caught signal SIGSEGV
    FAIL	sync/atomic	0.906s






## syscall



The compiler gave the following error when running the tests for this package:


    FAIL	syscall	0.000s
    # syscall_test
    /home/ron/.gvm/gos/go1.23/src/syscall/creds_test.go:53:19: undefined: net.FileConn
    /home/ron/.gvm/gos/go1.23/src/syscall/creds_test.go:60:19: undefined: net.FileConn
    /home/ron/.gvm/gos/go1.23/src/syscall/exec_linux_test.go:712:14: f.Chmod undefined (type *os.File has no field or method Chmod)
    /home/ron/.gvm/gos/go1.23/src/syscall/syscall_unix_test.go:183:16: undefined: net.FileConn
    /home/ron/.gvm/gos/go1.23/src/syscall/syscall_unix_test.go:238:14: undefined: net.UnixConn
    /home/ron/.gvm/gos/go1.23/src/syscall/syscall_unix_test.go:242:18: undefined: net.FileConn








## testing/fstest



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x0000000000222947: nil pointer dereference
    FAIL	testing/fstest	0.007s








## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	0.001s






## testing/slogtest



The compiler gave the following error when running the tests for this package:


    panic: runtime.Caller failed
    FAIL	testing/slogtest	0.001s










## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	0.001s








## time



The compiler gave the following error when running the tests for this package:


    FAIL	time	0.000s
    # time_test
    /home/ron/.gvm/gos/go1.23/src/time/sleep_test.go:866:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.23/src/time/sleep_test.go:867:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.23/src/time/sleep_test.go:868:4: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.23/src/time/sleep_test.go:939:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.23/src/time/sleep_test.go:940:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)
    /home/ron/.gvm/gos/go1.23/src/time/sleep_test.go:941:8: b.ReportMetric undefined (type *testing.B has no field or method ReportMetric)






## time/tzdata



The compiler gave the following error when this package was imported:


    /home/ron/Development/tinygo/tinygo-122/src/runtime/scheduler_any.go:24: linker could not find symbol time.registerLoadFromEmbeddedTZData















