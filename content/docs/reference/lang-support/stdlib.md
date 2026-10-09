
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Test results are for linux/amd64.

Some tests are skipped because they use functionality that TinyGo does not implement. Some packages are also tested with a larger stack than the default. This is the same procedure that the TinyGo CI uses. These are the skipped tests.

  * `TestExtraMethods`
  * `TestParseAndBytesRoundTrip/P256/Generic`
  * `TestAsValidation`
  * `TestUnmarshalNestingLimitSlice`
  * `TestUnmarshalNestingLimitStruct`

Package | Importable | Passes tests
--- | --- | --- |
archive/tar |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
archive/zip |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
bufio |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bufio)  | 
bytes |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#bytes)  | 
cmp |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/bzip2 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/flate |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/gzip |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/lzw |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
compress/zlib |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
container/heap |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
container/list |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
container/ring |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
context |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#context)  | 
crypto |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/aes |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/cipher |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptocipher)  | 
crypto/des |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/dsa |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ecdh |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ecdsa |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/ed25519 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptoed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/fips140 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hkdf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/hpke |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/mldsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptomldsa)  | 
crypto/mlkem |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/mlkem/mlkemtest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/pbkdf2 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptorsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosha256)  | 
crypto/sha3 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosha3)  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosha512)  | 
crypto/subtle |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptosubtle)  | 
crypto/tls |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
crypto/x509 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#cryptox509)  | 
crypto/x509/pkix |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#databasesql)  | 
database/sql/driver |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/buildinfo |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/dwarf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/elf |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/gosym |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/macho |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
debug/pe |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
encoding/json/jsontext |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingjsonjsontext)  | 
encoding/json/v2 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingjsonv2)  | 
encoding/pem |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#encodingpem)  | 
encoding/xml |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
errors |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
expvar |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
flag |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#flag)  | 
fmt |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#fmt)  | 
go/ast |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/build |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gobuild)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#gobuildconstraint)  | 
go/constant |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/doc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#godoc)  | 
go/doc/comment |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/format |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/importer |  [<span style="color: red">✗</span> no](#goimporter)  |  <span style="color: gray">✗</span> no  | 
go/parser |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goparser)  | 
go/printer |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#goprinter)  | 
go/scanner |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/token |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
go/types |  [<span style="color: red">✗</span> no](#gotypes)  |  <span style="color: gray">✗</span> no  | 
go/version |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
image/jpeg |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
image/png |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
index/suffixarray |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#io)  | 
io/fs |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
io/ioutil |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
iter |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
mime |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#mime)  | 
mime/multipart |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/cgi |  [<span style="color: red">✗</span> no](#nethttpcgi)  |  <span style="color: gray">✗</span> no  | 
net/http/cookiejar |  [<span style="color: red">✗</span> no](#nethttpcookiejar)  |  <span style="color: gray">✗</span> no  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#nethttpfcgi)  |  <span style="color: gray">✗</span> no  | 
net/http/httptest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/httptrace |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/httputil |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#nethttppprof)  |  <span style="color: gray">✗</span> no  | 
net/mail |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
net/netip |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netnetip)  | 
net/rpc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netrpc)  | 
net/rpc/jsonrpc |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netrpcjsonrpc)  | 
net/smtp |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#netsmtp)  | 
net/textproto |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#nettextproto)  | 
net/url |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
os |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
os/exec |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#osexec)  | 
os/signal |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#ossignal)  | 
os/user |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
path/filepath |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#pathfilepath)  | 
plugin |  [<span style="color: red">✗</span> no](#plugin)  |  <span style="color: gray">✗</span> no  | 
reflect |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#reflect)  | 
regexp |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#regexp)  | 
regexp/syntax |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
slices |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#slices)  | 
sort |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#sort)  | 
strconv |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strconv)  | 
strings |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#strings)  | 
structs |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
sync/atomic |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
syscall |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#syscall)  | 
testing |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
testing/cryptotest |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
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
unicode/utf16 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#unicodeutf16)  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#unicodeutf8)  | 
unique |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
uuid |  <span style="color: green">✔</span> yes  |  <span style="color: green">✔</span> yes  | 
weak |  <span style="color: green">✔</span> yes  |  [<span style="color: red">✗</span> no](#weak)  | 







## bufio



The compiler gave the following error when running the tests for this package:


    FAIL	bufio	0.000s
    # bufio_test
    /home/ron/.gvm/go/src/bufio/net_test.go:31:21: undefined: net.ListenUnix
    /home/ron/.gvm/go/src/bufio/net_test.go:70:25: undefined: net.DialUnix






## bytes



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestNewBufferShallow (0.00s)
        allocations occurred while shallow copying
    --- FAIL: TestReplace (0.00s)
        Replace("☺☻☹", "", "<>", -1) allocates 2.00 objects
    FAIL
    FAIL	bytes	15.437s
























## context



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAllocs (0.00s)
        WithTimeout(bg, 5*time.Millisecond) allocs = 9.000000 want 8
    FAIL
    FAIL	context	0.029s










## crypto/cipher



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	crypto/cipher	0.014s














## crypto/ed25519



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAllocations (0.02s)
        expected zero allocations, got 46.0
    FAIL
    FAIL	crypto/ed25519	0.063s


















## crypto/mldsa



The compiler gave the following error when running the tests for this package:


    [...no test output...]
















## crypto/rsa



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	crypto/rsa	0.009s








## crypto/sha256



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAllocations (0.00s)
        allocs = 12, want 0
    --- FAIL: TestAllocatonsWithTypeAsserts (0.00s)
        allocs = 6; want = 0
    FAIL
    FAIL	crypto/sha256	0.061s






## crypto/sha3



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAllocations (0.00s)
        --- FAIL: TestAllocations/New (0.00s)
            expected zero allocations, got 5.0
        --- FAIL: TestAllocations/NewSHAKE (0.00s)
            expected zero allocations, got 3.0
        --- FAIL: TestAllocations/Sum (0.00s)
            expected zero allocations, got 4.0
        --- FAIL: TestAllocations/SumSHAKE (0.00s)
            expected zero allocations, got 3.0
    FAIL
    FAIL	crypto/sha3	9.214s






## crypto/sha512



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAllocations (0.00s)
        allocs = 14, want 0
    FAIL
    FAIL	crypto/sha512	0.067s






## crypto/subtle



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	crypto/subtle	0.002s








## crypto/x509



The compiler gave the following error when running the tests for this package:


    [...no test output...]








## database/sql



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestGrabConnAllocs (0.00s)
    --- FAIL: TestGrabConnAllocs (0.00s)
        Conn.grabConn allocated 1 objects; want 0
    FAIL
    FAIL	database/sql	0.342s


































## encoding/binary



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAppendAllocs (0.00s)
        Append allocated 2 times instead of not allocating at all
    --- FAIL: TestSizeAllocs (0.00s)
        --- FAIL: TestSizeAllocs/binary.Struct (0.00s)
            Expected no allocations, got 1
        --- FAIL: TestSizeAllocs/*binary.Struct (0.00s)
            Expected no allocations, got 1
        --- FAIL: TestSizeAllocs/[]binary.Struct (0.00s)
            Expected no allocations, got 1
        --- FAIL: TestSizeAllocs/[]binary.Struct#01 (0.00s)
            Expected no allocations, got 1
        --- FAIL: TestSizeAllocs/[1]binary.Struct (0.00s)
            Expected no allocations, got 1
    FAIL
    FAIL	encoding/binary	0.019s
    [...more lines following...]








## encoding/gob



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: reflect.ArrayOf()
    FAIL	encoding/gob	2.610s








## encoding/json



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: reflect.SliceOf()
    FAIL	encoding/json	0.014s






## encoding/json/jsontext



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/json/jsontext	0.012s






## encoding/json/v2



The compiler gave the following error when running the tests for this package:


    FAIL	encoding/json/v2	0.000s
    # encoding/json/v2_test
    too many levels of pointers for typecode: *****bool
    too many levels of pointers for typecode: ******bool
    too many levels of pointers for typecode: *******bool
    too many levels of pointers for typecode: ********bool
    too many levels of pointers for typecode: *********bool
    too many levels of pointers for typecode: **********bool
    too many levels of pointers for typecode: ***********bool
    too many levels of pointers for typecode: *****bool
    too many levels of pointers for typecode: ******bool
    too many levels of pointers for typecode: *******bool
    too many levels of pointers for typecode: ********bool
    too many levels of pointers for typecode: *********bool
    too many levels of pointers for typecode: **********bool
    [...more lines following...]






## encoding/pem



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	encoding/pem	0.465s












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
    FAIL	fmt	0.002s








## go/build



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestDotSlashImport (0.00s)
        import ".": unknown compiler "tinygo"
    --- FAIL: TestLocalDirectory (0.00s)
        import ".": unknown compiler "tinygo"
    --- FAIL: TestImportCmd (0.00s)
        import "cmd/internal/objfile": unknown compiler "tinygo"
    --- FAIL: TestImportVendor (0.00s)
        cannot find vendored c/d from testdata src/a/b directory: import "c/d": unknown compiler "tinygo"
    --- FAIL: TestIssue23594 (0.00s)
        could not import testdata: import ".": unknown compiler "tinygo"
    --- FAIL: TestCgoImportsIgnored (0.00s)
        import ".": unknown compiler "tinygo"
    --- FAIL: TestAllTags (0.00s)
        import ".": unknown compiler "tinygo"
    --- FAIL: TestAllTagsNonSourceFile (0.00s)
    [...more lines following...]






## go/build/constraint



The compiler gave the following error when running the tests for this package:


    FAIL	go/build/constraint	0.008s








## go/doc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	go/doc	0.001s










## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/types](#gotypes)





## go/parser



The compiler gave the following error when running the tests for this package:


    FAIL	go/parser	0.033s






## go/printer



The compiler gave the following error when running the tests for this package:


    FAIL	go/printer	0.034s










## go/types


This package cannot be imported because the following dependencies cannot be compiled:

  * [hash/maphash](#hashmaphash)

















## hash/maphash



The compiler gave the following error when this package was imported:


    # internal/runtime/maps
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:298:39: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:305:40: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:327:25: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:336:42: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:346:26: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:350:32: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/group.go:362:32: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:283:21: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:290:17: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:312:32: undefined: abi.MapGroupSlots
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:422:28: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:426:35: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:445:38: undefined: abi.MapType
    /home/ron/.gvm/go/src/internal/runtime/maps/map.go:465:40: undefined: abi.MapType
    [...more lines following...]








## html/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	html/template	0.002s








## image/color



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/color	0.001s








## image/draw



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumIn()
    FAIL	image/draw	0.039s














## io



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestMultiWriterSingleChainFlatten (0.00s)
        multiWriter did not flatten chained multiWriters: expected writeDepth 12, got 4
    --- FAIL: TestMultiReaderFlatten (0.00s)
        multiReader did not flatten chained multiReaders: expected readDepth 3, got 1
    --- FAIL: TestMultiReaderFreesExhaustedReaders (5.00s)
        timeout waiting for collection of buf1
    FAIL
    FAIL	io	5.037s












## log



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAll (0.00s)
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^.*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/10/10 00:02:17.700244 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/10/10 00:02:17.700355 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/10/10 00:02:17.700399 ???:0: hello 23 world"
        log output should match "^XXX[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): hello 23 world$" is "XXX2026/10/10 00:02:17.700437 ???:0: hello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/10/10 00:02:17.700478 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] .*/[A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/10/10 00:02:17.700512 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/10/10 00:02:17.700627 ???:0: XXXhello 23 world"
        log output should match "^[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]\\.[0-9][0-9][0-9][0-9][0-9][0-9] [A-Za-z0-9_\\-]+\\.go:(67|69): XXXhello 23 world$" is "2026/10/10 00:02:17.700665 ???:0: XXXhello 23 world"
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










## math/big



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	math/big	4.936s






## math/bits



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestDivPanicOverflow (0.00s)
        Div should have panicked when y<=hi
    --- FAIL: TestDiv32PanicOverflow (0.00s)
        Div32 should have panicked when y<=hi
    --- FAIL: TestDiv64PanicOverflow (0.00s)
        Div64 should have panicked when y<=hi
    --- FAIL: TestDivPanicZero (0.00s)
        Div should have panicked when y==0
    --- FAIL: TestDiv64PanicZero (0.00s)
        Div64 should have panicked when y==0
    FAIL
    FAIL	math/bits	0.004s








## math/rand



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand	2.653s






## math/rand/v2



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	math/rand/v2	3.154s






## mime



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestLookupMallocs (0.01s)
        allocs = 8; want 0
    FAIL
    FAIL	mime	0.014s














## net/http/cgi



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-2311794891/main.go:2:8: package net/http/cgi is not in std (/home/ron/.cache/tinygo/goroot-7008f93e066cbb20d8039de3101dac378bc56bd4ed0ecafa578a369b60b78723/src/net/http/cgi)






## net/http/cookiejar



The compiler gave the following error when this package was imported:


    ../../../../../tmp/tinygo-test-375347846/main.go:2:8: package net/http/cookiejar is not in std (/home/ron/.cache/tinygo/goroot-7008f93e066cbb20d8039de3101dac378bc56bd4ed0ecafa578a369b60b78723/src/net/http/cookiejar)






## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http/cgi](#nethttpcgi)











## net/http/pprof



The compiler gave the following error when this package was imported:


    # net/http/pprof
    /home/ron/Development/tinygo/tinygo-122/src/net/http/pprof/pprof.go:378:18: undefined: goexperiment.GoroutineLeakProfile








## net/netip



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestNoAllocs (0.01s)
        --- FAIL: TestNoAllocs/Addr.IsGlobalUnicast (0.00s)
            allocs = 1; want 0
        --- FAIL: TestNoAllocs/Addr.IsInterfaceLocalMulticast (0.00s)
            allocs = 1; want 0
        --- FAIL: TestNoAllocs/Addr.IsLinkLocalMulticast (0.00s)
            allocs = 1; want 0
        --- FAIL: TestNoAllocs/Addr.IsLinkLocalUnicast (0.00s)
            allocs = 1; want 0
        --- FAIL: TestNoAllocs/Addr.IsPrivate (0.00s)
            allocs = 1; want 0
    --- FAIL: TestAddrStringAllocs (0.00s)
        --- FAIL: TestAddrStringAllocs/ipv4 (0.00s)
            allocs=2, want 1
        --- FAIL: TestAddrStringAllocs/ipv6 (0.00s)
    [...more lines following...]






## net/rpc



The compiler gave the following error when running the tests for this package:


    panic: type assert failed
    FAIL	net/rpc	0.003s






## net/rpc/jsonrpc



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).Method()
    FAIL	net/rpc/jsonrpc	0.001s






## net/smtp



The compiler gave the following error when running the tests for this package:


    FAIL	net/smtp	0.000s
    # net/smtp
    /home/ron/.gvm/go/src/net/smtp/smtp_test.go:622:19: undefined: tls.X509KeyPair
    /home/ron/.gvm/go/src/net/smtp/smtp_test.go:629:17: undefined: tls.Listen
    /home/ron/.gvm/go/src/net/smtp/smtp_test.go:631:17: undefined: tls.Listen
    /home/ron/.gvm/go/src/net/smtp/smtp_test.go:1040:29: cs.HandshakeComplete undefined (type tls.ConnectionState has no field or method HandshakeComplete)
    /home/ron/.gvm/go/src/net/smtp/smtp_test.go:1080:24: undefined: tls.X509KeyPair






## net/textproto



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCommonHeaders (0.00s)
        canonicalMIMEHeaderKey allocs = 1; want 0
    FAIL
    FAIL	net/textproto	0.010s










## os/exec



The compiler gave the following error when running the tests for this package:


    FAIL		0.000s
    package os/exec_test
    	imports internal/poll: build constraints exclude all Go files in /home/ron/.cache/tinygo/goroot-7008f93e066cbb20d8039de3101dac378bc56bd4ed0ecafa578a369b60b78723/src/internal/poll






## os/signal



The compiler gave the following error when running the tests for this package:


    FAIL	os/signal	0.000s
    /home/ron/.gvm/go/src/os/signal/signal_unix.go:61: linker could not find symbol os/signal.signal_ignored
    /home/ron/Development/tinygo/tinygo-122/src/syscall/syscall_linux.go:1127: linker could not find symbol syscall.runtime_doAllThreadsSyscall










## path/filepath



The compiler gave the following error when running the tests for this package:


    FAIL	path/filepath	0.002s






## plugin



The compiler gave the following error when this package was imported:


    # plugin
    /home/ron/.gvm/go/src/plugin/plugin_dlopen.go:10:6: not implemented: build constraints in #cgo line






## reflect



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestMapIterSet (0.00s)
        wanted 0 alloc, got 44
    --- FAIL: TestDeepEqualAllocs (0.01s)
        --- FAIL: TestDeepEqualAllocs/int8 (0.00s)
            DeepEqual(99, 99) allocated 2 times
        --- FAIL: TestDeepEqualAllocs/[]int8 (0.00s)
            DeepEqual([99], [99]) allocated 2 times
        --- FAIL: TestDeepEqualAllocs/int16 (0.00s)
            DeepEqual(99, 99) allocated 2 times
        --- FAIL: TestDeepEqualAllocs/[]int16 (0.00s)
            DeepEqual([99], [99]) allocated 2 times
        --- FAIL: TestDeepEqualAllocs/int32 (0.00s)
            DeepEqual(99, 99) allocated 2 times
        --- FAIL: TestDeepEqualAllocs/[]int32 (0.00s)
            DeepEqual([99], [99]) allocated 2 times
    [...more lines following...]






## regexp



The compiler gave the following error when running the tests for this package:


    FAIL	regexp	0.438s








## slices



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestInsert (0.00s)
        too many allocations inserting 50 elements: got 59, want less than 25
    --- FAIL: TestGrow (0.00s)
        Grow should allocate once when given insufficient capacity; allocated 2 times
    --- FAIL: TestConcat (0.00s)
        Concat([[1]]) allocated 2 times; want 1
        Concat([[1] [2]]) allocated 2 times; want 1
        Concat([[1] [] [2]]) allocated 2 times; want 1
    FAIL
    FAIL	slices	0.398s






## sort



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestSearchWrappersDontAlloc (0.00s)
        expected no allocs for runSearchWrappers, got 18
    FAIL
    FAIL	sort	0.752s






## strconv



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestCountMallocs (2.62s)
        AppendInt(localBuf[:0], 123, 10): 1 allocs, want <=0
        AppendFloat(localBuf[:0], 1.23, 'g', 5, 64): 1 allocs, want <=0
    --- FAIL: TestAllocationsFromBytes (0.00s)
        --- FAIL: TestAllocationsFromBytes/Atoi (0.00s)
            got 1 allocs, want 0 allocs
        --- FAIL: TestAllocationsFromBytes/ParseBool (0.00s)
            got 1 allocs, want 0 allocs
        --- FAIL: TestAllocationsFromBytes/ParseInt (0.00s)
            got 1 allocs, want 0 allocs
        --- FAIL: TestAllocationsFromBytes/ParseUint (0.00s)
            got 1 allocs, want 0 allocs
        --- FAIL: TestAllocationsFromBytes/ParseFloat (0.00s)
            got 1 allocs, want 0 allocs
        --- FAIL: TestAllocationsFromBytes/ParseComplex (0.00s)
    [...more lines following...]






## strings



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestBuilderGrow (0.01s)
        growLen=0: got 1 allocs during Write; want 0
        growLen=100: got 3 allocs during Write; want 1
        growLen=1000: got 3 allocs during Write; want 1
        growLen=10000: got 3 allocs during Write; want 1
        growLen=100000: got 3 allocs during Write; want 1
    --- FAIL: TestBuilderAllocs (0.00s)
        Builder allocs = 2; want 1
    --- FAIL: TestBuilderGrowSizeclasses (0.00s)
        unexpected amount of allocations: 3, want: 1
    --- FAIL: TestIndexRune (0.00s)
        expected no allocations, got 1.000000
    --- FAIL: TestReplace (0.00s)
        Replace("hello", "l", "L", -1) allocates 2.00 objects
        Replace("radar", "r", "<r>", -1) allocates 2.00 objects
    [...more lines following...]












## syscall



The compiler gave the following error when running the tests for this package:


    FAIL	syscall	0.000s
    # syscall_test
    /home/ron/.gvm/go/src/syscall/creds_test.go:53:19: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/creds_test.go:60:19: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/syscall_unix_test.go:183:16: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/syscall_unix_test.go:242:18: undefined: net.FileConn
    /home/ron/.gvm/go/src/syscall/syscall_unix_test.go:272:21: uc.WriteMsgUnix undefined (type *net.UnixConn has no field or method WriteMsgUnix)














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


    panic: deadlock: all goroutines in bubble are blocked
    FAIL	testing/synctest	0.202s










## text/template



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	text/template	0.002s








## time



The compiler gave the following error when running the tests for this package:


    panic: unimplemented: (reflect.Type).NumOut()
    FAIL	time	0.036s






## time/tzdata



The compiler gave the following error when this package was imported:


    /home/ron/Development/tinygo/tinygo-122/src/runtime/runtime.go:24: linker could not find symbol time.registerLoadFromEmbeddedTZData








## unicode/utf16



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestAllocationsDecode (0.00s)
        Decode allocated 1 times
        Decode allocated 1 times
        Decode allocated 1 times
        Decode allocated 1 times
    FAIL
    FAIL	unicode/utf16	0.002s






## unicode/utf8



The compiler gave the following error when running the tests for this package:


    --- FAIL: TestRuneCountNonASCIIAllocation (0.00s)
        unexpected RuneCount allocation, got 1, want 0
    FAIL
    FAIL	unicode/utf8	0.004s












## weak



The compiler gave the following error when running the tests for this package:


    [...no test output...]





