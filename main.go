package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	var dir string
	var pkgName string
	var outFile string

	flag.StringVar(&dir, "dir", ".", "")
	flag.StringVar(&pkgName, "package", "", "")
	flag.StringVar(&outFile, "out", path.Join(dir, "go-builder.go"), "")
	flag.Parse()

	structures := flag.Args()

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
