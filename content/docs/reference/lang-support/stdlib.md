
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
go/build |  [<span style="color: red">✗</span> no](#gobuild)  | 
go/build/constraint |  <span style="color: green">✔</span> yes  | 
go/constant |  <span style="color: green">✔</span> yes  | 
go/doc |  [<span style="color: red">✗</span> no](#godoc)  | 
go/format |  <span style="color: green">✔</span> yes  | 
go/importer |  [<span style="color: red">✗</span> no](#goimporter)  | 
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
net/http/cgi |  [<span style="color: red">✗</span> no](#nethttpcgi)  | 
net/http/cookiejar |  <span style="color: green">✔</span> yes  | 
net/http/fcgi |  [<span style="color: red">✗</span> no](#nethttpfcgi)  | 
net/http/httptest |  <span style="color: green">✔</span> yes  | 
net/http/httptrace |  <span style="color: green">✔</span> yes  | 
net/http/httputil |  <span style="color: green">✔</span> yes  | 
net/http/pprof |  [<span style="color: red">✗</span> no](#nethttppprof)  | 
net/mail |  <span style="color: green">✔</span> yes  | 
net/rpc |  [<span style="color: red">✗</span> no](#netrpc)  | 
net/rpc/jsonrpc |  [<span style="color: red">✗</span> no](#netrpcjsonrpc)  | 
net/smtp |  <span style="color: green">✔</span> yes  | 
net/textproto |  <span style="color: green">✔</span> yes  | 
net/url |  <span style="color: green">✔</span> yes  | 
os |  <span style="color: green">✔</span> yes  | 
os/exec |  [<span style="color: red">✗</span> no](#osexec)  | 
os/signal |  [<span style="color: red">✗</span> no](#ossignal)  | 
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





































































































## encoding/gob


The compiler gave the following error when this package was imported:

<pre># encoding/gob
/home/ron/.gvm/gos/go1.17/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/home/ron/.gvm/gos/go1.17/src/encoding/gob/decode.go:1118:30: srt.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.17/src/encoding/gob/encode.go:643:70: f.Index undefined (type reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.17/src/encoding/gob/type.go:867:9: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
/home/ron/.gvm/gos/go1.17/src/encoding/gob/type.go:870:21: rt.PkgPath undefined (type reflect.Type has no field or method PkgPath)
</pre>











## encoding/xml


The compiler gave the following error when this package was imported:

<pre># encoding/xml
/home/ron/.gvm/gos/go1.17/src/encoding/xml/typeinfo.go:114:29: f.Index undefined (type *reflect.StructField has no field or method Index)
/home/ron/.gvm/gos/go1.17/src/encoding/xml/typeinfo.go:318:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
/home/ron/.gvm/gos/go1.17/src/encoding/xml/typeinfo.go:319:14: typ.FieldByIndex undefined (type reflect.Type has no field or method FieldByIndex)
</pre>















## go/build


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/doc](#godoc)









## go/doc


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#texttemplate)







## go/importer


This package cannot be imported because the following dependencies cannot be compiled:

  * [go/build](#gobuild)





























## html/template


This package cannot be imported because the following dependencies cannot be compiled:

  * [text/template](#texttemplate)



















































## net/http/cgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [os/exec](#osexec)







## net/http/fcgi


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/http/cgi](#nethttpcgi)











## net/http/pprof


The compiler gave the following error when this package was imported:

<pre>../../../.cache/tinygo/goroot-go1.17-e499b35e58bbb925e7e3a1f14fb802fe799c3d6d9f463d80710a1fc85b74bf86-syscall/src/net/http/pprof/pprof.go:73:2: package runtime/trace is not in GOROOT (/home/ron/.cache/tinygo/goroot-go1.17-e499b35e58bbb925e7e3a1f14fb802fe799c3d6d9f463d80710a1fc85b74bf86-syscall/src/runtime/trace)
</pre>







## net/rpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [encoding/gob](#encodinggob)
  * [html/template](#htmltemplate)





## net/rpc/jsonrpc


This package cannot be imported because the following dependencies cannot be compiled:

  * [net/rpc](#netrpc)













## os/exec


The compiler gave the following error when this package was imported:

<pre># internal/syscall/execenv
/home/ron/.gvm/gos/go1.17/src/internal/syscall/execenv/execenv_default.go:19:17: Environ not declared by package syscall
</pre>





## os/signal


The compiler gave the following error when this package was imported:

<pre># os/signal
/home/ron/.gvm/gos/go1.17/src/os/signal/signal_unix.go:24:11: cannot use syscall.Signal(signal_recv()) (value of type syscall.Signal) as os.Signal value in argument to process: missing method Signal
/home/ron/.gvm/gos/go1.17/src/os/signal/signal_unix.go:38:7: sig (variable of type os.Signal) cannot have dynamic type syscall.Signal (missing method Signal)
</pre>





## os/user


The compiler gave the following error when this package was imported:

<pre># os/user
/home/ron/.gvm/gos/go1.17/src/os/user/lookup.go:15:41: undeclared name: current
/home/ron/.gvm/gos/go1.17/src/os/user/lookup.go:36:9: undeclared name: lookupUser
/home/ron/.gvm/gos/go1.17/src/os/user/lookup.go:45:9: undeclared name: lookupUserId
/home/ron/.gvm/gos/go1.17/src/os/user/lookup.go:51:9: undeclared name: lookupGroup
/home/ron/.gvm/gos/go1.17/src/os/user/lookup.go:57:9: undeclared name: lookupGroupId
/home/ron/.gvm/gos/go1.17/src/os/user/lookup.go:62:9: undeclared name: listGroups
</pre>





































## testing/quick


The compiler gave the following error when this package was imported:

<pre># testing/quick
/home/ron/.gvm/gos/go1.17/src/testing/quick/quick.go:273:11: fType.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.17/src/testing/quick/quick.go:276:11: fType.Out undefined (type reflect.Type has no field or method Out)
/home/ron/.gvm/gos/go1.17/src/testing/quick/quick.go:280:43: fType.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.17/src/testing/quick/quick.go:320:43: xType.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.17/src/testing/quick/quick.go:351:25: f.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.17/src/testing/quick/quick.go:353:95: f.In undefined (type reflect.Type has no field or method In)
</pre>









## text/template


The compiler gave the following error when this package was imported:

<pre># text/template
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:375:17: val.Type().ChanDir undefined (type reflect.Type has no field or method ChanDir)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:375:38: SendDir not declared by package reflect
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:381:20: val.Recv undefined (type reflect.Value has no field or method Recv)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:607:19: ptr.MethodByName undefined (type reflect.Value has no field or method MethodByName)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:614:33: receiver.Type().FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:649:21: etyp.FieldByName undefined (type reflect.Type has no field or method FieldByName)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:682:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:683:18: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:685:79: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:687:25: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:688:69: typ.NumIn undefined (type reflect.Type has no field or method NumIn)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:692:71: typ.NumOut undefined (type reflect.Type has no field or method NumOut)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:699:32: typ.In undefined (type reflect.Type has no field or method In)
/home/ron/.gvm/gos/go1.17/src/text/template/exec.go:702:9: typ.IsVariadic undefined (type reflect.Type has no field or method IsVariadic)
[...more lines following...]</pre>


















