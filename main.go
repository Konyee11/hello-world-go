package main

import (
	"fmt"

	"github.com/Konyee11/hello-world-go/subpkg"
)

func main() {
   fmt.Println(subpkg.Hello())
   fmt.Println(subpkg.Golang())
   fmt.Println(Goodbye())
}
