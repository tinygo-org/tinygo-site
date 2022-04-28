
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Package | Importable
--- | --- |
archive/tar |  <span style="color: green">✔</span> yes  | 
archive/zip |  <span style="color: green">✔</span> yes  | 
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
crypto/ecdsa |  <span style="color: green">✔</span> yes  | 
crypto/ed25519 |  <span style="color: green">✔</span> yes  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  | 
crypto/rand |  <span style="color: green">✔</span> yes  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  <span style="color: green">✔</span> yes  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  | 
crypto/subtle |  <span style="color: green">✔</span> yes  | 
crypto/tls |  <span style="color: green">✔</span> yes  | 
crypto/x509 |  <span style="color: green">✔</span> yes  | 
crypto/x509/pkix |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  | 
database/sql/driver |  <span style="color: green">✔</span> yes  | 
debug/buildinfo |  [<span style="color: red">✗</span> no](#debugbuildinfo)  | 
debug/dwarf |  <span style="color: green">✔</span> yes  | 
debug/elf |  <span style="color: green">✔</span> yes  | 
debug/gosym |  <span style="color: green">✔</span> yes  | 
debug/macho |  <span style="color: green">✔</span> yes  | 
debug/pe |  <span style="color: green">✔</span> yes  | 
debug/plan9obj |  <span style="color: green">✔</span> yes  | 
embed |  <span style="color: green">✔</span> yes  | 
encoding |  <span style="color: green">✔</span> yes  | 
encoding/ascii85 |  <span style="color: green">✔</span> yes  | 
encoding/asn1 |  <span style="color: green">✔</span> yes  | 
encoding/base32 |  <span style="color: green">✔</span> yes  | 
encoding/base64 |  <span style="color: green">✔</span> yes  | 
encoding/binary |  <span style="color: green">✔</span> yes  | 
encoding/csv |  <span style="color: green">✔</span> yes  | 
encoding/gob |  [<span style="color: red">✗</span> no](#encodinggob)  | 
encoding/hex |  <span style="color: green">✔</span> yes  | 
encoding/json |  <span style="color: green">✔</span> yes  | 
encoding/pem |  <span style="color: green">✔</span> yes  | 
encoding/xml |  [<span style="color: red">✗</span> no](#encodingxml)  | 
errors |  <span style="color: green">✔</span> yes  | 
expvar |  <span style="color: green">✔</span> yes  | 
flag |  <span style="color: green">✔</span> yes  | 
fmt |  <span style="color: green">✔</span> yes  | 
go/ast |  <span style="color: green">✔</span> yes  | 
go/build |  <span style="color: green">✔</span> yes  | 
go/build/constraint |  <span style="color: green">✔</span> yes  | 
go/constant |  <span style="color: green">✔</span> yes  | 
go/doc |  <span style="color: green">✔</span> yes  | 
go/format |  <span style="color: green">✔</span> yes  | 
go/importer |  <span style="color: green">✔</span> yes  | 
go/parser |  <span style="color: green">✔</span> yes  | 
go/printer |  <span style="color: green">✔</span> yes  | 
go/scanner |  <span style="color: green">✔</span> yes  | 
go/token |  <span style="color: green">✔</span> yes  | 
go/types |  <span style="color: green">✔</span> yes  | 
hash |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  | 
hash/crc64 |  <span style="color: green">✔</span> yes  | 
hash/fnv |  <span style="color: green">✔</span> yes  | 
hash/maphash |  <span style="color: green">✔</span> yes  | 
html |  <span style="color: green">✔</span> yes  | 
html/template |  <span style="color: green">✔</span> yes  | 
image |  <span style="color: green">✔</span> yes  | 
image/color |  <span style="color: green">✔</span> yes  | 
image/color/palette |  <span style="color: green">✔</span> yes  | 
image/draw |  <span style="color: green">✔</span> yes  | 
image/gif |  <span style="color: green">✔</span> yes  | 
image/jpeg |  <span style="color: green">✔</span> yes  | 
image/png |  <span style="color: green">✔</span> yes  | 
index/suffixarray |  <span style="color: green">✔</span> yes  | 
io |  <span style="color: green">✔</span> yes  | 
io/fs |  <span style="color: green">✔</span> yes  | 
io/ioutil |  <span style="color: green">✔</span> yes  | 
log |  <span style="color: green">✔</span> yes  | 
log/syslog |  <span style="color: green">✔</span> yes  | 
math |  <span style="color: green">✔</span> yes  | 
math/big |  <span style="color: green">✔</span> yes  | 
math/bits |  <span style="color: green">✔</span> yes  | 
math/cmplx |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  | 
mime |  <span style="color: green">✔</span> yes  | 
mime/multipart |  <span style="color: green">✔</span> yes  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  | 
net |  <span style="color: green">✔</span> yes  | 
net/http |  <span style="color: green">✔</span> yes  | 
net/http/cgi |  <span style="color: green">✔</span> yes  | 
net/http/cookiejar |  <span style="color: green">✔</span> yes  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#nethttpfcgi)  | 
net/http/httptest |  <span style="color: green">✔</span> yes  | 
net/http/httptrace |  <span style="color: green">✔</span> yes  | 
net/http/httputil |  <span style="color: green">✔</span> yes  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#nethttppprof)  | 
net/mail |  <span style="color: green">✔</span> yes  | 
net/netip |  <span style="color: green">✔</span> yes  | 
net/rpc |  [<span style="color: red">✗</span> no](#netrpc)  | 
net/rpc/jsonrpc |  [<span style="color: red">✗</span> no](#netrpcjsonrpc)  | 
net/smtp |  <span style="color: green">✔</span> yes  | 
net/textproto |  <span style="color: green">✔</span> yes  | 
net/url |  <span style="color: green">✔</span> yes  | 
os |  <span style="color: green">✔</span> yes  | 
os/exec |  <span style="color: green">✔</span> yes  | 
os/signal |  <span style="color: green">✔</span> yes  | 
os/user |  [<span style="color: red">✗</span> no](#osuser)  | 
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
testing/fstest |  <span style="color: green">✔</span> yes  | 
testing/iotest |  <span style="color: green">✔</span> yes  | 
testing/quick |  <span style="color: green">✔</span> yes  | 
text/scanner |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  | 
text/template |  <span style="color: green">✔</span> yes  | 
text/template/parse |  <span style="color: green">✔</span> yes  | 
time |  <span style="color: green">✔</span> yes  | 
time/tzdata |  <span style="color: green">✔</span> yes  | 
unicode |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  | 









































































## debug/buildinfo


The compiler gave the following error when this package was imported:

<pre># debug/buildinfo
/home/ron/.gvm/gos/go1.18.1/src/debug/buildinfo/buildinfo.go:78:19: ParseBuildInfo not declared by package debug
</pre>

































## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
/home/ron/.gvm/gos/go1.18.1/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
</pre>











## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
/home/ron/.gvm/gos/go1.18.1/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
/home/ron/.gvm/gos/go1.18.1/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
</pre>







































































































## net/http/fcgi


The compiler gave the following error when this package was imported:

<pre># net/http/fcgi
/home/ron/.gvm/gos/go1.18.1/src/net/http/fcgi/child.go:342:16: FileListener not declared by package net
</pre>











## net/http/pprof


The compiler gave the following error when this package was imported:

<pre>../../../.cache/tinygo/goroot-823f493b0dc9a9f0e80860fcc84322ec84e89131fcbbacca15b0be4a8879a3cb/src/net/http/pprof/pprof.go:73:2: package runtime/trace is not in GOROOT (/home/ron/.cache/tinygo/goroot-823f493b0dc9a9f0e80860fcc84322ec84e89131fcbbacca15b0be4a8879a3cb/src/runtime/trace)
</pre>









## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/gob](#encodinggob)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/rpc](#netrpc)

















## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
/home/ron/.gvm/gos/go1.18.1/src/os/user/lookup.go:15:41: undeclared name: current
/home/ron/.gvm/gos/go1.18.1/src/os/user/lookup.go:36:9: undeclared name: lookupUser
/home/ron/.gvm/gos/go1.18.1/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
/home/ron/.gvm/gos/go1.18.1/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
/home/ron/.gvm/gos/go1.18.1/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
/home/ron/.gvm/gos/go1.18.1/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>


























































