
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
crypto/dsa |  <span style="color: green">✔</span> yes  | 
crypto/ecdsa |  [<span style="color: red">✗</span> no](#crypto-ecdsa)  | 
crypto/ed25519 |  <span style="color: green">✔</span> yes  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  | 
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
go/constant |  <span style="color: green">✔</span> yes  | 
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
math/big |  <span style="color: green">✔</span> yes  | 
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
testing/iotest |  <span style="color: green">✔</span> yes  | 
testing/quick |  [<span style="color: red">✗</span> no](#testing-quick)  | 
text/scanner |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  | 
text/template |  [<span style="color: red">✗</span> no](#text-template)  | 
text/template/parse |  <span style="color: green">✔</span> yes  | 
time |  <span style="color: green">✔</span> yes  | 
time/tzdata |  <span style="color: green">✔</span> yes  | 
unicode |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  | 



## archive/tar


The compiler gave the following error when this package was imported:

<pre># archive/tar
/home/ron/.gvm/gos/go1.15/src/archive/tar/common.go:554:32: os.FileMode(fi.h.Mode).Perm undefined (type os.FileMode has no field or method Perm)
/home/ron/.gvm/gos/go1.15/src/archive/tar/common.go:636:15: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
/home/ron/.gvm/gos/go1.15/src/archive/tar/common.go:637:21: fm.Perm undefined (type os.FileMode has no field or method Perm)
/home/ron/.gvm/gos/go1.15/src/archive/tar/common.go:640:10: fm.IsRegular undefined (type os.FileMode has no field or method IsRegular)
</pre>





## archive/zip


The compiler gave the following error when this package was imported:

<pre># archive/zip
/home/ron/.gvm/gos/go1.15/src/archive/zip/struct.go:180:19: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
</pre>





































## crypto/ecdsa


The compiler gave the following error when this package was imported:

<pre># encoding/asn1
encoding/asn1/<init>: interp: unknown GEP

traceback:
encoding/asn1/<init>:
  %12 = getelementptr inbounds { %"encoding/asn1.BitString" }, { %"encoding/asn1.BitString" }* %11, i32 0, i32 0, !dbg !1934
</pre>

















## crypto/rsa


The compiler gave the following error when this package was imported:

<pre># crypto/rsa
/home/ron/.gvm/gos/go1.15/src/math/big/nat.go:74:11: interp: unknown GEP

traceback:
/home/ron/.gvm/gos/go1.15/src/math/big/nat.go:74:11:
  %16 = getelementptr inbounds i32, i32* %4, i32 0, !dbg !1742
/home/ron/.gvm/gos/go1.15/src/math/big/nat.go:84:19:
  %9 = call { i32*, i32, i32 } @"(math/big.nat).setWord"(i32* %6, i32 %7, i32 %8, i32 %3, i8* undef, i8* undef), !dbg !1749
/home/ron/.gvm/gos/go1.15/src/math/big/int.go:55:25:
  %17 = call { i32*, i32, i32 } @"(math/big.nat).setUint64"(i32* %14, i32 %15, i32 %16, i64 %2, i8* undef, i8* undef), !dbg !1755
/home/ron/.gvm/gos/go1.15/src/math/big/int.go:69:26:
  %2 = call %"math/big.Int"* @"(*math/big.Int).SetInt64"(%"math/big.Int"* %0, i64 %x, i8* undef, i8* undef), !dbg !1739
crypto/rsa/<init>:38:25:
  %288 = call %"math/big.Int"* @"math/big.NewInt"(i64 0, i8* undef, i8* undef), !dbg !1867
</pre>













## crypto/tls


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ecdsa](#crypto-ecdsa)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509](#crypto-x509)
  * [net](#net)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ecdsa](#crypto-ecdsa)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509/pkix](#crypto-x509-pkix)
  * [encoding/asn1](#encoding-asn1)
  * [net](#net)





## crypto/x509/pkix


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/asn1](#encoding-asn1)

























## encoding/asn1


The compiler gave the following error when this package was imported:

<pre># encoding/asn1
encoding/asn1/<init>: interp: unknown GEP

traceback:
encoding/asn1/<init>:
  %12 = getelementptr inbounds { %"encoding/asn1.BitString" }, { %"encoding/asn1.BitString" }* %11, i32 0, i32 0, !dbg !1837
</pre>













## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
/home/ron/.gvm/gos/go1.15/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/home/ron/.gvm/gos/go1.15/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.15/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.15/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.15/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.15/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.15/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.15/src/encoding/gob/type.go:867:9: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
/home/ron/.gvm/gos/go1.15/src/encoding/gob/type.go:870:21: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
</pre>







## encoding/json


The compiler gave the following error when this package was imported:

<pre># encoding/json
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:232:26: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:249:30: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:515:8: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:552:7: v.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:589:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:620:40: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:641:16: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:776:17: PtrTo not declared by package reflect
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:921:9: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:950:6: v.SetBytes undefined (type reflect.Value has no field or method SetBytes)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:957:9: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:990:9: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
/home/ron/.gvm/gos/go1.15/src/encoding/json/decode.go:1014:23: v.OverflowFloat undefined (type reflect.Value has no field or method OverflowFloat)
[...more lines following...]</pre>







## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
/home/ron/.gvm/gos/go1.15/src/encoding/xml/read.go:286:8: val.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.15/src/encoding/xml/read.go:402:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
/home/ron/.gvm/gos/go1.15/src/encoding/xml/read.go:665:7: dst.SetBytes undefined (type reflect.Value has no field or method SetBytes)
/home/ron/.gvm/gos/go1.15/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.15/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
/home/ron/.gvm/gos/go1.15/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
</pre>







## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [net/http](#net-http)











## go/build


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/doc](#go-doc)
  * [os/exec](#os-exec)







## go/doc


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#text-template)







## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/build](#go-build)
  * [go/types](#go-types)













## go/types


The compiler gave the following error when this package was imported:

<pre># go/types
/home/ron/Development/tinygo/tinygo/src/runtime/hashmap.go:203:21: interp: load from a bitcast

traceback:
/home/ron/Development/tinygo/tinygo/src/runtime/hashmap.go:203:21:
  %83 = load i8, i8* %82, !dbg !2170
/home/ron/Development/tinygo/tinygo/src/runtime/hashmap.go:335:19:
  %13 = call i1 @runtime.hashmapGet(%runtime.hashmap* %m, i8* %12, i8* %value, i32 %valueSize, i32 %11, i8* undef, i32 ptrtoint (%runtime.funcValueWithSignature* @"runtime.hashmapStringEqual$withSignature" to i32), i8* undef, i8* undef), !dbg !2059
/home/ron/.gvm/gos/go1.15/src/go/types/sizes.go:184:6:
  %13 = call i1 @runtime.hashmapStringGet(%runtime.hashmap* %9, i8* %11, i32 %12, i8* %hashmap.value.bitcast, i32 4, i8* undef, i8* null), !dbg !2055
go/types/<init>:202:24:
  %359 = call %runtime._interface @"go/types.SizesFor"(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @"go/types.init$string.1271", i32 0, i32 0), i32 2, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @"go/types.init$string.1272", i32 0, i32 0), i32 5, i8* undef, i8* undef), !dbg !2199
</pre>



















## html/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [text/template](#text-template)



























## log/syslog


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)

















## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/textproto](#net-textproto)







## net


The compiler gave the following error when this package was imported:

<pre># net
/home/ron/.gvm/gos/go1.15/src/net/parse.go:80:12: st.ModTime undefined (type os.FileInfo has no field or method ModTime)
/home/ron/.gvm/gos/go1.15/src/net/pipe.go:156:16: ErrDeadlineExceeded not declared by package os
/home/ron/.gvm/gos/go1.15/src/net/pipe.go:169:16: ErrDeadlineExceeded not declared by package os
/home/ron/.gvm/gos/go1.15/src/net/pipe.go:188:16: ErrDeadlineExceeded not declared by package os
/home/ron/.gvm/gos/go1.15/src/net/pipe.go:204:17: ErrDeadlineExceeded not declared by package os
</pre>





## net/http


This package cannot be imported because the following dependencies cannot be compiled:

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
  * [net/textproto](#net-textproto)
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
  * [net/textproto](#net-textproto)





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
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:130:14: Process not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:134:19: ProcessState not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:462:6: ProcessState not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:246:23: DevNull not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:258:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:291:27: DevNull not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:303:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:422:22: StartProcess not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:422:57: ProcAttr not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:584:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:626:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.15/src/os/exec/exec.go:651:20: Pipe not declared by package os
</pre>







## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
/home/ron/.gvm/gos/go1.15/src/os/user/lookup.go:15:41: undeclared name: current
/home/ron/.gvm/gos/go1.15/src/os/user/lookup.go:36:9: undeclared name: lookupUser
/home/ron/.gvm/gos/go1.15/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
/home/ron/.gvm/gos/go1.15/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
/home/ron/.gvm/gos/go1.15/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
/home/ron/.gvm/gos/go1.15/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>



































## testing/quick


The compiler gave the following error when this package was imported:

<pre># testing/quick
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:273:11: fType.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:276:11: fType.Out undefined (type reflect.Type has no field or method Out)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:280:43: fType.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:290:12: fVal.Call undefined (type reflect.Value has no field or method Call)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:320:43: xType.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:330:26: x.Call undefined (type reflect.Value has no field or method Call)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:331:26: y.Call undefined (type reflect.Value has no field or method Call)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:351:25: f.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.15/src/testing/quick/quick.go:353:95: f.In undefined (type reflect.Type has no field or method In)
</pre>









## text/template


The compiler gave the following error when this package was imported:

<pre># text/template
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:377:20: val.Recv undefined (type reflect.Value has no field or method Recv)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:541:25: cannot convert nil (untyped nil value) to reflect.Type
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:603:19: ptr.MethodByName undefined (type reflect.Value has no field or method MethodByName)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:610:33: receiver.Type().FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:645:21: etyp.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:678:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:679:18: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:681:79: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:683:25: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:684:69: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:688:71: typ.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:695:32: typ.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:698:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
/home/ron/.gvm/gos/go1.15/src/text/template/exec.go:699:18: typ.In undefined (type reflect.Type has no field or method In)
[...more lines following...]</pre>


















