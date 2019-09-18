
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
image/draw |  <span style="color: green">✔</span> yes  | 
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

  * [database/sql/driver](#database-sql-driver)





## database/sql/driver


The compiler gave the following error when this package was imported:

<pre># database/sql/driver
../../../../../usr/local/go/src/database/sql/driver/types.go:227:20: rv.Type().Elem().Implements undefined (type reflect.Type has no field or method Implements)
</pre>







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
[...more lines following...]</pre>







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
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:117:15: f.Tag.Get undefined (type string has no field or method Get)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:168:24: f.Tag.Get undefined (type string has no field or method Get)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:175:23: f.Tag.Get undefined (type string has no field or method Get)
../../../../../usr/local/go/src/encoding/xml/typeinfo.go:63:50: f.Tag.Get undefined (type string has no field or method Get)
../../../../../usr/local/go/src/encoding/xml/read.go:665:7: dst.SetBytes undefined (type reflect.Value has no field or method SetBytes)
../../../../../usr/local/go/src/encoding/xml/read.go:339:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:347:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:352:38: val.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:358:37: pv.Type().Implements undefined (type reflect.Type has no field or method Implements)
../../../../../usr/local/go/src/encoding/xml/read.go:398:17: Append not declared by package reflect
[...more lines following...]</pre>







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
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc002538798, 0x7fd5b1033e90, 0x7fd5b1033b10, 0x847752, 0x4, 0x0, 0x0, 0xc0034d38f8, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:279 +0x9174
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc002512780, 0x7fd5b0031e58, 0xc003500960, 0x2, 0x2, 0xc0034f4e50, 0x8, 0x847752, 0x4, 0x8da760, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc0025392b8, 0x7fd5b1032970, 0x7fd5b1032380, 0x0, 0x0, 0x0, 0x0, 0xc0034d38c0, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:449 +0x51b3
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc002512780, 0x7fd5b0031d68, 0xc0025393f8, 0x2, 0x2, 0xc0034f4e50, 0x8, 0x0, 0x0, 0x0, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
github.com/tinygo-org/tinygo/interp.(*Eval).Function(...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:87
github.com/tinygo-org/tinygo/interp.Run(0x2acf080, 0x2ac4f90, 0x0, 0x0, 0x0)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:74 +0x672
[...more lines following...]</pre>





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
  %7 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %6, 1, !dbg !1580
Use of instruction is not an instruction!
  %11 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %10, 1, !dbg !1580
Use of instruction is not an instruction!
  %15 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %14, 1, !dbg !1580
Use of instruction is not an instruction!
  %19 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %18, 1, !dbg !1580
Use of instruction is not an instruction!
  %23 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %22, 1, !dbg !1580
Use of instruction is not an instruction!
  %27 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %26, 1, !dbg !1580
Use of instruction is not an instruction!
  %31 = insertvalue %runtime._interface { i32 ptrtoint (%runtime.typeInInterface* @"typeInInterface:reflect/types.type:named:image/color.RGBA" to i32), i8* undef }, i8* %30, 1, !dbg !1580
Use of instruction is not an instruction!
[...more lines following...]</pre>







## image/gif


This package cannot be imported because the following dependencies cannot be compiled:

  * [image/color/palette](#image-color-palette)





## image/jpeg


The compiler gave the following error when this package was imported:

<pre>panic: todo: store

goroutine 1 [running]:
github.com/tinygo-org/tinygo/interp.(*Eval).hasLocalSideEffects(0xc0023f5e90, 0xc002694560, 0x7fece019a910, 0xc004304200)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:179 +0x1d2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc0023f5e90, 0x7fecec017cd8, 0x7fecec017cd8)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:95 +0x4aa
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc0023f5e90, 0x7fecec020a88, 0x7fecec020a88)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc0023f5e90, 0x7fecec021518, 0x1)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc0026952b8, 0x7fecefeb3060, 0x7fecefeb2cc0, 0x0, 0x0, 0x0, 0x0, 0xc003094f90, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:426 +0x4dc4
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc0023f5e90, 0x7fecec021428, 0xc0026953f8, 0x2, 0x2, 0xc005452640, 0xa, 0x0, 0x0, 0x0, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
[...more lines following...]</pre>





## image/png


The compiler gave the following error when this package was imported:

<pre>panic: todo: store

goroutine 1 [running]:
github.com/tinygo-org/tinygo/interp.(*Eval).hasLocalSideEffects(0xc004288720, 0xc006988560, 0x7f672cb55190, 0xc0043f6600)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:179 +0x1d2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004288720, 0x7f67380598f8, 0x7f67380598f8)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:95 +0x4aa
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004288720, 0x7f673803c9c8, 0x7f673803c9c8)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004288720, 0x7f673803d348, 0x1)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*frame).evalBasicBlock(0xc0069892b8, 0x7f672c394330, 0x7f672c393f30, 0x0, 0x0, 0x0, 0x0, 0xc006f738a8, 0x1, 0x1, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/frame.go:426 +0x4dc4
github.com/tinygo-org/tinygo/interp.(*Eval).function(0xc004288720, 0x7f673803d258, 0xc0069893f8, 0x2, 0x2, 0xc006f92650, 0x9, 0x0, 0x0, 0x0, ...)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/interp.go:104 +0x208
[...more lines following...]</pre>













## log/syslog


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)







## math/big


The compiler gave the following error when this package was imported:

<pre>panic: todo: store

goroutine 1 [running]:
github.com/tinygo-org/tinygo/interp.(*Eval).hasLocalSideEffects(0xc004ebea20, 0xc0035f1b10, 0x32162a0, 0xdc6c00)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:179 +0x1d2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004ebea20, 0x7f5c5003a158, 0x7f5c5003a158)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:95 +0x4aa
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004ebea20, 0x7f5c5003a098, 0x7f5c5003a098)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004ebea20, 0x7f5c5003a338, 0x7f5c5003a338)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004ebea20, 0x7f5c5003a658, 0x7f5c5003a658)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
github.com/tinygo-org/tinygo/interp.(*Eval).hasSideEffects(0xc004ebea20, 0x7f5c50037358, 0x7f5c50037358)
	/home/ayke/src/github.com/tinygo-org/tinygo/interp/scan.go:100 +0x4f2
[...more lines following...]</pre>











## mime


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/base64](#encoding-base64)





## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [mime](#mime)
  * [net/textproto](#net-textproto)







## net


The compiler gave the following error when this package was imported:

<pre># internal/singleflight
../../../../../usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
</pre>





## net/http


This package cannot be imported because the following dependencies cannot be compiled:

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
  * [encoding/base64](#encoding-base64)
  * [net](#net)
  * [net/textproto](#net-textproto)





## net/textproto


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)









## os/exec


The compiler gave the following error when this package was imported:

<pre># os/exec
../../../../../usr/local/go/src/os/exec/exec.go:125:14: Process not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:129:19: ProcessState not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:429:6: ProcessState not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:618:20: Pipe not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:593:20: Pipe not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:551:20: Pipe not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:392:22: StartProcess not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:392:57: ProcAttr not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:268:27: DevNull not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:280:20: Pipe not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:223:23: DevNull not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:235:20: Pipe not declared by package os
../../../../../usr/local/go/src/os/exec/exec.go:206:12: Environ not declared by package os
</pre>





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














