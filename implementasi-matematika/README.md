---
sort: 3
---

# Implementasi Matematika

## Apa itu Armstrong number?

Armstrong number adalah, jika masing-masing angka n diexponen dengan digit yang dimiliki memiliki jumlah yang sama dengan angka n.

Misal :

`n = 371`, pada n ada 3 digit jadi n^digit

Test `371, 3³ + 7³ + 1³ = 27 + 343 + 1 = 371`

Test `5, 5¹ = 5`

Test `1634, 1⁴ + 6⁴+ 3⁴+ 4⁴= 1 + 1296 + 81 + 256 = 1634`

371, 5, 1634 adalah angka armstrong karena angka awal sama dengan angka hasil jika, angka awal di exponen dengan digit yang dimiliki angka tersebut.

Notes: Semua angka berdigit 1 adalah armstrong number.

Berikut adalah program golang yang memeriksa apakah angka tersebut adalah angka Armstrong:

```go
package armstrong

import (
  "math"
  "strconv"
)

func cekArmstrong(angka int) bool{
  var digit_kanan int
  var sum int = 0
  var tempNum int = angka 

  // untuk mendapatkan jumlah digit
  // di dalam nomor
  length := float64(len(strconv.Itoa(angka)))

  // dapatkan digit paling kanan dan putuskan loop
  // setelah semua digit diulang

  for tempNum > 0 {
    digit_kanan = tempNum % 10
    sum += int(math.Pow(float64(digit_kanan), length))

    // perbarui digit input dikurangi yang diproses
    // paling kanan
    tempNum /= 10
  }
  
  return angka == sum
}
```