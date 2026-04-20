
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
crypto/aes |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoaes)  | 
crypto/cipher |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptocipher)  | 
crypto/des |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/dsa |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ecdh |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoecdh)  | 
crypto/ecdsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoecdsa)  | 
crypto/ed25519 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/fips140 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptofips140)  | 
crypto/hkdf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptohmac)  | 
crypto/hpke |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptohpke)  | 
crypto/md5 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/mlkem |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/mlkem/mlkemtest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
debug/gosym |  [<span style="color: red">✗</span> no](#debuggosym)  |  <span style="color: gray">✗</span> no  | 
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
errors |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
go/token |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
net/url |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
testing/cryptotest |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#testingcryptotest)  | 
testing/fstest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
    panic: runtime error at 0x0000000000280e67: nil pointer dereference
    FAIL	archive/tar	0.480s








## bufio



The compiler gave the following error when running the tests for this package:


    FAIL	bufio	0.000s
    # bufio_test
    /home/ron/.gvm/go/src/bufio/net_test.go:31:21: undefined: net.ListenUnix
    /home/ron/.gvm/go/src/bufio/net_test.go:70:25: undefined: net.DialUnix






## bytes



The compiler gave the following error when running the tests for this package:


    GC Warning: Failed to expand heap by 9007199254741120 KiB
    GC Warning: Failed to expand heap by 9007199254740992 KiB
    GC Warning: Out of Memory! Heap size: 37 MiB. Returning NULL!
    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/lib/musl/src/string/memset.c:14:7]
    panic: runtime error at 0x0000000000235be9: caught signal SIGSEGV
    FAIL	bytes	6.595s












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
    FAIL	compress/gzip	0.603s


















## crypto



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x00000000002672ec: nil pointer dereference
    FAIL	crypto	0.017s






## crypto/aes



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/aes	0.003s






## crypto/cipher



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/cipher	0.004s










## crypto/ecdh



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestLinker (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        [ build -o hello.exe hello.go]: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        [./hello.exe]: directory setting not implemented
        FailNow is incomplete, requires runtime.Goexit()
        unexpected output: 
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        [ tool nm hello.exe]: exec: no command
    [...more lines following...]






## crypto/ecdsa



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/crypto/ecdsa/ecdsa_test.go:563:6]
    panic: runtime error at 0x00000000002c272b: index out of range
    FAIL	crypto/ecdsa	0.916s






## crypto/ed25519



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestEd25519Vectors (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
        FailNow is incomplete, requires runtime.Goexit()
         env GOMODCACHE: exec: no command
        FailNow is incomplete, requires runtime.Goexit()
        fetching filippo.io/mostly-harmless/ed25519vectors@v0.0.0-20210322192420-30a2d7243a94
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        platform cannot run go tool
    [...more lines following...]








## crypto/fips140



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestWithoutEnforcement (0.00s)
        test requires FIPS 140 mode: FIPS 140-3 mode is incompatible with the purego build tag
        SkipNow is incomplete, requires runtime.Goexit()
        running with GODEBUG=fips140=only:
        fips140=only subprocess failed: files setting not implemented
    FAIL
    FAIL	crypto/fips140	0.001s








## crypto/hmac



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/hmac	0.012s






## crypto/hpke



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at <Go interface method>]
    panic: runtime error at 0x000000000034bbbe: nil pointer dereference
    FAIL	crypto/hpke	3.874s
















## crypto/rc4



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/crypto/rc4/rc4.go]
    panic: runtime error at 0x0000000000260059: index out of range
    FAIL	crypto/rc4	0.019s






## crypto/rsa



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	crypto/rsa	0.011s








## crypto/sha256



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/sha256	0.016s








## crypto/sha512



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/sha512	0.004s






## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	0.001s








## crypto/x509



The compiler gave the following error when running the tests for this package:


    FAIL	crypto/x509	0.000s
    # crypto/x509
    /home/ron/.gvm/go/src/crypto/x509/verify_test.go:1549:31: undefined: rand.Text








## database/sql



The compiler gave the following error when running the tests for this package:


    FAIL	database/sql	0.000s
    # database/sql
    /home/ron/.gvm/go/src/database/sql/sql_test.go:4839:4: b.SetParallelism undefined (type *testing.B has no field or method SetParallelism)








## debug/buildinfo



The compiler gave the following error when this package was imported:


    # debug/buildinfo
    /home/ron/.gvm/go/src/debug/buildinfo/buildinfo.go:93:19: undefined: debug.ParseBuildInfo










## debug/gosym



