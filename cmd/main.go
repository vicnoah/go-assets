package main

import (
	"fmt"
	"github.com/vicnoah/go-assets"
	"os"
)

func main() {
	g := assets.Generator{
		PackageName:  "mfs",
		VariableName: "Assets",
		StripPrefix:  "as",
	}

	path := "cmd"

	if err := g.Add(path); err != nil {
		panic(err)
	}

	err := g.Write(os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
