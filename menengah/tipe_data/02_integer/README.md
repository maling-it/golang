# Integer

Golang mendukung tipe data integer secara ekstensif yang antara lain adalah

- ``int8`` adalah integer bertanda 8-bit yang rentangnya ``-128`` hingga ``127``
- ``int16`` adalah integer bertanda 16-bit yang jangkaunnya ``-32768`` hingga ``32767``
- ``int32`` integer bertanda 32-bit yang rentangnya ``-2147483648`` hingga ``2147483647``
- ``int64`` integer yang bertanda 64-bit yang rentangnya ``-9223372036854775808`` hingga ``9223372036854775807``

**unsigned int pada go**

- ``uint8`` 8-bit unsigned integer yang rentangnya ``0`` hingga ``255``
- ``uint16`` 16-bit unigned integer yang rentangnya ``0`` hingga ``65535``
- ``uint32`` integer yang tidak bertanda 32-bit yang rentangnya ``0`` hingga ``4294967295``
- ``uint64`` integer 64-bit unsigned yang rentangnya ``0`` hingga ``18446744073709551615``

## integer overflow pada golang

jika kita menetapkan jenis dan kemudian menggunakan angka yang lebih besar dari rentang jenis untuk menetapkannya, itu akan gagal. contoh

```golang
package main

import "fmt"

func main() {
    var x uint8
        fmt.Println("throw integer overflow")
    x = 267  // range dari uint8 adalah 0 - 255
}
```

## type conversion pada golang

jika kita mengonversi ke jenis yang memiliki rentang lebih renda dari rentang kita saat ini, kehilangan data akan terjadi. kita melakukan typecast dengan langsung menggunakan nama variabel sebagai fungsi untuk mengkonveri tipe.

```golang
package main

import "fmt"

func main() {
    var x int32
    var y uint32  // 0 sampai 4294967295
    var z uint8   // 0 sampai 255

    x = 26700
    
    y = uint32(x)   // data ditampilkan secara penuh karena size terpenuhi

    z = uint8(x)    // data hilang karena size sudah melewati batas

    fmt.Println(y, z)
}
```
