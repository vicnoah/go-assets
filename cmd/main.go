package main

import (
	"os"

	"github.com/vicnoah/go-assets"
)

func main() {
	/* 	g := assets.Generator{
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
	   	} */
	g := assets.Generator{}

	if err := g.Add("."); err != nil {
		panic(err)
	}

	// This will write a go file to standard out. The generated go file
	// will reside in the g.PackageName package and will contain a
	// single variable g.VariableName of type assets.FileSystem containing
	// the whole file system.
	g.Write(os.Stdout)
}
