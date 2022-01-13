package main

import (
	"fmt"
	"github.com/xFayre/go-hello-world/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes(morestrings.ReverseRunes("Hello, world.")))
}