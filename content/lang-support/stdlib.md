
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
database/sql |  [<span style="color: red">✗</span> no](#database-sql)  | 
database/sql/driver |  <span style="color: green">✔</span> yes  | 
debug/dwarf |  <span style="color: green">✔</span> yes  | 
debug/elf |  [<span style="color: red">✗</span> no](#debug-elf)  | 
debug/gosym |  <span style="color: green">✔</span> yes  | 
debug/macho |  [<span style="color: red">✗</span> no](#debug-macho)  | 
debug/pe |  [<span style="color: red">✗</span> no](#debug-pe)  | 
debug/plan9obj |  [<span style="color: red">✗</span> no](#debug-plan9obj)  | 
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
flag |  [<span style="color: red">✗</span> no](#flag)  | 
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
html |  <span style="color: green">✔</span> yes | 
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
mime |  [<span style="color: red">✗</span> no](#mime)  | 
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
os/signal |  [<span style="color: red">✗</span> no](#os-signal)  | 
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
testing |  [<span style="color: red">✗</span> no](#testing)  | 
testing/iotest |  <span style="color: green">✔</span> yes  | 
testing/quick |  [<span style="color: red">✗</span> no](#testing-quick)  | 
text/scanner |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  | 
text/template |  [<span style="color: red">✗</span> no](#text-template)  | 
text/template/parse |  [<span style="color: red">✗</span> no](#text-template-parse)  | 
time |  <span style="color: green">✔</span> yes  | 
unicode |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/crypto/chacha20poly1305 |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/crypto/cryptobyte |  [<span style="color: red">✗</span> no](#vendor-golang.org-x-crypto-cryptobyte)  | 
vendor/golang.org/x/crypto/cryptobyte/asn1 |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/crypto/curve25519 |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/crypto/hkdf |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/crypto/poly1305 |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/net/dns/dnsmessage |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/net/http/httpguts |  [<span style="color: red">✗</span> no](#vendor-golang.org-x-net-http-httpguts)  | 
vendor/golang.org/x/net/http/httpproxy |  [<span style="color: red">✗</span> no](#vendor-golang.org-x-net-http-httpproxy)  | 
vendor/golang.org/x/net/http2/hpack |  [<span style="color: red">✗</span> no](#vendor-golang.org-x-net-http2-hpack)  | 
vendor/golang.org/x/net/idna |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/net/nettest |  [<span style="color: red">✗</span> no](#vendor-golang.org-x-net-nettest)  | 
vendor/golang.org/x/sys/cpu |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/text/secure/bidirule |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/text/transform |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/text/unicode/bidi |  <span style="color: green">✔</span> yes  | 
vendor/golang.org/x/text/unicode/norm |  <span style="color: green">✔</span> yes  | 



## archive/tar


The compiler gave the following error when this package was imported:

<pre># archive/tar
../../.gvm/gos/go1.14/src/archive/tar/common.go:445:15: DeepEqual not declared by package reflect
../../.gvm/gos/go1.14/src/archive/tar/common.go:554:32: os.FileMode(fi.h.Mode).Perm undefined (type os.FileMode has no field or method Perm)
../../.gvm/gos/go1.14/src/archive/tar/common.go:636:15: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
../../.gvm/gos/go1.14/src/archive/tar/common.go:637:21: fm.Perm undefined (type os.FileMode has no field or method Perm)
../../.gvm/gos/go1.14/src/archive/tar/common.go:640:10: fm.IsRegular undefined (type os.FileMode has no field or method IsRegular)</pre>





## archive/zip


The compiler gave the following error when this package was imported:

<pre># archive/zip
../../.gvm/gos/go1.14/src/archive/zip/reader.go:60:19: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to r.init: missing method ReadAt
../../.gvm/gos/go1.14/src/archive/zip/struct.go:180:19: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
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
  * [vendor/golang.org/x/crypto/cryptobyte](#vendor-golang.org-x-crypto-cryptobyte)
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
  * [vendor/golang.org/x/crypto/cryptobyte](#vendor-golang.org-x-crypto-cryptobyte)
  * [math/big](#math-big)
  * [net](#net)





## crypto/x509/pkix


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)





## database/sql


The compiler gave the following error when this package was imported:

<pre># database/sql
../../.gvm/gos/go1.14/src/database/sql/sql.go:590:7: Locker not declared by package sync
../../.gvm/gos/go1.14/src/database/sql/sql.go:3159:7: Locker not declared by package sync
../../.gvm/gos/go1.14/src/database/sql/sql.go:3181:23: Locker not declared by package sync
../../.gvm/gos/go1.14/src/database/sql/convert.go:411:41: sv.Type().ConvertibleTo undefined (type reflect.Type has no field or method ConvertibleTo)
../../.gvm/gos/go1.14/src/database/sql/convert.go:412:13: sv.Convert undefined (type reflect.Value has no field or method Convert)
../../.gvm/gos/go1.14/src/database/sql/sql.go:599:5: ds.Lock undefined (type *driverStmt has no field or method Lock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:600:11: ds.Unlock undefined (type *driverStmt has no field or method Unlock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:2427:5: ds.Lock undefined (type *driverStmt has no field or method Lock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:2428:11: ds.Unlock undefined (type *driverStmt has no field or method Unlock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:2598:5: ds.Lock undefined (type *driverStmt has no field or method Lock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:2599:11: ds.Unlock undefined (type *driverStmt has no field or method Unlock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:2744:22: rs.closemu.RLocker undefined (type sync.RWMutex has no field or method RLocker)
../../.gvm/gos/go1.14/src/database/sql/sql.go:3164:5: dr.Lock undefined (type driverResult has no field or method Lock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:3165:11: dr.Unlock undefined (type driverResult has no field or method Unlock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:3170:5: dr.Lock undefined (type driverResult has no field or method Lock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:3171:11: dr.Unlock undefined (type driverResult has no field or method Unlock)
../../.gvm/gos/go1.14/src/database/sql/sql.go:3177:29: Stack not declared by package runtime</pre>









## debug/elf


The compiler gave the following error when this package was imported:

<pre># debug/elf
../../.gvm/gos/go1.14/src/debug/elf/file.go:206:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
../../.gvm/gos/go1.14/src/debug/elf/reader.go:73:16: ErrInvalid not declared by package os
../../.gvm/gos/go1.14/src/debug/elf/reader.go:81:16: ErrInvalid not declared by package os
</pre>







## debug/macho


The compiler gave the following error when this package was imported:

<pre># debug/macho
../../.gvm/gos/go1.14/src/debug/macho/fat.go:130:24: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFatFile: missing method ReadAt
../../.gvm/gos/go1.14/src/debug/macho/file.go:205:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
</pre>





## debug/pe


The compiler gave the following error when this package was imported:

<pre># debug/pe
../../.gvm/gos/go1.14/src/debug/pe/file.go:40:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
</pre>





## debug/plan9obj


The compiler gave the following error when this package was imported:

<pre># debug/plan9obj
../../.gvm/gos/go1.14/src/debug/plan9obj/file.go:103:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
</pre>









## encoding/asn1


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)













## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:466:18: value.FieldByIndex undefined (type reflect.Value has no field or method FieldByIndex)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:564:19: mtyp.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:568:27: mtyp.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:569:28: mtyp.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:575:9: value.SetMapIndex undefined (type reflect.Value has no field or method SetMapIndex)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:825:35: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:1019:31: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
../../.gvm/gos/go1.14/src/encoding/gob/decode.go:1202:17: base.Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/gob/encode.go:320:18: value.FieldByIndex undefined (type reflect.Value has no field or method FieldByIndex)
../../.gvm/gos/go1.14/src/encoding/gob/encode.go:562:34: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
../../.gvm/gos/go1.14/src/encoding/gob/encoder.go:128:29: st.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
../../.gvm/gos/go1.14/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/gob/type.go:496:34: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/gob/type.go:514:37: t.Elem().Name undefined (type reflect.Type has no field or method Name)
[...more lines following...]</pre>







## encoding/json


The compiler gave the following error when this package was imported:

<pre># encoding/json
../../.gvm/gos/go1.14/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
../../.gvm/gos/go1.14/src/encoding/json/decode.go:235:26: cannot convert nil (untyped nil value) to reflect.Type
../../.gvm/gos/go1.14/src/encoding/json/decode.go:252:30: cannot convert nil (untyped nil value) to reflect.Type
../../.gvm/gos/go1.14/src/encoding/json/decode.go:255:39: d.errorContext.Struct.Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:457:41: v.Type().Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:491:15: v.Type().NumMethod undefined (type reflect.Type has no field or method NumMethod)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:532:8: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:565:13: Copy not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/json/decode.go:569:7: v.SetLen undefined (type reflect.Value has no field or method SetLen)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:606:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:637:40: v.NumMethod undefined (type reflect.Value has no field or method NumMethod)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:653:12: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:658:16: PtrTo not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/json/decode.go:658:24: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:665:18: MakeMap not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/json/decode.go:790:12: t.Key undefined (type reflect.Type has no field or method Key)
../../.gvm/gos/go1.14/src/encoding/json/decode.go:793:17: PtrTo not declared by package reflect
[...more lines following...]</pre>







## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
../../.gvm/gos/go1.14/src/encoding/xml/marshal.go:490:15: typ.Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/xml/marshal.go:642:16: typ.Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/xml/marshal.go:643:26: typ.Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/xml/marshal.go:647:33: typ.Elem().Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/xml/marshal.go:770:12: Copy not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/xml/read.go:193:7: t.Name undefined (type reflect.Type has no field or method Name)
../../.gvm/gos/go1.14/src/encoding/xml/read.go:282:19: Append not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/xml/read.go:286:8: val.SetLen undefined (type reflect.Value has no field or method SetLen)
../../.gvm/gos/go1.14/src/encoding/xml/read.go:398:17: Append not declared by package reflect
../../.gvm/gos/go1.14/src/encoding/xml/read.go:402:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
../../.gvm/gos/go1.14/src/encoding/xml/read.go:665:7: dst.SetBytes undefined (type reflect.Value has no field or method SetBytes)
../../.gvm/gos/go1.14/src/encoding/xml/typeinfo.go:63:50: f.Tag.Get undefined (type string has no field or method Get)
../../.gvm/gos/go1.14/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
[...more lines following...]</pre>







## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [net/http](#net-http)





## flag


The compiler gave the following error when this package was imported:

<pre># flag
../../.gvm/gos/go1.14/src/flag/flag.go:1010:33: Args not declared by package os
../../.gvm/gos/go1.14/src/flag/flag.go:582:57: Args not declared by package os
../../.gvm/gos/go1.14/src/flag/flag.go:999:23: Args not declared by package os
</pre>









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
  * [text/template/parse](#text-template-parse)





















## log/syslog


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)







## math/big


The compiler gave the following error when this package was imported:

<pre># math/big
../../.gvm/gos/go1.14/src/math/bits.go: interp: branch on a non-constant
</pre>











## mime


The compiler gave the following error when this package was imported:

<pre># mime
../../.gvm/gos/go1.14/src/mime/type.go:25:4: m.Range undefined (type *sync.Map has no field or method Range)
</pre>





## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [mime](#mime)
  * [net/textproto](#net-textproto)







## net


The compiler gave the following error when this package was imported:

<pre># internal/singleflight
../../.gvm/gos/go1.14/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
</pre>





## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [crypto/tls](#crypto-tls)
  * [vendor/golang.org/x/net/http/httpguts](#vendor-golang.org-x-net-http-httpguts)
  * [vendor/golang.org/x/net/http/httpproxy](#vendor-golang.org-x-net-http-httpproxy)
  * [vendor/golang.org/x/net/http2/hpack](#vendor-golang.org-x-net-http2-hpack)
  * [mime](#mime)
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
  * [flag](#flag)
  * [vendor/golang.org/x/net/http/httpguts](#vendor-golang.org-x-net-http-httpguts)
  * [net](#net)
  * [net/http](#net-http)





## net/http/httptrace


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net](#net)
  * [net/textproto](#net-textproto)





## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [vendor/golang.org/x/net/http/httpguts](#vendor-golang.org-x-net-http-httpguts)
  * [net](#net)
  * [net/http](#net-http)
  * [net/textproto](#net-textproto)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [html/template](#html-template)
  * [net/http](#net-http)





## net/mail


This package cannot be imported because the following dependencies cannot be compiled:

  * [mime](#mime)
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
../../.gvm/gos/go1.14/src/os/exec/exec.go:129:14: Process not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:133:19: ProcessState not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:457:6: ProcessState not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:229:12: Environ not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:245:23: DevNull not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:257:20: Pipe not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:290:27: DevNull not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:302:20: Pipe not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:417:22: StartProcess not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:417:57: ProcAttr not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:579:20: Pipe not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:621:20: Pipe not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:646:20: Pipe not declared by package os
../../.gvm/gos/go1.14/src/os/exec/exec.go:783:38: Getenv not declared by package os
</pre>





## os/signal


The compiler gave the following error when this package was imported:

<pre># os/signal
../../.gvm/gos/go1.14/src/os/signal/signal.go:15:18: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:27:14: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:50:23: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:85:23: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:90:21: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:120:25: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:120:43: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:170:22: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:177:23: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:229:21: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal_unix.go:37:20: Signal not declared by package os
../../.gvm/gos/go1.14/src/os/signal/signal.go:137:36: Signal not declared by package os
</pre>





## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
../../.gvm/gos/go1.14/src/os/user/lookup.go:15:41: undeclared name: current
../../.gvm/gos/go1.14/src/os/user/lookup.go:36:9: undeclared name: lookupUser
../../.gvm/gos/go1.14/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
../../.gvm/gos/go1.14/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
../../.gvm/gos/go1.14/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
../../.gvm/gos/go1.14/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>































## testing


This package cannot be imported because the following dependencies cannot be compiled:

  * [flag](#flag)







## testing/quick


This package cannot be imported because the following dependencies cannot be compiled:

  * [flag](#flag)









## text/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template/parse](#text-template-parse)





## text/template/parse


The compiler gave the following error when this package was imported:

<pre># text/template/parse
../../.gvm/gos/go1.14/src/text/template/parse/parse.go:188:26: Error not declared by package runtime
</pre>

















## vendor/golang.org/x/crypto/cryptobyte


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)















## vendor/golang.org/x/net/http/httpguts


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/textproto](#net-textproto)





## vendor/golang.org/x/net/http/httpproxy


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)





## vendor/golang.org/x/net/http2/hpack


The compiler gave the following error when this package was imported:

<pre>../../.gvm/gos/go1.14/src/vendor/golang.org/x/net/http2/hpack/tables.go:56:28: interp: unknown GEP
</pre>







## vendor/golang.org/x/net/nettest


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [os/exec](#os-exec)
  * [testing](#testing)