The compiler gave the following error when this package was imported:


    # debug/gosym
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:221:17: undefined: abi.PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:222:17: undefined: abi.PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:225:22: undefined: abi.Go12PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:227:22: undefined: abi.Go12PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:229:22: undefined: abi.Go116PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:231:22: undefined: abi.Go116PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:233:22: undefined: abi.Go118PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:235:22: undefined: abi.Go118PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:237:22: undefined: abi.Go120PCLnTabMagic
    /home/ron/.gvm/go/src/debug/gosym/pclntab.go:239:22: undefined: abi.Go120PCLnTabMagic








## debug/pe



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/debug/pe/file_test.go]
    panic: runtime error at 0x000000000028b6f2: index out of range
    FAIL	debug/pe	0.005s




















## encoding/binary



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/encoding/binary/binary.go:116:7]
    panic: runtime error at 0x000000000025c0b5: index out of range
    FAIL	encoding/binary	0.006s








## encoding/gob



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/src/internal/reflectlite/type.go:389:16]
    panic: runtime error at 0x000000000024c3f9: caught signal SIGSEGV
    FAIL	encoding/gob	0.007s








## encoding/json



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/json	0.000s
    /home/ron/.gvm/go/src/testing/synctest/synctest.go:290: linker could not find symbol internal/synctest.Run
    /home/ron/.gvm/go/src/testing/synctest/synctest.go:291: linker could not find symbol testing/synctest.testingSynctestTest






## encoding/pem



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	encoding/pem	0.386s












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
    FAIL	fmt	0.001s








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
        go/build: go list cmd/internal/objfile: files setting not implemented
            
            
        FailNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## go/constant



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/go/constant/value.go]
    panic: runtime error at 0x000000000026b5b6: divide by zero
    FAIL	go/constant	0.006s






## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	0.001s






## go/doc/comment



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestStd (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
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
    panic: runtime error at 0x00000000003420ef: nil pointer dereference
    FAIL	go/importer	0.007s






## go/parser



The compiler gave the following error when running the tests for this package:


    FAIL	go/parser	0.068s












## go/types



The compiler gave the following error when running the tests for this package:


    [...no test output...]












## hash/crc32



The compiler gave the following error when running the tests for this package:


    panic: not available
    FAIL	hash/crc32	0.019s










## hash/maphash



The compiler gave the following error when running the tests for this package:


    FAIL	hash/maphash	0.000s
    # internal/runtime/maps
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:289:39: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:296:40: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:318:25: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:327:42: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:337:26: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:341:32: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:353:32: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:264:21: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:271:17: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:293:32: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:403:28: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:407:35: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:426:38: undefined: abi.MapType
    [...more lines following...]








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	0.001s








## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	0.001s








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	0.038s








## image/jpeg



The compiler gave the following error when running the tests for this package:


    FAIL	image/jpeg	0.000s
    /home/ron/.gvm/go/src/image/jpeg/reader_test.go:253: linker could not find symbol runtime/debug.SetTraceback










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
    FAIL	io	5.029s










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
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/04/20 18:19:10.490380 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/04/20 18:19:10.490497 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/04/20 18:19:10.490527 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/04/20 18:19:10.490551 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/04/20 18:19:10.490666 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/04/20 18:19:10.490693 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/04/20 18:19:10.490802 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/04/20 18:19:10.490830 ???:0: XXXhello 23 world"
    [...more lines following...]






## log/slog



The compiler gave the following error when running the tests for this package:


    panic: runtime.Caller failed
    FAIL	log/slog	0.001s






## log/syslog



The compiler gave the following error when running the tests for this package:


    FAIL	log/syslog	0.000s
    # log/syslog
    /home/ron/.gvm/go/src/log/syslog/syslog_test.go:35:21: oe.Temporary undefined (type *net.OpError has no field or method Temporary)
    /home/ron/.gvm/go/src/log/syslog/syslog_test.go:110:15: undefined: net.ListenPacket










## math/big



The compiler gave the following error when running the tests for this package:


    [...no test output...]






## math/bits



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/math/bits/bits.go:510:21]
    panic: runtime error at 0x000000000022b692: divide by zero
    FAIL	math/bits	0.003s








## math/rand



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand	1.939s






## math/rand/v2



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand/v2	3.065s
















## net/http/cgi



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-3366960319/main.go:2:8: package net/http/cgi is not in std (/home/ron/.cache/tinygo/goroot-3f907f422909f8c14177c762d0377d28008de1630656826539080e09fca7aa2c/src/net/http/cgi)






