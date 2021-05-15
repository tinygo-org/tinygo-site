
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
crypto/ed25519 |  [<span style="color: red">✗</span> no](#cryptoed25519)  | 
crypto/elliptic |  <span style="color: green">✔</span> yes  | 
crypto/hmac |  <span style="color: green">✔</span> yes  | 
crypto/md5 |  <span style="color: green">✔</span> yes  | 
crypto/rand |  [<span style="color: red">✗</span> no](#cryptorand)  | 
crypto/rc4 |  <span style="color: green">✔</span> yes  | 
crypto/rsa |  [<span style="color: red">✗</span> no](#cryptorsa)  | 
crypto/sha1 |  <span style="color: green">✔</span> yes  | 
crypto/sha256 |  <span style="color: green">✔</span> yes  | 
crypto/sha512 |  <span style="color: green">✔</span> yes  | 
crypto/subtle |  <span style="color: green">✔</span> yes  | 
crypto/tls |  [<span style="color: red">✗</span> no](#cryptotls)  | 
crypto/x509 |  [<span style="color: red">✗</span> no](#cryptox509)  | 
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
encoding/gob |  [<span style="color: red">✗</span> no](#encodinggob)  | 
encoding/hex |  <span style="color: green">✔</span> yes  | 
encoding/json |  <span style="color: green">✔</span> yes  | 
encoding/pem |  <span style="color: green">✔</span> yes  | 
encoding/xml |  [<span style="color: red">✗</span> no](#encodingxml)  | 
errors |  <span style="color: green">✔</span> yes  | 
expvar |  [<span style="color: red">✗</span> no](#expvar)  | 
flag |  <span style="color: green">✔</span> yes  | 
fmt |  <span style="color: green">✔</span> yes  | 
go/ast |  <span style="color: green">✔</span> yes  | 
go/build |  [<span style="color: red">✗</span> no](#gobuild)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  | 
go/constant |  <span style="color: green">✔</span> yes  | 
go/doc |  [<span style="color: red">✗</span> no](#godoc)  | 
go/format |  [<span style="color: red">✗</span> no](#goformat)  | 
go/importer |  [<span style="color: red">✗</span> no](#goimporter)  | 
go/parser |  [<span style="color: red">✗</span> no](#goparser)  | 
go/printer |  <span style="color: green">✔</span> yes  | 
go/scanner |  <span style="color: green">✔</span> yes  | 
go/token |  <span style="color: green">✔</span> yes  | 
go/types |  [<span style="color: red">✗</span> no](#gotypes)  | 
hash |  <span style="color: green">✔</span> yes  | 
hash/adler32 |  <span style="color: green">✔</span> yes  | 
hash/crc32 |  <span style="color: green">✔</span> yes  | 
hash/crc64 |  <span style="color: green">✔</span> yes  | 
hash/fnv |  <span style="color: green">✔</span> yes  | 
hash/maphash |  <span style="color: green">✔</span> yes  | 
html |  <span style="color: green">✔</span> yes  | 
html/template |  [<span style="color: red">✗</span> no](#htmltemplate)  | 
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
log/syslog |  [<span style="color: red">✗</span> no](#logsyslog)  | 
math |  <span style="color: green">✔</span> yes  | 
math/big |  <span style="color: green">✔</span> yes  | 
math/bits |  <span style="color: green">✔</span> yes  | 
math/cmplx |  <span style="color: green">✔</span> yes  | 
math/rand |  <span style="color: green">✔</span> yes  | 
mime |  <span style="color: green">✔</span> yes  | 
mime/multipart |  [<span style="color: red">✗</span> no](#mimemultipart)  | 
mime/quotedprintable |  <span style="color: green">✔</span> yes  | 
net |  [<span style="color: red">✗</span> no](#net)  | 
net/http |  [<span style="color: red">✗</span> no](#nethttp)  | 
net/http/cgi |  [<span style="color: red">✗</span> no](#nethttpcgi)  | 
net/http/cookiejar |  [<span style="color: red">✗</span> no](#nethttpcookiejar)  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#nethttpfcgi)  | 
net/http/httptest |  [<span style="color: red">✗</span> no](#nethttphttptest)  | 
net/http/httptrace |  [<span style="color: red">✗</span> no](#nethttphttptrace)  | 
net/http/httputil |  [<span style="color: red">✗</span> no](#nethttphttputil)  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#nethttppprof)  | 
net/mail |  [<span style="color: red">✗</span> no](#netmail)  | 
net/rpc |  [<span style="color: red">✗</span> no](#netrpc)  | 
net/rpc/jsonrpc |  [<span style="color: red">✗</span> no](#netrpcjsonrpc)  | 
net/smtp |  [<span style="color: red">✗</span> no](#netsmtp)  | 
net/textproto |  [<span style="color: red">✗</span> no](#nettextproto)  | 
net/url |  <span style="color: green">✔</span> yes  | 
os |  <span style="color: green">✔</span> yes  | 
os/exec |  [<span style="color: red">✗</span> no](#osexec)  | 
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
testing/quick |  [<span style="color: red">✗</span> no](#testingquick)  | 
text/scanner |  <span style="color: green">✔</span> yes  | 
text/tabwriter |  <span style="color: green">✔</span> yes  | 
text/template |  [<span style="color: red">✗</span> no](#texttemplate)  | 
text/template/parse |  <span style="color: green">✔</span> yes  | 
time |  <span style="color: green">✔</span> yes  | 
time/tzdata |  <span style="color: green">✔</span> yes  | 
unicode |  <span style="color: green">✔</span> yes  | 
unicode/utf16 |  <span style="color: green">✔</span> yes  | 
unicode/utf8 |  <span style="color: green">✔</span> yes  | 
unsafe |  <span style="color: green">✔</span> yes  | 









































## crypto/ed25519


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#cryptorand)











## crypto/rand


The compiler gave the following error when this package was imported:

<pre>fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x4c3675e]

runtime stack:
runtime.throw(0x50dc149, 0x2a)
	/home/ron/.gvm/gos/go1.16.3/src/runtime/panic.go:1117 +0x72
runtime.sigpanic()
	/home/ron/.gvm/gos/go1.16.3/src/runtime/signal_unix.go:718 +0x2e5

goroutine 21 [syscall]:
runtime.cgocall(0xdd6800, 0xc003d939c0, 0xc0038589a0)
	/home/ron/.gvm/gos/go1.16.3/src/runtime/cgocall.go:154 +0x5b fp=0xc003d93990 sp=0xc003d93958 pc=0xa7db3b
tinygo.org/x/go-llvm._Cfunc_LLVMVerifyModule(0x7f434801c7f0, 0xc000000001, 0xc0038589a0, 0x0)
	_cgo_gotypes.go:9591 +0x48 fp=0xc003d939c0 sp=0xc003d93990 pc=0xcd4168
tinygo.org/x/go-llvm.VerifyModule.func1(0x7f434801c7f0, 0xc000000001, 0xc0038589a0, 0x7f4349080d58)
[...more lines following...]</pre>







## crypto/rsa


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#cryptorand)













## crypto/tls


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#cryptoed25519)
  * [crypto/rand](#cryptorand)
  * [crypto/rsa](#cryptorsa)
  * [crypto/x509](#cryptox509)
  * [net](#net)





## crypto/x509


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/ed25519](#cryptoed25519)
  * [crypto/rsa](#cryptorsa)
  * [net](#net)







































## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
/home/ron/.gvm/gos/go1.16.3/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/home/ron/.gvm/gos/go1.16.3/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.16.3/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.16.3/src/encoding/gob/type.go:867:9: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
/home/ron/.gvm/gos/go1.16.3/src/encoding/gob/type.go:870:21: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
</pre>











## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
/home/ron/.gvm/gos/go1.16.3/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.16.3/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
/home/ron/.gvm/gos/go1.16.3/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
</pre>







## expvar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)











## go/build


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/doc](#godoc)
  * [go/parser](#goparser)









## go/doc


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#texttemplate)





## go/format


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/parser](#goparser)





## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/build](#gobuild)
  * [go/types](#gotypes)





## go/parser


The compiler gave the following error when this package was imported:

<pre># go/parser
/home/ron/.gvm/gos/go1.16.3/src/go/parser/interface.go:136:18: ReadDir not declared by package os
</pre>











## go/types


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/parser](#goparser)



















## html/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#texttemplate)





























## log/syslog


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)

















## mime/multipart


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#cryptorand)
  * [net/textproto](#nettextproto)







## net


The compiler gave the following error when this package was imported:

<pre># net
/home/ron/.gvm/gos/go1.16.3/src/net/pipe.go:156:16: ErrDeadlineExceeded not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/net/pipe.go:169:16: ErrDeadlineExceeded not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/net/pipe.go:188:16: ErrDeadlineExceeded not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/net/pipe.go:204:17: ErrDeadlineExceeded not declared by package os
</pre>





## net/http


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/rand](#cryptorand)
  * [crypto/tls](#cryptotls)
  * [mime/multipart](#mimemultipart)
  * [net](#net)
  * [net/http/httptrace](#nethttphttptrace)
  * [net/textproto](#nettextproto)





## net/http/cgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [net](#net)
  * [net/http](#nethttp)
  * [net/textproto](#nettextproto)
  * [os/exec](#osexec)





## net/http/cookiejar


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/http](#nethttp)





## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/http](#nethttp)
  * [net/http/cgi](#nethttpcgi)





## net/http/httptest


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [crypto/x509](#cryptox509)
  * [net](#net)
  * [net/http](#nethttp)
  * [net/textproto](#nettextproto)





## net/http/httptrace


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [net](#net)
  * [net/textproto](#nettextproto)





## net/http/httputil


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/http](#nethttp)
  * [net/textproto](#nettextproto)





## net/http/pprof


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http](#nethttp)





## net/mail


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/textproto](#nettextproto)





## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/gob](#encodinggob)
  * [html/template](#htmltemplate)
  * [net](#net)
  * [net/http](#nethttp)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)
  * [net/rpc](#netrpc)





## net/smtp


This package cannot be imported because the following dependencies cannot be compiled:

  * [crypto/tls](#cryptotls)
  * [net](#net)
  * [net/textproto](#nettextproto)





## net/textproto


This package cannot be imported because the following dependencies cannot be compiled:

  * [net](#net)









## os/exec


The compiler gave the following error when this package was imported:

<pre># os/exec
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:130:14: Process not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:134:19: ProcessState not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:462:6: ProcessState not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:246:23: DevNull not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:258:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:291:27: DevNull not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:303:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:422:22: StartProcess not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:422:57: ProcAttr not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:584:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:626:20: Pipe not declared by package os
/home/ron/.gvm/gos/go1.16.3/src/os/exec/exec.go:651:20: Pipe not declared by package os
</pre>







## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
/home/ron/.gvm/gos/go1.16.3/src/os/user/lookup.go:15:41: undeclared name: current
/home/ron/.gvm/gos/go1.16.3/src/os/user/lookup.go:36:9: undeclared name: lookupUser
/home/ron/.gvm/gos/go1.16.3/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
/home/ron/.gvm/gos/go1.16.3/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
/home/ron/.gvm/gos/go1.16.3/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
/home/ron/.gvm/gos/go1.16.3/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>





































## testing/quick


The compiler gave the following error when this package was imported:

<pre># testing/quick
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:273:11: fType.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:276:11: fType.Out undefined (type reflect.Type has no field or method Out)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:280:43: fType.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:290:12: fVal.Call undefined (type reflect.Value has no field or method Call)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:320:43: xType.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:330:26: x.Call undefined (type reflect.Value has no field or method Call)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:331:26: y.Call undefined (type reflect.Value has no field or method Call)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:351:25: f.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.16.3/src/testing/quick/quick.go:353:95: f.In undefined (type reflect.Type has no field or method In)
</pre>









## text/template


The compiler gave the following error when this package was imported:

<pre># text/template
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:376:17: val.Type().ChanDir undefined (type reflect.Type has no field or method ChanDir)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:376:38: SendDir not declared by package reflect
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:382:20: val.Recv undefined (type reflect.Value has no field or method Recv)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:608:19: ptr.MethodByName undefined (type reflect.Value has no field or method MethodByName)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:615:33: receiver.Type().FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:650:21: etyp.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:683:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:684:18: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:686:79: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:688:25: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:689:69: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:693:71: typ.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:700:32: typ.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.16.3/src/text/template/exec.go:703:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
[...more lines following...]</pre>


















