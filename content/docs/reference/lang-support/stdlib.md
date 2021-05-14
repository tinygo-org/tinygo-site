
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
crypto/ed25519 |  [<span style="color: red">✗</span> no](#crypto-ed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  | 
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
crypto/x509/pkix |  <span style="color: green">✔</span> yes  | 
database/sql |  <span style="color: green">✔</span> yes  | 
database/sql/driver |  <span style="color: green">✔</span> yes  | 
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
encoding/gob |  [<span style="color: red">✗</span> no](#encoding-gob)  | 
encoding/hex |  <span style="color: green">✔</span> yes  | 
encoding/json |  <span style="color: green">✔</span> yes  | 
encoding/pem |  <span style="color: green">✔</span> yes  | 
encoding/xml |  [<span style="color: red">✗</span> no](#encoding-xml)  | 
errors |  <span style="color: green">✔</span> yes  | 
expvar |  [<span style="color: red">✗</span> no](#expvar)  | 
flag |  <span style="color: green">✔</span> yes  | 
fmt |  <span style="color: green">✔</span> yes  | 
go/ast |  <span style="color: green">✔</span> yes  | 
go/build |  [<span style="color: red">✗</span> no](#go-build)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  | 
go/constant |  <span style="color: green">✔</span> yes  | 
go/doc |  [<span style="color: red">✗</span> no](#go-doc)  | 
go/format |  [<span style="color: red">✗</span> no](#go-format)  | 
go/importer |  [<span style="color: red">✗</span> no](#go-importer)  | 
go/parser |  [<span style="color: red">✗</span> no](#go-parser)  | 
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
io/fs |  <span style="color: green">✔</span> yes  | 
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
net |  <span style="color: green">✔</span> yes  | 
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
net/url |  [<span style="color: red">✗</span> no](#net-url)  | 
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
testing/fstest |  <span style="color: green">✔</span> yes  | 
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









































## crypto/ed25519


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)











## crypto/rand


The compiler gave the following error when this package was imported:

<pre>fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x4a1ed8e]

runtime stack:
runtime.throw(0x4ea9aba, 0x2a)
	/usr/local/go/src/runtime/panic.go:1116 +0x72
runtime.sigpanic()
	/usr/local/go/src/runtime/signal_unix.go:701 +0x46a

goroutine 10 [syscall]:
runtime.cgocall(0xea6d70, 0xc00435b990, 0xc0025d4508)
	/usr/local/go/src/runtime/cgocall.go:133 +0x5b fp=0xc00435b960 sp=0xc00435b928 pc=0xad570b
tinygo.org/x/go-llvm._Cfunc_LLVMVerifyModule(0x7f1f2801c830, 0xc000000001, 0xc0025d4508, 0x0)
	_cgo_gotypes.go:9557 +0x4d fp=0xc00435b990 sp=0xc00435b960 pc=0xd7fb5d
tinygo.org/x/go-llvm.VerifyModule.func1(0x7f1f2801c830, 0xc000000001, 0xc0025d4508, 0x1)
[...more lines following...]</pre>







## crypto/rsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)













## crypto/tls


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#crypto-ed25519)
  * [crypto/rand](#crypto-rand)
  * [crypto/rsa](#crypto-rsa)
  * [crypto/x509](#crypto-x509)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#crypto-ed25519)
  * [crypto/rsa](#crypto-rsa)
  * [net/url](#net-url)







































## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
/opt/go-1.16.4/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/opt/go-1.16.4/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/opt/go-1.16.4/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
/opt/go-1.16.4/src/encoding/gob/type.go:867:9: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
/opt/go-1.16.4/src/encoding/gob/type.go:870:21: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
</pre>











## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
/opt/go-1.16.4/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
/opt/go-1.16.4/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
/opt/go-1.16.4/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
</pre>







## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#net-http)











## go/build


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/doc](#go-doc)
  * [go/parser](#go-parser)









## go/doc


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#text-template)





## go/format


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/parser](#go-parser)





## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/build](#go-build)
  * [go/types](#go-types)





## go/parser


The compiler gave the following error when this package was imported:

<pre># go/parser
/opt/go-1.16.4/src/go/parser/interface.go:136:18: ReadDir not declared by package os
</pre>











## go/types


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/parser](#go-parser)



















## html/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#text-template)





























## log/syslog


The compiler gave the following error when this package was imported:

<pre># log/syslog
/opt/go-1.16.4/src/log/syslog/syslog.go:100:12: Conn not declared by package net
/opt/go-1.16.4/src/log/syslog/syslog.go:161:13: Conn not declared by package net
/opt/go-1.16.4/src/log/syslog/syslog.go:162:16: Dial not declared by package net
/opt/go-1.16.4/src/log/syslog/syslog_unix.go:22:21: Dial not declared by package net
</pre>

















## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [net/textproto](#net-textproto)









## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#crypto-rand)
  * [crypto/tls](#crypto-tls)
  * [mime/multipart](#mime-multipart)
  * [net/http/httptrace](#net-http-httptrace)
  * [net/textproto](#net-textproto)
  * [net/url](#net-url)





## net/http/cgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net/http](#net-http)
  * [net/textproto](#net-textproto)
  * [net/url](#net-url)
  * [os/exec](#os-exec)





## net/http/cookiejar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#net-http)
  * [net/url](#net-url)





## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#net-http)
  * [net/http/cgi](#net-http-cgi)





## net/http/httptest


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [crypto/x509](#crypto-x509)
  * [net/http](#net-http)
  * [net/textproto](#net-textproto)





## net/http/httptrace


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net/textproto](#net-textproto)





## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#net-http)
  * [net/textproto](#net-textproto)
  * [net/url](#net-url)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#net-http)
  * [net/url](#net-url)





## net/mail


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/textproto](#net-textproto)





## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/gob](#encoding-gob)
  * [html/template](#html-template)
  * [net/http](#net-http)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/rpc](#net-rpc)





## net/smtp


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#crypto-tls)
  * [net/textproto](#net-textproto)





## net/textproto


The compiler gave the following error when this package was imported:

<pre>/tmp/tinygo-test-224097129/main.go:2:8: package net/textproto is not in GOROOT (/home/fgsch/.cache/tinygo/goroot-go1.16.4-1fabd5e82c0fa87fc2690e2236a57bac63173c5dadf42e03c38f73017a16208e/src/net/textproto)
</pre>





## net/url


The compiler gave the following error when this package was imported:

<pre>/tmp/tinygo-test-420930484/main.go:2:8: package net/url is not in GOROOT (/home/fgsch/.cache/tinygo/goroot-go1.16.4-1fabd5e82c0fa87fc2690e2236a57bac63173c5dadf42e03c38f73017a16208e/src/net/url)
</pre>







## os/exec


The compiler gave the following error when this package was imported:

<pre># os/exec
/opt/go-1.16.4/src/os/exec/exec.go:130:14: Process not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:134:19: ProcessState not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:462:6: ProcessState not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:246:23: DevNull not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:258:20: Pipe not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:291:27: DevNull not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:303:20: Pipe not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:422:22: StartProcess not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:422:57: ProcAttr not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:584:20: Pipe not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:626:20: Pipe not declared by package os
/opt/go-1.16.4/src/os/exec/exec.go:651:20: Pipe not declared by package os
</pre>







## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
/opt/go-1.16.4/src/os/user/lookup.go:15:41: undeclared name: current
/opt/go-1.16.4/src/os/user/lookup.go:36:9: undeclared name: lookupUser
/opt/go-1.16.4/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
/opt/go-1.16.4/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
/opt/go-1.16.4/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
/opt/go-1.16.4/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>





































## testing/quick


The compiler gave the following error when this package was imported:

<pre># testing/quick
/opt/go-1.16.4/src/testing/quick/quick.go:273:11: fType.NumOut undefined (type reflect.Type has no field or method NumOut)
/opt/go-1.16.4/src/testing/quick/quick.go:276:11: fType.Out undefined (type reflect.Type has no field or method Out)
/opt/go-1.16.4/src/testing/quick/quick.go:280:43: fType.NumIn undefined (type reflect.Type has no field or method NumIn)
/opt/go-1.16.4/src/testing/quick/quick.go:290:12: fVal.Call undefined (type reflect.Value has no field or method Call)
/opt/go-1.16.4/src/testing/quick/quick.go:320:43: xType.NumIn undefined (type reflect.Type has no field or method NumIn)
/opt/go-1.16.4/src/testing/quick/quick.go:330:26: x.Call undefined (type reflect.Value has no field or method Call)
/opt/go-1.16.4/src/testing/quick/quick.go:331:26: y.Call undefined (type reflect.Value has no field or method Call)
/opt/go-1.16.4/src/testing/quick/quick.go:351:25: f.In undefined (type reflect.Type has no field or method In)
/opt/go-1.16.4/src/testing/quick/quick.go:353:95: f.In undefined (type reflect.Type has no field or method In)
</pre>









## text/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/url](#net-url)


















