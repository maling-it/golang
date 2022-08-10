package main

import (
	"fmt"
	"time"
)

func main() {
	go func(){fmt.Println("menjalankan fungsi goroutine")}()
	fmt.Println("malingit")
	time.Sleep(1 * time.Second)
	fmt.Println("selesai")
}

//func start() {
//	fmt.Println("menjalankan fungsi goroutine")
//}