## net/http/cookiejar



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-1383193711/main.go:2:8: package net/http/cookiejar is not in std (/home/ron/.cache/tinygo/goroot-3f907f422909f8c14177c762d0377d28008de1630656826539080e09fca7aa2c/src/net/http/cookiejar)






## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http/cgi](#nethttpcgi)









## net/http/httputil



The compiler gave the following error when this package was imported:


    # net/http/httputil
    /home/ron/Development/tinygo/tinygo-122/src/net/http/httputil/dump.go:119:3: unknown field Dial in struct literal of type http.Transport
    /home/ron/Development/tinygo/tinygo-122/src/net/http/httputil/dump.go:123:10: t.CloseIdleConnections undefined (type *http.Transport has no field or method CloseIdleConnections)
    /home/ron/Development/tinygo/tinygo-122/src/net/http/httputil/reverseproxy.go:620:8: undefined: http.NewResponseController
    /home/ron/Development/tinygo/tinygo-122/src/net/http/httputil/reverseproxy.go:701:18: undefined: http.NewResponseController
    /home/ron/Development/tinygo/tinygo-122/src/net/http/httputil/reverseproxy.go:837:13: undefined: http.NewResponseController






## net/http/pprof



The compiler gave the following error when this package was imported:


    # net/http/pprof
    /home/ron/Development/tinygo/tinygo-122/src/net/http/pprof/pprof.go:129:14: undefined: http.NewResponseController








## net/netip



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestInlining (0.00s)
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
        SkipNow is incomplete, requires runtime.Goexit()
        skipping test: 'go build' unavailable: /home/ron/.gvm/go/bin/go tool -n compile: files setting not implemented
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


    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/src/runtime/interface.go:89:15]
    panic: runtime error at 0x000000000026ccd5: type assert failed
    FAIL	net/rpc	0.004s






## net/rpc/jsonrpc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	net/rpc/jsonrpc	0.001s






## net/smtp



The compiler gave the following error when this package was imported:


    # net/smtp
    /home/ron/.gvm/go/src/net/smtp/smtp.go:72:24: undefined: tls.Conn
    /home/ron/.gvm/go/src/net/smtp/smtp.go:172:25: undefined: tls.Conn












## os/exec



The compiler gave the following error when running the tests for this package:


    FAIL	os/exec	0.000s
    # os/exec_test
    /home/ron/.gvm/go/src/os/exec/exec_test.go:260:17: undefined: net.FileListener
    /home/ron/.gvm/go/src/os/exec/exec_test.go:405:92: cmd.ProcessState.Pid undefined (type *os.ProcessState has no field or method Pid)
    /home/ron/.gvm/go/src/os/exec/exec_test.go:704:34: ln.(*net.TCPListener).File undefined (type *net.TCPListener has no field or method File)
    /home/ron/.gvm/go/src/os/exec/exec_test.go:709:18: undefined: net.FileListener
    /home/ron/.gvm/go/src/os/exec/exec_test.go:720:5: ts.StartTLS undefined (type *httptest.Server has no field or method StartTLS)
    /home/ron/.gvm/go/src/os/exec/exec_test.go:809:35: ln.(*net.TCPListener).File undefined (type *net.TCPListener has no field or method File)






## os/signal



The compiler gave the following error when running the tests for this package:


    FAIL	os/signal	0.000s
    /home/ron/.gvm/go/src/os/signal/signal_unix.go:61: linker could not find symbol os/signal.signal_ignored
    /home/ron/Development/tinygo/tinygo-122/src/syscall/syscall_linux.go:1127: linker could not find symbol syscall.runtime_doAllThreadsSyscall










## path/filepath



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestWindowsGlob (0.00s)
        skipping windows specific test
        SkipNow is incomplete, requires runtime.Goexit()
        tmpDir path "/tmp/TestWindowsGlob251716594/000" must have drive letter in it
        FailNow is incomplete, requires runtime.Goexit()
        Glob("/tmp/TestWindowsGlob251716594/000\\a") returns [], but ["/tmp/TestWindowsGlob251716594/000\\a"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\a") returns [], but ["/tmp/TestWindowsGlob251716594/000\\a"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\a") returns [], but ["/tmp/TestWindowsGlob251716594/000\\a"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\b") returns [], but ["/tmp/TestWindowsGlob251716594/000\\b"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\b") returns [], but ["/tmp/TestWindowsGlob251716594/000\\b"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\b") returns [], but ["/tmp/TestWindowsGlob251716594/000\\b"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\*") returns [], but ["/tmp/TestWindowsGlob251716594/000\\a" "/tmp/TestWindowsGlob251716594/000\\b" "/tmp/TestWindowsGlob251716594/000\\dir"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\*") returns [], but ["/tmp/TestWindowsGlob251716594/000\\a" "/tmp/TestWindowsGlob251716594/000\\b" "/tmp/TestWindowsGlob251716594/000\\dir"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\*") returns [], but ["/tmp/TestWindowsGlob251716594/000\\a" "/tmp/TestWindowsGlob251716594/000\\b" "/tmp/TestWindowsGlob251716594/000\\dir"] expected
        Glob("/tmp/TestWindowsGlob251716594/000\\d*") returns [], but ["/tmp/TestWindowsGlob251716594/000\\dir"] expected
    [...more lines following...]






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    /home/ron/.gvm/go/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line












