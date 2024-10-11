package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	Main2()
	fmt.Print(morestrings.ReverseRunes("Hello, world.") + "\n")
	fmt.Print(cmp.Diff("Hello, world.", "Hello, Go.") + "\n")
}
