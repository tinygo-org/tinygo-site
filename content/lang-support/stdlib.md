
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
context |  [<span style="color: red">✗</span> no](#context)  | 
crypto |  <span style="color: green">✔</span> yes  | 
crypto/aes |  <span style="color: green">✔</span> yes  | 
crypto/cipher |  <span style="color: green">✔</span> yes  | 
crypto/des |  <span style="color: green">✔</span> yes  | 
crypto/dsa |  [<span style="color: red">✗</span> no](#crypto-dsa)  | 
crypto/ecdsa |  [<span style="color: red">✗</span> no](#crypto-ecdsa)  | 
crypto/elliptic |  [<span style="color: red">✗</span> no](#crypto-elliptic)  | 
crypto/hmac |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  [<span style="color: red">✗</span> no](#crypto-md5)  | 
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
database/sql/driver |  [<span style="color: red">✗</span> no](#database-sql-driver)  | 
debug/dwarf |  <span style="color: green">✔</span> yes  | 
debug/elf |  [<span style="color: red">✗</span> no](#debug-elf)  | 
debug/gosym |  <span style="color: green">✔</span> yes  | 
debug/macho |  [<span style="color: red">✗</span> no](#debug-macho)  | 
debug/pe |  [<span style="color: red">✗</span> no](#debug-pe)  | 
debug/plan9obj |  [<span style="color: red">✗</span> no](#debug-plan9obj)  | 
encoding |  <span style="color: green">✔</span> yes  | 
encoding/ascii85 |  <span style="color: green">✔</span> yes  | 
encoding/asn1 |  [<span style="color: red">✗</span> no](#encoding-asn1)  | 
encoding/base32 |  [<span style="color: red">✗</span> no](#encoding-base32)  | 
encoding/base64 |  [<span style="color: red">✗</span> no](#encoding-base64)  | 
encoding/binary |  <span style="color: green">✔</span> yes  | 
encoding/csv |  <span style="color: green">✔</span> yes  | 
encoding/gob |  [<span style="color: red">✗</span> no](#encoding-gob)  | 
encoding/hex |  <span style="color: green">✔</span> yes  | 
encoding/json |  [<span style="color: red">✗</span> no](#encoding-json)  | 
encoding/pem |  [<span style="color: red">✗</span> no](#encoding-pem)  | 
encoding/xml |  [<span style="color: red">✗</span> no](#encoding-xml)  | 
errors |  <span style="color: green">✔</span> yes  | 
expvar |  [<span style="color: red">✗</span> no](#expvar)  | 
flag |  [<span style="color: red">✗</span> no](#flag)  | 
fmt |  <span style="color: green">✔</span> yes  | 
go/ast |  [<span style="color: red">✗</span> no](#go-ast)  | 
go/build |  [<span style="color: red">✗</span> no](#go-build)  | 
go/constant |  [<span style="color: red">✗</span> no](#go-constant)  | 
go/doc |  [<span style="color: red">✗</span> no](#go-doc)  | 
go/format |  [<span style="color: red">✗</span> no](#go-format)  | 
go/importer |  [<span style="color: red">✗</span> no](#go-importer)  | 
go/parser |  [<span style="color: red">✗</span> no](#go-parser)  | 
go/printer |  [<span style="color: red">✗</span> no](#go-printer)  | 
go/scanner |  [<span style="color: red">✗</span> no](#go-scanner)  | 
go/token |  [<span style="color: red">✗</span> no](#go-token)  | 
go/types |  [<span style="color: red">✗</span> no](#go-types)  | 
hash |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  | 
hash/crc64 |  <span style="color: green">✔</span> yes  | 
hash/fnv |  <span style="color: green">✔</span> yes  | 
html |  [<span style="color: red">✗</span> no](#html)  | 
html/template |  [<span style="color: red">✗</span> no](#html-template)  | 
image |  <span style="color: green">✔</span> yes  | 
image/color |  <span style="color: green">✔</span> yes  | 
image/color/palette |  [<span style="color: red">✗</span> no](#image-color-palette)  | 
image/draw |  [<span style="color: red">✗</span> no](#image-draw)  | 
image/gif |  [<span style="color: red">✗</span> no](#image-gif)  | 
image/jpeg |  [<span style="color: red">✗</span> no](#image-jpeg)  | 
image/png |  [<span style="color: red">✗</span> no](#image-png)  | 
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



## archive/tar


The compiler gave the following error when this package was imported:

<pre># archive/tar
../../../../../usr/local/go/src/archive/tar/common.go:636:15: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
../../../../../usr/local/go/src/archive/tar/common.go:637:21: fm.Perm undefined (type os.FileMode has no field or method Perm)
../../../../../usr/local/go/src/archive/tar/common.go:640:10: fm.IsRegular undefined (type os.FileMode has no field or method IsRegular)
../../../../../usr/local/go/src/archive/tar/common.go:554:32: os.FileMode(fi.h.Mode).Perm undefined (type os.FileMode has no field or method Perm)
../../../../../usr/local/go/src/archive/tar/common.go:445:15: DeepEqual not declared by package reflect
</pre>





## archive/zip


The compiler gave the following error when this package was imported:

<pre># archive/zip
../../../../../usr/local/go/src/archive/zip/register.go:106:21: Map not declared by package sync
../../../../../usr/local/go/src/archive/zip/register.go:107:21: Map not declared by package sync
../../../../../usr/local/go/src/archive/zip/struct.go:175:19: fi.ModTime undefined (type os.FileInfo has no field or method ModTime)
../../../../../usr/local/go/src/archive/zip/reader.go:61:19: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to r.init: missing method ReadAt
</pre>

























## context


The compiler gave the following error when this package was imported:

<pre># context
../../../../../usr/local/go/src/context/context.go:471:26: reflect.TypeOf(key).Comparable undefined (type reflect.Type has no field or method Comparable)
</pre>













## crypto/dsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)





## crypto/ecdsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/elliptic](#crypto-elliptic)
  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)





## crypto/elliptic


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)







## crypto/md5


The compiler gave the following error when this package was imported:

<pre>../../../../../usr/local/go/src/crypto/md5/md5block.go:21:8: todo: full slice expressions (with max): []byte
</pre>





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
  * [crypto/elliptic](#crypto-elliptic)
  * [crypto/md5](#crypto-md5)
  * [crypto/rand](#crypto-rand)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509](#crypto-x509)
  * [encoding/asn1](#encoding-asn1)
  * [encoding/pem](#encoding-pem)
  * [math/big](#math-big)
  * [net](#net)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/dsa](#crypto-dsa)
  * [crypto/ecdsa](#crypto-ecdsa)
  * [crypto/elliptic](#crypto-elliptic)
  * [crypto/md5](#crypto-md5)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509/pkix](#crypto-x509-pkix)
  * [encoding/asn1](#encoding-asn1)
  * [encoding/pem](#encoding-pem)
  * [math/big](#math-big)
  * [net](#net)





## crypto/x509/pkix


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/asn1](#encoding-asn1)
  * [math/big](#math-big)





## database/sql


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)
  * [database/sql/driver](#database-sql-driver)





## database/sql/driver


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)







## debug/elf


The compiler gave the following error when this package was imported:

<pre># debug/elf
../../../../../usr/local/go/src/debug/elf/reader.go:73:16: ErrInvalid not declared by package os
../../../../../usr/local/go/src/debug/elf/reader.go:81:16: ErrInvalid not declared by package os
../../../../../usr/local/go/src/debug/elf/file.go:201:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
</pre>







## debug/macho


The compiler gave the following error when this package was imported:

<pre># debug/macho
../../../../../usr/local/go/src/debug/macho/file.go:205:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
../../../../../usr/local/go/src/debug/macho/fat.go:130:24: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFatFile: missing method ReadAt
</pre>





## debug/pe


The compiler gave the following error when this package was imported:

<pre># debug/pe
../../../../../usr/local/go/src/debug/pe/file.go:40:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
</pre>





## debug/plan9obj


The compiler gave the following error when this package was imported:

<pre># debug/plan9obj
../../../../../usr/local/go/src/debug/plan9obj/file.go:103:21: cannot use f (variable of type *os.File) as io.ReaderAt value in argument to NewFile: missing method ReadAt
</pre>









## encoding/asn1


This package cannot be imported because the following dependencies cannot be compiled:

  * [math/big](#math-big)





## encoding/base32


The compiler gave the following error when this package was imported:

<pre>error: interp: branch on a non-const-propagated constant expression
</pre>





## encoding/base64


The compiler gave the following error when this package was imported:

<pre>error: interp: branch on a non-const-propagated constant expression
</pre>









## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
../../../../../usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
../../../../../usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
../../../../../usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
../../../../../usr/local/go/src/encoding/gob/type.go:844:8: rt.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:866:8: rt.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:867:9: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
../../../../../usr/local/go/src/encoding/gob/type.go:868:21: rt.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:870:21: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
../../../../../usr/local/go/src/encoding/gob/type.go:870:42: rt.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:715:28: rt.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:722:31: rt.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:496:34: t.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/type.go:514:37: t.Elem().Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:531:17: typ.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
../../../../../usr/local/go/src/encoding/gob/type.go:126:9: rt.Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
../../../../../usr/local/go/src/encoding/gob/encoder.go:128:29: st.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
../../../../../usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
../../../../../usr/local/go/src/encoding/gob/encode.go:562:34: t.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/encode.go:320:18: value.FieldByIndex undefined (type reflect.Value has no field or method FieldByIndex)
../../../../../usr/local/go/src/encoding/gob/decode.go:1202:17: base.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
../../../../../usr/local/go/src/encoding/gob/decode.go:1019:31: t.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
../../../../../usr/local/go/src/encoding/gob/decode.go:825:35: t.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
../../../../../usr/local/go/src/encoding/gob/decode.go:564:19: mtyp.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/decode.go:568:27: mtyp.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/decode.go:569:28: mtyp.Key undefined (type reflect.Type has no field or method Key)
../../../../../usr/local/go/src/encoding/gob/decode.go:575:9: value.SetMapIndex undefined (type reflect.Value has no field or method SetMapIndex)
../../../../../usr/local/go/src/encoding/gob/decode.go:466:18: value.FieldByIndex undefined (type reflect.Value has no field or method FieldByIndex)
</pre>







## encoding/json


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/base64](#encoding-base64)





## encoding/pem


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/base64](#encoding-base64)





## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:47:19: Map not declared by package sync
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:117:11: f.Tag undefined (type *reflect.StructField has no field or method Tag)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:168:20: f.Tag undefined (type *reflect.StructField has no field or method Tag)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:175:19: f.Tag undefined (type *reflect.StructField has no field or method Tag)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:63:10: f.PkgPath undefined (type reflect.StructField has no field or method PkgPath)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:63:30: f.Anonymous undefined (type reflect.StructField has no field or method Anonymous)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:63:46: f.Tag undefined (type reflect.StructField has no field or method Tag)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:68:9: f.Anonymous undefined (type reflect.StructField has no field or method Anonymous)
../../../../../usr/local/go/src/encoding/xml/read.go:665:7: dst.SetBytes undefined (type reflect.Value has no field or method SetBytes)
../../../../../usr/local/go/src/encoding/xml/read.go:339:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:347:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:352:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:358:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:398:17: Append not declared by package reflect
../../../../../usr/local/go/src/encoding/xml/read.go:402:6: v.SetLen undefined (type reflect.Value has no field or method SetLen)
../../../../../usr/local/go/src/encoding/xml/read.go:562:70: saveData.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:571:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:253:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:260:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:266:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:273:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:282:19: Append not declared by package reflect
../../../../../usr/local/go/src/encoding/xml/read.go:286:8: val.SetLen undefined (type reflect.Value has no field or method SetLen)
../../../../../usr/local/go/src/encoding/xml/read.go:193:7: t.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/xml/marshal.go:817:38: vf.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:829:39: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:770:12: Copy not declared by package reflect
../../../../../usr/local/go/src/encoding/xml/marshal.go:642:16: typ.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/xml/marshal.go:643:26: typ.Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/xml/marshal.go:647:33: typ.Elem().Name undefined (type reflect.Type has no field or method Name)
../../../../../usr/local/go/src/encoding/xml/marshal.go:548:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:561:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:573:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:584:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:431:31: typ.Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:436:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:442:31: typ.Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:447:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/marshal.go:490:15: typ.Name undefined (type reflect.Type has no field or method Name)
</pre>







## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [net/http](#net-http)





## flag


The compiler gave the following error when this package was imported:

<pre># flag
../../../../../usr/local/go/src/flag/flag.go:1009:33: Args not declared by package os
../../../../../usr/local/go/src/flag/flag.go:998:23: Args not declared by package os
../../../../../usr/local/go/src/flag/flag.go:581:57: Args not declared by package os
</pre>







## go/ast


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/scanner](#go-scanner)
  * [go/token](#go-token)





## go/build


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/ast](#go-ast)
  * [go/doc](#go-doc)
  * [go/parser](#go-parser)
  * [go/token](#go-token)
  * [os/exec](#os-exec)





## go/constant


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/token](#go-token)
  * [math/big](#math-big)





## go/doc


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/ast](#go-ast)
  * [go/token](#go-token)
  * [text/template](#text-template)





## go/format


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/ast](#go-ast)
  * [go/parser](#go-parser)
  * [go/printer](#go-printer)
  * [go/token](#go-token)





## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/build](#go-build)
  * [go/token](#go-token)
  * [go/types](#go-types)





## go/parser


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/ast](#go-ast)
  * [go/scanner](#go-scanner)
  * [go/token](#go-token)





## go/printer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/ast](#go-ast)
  * [go/token](#go-token)





## go/scanner


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/token](#go-token)





## go/token


The compiler gave the following error when this package was imported:

<pre>panic: interface conversion: interp.Value is *interp.LocalValue, not *interp.MapValue

goroutine 1 [running]:
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc0028287f0, 0x7fdc34fb1130, 0x7fdc34fb0df0, 0x83fad2, 0x4, 0x0, 0x0, 0xc0033e2608, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:279 +0x9174
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc0023cb860, 0x7fdc34030be8, 0xc0033f4620, 0x2, 0x2, 0xc0033e8470, 0x8, 0x83fad2, 0x4, 0x8d16c0, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc002829310, 0x7fdc34fafc20, 0x7fdc34faf630, 0x0, 0x0, 0x0, 0x0, 0xc0033e25d0, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:449 +0x51b3
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc0023cb860, 0x7fdc34030af8, 0xc002829450, 0x2, 0x2, 0xc0033e8470, 0x8, 0x0, 0x0, 0x0, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
github.com/tinygo-org/tinygo/interp.(*Eval).Function(...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:87
github.com/tinygo-org/tinygo/interp.Run(0x17dd080, 0x17d2f90, 0x0, 0x0, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:74 +0x672
main.Compile(0x7ffe7330621c, 0x13, 0x7ffe73306200, 0x1b, 0xc000138180, 0xc002829ee0, 0xc002829d48, 0x0, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/main.go:140 +0xc5b
main.Build(0x7ffe7330621c, 0x13, 0x7ffe73306200, 0x1b, 0x83fdc6, 0x4, 0xc00008bee0, 0xc0000f01b0, 0x7fdc46bb0ba8)
	/home/ayke/src/github.com/tinygo-org/tinygo/main.go:337 +0xe2
main.main()
	/home/ayke/src/github.com/tinygo-org/tinygo/main.go:696 +0xf4a
</pre>





## go/types


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/ast](#go-ast)
  * [go/constant](#go-constant)
  * [go/parser](#go-parser)
  * [go/token](#go-token)















## html


The compiler gave the following error when this package was imported:

<pre>error: interp: branch on a non-const-propagated constant expression
</pre>





## html/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/json](#encoding-json)
  * [html](#html)
  * [text/template](#text-template)
  * [text/template/parse](#text-template-parse)









## image/color/palette


The compiler gave the following error when this package was imported:

<pre>Use of instruction is not an instruction!
  %7 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %6, 1, !dbg !1562
Use of instruction is not an instruction!
  %11 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %10, 1, !dbg !1562
Use of instruction is not an instruction!
  %15 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %14, 1, !dbg !1562
Use of instruction is not an instruction!
  %19 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %18, 1, !dbg !1562
Use of instruction is not an instruction!
  %23 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %22, 1, !dbg !1562
Use of instruction is not an instruction!
  %27 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %26, 1, !dbg !1562
Use of instruction is not an instruction!
  %31 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %30, 1, !dbg !1562
Use of instruction is not an instruction!
  %35 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %34, 1, !dbg !1562
Use of instruction is not an instruction!
  %39 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %38, 1, !dbg !1562
Use of instruction is not an instruction!
  %43 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %42, 1, !dbg !1562
Use of instruction is not an instruction!
  %47 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %46, 1, !dbg !1562
Use of instruction is not an instruction!
  %51 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %50, 1, !dbg !1562
Use of instruction is not an instruction!
  %55 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %54, 1, !dbg !1562
Use of instruction is not an instruction!
  %59 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %58, 1, !dbg !1562
Use of instruction is not an instruction!
  %63 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %62, 1, !dbg !1562
Use of instruction is not an instruction!
  %67 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %66, 1, !dbg !1562
Use of instruction is not an instruction!
  %71 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %70, 1, !dbg !1562
Use of instruction is not an instruction!
  %75 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %74, 1, !dbg !1562
Use of instruction is not an instruction!
  %79 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %78, 1, !dbg !1562
Use of instruction is not an instruction!
  %83 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %82, 1, !dbg !1562
Use of instruction is not an instruction!
  %87 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %86, 1, !dbg !1562
Use of instruction is not an instruction!
  %91 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %90, 1, !dbg !1562
Use of instruction is not an instruction!
  %95 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %94, 1, !dbg !1562
Use of instruction is not an instruction!
  %99 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %98, 1, !dbg !1562
Use of instruction is not an instruction!
  %103 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %102, 1, !dbg !1562
Use of instruction is not an instruction!
  %107 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %106, 1, !dbg !1562
Use of instruction is not an instruction!
  %111 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %110, 1, !dbg !1562
Use of instruction is not an instruction!
  %115 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %114, 1, !dbg !1562
Use of instruction is not an instruction!
  %119 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %118, 1, !dbg !1562
Use of instruction is not an instruction!
  %123 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %122, 1, !dbg !1562
Use of instruction is not an instruction!
  %127 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %126, 1, !dbg !1562
Use of instruction is not an instruction!
  %131 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %130, 1, !dbg !1562
Use of instruction is not an instruction!
  %135 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %134, 1, !dbg !1562
Use of instruction is not an instruction!
  %139 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %138, 1, !dbg !1562
Use of instruction is not an instruction!
  %143 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %142, 1, !dbg !1562
Use of instruction is not an instruction!
  %147 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %146, 1, !dbg !1562
Use of instruction is not an instruction!
  %151 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %150, 1, !dbg !1562
Use of instruction is not an instruction!
  %155 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %154, 1, !dbg !1562
Use of instruction is not an instruction!
  %159 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %158, 1, !dbg !1562
Use of instruction is not an instruction!
  %163 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %162, 1, !dbg !1562
Use of instruction is not an instruction!
  %167 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %166, 1, !dbg !1562
Use of instruction is not an instruction!
  %171 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %170, 1, !dbg !1562
Use of instruction is not an instruction!
  %175 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %174, 1, !dbg !1562
Use of instruction is not an instruction!
  %179 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %178, 1, !dbg !1562
Use of instruction is not an instruction!
  %183 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %182, 1, !dbg !1562
Use of instruction is not an instruction!
  %187 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %186, 1, !dbg !1562
Use of instruction is not an instruction!
  %191 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %190, 1, !dbg !1562
Use of instruction is not an instruction!
  %195 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %194, 1, !dbg !1562
Use of instruction is not an instruction!
  %199 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %198, 1, !dbg !1562
Use of instruction is not an instruction!
  %203 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %202, 1, !dbg !1562
Use of instruction is not an instruction!
  %207 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %206, 1, !dbg !1562
Use of instruction is not an instruction!
  %211 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %210, 1, !dbg !1562
Use of instruction is not an instruction!
  %215 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %214, 1, !dbg !1562
Use of instruction is not an instruction!
  %219 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %218, 1, !dbg !1562
Use of instruction is not an instruction!
  %223 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %222, 1, !dbg !1562
Use of instruction is not an instruction!
  %227 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %226, 1, !dbg !1562
Use of instruction is not an instruction!
  %231 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %230, 1, !dbg !1562
Use of instruction is not an instruction!
  %235 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %234, 1, !dbg !1562
Use of instruction is not an instruction!
  %239 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %238, 1, !dbg !1562
Use of instruction is not an instruction!
  %243 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %242, 1, !dbg !1562
Use of instruction is not an instruction!
  %247 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %246, 1, !dbg !1562
Use of instruction is not an instruction!
  %251 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %250, 1, !dbg !1562
Use of instruction is not an instruction!
  %255 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %254, 1, !dbg !1562
Use of instruction is not an instruction!
  %259 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %258, 1, !dbg !1562
Use of instruction is not an instruction!
  %263 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %262, 1, !dbg !1562
Use of instruction is not an instruction!
  %267 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %266, 1, !dbg !1562
Use of instruction is not an instruction!
  %271 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %270, 1, !dbg !1562
Use of instruction is not an instruction!
  %275 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %274, 1, !dbg !1562
Use of instruction is not an instruction!
  %279 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %278, 1, !dbg !1562
Use of instruction is not an instruction!
  %283 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %282, 1, !dbg !1562
Use of instruction is not an instruction!
  %287 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %286, 1, !dbg !1562
Use of instruction is not an instruction!
  %291 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %290, 1, !dbg !1562
Use of instruction is not an instruction!
  %295 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %294, 1, !dbg !1562
Use of instruction is not an instruction!
  %299 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %298, 1, !dbg !1562
Use of instruction is not an instruction!
  %303 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %302, 1, !dbg !1562
Use of instruction is not an instruction!
  %307 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %306, 1, !dbg !1562
Use of instruction is not an instruction!
  %311 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %310, 1, !dbg !1562
Use of instruction is not an instruction!
  %315 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %314, 1, !dbg !1562
Use of instruction is not an instruction!
  %319 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %318, 1, !dbg !1562
Use of instruction is not an instruction!
  %323 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %322, 1, !dbg !1562
Use of instruction is not an instruction!
  %327 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %326, 1, !dbg !1562
Use of instruction is not an instruction!
  %331 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %330, 1, !dbg !1562
Use of instruction is not an instruction!
  %335 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %334, 1, !dbg !1562
Use of instruction is not an instruction!
  %339 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %338, 1, !dbg !1562
Use of instruction is not an instruction!
  %343 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %342, 1, !dbg !1562
Use of instruction is not an instruction!
  %347 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %346, 1, !dbg !1562
Use of instruction is not an instruction!
  %351 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %350, 1, !dbg !1562
Use of instruction is not an instruction!
  %355 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %354, 1, !dbg !1562
Use of instruction is not an instruction!
  %359 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %358, 1, !dbg !1562
Use of instruction is not an instruction!
  %363 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %362, 1, !dbg !1562
Use of instruction is not an instruction!
  %367 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %366, 1, !dbg !1562
Use of instruction is not an instruction!
  %371 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %370, 1, !dbg !1562
Use of instruction is not an instruction!
  %375 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %374, 1, !dbg !1562
Use of instruction is not an instruction!
  %379 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %378, 1, !dbg !1562
Use of instruction is not an instruction!
  %383 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %382, 1, !dbg !1562
Use of instruction is not an instruction!
  %387 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %386, 1, !dbg !1562
Use of instruction is not an instruction!
  %391 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %390, 1, !dbg !1562
Use of instruction is not an instruction!
  %395 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %394, 1, !dbg !1562
Use of instruction is not an instruction!
  %399 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %398, 1, !dbg !1562
Use of instruction is not an instruction!
  %403 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %402, 1, !dbg !1562
Use of instruction is not an instruction!
  %407 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %406, 1, !dbg !1562
Use of instruction is not an instruction!
  %411 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %410, 1, !dbg !1562
Use of instruction is not an instruction!
  %415 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %414, 1, !dbg !1562
Use of instruction is not an instruction!
  %419 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %418, 1, !dbg !1562
Use of instruction is not an instruction!
  %423 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %422, 1, !dbg !1562
Use of instruction is not an instruction!
  %427 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %426, 1, !dbg !1562
Use of instruction is not an instruction!
  %431 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %430, 1, !dbg !1562
Use of instruction is not an instruction!
  %435 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %434, 1, !dbg !1562
Use of instruction is not an instruction!
  %439 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %438, 1, !dbg !1562
Use of instruction is not an instruction!
  %443 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %442, 1, !dbg !1562
Use of instruction is not an instruction!
  %447 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %446, 1, !dbg !1562
Use of instruction is not an instruction!
  %451 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %450, 1, !dbg !1562
Use of instruction is not an instruction!
  %455 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %454, 1, !dbg !1562
Use of instruction is not an instruction!
  %459 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %458, 1, !dbg !1562
Use of instruction is not an instruction!
  %463 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %462, 1, !dbg !1562
Use of instruction is not an instruction!
  %467 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %466, 1, !dbg !1562
Use of instruction is not an instruction!
  %471 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %470, 1, !dbg !1562
Use of instruction is not an instruction!
  %475 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %474, 1, !dbg !1562
Use of instruction is not an instruction!
  %479 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %478, 1, !dbg !1562
Use of instruction is not an instruction!
  %483 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %482, 1, !dbg !1562
Use of instruction is not an instruction!
  %487 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %486, 1, !dbg !1562
Use of instruction is not an instruction!
  %491 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %490, 1, !dbg !1562
Use of instruction is not an instruction!
  %495 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %494, 1, !dbg !1562
Use of instruction is not an instruction!
  %499 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %498, 1, !dbg !1562
Use of instruction is not an instruction!
  %503 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %502, 1, !dbg !1562
Use of instruction is not an instruction!
  %507 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %506, 1, !dbg !1562
Use of instruction is not an instruction!
  %511 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %510, 1, !dbg !1562
Use of instruction is not an instruction!
  %515 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %514, 1, !dbg !1562
Use of instruction is not an instruction!
  %519 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %518, 1, !dbg !1562
Use of instruction is not an instruction!
  %523 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %522, 1, !dbg !1562
Use of instruction is not an instruction!
  %527 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %526, 1, !dbg !1562
Use of instruction is not an instruction!
  %531 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %530, 1, !dbg !1562
Use of instruction is not an instruction!
  %535 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %534, 1, !dbg !1562
Use of instruction is not an instruction!
  %539 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %538, 1, !dbg !1562
Use of instruction is not an instruction!
  %543 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %542, 1, !dbg !1562
Use of instruction is not an instruction!
  %547 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %546, 1, !dbg !1562
Use of instruction is not an instruction!
  %551 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %550, 1, !dbg !1562
Use of instruction is not an instruction!
  %555 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %554, 1, !dbg !1562
Use of instruction is not an instruction!
  %559 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %558, 1, !dbg !1562
Use of instruction is not an instruction!
  %563 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %562, 1, !dbg !1562
Use of instruction is not an instruction!
  %567 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %566, 1, !dbg !1562
Use of instruction is not an instruction!
  %571 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %570, 1, !dbg !1562
Use of instruction is not an instruction!
  %575 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %574, 1, !dbg !1562
Use of instruction is not an instruction!
  %579 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %578, 1, !dbg !1562
Use of instruction is not an instruction!
  %583 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %582, 1, !dbg !1562
Use of instruction is not an instruction!
  %587 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %586, 1, !dbg !1562
Use of instruction is not an instruction!
  %591 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %590, 1, !dbg !1562
Use of instruction is not an instruction!
  %595 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %594, 1, !dbg !1562
Use of instruction is not an instruction!
  %599 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %598, 1, !dbg !1562
Use of instruction is not an instruction!
  %603 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %602, 1, !dbg !1562
Use of instruction is not an instruction!
  %607 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %606, 1, !dbg !1562
Use of instruction is not an instruction!
  %611 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %610, 1, !dbg !1562
Use of instruction is not an instruction!
  %615 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %614, 1, !dbg !1562
Use of instruction is not an instruction!
  %619 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %618, 1, !dbg !1562
Use of instruction is not an instruction!
  %623 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %622, 1, !dbg !1562
Use of instruction is not an instruction!
  %627 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %626, 1, !dbg !1562
Use of instruction is not an instruction!
  %631 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %630, 1, !dbg !1562
Use of instruction is not an instruction!
  %635 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %634, 1, !dbg !1562
Use of instruction is not an instruction!
  %639 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %638, 1, !dbg !1562
Use of instruction is not an instruction!
  %643 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %642, 1, !dbg !1562
Use of instruction is not an instruction!
  %647 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %646, 1, !dbg !1562
Use of instruction is not an instruction!
  %651 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %650, 1, !dbg !1562
Use of instruction is not an instruction!
  %655 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %654, 1, !dbg !1562
Use of instruction is not an instruction!
  %659 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %658, 1, !dbg !1562
Use of instruction is not an instruction!
  %663 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %662, 1, !dbg !1562
Use of instruction is not an instruction!
  %667 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %666, 1, !dbg !1562
Use of instruction is not an instruction!
  %671 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %670, 1, !dbg !1562
Use of instruction is not an instruction!
  %675 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %674, 1, !dbg !1562
Use of instruction is not an instruction!
  %679 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %678, 1, !dbg !1562
Use of instruction is not an instruction!
  %683 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %682, 1, !dbg !1562
Use of instruction is not an instruction!
  %687 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %686, 1, !dbg !1562
Use of instruction is not an instruction!
  %691 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %690, 1, !dbg !1562
Use of instruction is not an instruction!
  %695 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %694, 1, !dbg !1562
Use of instruction is not an instruction!
  %699 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %698, 1, !dbg !1562
Use of instruction is not an instruction!
  %703 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %702, 1, !dbg !1562
Use of instruction is not an instruction!
  %707 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %706, 1, !dbg !1562
Use of instruction is not an instruction!
  %711 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %710, 1, !dbg !1562
Use of instruction is not an instruction!
  %715 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %714, 1, !dbg !1562
Use of instruction is not an instruction!
  %719 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %718, 1, !dbg !1562
Use of instruction is not an instruction!
  %723 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %722, 1, !dbg !1562
Use of instruction is not an instruction!
  %727 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %726, 1, !dbg !1562
Use of instruction is not an instruction!
  %731 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %730, 1, !dbg !1562
Use of instruction is not an instruction!
  %735 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %734, 1, !dbg !1562
Use of instruction is not an instruction!
  %739 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %738, 1, !dbg !1562
Use of instruction is not an instruction!
  %743 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %742, 1, !dbg !1562
Use of instruction is not an instruction!
  %747 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %746, 1, !dbg !1562
Use of instruction is not an instruction!
  %751 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %750, 1, !dbg !1562
Use of instruction is not an instruction!
  %755 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %754, 1, !dbg !1562
Use of instruction is not an instruction!
  %759 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %758, 1, !dbg !1562
Use of instruction is not an instruction!
  %763 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %762, 1, !dbg !1562
Use of instruction is not an instruction!
  %767 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %766, 1, !dbg !1562
Use of instruction is not an instruction!
  %771 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %770, 1, !dbg !1562
Use of instruction is not an instruction!
  %775 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %774, 1, !dbg !1562
Use of instruction is not an instruction!
  %779 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %778, 1, !dbg !1562
Use of instruction is not an instruction!
  %783 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %782, 1, !dbg !1562
Use of instruction is not an instruction!
  %787 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %786, 1, !dbg !1562
Use of instruction is not an instruction!
  %791 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %790, 1, !dbg !1562
Use of instruction is not an instruction!
  %795 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %794, 1, !dbg !1562
Use of instruction is not an instruction!
  %799 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %798, 1, !dbg !1562
Use of instruction is not an instruction!
  %803 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %802, 1, !dbg !1562
Use of instruction is not an instruction!
  %807 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %806, 1, !dbg !1562
Use of instruction is not an instruction!
  %811 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %810, 1, !dbg !1562
Use of instruction is not an instruction!
  %815 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %814, 1, !dbg !1562
Use of instruction is not an instruction!
  %819 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %818, 1, !dbg !1562
Use of instruction is not an instruction!
  %823 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %822, 1, !dbg !1562
Use of instruction is not an instruction!
  %827 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %826, 1, !dbg !1562
Use of instruction is not an instruction!
  %831 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %830, 1, !dbg !1562
Use of instruction is not an instruction!
  %835 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %834, 1, !dbg !1562
Use of instruction is not an instruction!
  %839 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %838, 1, !dbg !1562
Use of instruction is not an instruction!
  %843 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %842, 1, !dbg !1562
Use of instruction is not an instruction!
  %847 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %846, 1, !dbg !1562
Use of instruction is not an instruction!
  %851 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %850, 1, !dbg !1562
Use of instruction is not an instruction!
  %855 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %854, 1, !dbg !1562
Use of instruction is not an instruction!
  %859 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %858, 1, !dbg !1562
Use of instruction is not an instruction!
  %863 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %862, 1, !dbg !1562
Use of instruction is not an instruction!
  %867 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %866, 1, !dbg !1562
Use of instruction is not an instruction!
  %871 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %870, 1, !dbg !1562
Use of instruction is not an instruction!
  %875 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %874, 1, !dbg !1562
Use of instruction is not an instruction!
  %879 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %878, 1, !dbg !1562
Use of instruction is not an instruction!
  %883 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %882, 1, !dbg !1562
Use of instruction is not an instruction!
  %887 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %886, 1, !dbg !1562
Use of instruction is not an instruction!
  %891 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %890, 1, !dbg !1562
Use of instruction is not an instruction!
  %895 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %894, 1, !dbg !1562
Use of instruction is not an instruction!
  %899 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %898, 1, !dbg !1562
Use of instruction is not an instruction!
  %903 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %902, 1, !dbg !1562
Use of instruction is not an instruction!
  %907 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %906, 1, !dbg !1562
Use of instruction is not an instruction!
  %911 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %910, 1, !dbg !1562
Use of instruction is not an instruction!
  %915 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %914, 1, !dbg !1562
Use of instruction is not an instruction!
  %919 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %918, 1, !dbg !1562
Use of instruction is not an instruction!
  %923 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %922, 1, !dbg !1562
Use of instruction is not an instruction!
  %927 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %926, 1, !dbg !1562
Use of instruction is not an instruction!
  %931 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %930, 1, !dbg !1562
Use of instruction is not an instruction!
  %935 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %934, 1, !dbg !1562
Use of instruction is not an instruction!
  %939 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %938, 1, !dbg !1562
Use of instruction is not an instruction!
  %943 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %942, 1, !dbg !1562
Use of instruction is not an instruction!
  %947 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %946, 1, !dbg !1562
Use of instruction is not an instruction!
  %951 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %950, 1, !dbg !1562
Use of instruction is not an instruction!
  %955 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %954, 1, !dbg !1562
Use of instruction is not an instruction!
  %959 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %958, 1, !dbg !1562
Use of instruction is not an instruction!
  %963 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %962, 1, !dbg !1562
Use of instruction is not an instruction!
  %967 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %966, 1, !dbg !1562
Use of instruction is not an instruction!
  %971 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %970, 1, !dbg !1562
Use of instruction is not an instruction!
  %975 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %974, 1, !dbg !1562
Use of instruction is not an instruction!
  %979 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %978, 1, !dbg !1562
Use of instruction is not an instruction!
  %983 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %982, 1, !dbg !1562
Use of instruction is not an instruction!
  %987 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %986, 1, !dbg !1562
Use of instruction is not an instruction!
  %991 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %990, 1, !dbg !1562
Use of instruction is not an instruction!
  %995 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %994, 1, !dbg !1562
Use of instruction is not an instruction!
  %999 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %998, 1, !dbg !1562
Use of instruction is not an instruction!
  %1003 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1002, 1, !dbg !1562
Use of instruction is not an instruction!
  %1007 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1006, 1, !dbg !1562
Use of instruction is not an instruction!
  %1011 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1010, 1, !dbg !1562
Use of instruction is not an instruction!
  %1015 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1014, 1, !dbg !1562
Use of instruction is not an instruction!
  %1019 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1018, 1, !dbg !1562
Use of instruction is not an instruction!
  %1023 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1022, 1, !dbg !1562
Use of instruction is not an instruction!
  %1027 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1026, 1, !dbg !1562
Use of instruction is not an instruction!
  %1031 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1030, 1, !dbg !1562
Use of instruction is not an instruction!
  %1035 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1034, 1, !dbg !1562
Use of instruction is not an instruction!
  %1039 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1038, 1, !dbg !1562
Use of instruction is not an instruction!
  %1043 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1042, 1, !dbg !1562
Use of instruction is not an instruction!
  %1047 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1046, 1, !dbg !1562
Use of instruction is not an instruction!
  %1051 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1050, 1, !dbg !1562
Use of instruction is not an instruction!
  %1055 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1054, 1, !dbg !1562
Use of instruction is not an instruction!
  %1059 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1058, 1, !dbg !1562
Use of instruction is not an instruction!
  %1063 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1062, 1, !dbg !1562
Use of instruction is not an instruction!
  %1067 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1066, 1, !dbg !1562
Use of instruction is not an instruction!
  %1071 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1070, 1, !dbg !1562
Use of instruction is not an instruction!
  %1075 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1074, 1, !dbg !1562
Use of instruction is not an instruction!
  %1079 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1078, 1, !dbg !1562
Use of instruction is not an instruction!
  %1083 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1082, 1, !dbg !1562
Use of instruction is not an instruction!
  %1087 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1086, 1, !dbg !1562
Use of instruction is not an instruction!
  %1091 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1090, 1, !dbg !1562
Use of instruction is not an instruction!
  %1095 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1094, 1, !dbg !1562
Use of instruction is not an instruction!
  %1099 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1098, 1, !dbg !1562
Use of instruction is not an instruction!
  %1103 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1102, 1, !dbg !1562
Use of instruction is not an instruction!
  %1107 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1106, 1, !dbg !1562
Use of instruction is not an instruction!
  %1111 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1110, 1, !dbg !1562
Use of instruction is not an instruction!
  %1115 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1114, 1, !dbg !1562
Use of instruction is not an instruction!
  %1119 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1118, 1, !dbg !1562
Use of instruction is not an instruction!
  %1123 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1122, 1, !dbg !1562
Use of instruction is not an instruction!
  %1127 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1126, 1, !dbg !1562
Use of instruction is not an instruction!
  %1131 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1130, 1, !dbg !1562
Use of instruction is not an instruction!
  %1135 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1134, 1, !dbg !1562
Use of instruction is not an instruction!
  %1139 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1138, 1, !dbg !1562
Use of instruction is not an instruction!
  %1143 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1142, 1, !dbg !1562
Use of instruction is not an instruction!
  %1147 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1146, 1, !dbg !1562
Use of instruction is not an instruction!
  %1151 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1150, 1, !dbg !1562
Use of instruction is not an instruction!
  %1155 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1154, 1, !dbg !1562
Use of instruction is not an instruction!
  %1159 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1158, 1, !dbg !1562
Use of instruction is not an instruction!
  %1163 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1162, 1, !dbg !1562
Use of instruction is not an instruction!
  %1167 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1166, 1, !dbg !1562
Use of instruction is not an instruction!
  %1171 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1170, 1, !dbg !1562
Use of instruction is not an instruction!
  %1175 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1174, 1, !dbg !1562
Use of instruction is not an instruction!
  %1179 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1178, 1, !dbg !1562
Use of instruction is not an instruction!
  %1183 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1182, 1, !dbg !1562
Use of instruction is not an instruction!
  %1187 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1186, 1, !dbg !1562
Use of instruction is not an instruction!
  %1191 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1190, 1, !dbg !1562
Use of instruction is not an instruction!
  %1195 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1194, 1, !dbg !1562
Use of instruction is not an instruction!
  %1199 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1198, 1, !dbg !1562
Use of instruction is not an instruction!
  %1203 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1202, 1, !dbg !1562
Use of instruction is not an instruction!
  %1207 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1206, 1, !dbg !1562
Use of instruction is not an instruction!
  %1211 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1210, 1, !dbg !1562
Use of instruction is not an instruction!
  %1215 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1214, 1, !dbg !1562
Use of instruction is not an instruction!
  %1219 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1218, 1, !dbg !1562
Use of instruction is not an instruction!
  %1223 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1222, 1, !dbg !1562
Use of instruction is not an instruction!
  %1227 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1226, 1, !dbg !1562
Use of instruction is not an instruction!
  %1231 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1230, 1, !dbg !1562
Use of instruction is not an instruction!
  %1235 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1234, 1, !dbg !1562
Use of instruction is not an instruction!
  %1239 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1238, 1, !dbg !1562
Use of instruction is not an instruction!
  %1243 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1242, 1, !dbg !1562
Use of instruction is not an instruction!
  %1247 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1246, 1, !dbg !1562
Use of instruction is not an instruction!
  %1251 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1250, 1, !dbg !1562
Use of instruction is not an instruction!
  %1255 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1254, 1, !dbg !1562
Use of instruction is not an instruction!
  %1259 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1258, 1, !dbg !1562
Use of instruction is not an instruction!
  %1263 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1262, 1, !dbg !1562
Use of instruction is not an instruction!
  %1267 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1266, 1, !dbg !1562
Use of instruction is not an instruction!
  %1271 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1270, 1, !dbg !1562
Use of instruction is not an instruction!
  %1275 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1274, 1, !dbg !1562
Use of instruction is not an instruction!
  %1279 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1278, 1, !dbg !1562
Use of instruction is not an instruction!
  %1283 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1282, 1, !dbg !1562
Use of instruction is not an instruction!
  %1287 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1286, 1, !dbg !1562
Use of instruction is not an instruction!
  %1291 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1290, 1, !dbg !1562
Use of instruction is not an instruction!
  %1295 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1294, 1, !dbg !1562
Use of instruction is not an instruction!
  %1299 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1298, 1, !dbg !1562
Use of instruction is not an instruction!
  %1303 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1302, 1, !dbg !1562
Use of instruction is not an instruction!
  %1307 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1306, 1, !dbg !1562
Use of instruction is not an instruction!
  %1311 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1310, 1, !dbg !1562
Use of instruction is not an instruction!
  %1315 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1314, 1, !dbg !1562
Use of instruction is not an instruction!
  %1319 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1318, 1, !dbg !1562
Use of instruction is not an instruction!
  %1323 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1322, 1, !dbg !1562
Use of instruction is not an instruction!
  %1327 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1326, 1, !dbg !1562
Use of instruction is not an instruction!
  %1331 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1330, 1, !dbg !1562
Use of instruction is not an instruction!
  %1335 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1334, 1, !dbg !1562
Use of instruction is not an instruction!
  %1339 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1338, 1, !dbg !1562
Use of instruction is not an instruction!
  %1343 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1342, 1, !dbg !1562
Use of instruction is not an instruction!
  %1347 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1346, 1, !dbg !1562
Use of instruction is not an instruction!
  %1351 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1350, 1, !dbg !1562
Use of instruction is not an instruction!
  %1355 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1354, 1, !dbg !1562
Use of instruction is not an instruction!
  %1359 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1358, 1, !dbg !1562
Use of instruction is not an instruction!
  %1363 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1362, 1, !dbg !1562
Use of instruction is not an instruction!
  %1367 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1366, 1, !dbg !1562
Use of instruction is not an instruction!
  %1371 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1370, 1, !dbg !1562
Use of instruction is not an instruction!
  %1375 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1374, 1, !dbg !1562
Use of instruction is not an instruction!
  %1379 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1378, 1, !dbg !1562
Use of instruction is not an instruction!
  %1383 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1382, 1, !dbg !1562
Use of instruction is not an instruction!
  %1387 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1386, 1, !dbg !1562
Use of instruction is not an instruction!
  %1391 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1390, 1, !dbg !1562
Use of instruction is not an instruction!
  %1395 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1394, 1, !dbg !1562
Use of instruction is not an instruction!
  %1399 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1398, 1, !dbg !1562
Use of instruction is not an instruction!
  %1403 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1402, 1, !dbg !1562
Use of instruction is not an instruction!
  %1407 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1406, 1, !dbg !1562
Use of instruction is not an instruction!
  %1411 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1410, 1, !dbg !1562
Use of instruction is not an instruction!
  %1415 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1414, 1, !dbg !1562
Use of instruction is not an instruction!
  %1419 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1418, 1, !dbg !1562
Use of instruction is not an instruction!
  %1423 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1422, 1, !dbg !1562
Use of instruction is not an instruction!
  %1427 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1426, 1, !dbg !1562
Use of instruction is not an instruction!
  %1431 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1430, 1, !dbg !1562
Use of instruction is not an instruction!
  %1435 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1434, 1, !dbg !1562
Use of instruction is not an instruction!
  %1439 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1438, 1, !dbg !1562
Use of instruction is not an instruction!
  %1443 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1442, 1, !dbg !1562
Use of instruction is not an instruction!
  %1447 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1446, 1, !dbg !1562
Use of instruction is not an instruction!
  %1451 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1450, 1, !dbg !1562
Use of instruction is not an instruction!
  %1455 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1454, 1, !dbg !1562
Use of instruction is not an instruction!
  %1459 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1458, 1, !dbg !1562
Use of instruction is not an instruction!
  %1463 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1462, 1, !dbg !1562
Use of instruction is not an instruction!
  %1467 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1466, 1, !dbg !1562
Use of instruction is not an instruction!
  %1471 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1470, 1, !dbg !1562
Use of instruction is not an instruction!
  %1475 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1474, 1, !dbg !1562
Use of instruction is not an instruction!
  %1479 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1478, 1, !dbg !1562
Use of instruction is not an instruction!
  %1483 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1482, 1, !dbg !1562
Use of instruction is not an instruction!
  %1487 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1486, 1, !dbg !1562
Use of instruction is not an instruction!
  %1491 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1490, 1, !dbg !1562
Use of instruction is not an instruction!
  %1495 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1494, 1, !dbg !1562
Use of instruction is not an instruction!
  %1499 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1498, 1, !dbg !1562
Use of instruction is not an instruction!
  %1503 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1502, 1, !dbg !1562
Use of instruction is not an instruction!
  %1507 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1506, 1, !dbg !1562
Use of instruction is not an instruction!
  %1511 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1510, 1, !dbg !1562
Use of instruction is not an instruction!
  %1515 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1514, 1, !dbg !1562
Use of instruction is not an instruction!
  %1519 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1518, 1, !dbg !1562
Use of instruction is not an instruction!
  %1523 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1522, 1, !dbg !1562
Use of instruction is not an instruction!
  %1527 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1526, 1, !dbg !1562
Use of instruction is not an instruction!
  %1531 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1530, 1, !dbg !1562
Use of instruction is not an instruction!
  %1535 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1534, 1, !dbg !1562
Use of instruction is not an instruction!
  %1539 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1538, 1, !dbg !1562
Use of instruction is not an instruction!
  %1543 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1542, 1, !dbg !1562
Use of instruction is not an instruction!
  %1547 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1546, 1, !dbg !1562
Use of instruction is not an instruction!
  %1551 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1550, 1, !dbg !1562
Use of instruction is not an instruction!
  %1555 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1554, 1, !dbg !1562
Use of instruction is not an instruction!
  %1559 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1558, 1, !dbg !1562
Use of instruction is not an instruction!
  %1563 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1562, 1, !dbg !1562
Use of instruction is not an instruction!
  %1567 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1566, 1, !dbg !1562
Use of instruction is not an instruction!
  %1571 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1570, 1, !dbg !1562
Use of instruction is not an instruction!
  %1575 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1574, 1, !dbg !1562
Use of instruction is not an instruction!
  %1579 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1578, 1, !dbg !1562
Use of instruction is not an instruction!
  %1583 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1582, 1, !dbg !1562
Use of instruction is not an instruction!
  %1587 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1586, 1, !dbg !1562
Use of instruction is not an instruction!
  %1591 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1590, 1, !dbg !1562
Use of instruction is not an instruction!
  %1595 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1594, 1, !dbg !1562
Use of instruction is not an instruction!
  %1599 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1598, 1, !dbg !1562
Use of instruction is not an instruction!
  %1603 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1602, 1, !dbg !1562
Use of instruction is not an instruction!
  %1607 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1606, 1, !dbg !1562
Use of instruction is not an instruction!
  %1611 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1610, 1, !dbg !1562
Use of instruction is not an instruction!
  %1615 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1614, 1, !dbg !1562
Use of instruction is not an instruction!
  %1619 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1618, 1, !dbg !1562
Use of instruction is not an instruction!
  %1623 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1622, 1, !dbg !1562
Use of instruction is not an instruction!
  %1627 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1626, 1, !dbg !1562
Use of instruction is not an instruction!
  %1631 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1630, 1, !dbg !1562
Use of instruction is not an instruction!
  %1635 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1634, 1, !dbg !1562
Use of instruction is not an instruction!
  %1639 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1638, 1, !dbg !1562
Use of instruction is not an instruction!
  %1643 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1642, 1, !dbg !1562
Use of instruction is not an instruction!
  %1647 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1646, 1, !dbg !1562
Use of instruction is not an instruction!
  %1651 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1650, 1, !dbg !1562
Use of instruction is not an instruction!
  %1655 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1654, 1, !dbg !1562
Use of instruction is not an instruction!
  %1659 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1658, 1, !dbg !1562
Use of instruction is not an instruction!
  %1663 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1662, 1, !dbg !1562
Use of instruction is not an instruction!
  %1667 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1666, 1, !dbg !1562
Use of instruction is not an instruction!
  %1671 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1670, 1, !dbg !1562
Use of instruction is not an instruction!
  %1675 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1674, 1, !dbg !1562
Use of instruction is not an instruction!
  %1679 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1678, 1, !dbg !1562
Use of instruction is not an instruction!
  %1683 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1682, 1, !dbg !1562
Use of instruction is not an instruction!
  %1687 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1686, 1, !dbg !1562
Use of instruction is not an instruction!
  %1691 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1690, 1, !dbg !1562
Use of instruction is not an instruction!
  %1695 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1694, 1, !dbg !1562
Use of instruction is not an instruction!
  %1699 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1698, 1, !dbg !1562
Use of instruction is not an instruction!
  %1703 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1702, 1, !dbg !1562
Use of instruction is not an instruction!
  %1707 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1706, 1, !dbg !1562
Use of instruction is not an instruction!
  %1711 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1710, 1, !dbg !1562
Use of instruction is not an instruction!
  %1715 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1714, 1, !dbg !1562
Use of instruction is not an instruction!
  %1719 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1718, 1, !dbg !1562
Use of instruction is not an instruction!
  %1723 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1722, 1, !dbg !1562
Use of instruction is not an instruction!
  %1727 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1726, 1, !dbg !1562
Use of instruction is not an instruction!
  %1731 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1730, 1, !dbg !1562
Use of instruction is not an instruction!
  %1735 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1734, 1, !dbg !1562
Use of instruction is not an instruction!
  %1739 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1738, 1, !dbg !1562
Use of instruction is not an instruction!
  %1743 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1742, 1, !dbg !1562
Use of instruction is not an instruction!
  %1747 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1746, 1, !dbg !1562
Use of instruction is not an instruction!
  %1751 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1750, 1, !dbg !1562
Use of instruction is not an instruction!
  %1755 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1754, 1, !dbg !1562
Use of instruction is not an instruction!
  %1759 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1758, 1, !dbg !1562
Use of instruction is not an instruction!
  %1763 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1762, 1, !dbg !1562
Use of instruction is not an instruction!
  %1767 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1766, 1, !dbg !1562
Use of instruction is not an instruction!
  %1771 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1770, 1, !dbg !1562
Use of instruction is not an instruction!
  %1775 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1774, 1, !dbg !1562
Use of instruction is not an instruction!
  %1779 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1778, 1, !dbg !1562
Use of instruction is not an instruction!
  %1783 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1782, 1, !dbg !1562
Use of instruction is not an instruction!
  %1787 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1786, 1, !dbg !1562
Use of instruction is not an instruction!
  %1791 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1790, 1, !dbg !1562
Use of instruction is not an instruction!
  %1795 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1794, 1, !dbg !1562
Use of instruction is not an instruction!
  %1799 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1798, 1, !dbg !1562
Use of instruction is not an instruction!
  %1803 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1802, 1, !dbg !1562
Use of instruction is not an instruction!
  %1807 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1806, 1, !dbg !1562
Use of instruction is not an instruction!
  %1811 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1810, 1, !dbg !1562
Use of instruction is not an instruction!
  %1815 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1814, 1, !dbg !1562
Use of instruction is not an instruction!
  %1819 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1818, 1, !dbg !1562
Use of instruction is not an instruction!
  %1823 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1822, 1, !dbg !1562
Use of instruction is not an instruction!
  %1827 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1826, 1, !dbg !1562
Use of instruction is not an instruction!
  %1831 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1830, 1, !dbg !1562
Use of instruction is not an instruction!
  %1835 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1834, 1, !dbg !1562
Use of instruction is not an instruction!
  %1839 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1838, 1, !dbg !1562
Use of instruction is not an instruction!
  %1843 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1842, 1, !dbg !1562
Use of instruction is not an instruction!
  %1847 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1846, 1, !dbg !1562
Use of instruction is not an instruction!
  %1851 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1850, 1, !dbg !1562
Use of instruction is not an instruction!
  %1855 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1854, 1, !dbg !1562
Use of instruction is not an instruction!
  %1859 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1858, 1, !dbg !1562
Use of instruction is not an instruction!
  %1863 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1862, 1, !dbg !1562
Use of instruction is not an instruction!
  %1867 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1866, 1, !dbg !1562
Use of instruction is not an instruction!
  %1871 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1870, 1, !dbg !1562
Use of instruction is not an instruction!
  %1875 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1874, 1, !dbg !1562
Use of instruction is not an instruction!
  %1879 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1878, 1, !dbg !1562
Use of instruction is not an instruction!
  %1883 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1882, 1, !dbg !1562
Use of instruction is not an instruction!
  %1887 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1886, 1, !dbg !1562
Use of instruction is not an instruction!
  %1891 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:type:struct:~image/color.RGBA:{basic:uint8,basic:uint8,basic:uint8,basic:uint8}" to i32), i8* undef }, i8* %1890, 1, !dbg !1562
error: verification error after interpreting runtime.initAll
</pre>





## image/draw


The compiler gave the following error when this package was imported:

<pre>../../../../../usr/local/go/src/image/image.go:300:12: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:83:12: todo: full slice expressions (with max): []uint8
</pre>





## image/gif


This package cannot be imported because the following dependencies cannot be compiled:

  * [image/color/palette](#image-color-palette)
  * [image/draw](#image-draw)





## image/jpeg


The compiler gave the following error when this package was imported:

<pre>../../../../../usr/local/go/src/image/ycbcr.go:176:20: todo: full slice expressions (with max): []byte
../../../../../usr/local/go/src/image/ycbcr.go:177:20: todo: full slice expressions (with max): []byte
../../../../../usr/local/go/src/image/ycbcr.go:178:20: todo: full slice expressions (with max): []byte
../../../../../usr/local/go/src/image/internal/imageutil/impl.go:86:17: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/internal/imageutil/impl.go:143:17: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/internal/imageutil/impl.go:200:17: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/internal/imageutil/impl.go:256:17: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:83:12: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:99:12: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:111:12: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:865:12: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:881:12: todo: full slice expressions (with max): []uint8
../../../../../usr/local/go/src/image/image.go:893:12: todo: full slice expressions (with max): []uint8
</pre>





## image/png


The compiler gave the following error when this package was imported:

<pre>runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow

runtime stack:
runtime.throw(0x8442e3, 0xe)
	/usr/local/go/src/runtime/panic.go:617 +0x72
runtime.newstack()
	/usr/local/go/src/runtime/stack.go:1041 +0x6f0
runtime.morestack()
	/usr/local/go/src/runtime/asm_amd64.s:429 +0x8f

goroutine 1 [running]:
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0860, 0xd812e0, 0x0, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:60 +0x11b1 fp=0xc02c0003c0 sp=0xc02c0003b8 pc=0x724901
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0a20, 0xc000e5c250, 0x1, 0xc04e101530)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:137 +0xd39 fp=0xc02c000570 sp=0xc02c0003c0 pc=0x724489
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000e43d40, 0x2, 0xc04e132160)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c000720 sp=0xc02c000570 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49140, 0x13, 0xc04e123900)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:117 +0x737 fp=0xc02c0008d0 sp=0xc02c000720 pc=0x723e87
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c000a80 sp=0xc02c0008d0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101520)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c000c30 sp=0xc02c000a80 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e130be0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c000de0 sp=0xc02c000c30 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c000f90 sp=0xc02c000de0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e1014c0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c001140 sp=0xc02c000f90 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e130960, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c0012f0 sp=0xc02c001140 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0014a0 sp=0xc02c0012f0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101460)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c001650 sp=0xc02c0014a0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1306e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c001800 sp=0xc02c001650 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0019b0 sp=0xc02c001800 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101400)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c001b60 sp=0xc02c0019b0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e130460, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c001d10 sp=0xc02c001b60 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c001ec0 sp=0xc02c001d10 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e1013a0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c002070 sp=0xc02c001ec0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1301e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c002220 sp=0xc02c002070 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0023d0 sp=0xc02c002220 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101340)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c002580 sp=0xc02c0023d0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e125f60, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c002730 sp=0xc02c002580 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0028e0 sp=0xc02c002730 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e1012e0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c002a90 sp=0xc02c0028e0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e125ce0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c002c40 sp=0xc02c002a90 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c002df0 sp=0xc02c002c40 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101280)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c002fa0 sp=0xc02c002df0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e125a60, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c003150 sp=0xc02c002fa0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c003300 sp=0xc02c003150 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101220)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0034b0 sp=0xc02c003300 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1257e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c003660 sp=0xc02c0034b0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c003810 sp=0xc02c003660 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e1011c0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0039c0 sp=0xc02c003810 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e125560, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c003b70 sp=0xc02c0039c0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c003d20 sp=0xc02c003b70 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101160)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c003ed0 sp=0xc02c003d20 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1252e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c004080 sp=0xc02c003ed0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c004230 sp=0xc02c004080 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101100)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0043e0 sp=0xc02c004230 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e125060, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c004590 sp=0xc02c0043e0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c004740 sp=0xc02c004590 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e1010a0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0048f0 sp=0xc02c004740 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e124de0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c004aa0 sp=0xc02c0048f0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c004c50 sp=0xc02c004aa0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e101040)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c004e00 sp=0xc02c004c50 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e124b60, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c004fb0 sp=0xc02c004e00 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c005160 sp=0xc02c004fb0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100fe0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c005310 sp=0xc02c005160 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1248e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c0054c0 sp=0xc02c005310 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c005670 sp=0xc02c0054c0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100f80)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c005820 sp=0xc02c005670 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e124660, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c0059d0 sp=0xc02c005820 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c005b80 sp=0xc02c0059d0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100f20)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c005d30 sp=0xc02c005b80 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1243e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c005ee0 sp=0xc02c005d30 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c006090 sp=0xc02c005ee0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100ec0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c006240 sp=0xc02c006090 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e124160, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c0063f0 sp=0xc02c006240 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0065a0 sp=0xc02c0063f0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100e60)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c006750 sp=0xc02c0065a0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e117ee0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c006900 sp=0xc02c006750 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c006ab0 sp=0xc02c006900 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100e00)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c006c60 sp=0xc02c006ab0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e117c60, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c006e10 sp=0xc02c006c60 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c006fc0 sp=0xc02c006e10 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100da0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c007170 sp=0xc02c006fc0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1179e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c007320 sp=0xc02c007170 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0074d0 sp=0xc02c007320 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100d40)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c007680 sp=0xc02c0074d0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e117760, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c007830 sp=0xc02c007680 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c0079e0 sp=0xc02c007830 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100ce0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c007b90 sp=0xc02c0079e0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1174e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c007d40 sp=0xc02c007b90 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c007ef0 sp=0xc02c007d40 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100c80)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0080a0 sp=0xc02c007ef0 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e117260, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c008250 sp=0xc02c0080a0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c008400 sp=0xc02c008250 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100c20)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0085b0 sp=0xc02c008400 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e116fe0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c008760 sp=0xc02c0085b0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c008910 sp=0xc02c008760 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100bc0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c008ac0 sp=0xc02c008910 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e116d60, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c008c70 sp=0xc02c008ac0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c008e20 sp=0xc02c008c70 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100b60)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c008fd0 sp=0xc02c008e20 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e116ae0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c009180 sp=0xc02c008fd0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c009330 sp=0xc02c009180 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100b00)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0094e0 sp=0xc02c009330 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e116860, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c009690 sp=0xc02c0094e0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c009840 sp=0xc02c009690 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100aa0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c0099f0 sp=0xc02c009840 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1165e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c009ba0 sp=0xc02c0099f0 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c009d50 sp=0xc02c009ba0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100a40)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c009f00 sp=0xc02c009d50 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e116360, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c00a0b0 sp=0xc02c009f00 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c00a260 sp=0xc02c00a0b0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e1009e0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c00a410 sp=0xc02c00a260 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e1160e0, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c00a5c0 sp=0xc02c00a410 pc=0x7238e2
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c0960, 0xc000a49320, 0x6, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:145 +0x536 fp=0xc02c00a770 sp=0xc02c00a5c0 pc=0x723c86
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09a0, 0xc0009f84e0, 0x1, 0xc04e100980)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:125 +0xe2b fp=0xc02c00a920 sp=0xc02c00a770 pc=0x72457b
github.com/tinygo-org/tinygo/compiler.getTypeCodeName(0x8c09e0, 0xc000a494d0, 0xc04e10fe60, 0x13)
	/home/ayke/src/github.com/tinygo-org/tinygo/compiler/interface.go:129 +0x192 fp=0xc02c00aad0 sp=0xc02c00a920 pc=0x7238e2
...additional frames elided...

goroutine 5 [syscall]:
os/signal.signal_recv(0x0)
	/usr/local/go/src/runtime/sigqueue.go:139 +0x9c
os/signal.loop()
	/usr/local/go/src/os/signal/signal_unix.go:23 +0x22
created by os/signal.init.0
	/usr/local/go/src/os/signal/signal_unix.go:29 +0x41
</pre>













## log/syslog


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)







## math/big


The compiler gave the following error when this package was imported:

<pre>panic: todo: store

goroutine 1 [running]:
github.com/tinygo-org/tinygo/interp.(*Eval).hasLocalSideEffects(0xc00754fa10, 0xc005ca1b68, 0x7f43c44b3350, 0xdbab00)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:179 +0x1d2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc037e18, 0x7f43cc037e18)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:95 +0x4aa
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc037d58, 0x7f43cc037d58)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc037ff8, 0x7f43cc037ff8)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc038318, 0x7f43cc038318)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc0360c8, 0x7f43cc0360c8)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc0383d8, 0x7f43cc0383d8)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03ac48, 0x7f43cc03ac48)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03c228, 0x7f43cc03c228)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03c128, 0x7f43cc03c128)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03c038, 0x7f43cc03c038)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03bf38, 0x7f43cc03bf38)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03be48, 0x7f43cc03be48)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc03bd48, 0x7f43cc03bd48)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc00754fa10, 0x7f43cc033e08, 0x2)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc005ca3310, 0x7f43c441dce0, 0x7f43c441d910, 0x0, 0x0, 0x0, 0x0, 0xc0024aa840, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:426 +0x4dc4
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc00754fa10, 0x7f43cc036698, 0xc005ca3450, 0x2, 0x2, 0xc004690c20, 0x8, 0x0, 0x0, 0x0, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
github.com/tinygo-org/tinygo/interp.(*Eval).Function(...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:87
github.com/tinygo-org/tinygo/interp.Run(0x11ae080, 0x11a3f90, 0x0, 0x0, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:74 +0x672
main.Compile(0x7ffc2d2c221c, 0x13, 0x7ffc2d2c2200, 0x1b, 0xc0000d8480, 0xc005ca3ee0, 0xc005ca3d48, 0x0, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/main.go:140 +0xc5b
main.Build(0x7ffc2d2c221c, 0x13, 0x7ffc2d2c2200, 0x1b, 0x83fdc6, 0x4, 0xc00008bee0, 0xc0000ea1f0, 0x7f43e4572ba8)
	/home/ayke/src/github.com/tinygo-org/tinygo/main.go:337 +0xe2
main.main()
	/home/ayke/src/github.com/tinygo-org/tinygo/main.go:696 +0xf4a
</pre>











## mime


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/base64](#encoding-base64)





## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [mime](#mime)
  * [net/textproto](#net-textproto)







## net


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)





## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)
  * [crypto/rand](#crypto-rand)
  * [crypto/tls](#crypto-tls)
  * [encoding/base64](#encoding-base64)
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

  * [context](#context)
  * [net](#net)
  * [net/http](#net-http)
  * [net/http/cgi](#net-http-cgi)





## net/http/httptest


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [crypto/x509](#crypto-x509)
  * [flag](#flag)
  * [net](#net)
  * [net/http](#net-http)





## net/http/httptrace


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)
  * [crypto/tls](#crypto-tls)
  * [net](#net)
  * [net/textproto](#net-textproto)





## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)
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

  * [crypto/md5](#crypto-md5)
  * [crypto/tls](#crypto-tls)
  * [encoding/base64](#encoding-base64)
  * [net](#net)
  * [net/textproto](#net-textproto)





## net/textproto


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)









## os/exec


This package cannot be imported because the following dependencies cannot be compiled:

  * [context](#context)





## os/signal


The compiler gave the following error when this package was imported:

<pre># os/signal
../../../../../usr/local/go/src/os/signal/signal.go:15:18: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:27:14: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:50:23: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:85:23: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:90:21: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:111:25: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:111:43: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:155:22: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:162:23: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:214:21: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal_unix.go:36:20: Signal not declared by package os
../../../../../usr/local/go/src/os/signal/signal.go:122:36: Signal not declared by package os
</pre>





## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
../../../../../usr/local/go/src/os/user/lookup.go:62:9: undeclared name: listGroups
../../../../../usr/local/go/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
../../../../../usr/local/go/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
../../../../../usr/local/go/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
../../../../../usr/local/go/src/os/user/lookup.go:36:9: undeclared name: lookupUser
../../../../../usr/local/go/src/os/user/lookup.go:15:41: undeclared name: current
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
../../../../../usr/local/go/src/text/template/parse/parse.go:193:26: Error not declared by package runtime
</pre>