## slices



The compiler gave the following error when running the tests for this package:


    [...no test output...]






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountStableOps (1.06s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
        Stable      100 elements:         930 Swap,        830 Less
        Stable      300 elements:        3916 Swap,       3005 Less
        Stable     1000 elements:       19565 Swap,      12317 Less
        Stable     3000 elements:       82318 Swap,      43024 Less
        Stable    10000 elements:      348714 Swap,     165881 Less
        Stable    30000 elements:     1337530 Swap,     558098 Less
        Stable   100000 elements:     5791059 Swap,    2085637 Less
        Stable   300000 elements:    20465506 Swap,    6870726 Less
        Stable  1000000 elements:    84482753 Swap,   25125401 Less
    --- FAIL: TestCountSortOps (0.61s)
        Counting skipped as non-verbose mode.
        SkipNow is incomplete, requires runtime.Goexit()
    [...more lines following...]








## strings



The compiler gave the following error when running the tests for this package:


    GC Warning: Failed to expand heap by 9007199254741120 KiB
    GC Warning: Failed to expand heap by 9007199254740992 KiB
    GC Warning: Out of Memory! Heap size: 11 MiB. Returning NULL!
    [tinygo: panic at /home/ron/Development/tinygo/tinygo-122/lib/musl/src/string/memset.c:14:7]
    panic: runtime error at 0x00000000002357d9: caught signal SIGSEGV
    FAIL	strings	7.551s










## sync/atomic



The compiler gave the following error when running the tests for this package:


    [tinygo: panic at /home/ron/.gvm/go/src/sync/atomic/doc.go:93:6]
    panic: runtime error at 0x000000000025fa33: caught signal SIGSEGV
    FAIL	sync/atomic	2.845s






## syscall



The compiler gave the following error when running the tests for this package:


    FAIL	syscall	0.000s
    # syscall_test
    /home/ron/.gvm/go/src/syscall/creds_test.go:53:19: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/creds_test.go:60:19: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/syscall_unix_test.go:183:16: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/syscall_unix_test.go:238:14: undefined: net.UnixConn
    /home/ron/.gvm/go/src/syscall/syscall_unix_test.go:242:18: undefined: net.FileConn








## testing/cryptotest



The compiler gave the following error when running the tests for this package:


    FAIL	testing/cryptotest	0.000s
    /home/ron/.gvm/go/src/testing/cryptotest/rand.go:52: linker could not find symbol testing.checkParallel










## testing/quick



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	testing/quick	0.001s






## testing/slogtest



The compiler gave the following error when running the tests for this package:


    panic: runtime.Caller failed
    FAIL	testing/slogtest	0.001s






## testing/synctest



The compiler gave the following error when running the tests for this package:


    FAIL	testing/synctest	0.000s
    # testing/synctest_test
    /home/ron/.gvm/go/src/testing/synctest/example_test.go:109:4: unknown field DialContext in struct literal of type http.Transport
    /home/ron/.gvm/go/src/testing/synctest/example_test.go:113:4: unknown field ExpectContinueTimeout in struct literal of type http.Transport










## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	0.001s








## time



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	time	0.008s






## time/tzdata



The compiler gave the following error when this package was imported:


    /home/ron/Development/tinygo/tinygo-122/src/runtime/scheduler_threads.go:27: linker could not find symbol time.registerLoadFromEmbeddedTZData
















## weak



The compiler gave the following error when running the tests for this package:


    FAIL	weak	0.000s
    /home/ron/.gvm/go/src/weak/pointer.go: linker could not find symbol weak.runtime_makeStrongFromWeak
    /home/ron/.gvm/go/src/weak/pointer.go: linker could not find symbol weak.runtime_makeStrongFromWeak





