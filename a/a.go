package main

import (
	"fmt"
	"github.com/norenstrakh-r7/submodules/b"
)

//Name name imported from b
const Name = b.Name

func main() {
	fmt.Println(Name)
}
