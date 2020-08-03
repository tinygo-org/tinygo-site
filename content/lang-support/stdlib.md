
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Package | Importable
--- | --- |
archive/tar |  [<span style="color: red">✗</span> no](#archive-tar)  | 
archive/zip |  [<span style="color: red">✗</span> no](#archive-zip)  | 
bufio |  <span style="color: green">✔</span> yes  | 
bytes |  <span style="color: green">✔</span> yes  | 
compress/bzip2 |  <span style="color: green">✔</span> yes  | 
compress/flate |  <span style="color: green">✔</span> yes  | 
compress/gzip |  <span style="color: green">✔</span> yes  | 
compress/lzw |  <span style="color: green">✔</span> yes  | 
compress/zlib |  <span style="color: green">✔</span> yes  | 
container/heap |  <span style="color: green">✔</span> yes  | 
container/list |  <span style="color: green">✔</span> yes  | 
container/ring |  <span style="color: green">✔</span> yes  | 
context |  <span style="color: green">✔</span> yes  | 
crypto |  <span style="color: green">✔</span> yes  | 
crypto/aes |  <span style="color: green">✔</span> yes  | 
crypto/cipher |  <span style="color: green">✔</span> yes  | 
crypto/des |  <span style="color: green">✔</span> yes  | 
crypto/dsa |  [<span style="color: red">✗</span> no](#crypto-dsa)  | 
crypto/ecdsa |  [<span style="color: red">✗</span> no](#crypto-ecdsa)  | 
crypto/ed25519 |  [<span style="color: red">✗</span> no](#crypto-ed25519)  | 
crypto/elliptic |  [<span style="color: red">✗</span> no](#crypto-elliptic)  | 
crypto/hmac |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  | 
crypto/rand |  [<span style="color: red">✗</span> no](#crypto-rand)  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  [<span style="color: red">✗</span> no](#crypto-rsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  | 
crypto/subtle |  <span style="color: green">✔</span> yes  | 
crypto/tls |  [<span style="color: red">✗</span> no](#crypto-tls)  | 
crypto/x509 |  [<span style="color: red">✗</span> no](#crypto-x509)  | 
crypto/x509/pkix |  [<span style="color: red">✗</span> no](#crypto-x509-pkix)  | 
database/sql |  <span style="color: green">✔</span> yes  | 
database/sql/driver |  <span style="color: green">✔</span> yes  | 
debug/dwarf |  <span style="color: green">✔</span> yes  | 
debug/elf |  <span style="color: green">✔</span> yes  | 
debug/gosym |  <span style="color: green">✔</span> yes  | 
debug/macho |  <span style="color: green">✔</span> yes  | 
debug/pe |  <span style="color: green">✔</span> yes  | 
debug/plan9obj |  <span style="color: green">✔</span> yes  | 
encoding |  <span style="color: green">✔</span> yes  | 
encoding/ascii85 |  <span style="color: green">✔</span> yes  | 
encoding/asn1 |  [<span style="color: red">✗</span> no](#encoding-asn1)  | 
encoding/base32 |  <span style="color: green">✔</span> yes  | 
encoding/base64 |  <span style="color: green">✔</span> yes  | 
encoding/binary |  <span style="color: green">✔</span> yes  | 
encoding/csv |  <span style="color: green">✔</span> yes  | 
encoding/gob |  [<span style="color: red">✗</span> no](#encoding-gob)  | 
encoding/hex |  <span style="color: green">✔</span> yes  | 
encoding/json |  [<span style="color: red">✗</span> no](#encoding-json)  | 
encoding/pem |  <span style="color: green">✔</span> yes  | 
encoding/xml |  [<span style="color: red">✗</span> no](#encoding-xml)  | 
errors |  <span style="color: green">✔</span> yes  | 
expvar |  [<span style="color: red">✗</span> no](#expvar)  | 
flag |  <span style="color: green">✔</span> yes  | 
fmt |  <span style="color: green">✔</span> yes  | 
go/ast |  <span style="color: green">✔</span> yes  | 
go/build |  [<span style="color: red">✗</span> no](#go-build)  | 
go/constant |  [<span style="color: red">✗</span> no](#go-constant)  | 
go/doc |  [<span style="color: red">✗</span> no](#go-doc)  | 
go/format |  <span style="color: green">✔</span> yes  | 
go/importer |  [<span style="color: red">✗</span> no](#go-importer)  | 
go/parser |  <span style="color: green">✔</span> yes  | 
go/printer |  <span style="color: green">✔</span> yes  | 
go/scanner |  <span style="color: green">✔</span> yes  | 
go/token |  <span style="color: green">✔</span> yes  | 
go/types |  [<span style="color: red">✗</span> no](#go-types)  | 
hash |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  | 
hash/crc64 |  <span style="color: green">✔</span> yes  | 
hash/fnv |  <span style="color: green">✔</span> yes  | 
hash/maphash |  <span style="color: green">✔</span> yes  | 
html |  <span style="color: green">✔</span> yes  | 
html/template |  [<span style="color: red">✗</span> no](#html-template)  | 
image |  <span style="color: green">✔</span> yes  | 
image/color |  <span style="color: green">✔</span> yes  | 
image/color/palette |  <span style="color: green">✔</span> yes  | 
image/draw |  <span style="color: green">✔</span> yes  | 
image/gif |  <span style="color: green">✔</span> yes  | 
image/jpeg |  <span style="color: green">✔</span> yes  | 
image/png |  <span style="color: green">✔</span> yes  | 
index/suffixarray |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  | 
io/ioutil |  <span style="color: green">✔</span> yes  | 
log |  <span style="color: green">✔</span> yes  | 
log/syslog |  [<span style="color: red">✗</span> no](#log-syslog)  | 
math |  <span style="color: green">✔</span> yes  | 
math/big |  [<span style="color: red">✗</span> no](#math-big)  | 
math/bits |  <span style="color: green">✔</span> yes  | 
math/cmplx |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  | 
mime |  <span style="color: green">✔</span> yes  | 
mime/multipart |  [<span style="color: red">✗</span> no](#mime-multipart)  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  | 
net |  [<span style="color: red">✗</span> no](#net)  | 
net/http |  [<span style="color: red">✗</span> no](#net-http)  | 
net/http/cgi |  [<span style="color: red">✗</span> no](#net-http-cgi)  | 
net/http/cookiejar |  [<span style="color: red">✗</span> no](#net-http-cookiejar)  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#net-http-fcgi)  | 
net/http/httptest |  [<span style="color: red">✗</span> no](#net-http-httptest)  | 
net/http/httptrace |  [<span style="color: red">✗</span> no](#net-http-httptrace)  | 
net/http/httputil |  [<span style="color: red">✗</span> no](#net-http-httputil)  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#net-http-pprof)  | 
net/mail |  [<span style="color: red">✗</span> no](#net-mail)  | 
net/rpc |  [<span style="color: red">✗</span> no](#net-rpc)  | 
net/rpc/jsonrpc |  [<span style="color: red">✗</span> no](#net-rpc-jsonrpc)  | 
net/smtp |  [<span style="color: red">✗</span> no](#net-smtp)  | 
net/textproto |  [<span style="color: red">✗</span> no](#net-textproto)  | 
net/url |  <span style="color: green">✔</span> yes  | 
os |  <span style="color: green">✔</span> yes  | 
os/exec |  [<span style="color: red">✗</span> no](#os-exec)  | 
os/signal |  <span style="color: green">✔</span> yes  | 
os/user |  [<span style="color: red">✗</span> no](#os-user)  | 
path |  <span style="color: green">✔</span> yes  | 
path/filepath |  <span style="color: green">✔</span> yes  | 
plugin |  <span style="color: green">✔</span> yes  | 
reflect |  <span style="color: green">✔</span> yes  | 
regexp |  <span style="color: green">✔</span> yes  | 
regexp/syntax |  <span style="color: green">✔</span> yes  | 
sort |  <span style="color: green">✔</span> yes  | 
strconv |  <span style="color: green">✔</span> yes  | 
strings |  <span style="color: green">✔</span> yes  | 
sync |  <span style="color: green">✔</span> yes  | 
sync/atomic |  <span style="color: green">✔</span> yes  | 
syscall |  <span style="color: green">✔</span> yes  | 
syscall/js |  <span style="color: green">✔</span> yes  | 
testing |  <span style="color: green">✔</span> yes  | 
testing/iotest |  [<span style="color: red">✗</span> no](#testing-iotest)  | 
testing/quick |  [<span style="color: red">✗</span> no](#testing-quick)  | 
text/scanner |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  | 
text/template |  [<span style="color: red">✗</span> no](#text-template)  | 
text/template/parse |  <span style="color: green">✔</span> yes  | 
time |  <span style="color: green">✔</span> yes  | 
unicode |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  | 



## archive/tar


The compiler gave the following error when this package was imported:

<pre># archive/tar
/home/ron/.gvm/gos/go1.14.2/src/archive/tar/common.go:554:32: os.FileMode(fi.h.Mode).Perm undefined (type os.FileMode has no field or method Perm)
/home/ron/.gvm/gos/go1.14.2/src/archive/tar/common.go:636:15: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
/home/ron/.gvm/gos/go1.14.2/src/archive/tar/common.go:637:21: fm.Perm undefined (type os.FileMode has no field or method Perm)
/home/ron/.gvm/gos/go1.14.2/src/archive/tar/common.go:640:10: fm.IsRegular undefined (type os.FileMode has no field or method IsRegular)
</pre>





## archive/zip


The compiler gave the following error when this package was imported:

<pre># archive/zip
/home/ron/.gvm/gos/go1.14.2/src/archive/zip/struct.go:180:19: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
</pre>



































## crypto/dsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)





## crypto/ecdsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/elliptic](#crypto-elliptic)
  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)





## crypto/ed25519


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)





## crypto/elliptic


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)









## crypto/rand


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)







## crypto/rsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [math/big](#math-big)













## crypto/tls


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ecdsa](#crypto-ecdsa)
  * [crypto/ed25519](#crypto-ed25519)
  * [crypto/elliptic](#crypto-elliptic)
  * [crypto/rand](#crypto-rand)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509](#crypto-x509)
  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)
  * [net](#net)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/dsa](#crypto-dsa)
  * [crypto/ecdsa](#crypto-ecdsa)
  * [crypto/ed25519](#crypto-ed25519)
  * [crypto/elliptic](#crypto-elliptic)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509/pkix](#crypto-x509-pkix)
  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)
  * [net](#net)





## crypto/x509/pkix


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)

























## encoding/asn1


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)













## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/type.go:867:9: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
/home/ron/.gvm/gos/go1.14.2/src/encoding/gob/type.go:870:21: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
</pre>







## encoding/json


The compiler gave the following error when this package was imported:

<pre># encoding/json
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:235:26: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:252:30: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:532:8: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:569:7: v.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:606:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:637:40: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:658:16: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:793:17: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:938:9: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:967:6: v.SetBytes undefined (type reflect.Value has no field or method SetBytes)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:974:9: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:1007:9: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.14.2/src/encoding/json/decode.go:1031:23: v.OverflowFloat undefined (type reflect.Value has no field or method OverflowFloat)
[...more lines following...]</pre>







## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
/home/ron/.gvm/gos/go1.14.2/src/encoding/xml/read.go:286:8: val.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.14.2/src/encoding/xml/read.go:402:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.14.2/src/encoding/xml/read.go:665:7: dst.SetBytes undefined (type reflect.Value has no field or method SetBytes)
/home/ron/.gvm/gos/go1.14.2/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.14.2/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
/home/ron/.gvm/gos/go1.14.2/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
</pre>







## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [net/http](#net-http)











## go/build


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/doc](#go-doc)
  * [os/exec](#os-exec)





## go/constant


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)





## go/doc


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#text-template)







## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/build](#go-build)
  * [go/types](#go-types)













## go/types


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/constant](#go-constant)



















## html/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [text/template](#text-template)



























## log/syslog


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)







## math/big


The compiler gave the following error when this package was imported:

<pre># math/big
/home/ron/.gvm/gos/go1.14.2/src/math/bits.go: interp: branch on a non-constant

traceback:
/home/ron/.gvm/gos/go1.14.2/src/math/frexp.go:26:20:
  %7 = call { double, i32 } @math.normalize(double %f, i8* undef, i8* undef), !dbg !1745
/home/ron/.gvm/pkgsets/go1.14.2/global/src/github.com/tinygo-org/tinygo/src/runtime/math.go:129:62:
  %0 = call { double, i32 } @math.frexp(double %x, i8* undef, i8* undef), !dbg !1732
/home/ron/.gvm/gos/go1.14.2/src/math/big/float.go:560:26:
  %44 = call { double, i32 } @math.Frexp(double %x, i8* undef, i8* undef), !dbg !1776
/home/ron/.gvm/gos/go1.14.2/src/math/big/float.go:92:30:
  %5 = call %"math/big.Float"* @"(*math/big.Float).SetFloat64"(%"math/big.Float"* %3, double %x, i8* undef, i8* undef), !dbg !1736
math/big/<init>:10:18:
  %51 = call %"math/big.Float"* @"math/big.NewFloat"(double 3.000000e+00, i8* undef, i8* undef), !dbg !1792
</pre>













## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [net/textproto](#net-textproto)







## net


The compiler gave the following error when this package was imported:

<pre># net
/home/ron/.gvm/gos/go1.14.2/src/net/parse.go:80:12: st.ModTime undefined (type os.FileInfo has no field or method ModTime)
</pre>





## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [crypto/tls](#crypto-tls)
  * [mime/multipart](#mime-multipart)
  * [net](#net)
  * [net/http/httptrace](#net-http-httptrace)
  * [net/textproto](#net-textproto)





## net/http/cgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net](#net)
  * [net/http](#net-http)
  * [os/exec](#os-exec)





## net/http/cookiejar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/http](#net-http)





## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/http](#net-http)
  * [net/http/cgi](#net-http-cgi)





## net/http/httptest


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [crypto/x509](#crypto-x509)
  * [net](#net)
  * [net/http](#net-http)





## net/http/httptrace


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net](#net)
  * [net/textproto](#net-textproto)





## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/http](#net-http)
  * [net/textproto](#net-textproto)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [html/template](#html-template)
  * [net/http](#net-http)





## net/mail


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/textproto](#net-textproto)





## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/gob](#encoding-gob)
  * [html/template](#html-template)
  * [net](#net)
  * [net/http](#net-http)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [net](#net)
  * [net/rpc](#net-rpc)





## net/smtp


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net](#net)
  * [net/textproto](#net-textproto)





## net/textproto


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)









## os/exec


The compiler gave the following error when this package was imported:

<pre># os/exec
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:130:14: Process not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:134:19: ProcessState not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:462:6: ProcessState not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:246:23: DevNull not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:258:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:291:27: DevNull not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:303:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:422:22: StartProcess not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:422:57: ProcAttr not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:584:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:626:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.14.2/src/os/exec/exec.go:651:20: Pipe not declared by package os
</pre>







## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
/home/ron/.gvm/gos/go1.14.2/src/os/user/lookup.go:15:41: undeclared name: current
/home/ron/.gvm/gos/go1.14.2/src/os/user/lookup.go:36:9: undeclared name: lookupUser
/home/ron/.gvm/gos/go1.14.2/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
/home/ron/.gvm/gos/go1.14.2/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
/home/ron/.gvm/gos/go1.14.2/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
/home/ron/.gvm/gos/go1.14.2/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>

































## testing/iotest


The compiler gave the following error when this package was imported:

<pre># testing/iotest
../../../../../tmp/tinygo-test.go:2:8: cannot find package "testing/iotest" in any of:
	/home/ron/.cache/tinygo/goroot-go1.14.2-5e37b0c7857e005e0b4205e9c21953b65d95fc73862a09032b80e833cddceb46/src/testing/iotest (from $GOROOT)
	/home/ron/.gvm/pkgsets/go1.14.2/global/src/testing/iotest (from $GOPATH)
</pre>





## testing/quick


The compiler gave the following error when this package was imported:

<pre># testing/quick
../../../../../tmp/tinygo-test.go:2:8: cannot find package "testing/quick" in any of:
	/home/ron/.cache/tinygo/goroot-go1.14.2-5e37b0c7857e005e0b4205e9c21953b65d95fc73862a09032b80e833cddceb46/src/testing/quick (from $GOROOT)
	/home/ron/.gvm/pkgsets/go1.14.2/global/src/testing/quick (from $GOPATH)
</pre>









## text/template


The compiler gave the following error when this package was imported:

<pre># text/template
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:377:20: val.Recv undefined (type reflect.Value has no field or method Recv)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:541:25: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:603:19: ptr.MethodByName undefined (type reflect.Value has no field or method MethodByName)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:610:33: receiver.Type().FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:645:21: etyp.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:678:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:679:18: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:681:79: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:683:25: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:684:69: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:688:71: typ.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:695:32: typ.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:698:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
/home/ron/.gvm/gos/go1.14.2/src/text/template/exec.go:699:18: typ.In undefined (type reflect.Type has no field or method In)
[...more lines following...]</pre>
















