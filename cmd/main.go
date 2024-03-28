package main

import (
	"fmt"

	"github.com/imjowend/hex-arch-golang/internal/adapters/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
