package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var dir string
	var pkgName string
	var outFile string

	flag.StringVar(&outFile, "o", "go-builder.go", "")
	flag.Parse()

	if dir = flag.Arg(0); dir == "" {
		dir = "."
	}

	structures := strings.Split(flag.Arg(1), ",")

	pkg, err := readAllFile(dir, structures)
	if err != nil {
		panic(err)
	}

	if pkgName != "" {
		pkg.Package = pkgName
	}

	builderFile, err := createBuilder(pkg)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(outFile, builderFile, os.ModePerm); err != nil {
		panic(err)
	}
}
