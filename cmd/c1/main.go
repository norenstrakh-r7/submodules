package main

import (
	"fmt"
	"github.com/norenstrakh-r7/submodules/b"

	// this shows that name of the folder doesn't dictate the name of the
	// the module, we still need to import by folder name
	// but variables that are imported should be referenced by the
	// module name, for example module C objects should be referenced
	// as c.Name , not c.Name
	// go will try to find the module for the package, for example:
	// 		finding module for package github.com/norenstrakh-r7/submodules/c
	"github.com/norenstrakh-r7/submodules/c"
)


const Name = "cmd.Name" + c.Name // c.Name references name in submodules/c folder


func main()  {
	fmt.Println(b.Name)
}
