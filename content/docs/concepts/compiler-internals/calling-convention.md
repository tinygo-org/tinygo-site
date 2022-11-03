---
title: "Calling convention"
weight: 7
---

Go uses a stack-based calling convention and passes a pointer to the argument list as the first argument in the function. There were/are [plans to switch to a register-based calling convention](https://github.com/golang/go/issues/18597) but they're now on hold.

TinyGo, however, uses a register based calling convention. In fact it is somewhat compatible with the C calling convention but with a few quirks:

  * Struct parameters are split into separate arguments, if the number of fields (after flattening recursively) is 3 or lower. This is similar to the [Swift calling convention](https://github.com/apple/swift/blob/master/docs/CallingConvention.rst#physical-conventions). In the case of TinyGo, the size of each field does not matter, a field can even be an array:

        {i8*, i32}           -> i8*, i32
        {{i8*, i32}, i16}    -> i8*, i32, i16
        {{i64}}              -> i64
        {}                   ->
        {i8*, i32, i8, i8}   -> {i8*, i32, i8, i8}
        {{i8*, i32, i8}, i8} -> {{i8*, i32, i8}, i8}

    Note that all native Go data types that are lowered to aggregate types in LLVM are expanded this way: `string`, slices, interfaces, and fat function pointers. This avoids some overhead in the C calling convention and makes the work of the LLVM optimizers easier.

  * Some environments don't support `i64` parameters and return values (like old Node.js versions). Therefore, when targeting browsers, TinyGo will use a special ABI where these parameters and return values are returned with `i64`, allocating the value on the C stack. In other words, imported functions are called with a 64-bit integer on the stack and exported functions must be called with a pointer to a 64-bit integer somewhere in linear memory.

    WASI doesn't have this limitation so it will always use plain `i64` values.

    We will eventually remove support for this workaround, as it is only really needed for very old Node.js versions.

  * The WebAssembly target does not return variables directly that cannot be handled by JavaScript (see above about `i64`, also `struct`, `i64`, multiple return values, etc). Instead, they are stored into a pointer passed as the first parameter by the caller.

    This is the calling convention as implemented by LLVM, with the extension that `i64` return values are returned in the same way as aggregate types.

  * Non-exported functions have two extra `i8*` (roughly equivalent to
    `unsafe.Pointer`) parameters appended at the end of the argument list. The
    first extra parameter is the context for when the function is a function
    pointer. If it is known not to be a closure, this value may be left
    undefined. The second extra parameter is the parent goroutine handle, which
    is used only in blocking functions to return to the parent, otherwise it is
    also unused and may be left undefined.

This calling convention may change in the future. Changes will be documented here. However, even though it may change, it is expected that function signatures that only contain integers and pointers will remain stable.
