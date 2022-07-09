package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
	"strings"
)

func readAllFile(dir string) (*Package, error) {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	pkg := Package{}

loopFile:
	for _, file := range dirFiles {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		ff, err := parser.ParseFile(fset, path.Join(dir, file.Name()), nil, parser.ParseComments)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, comment := range ff.Comments {
			if strings.Contains(strings.ToLower(comment.Text()), "code generated by go-builder. do not edit") {
				continue loopFile
			}
		}

		if strings.HasSuffix(ff.Name.Name, "_test") {
			continue
		}

		pkg.FileName = append(pkg.FileName, file.Name())
		pkg.Package = ff.Name.Name

		if err := readFiles(fset, ff, &pkg); err != nil {
			return nil, err
		}
	}

	return &pkg, nil
}

func readFiles(fset *token.FileSet, file *ast.File, pkg *Package) error {
	for _, o := range file.Scope.Objects {
		structType := getStructType(o)
		if structType == nil {
			continue
		}

		stru := &Struct{Name: o.Name}
		pkg.Structures = append(pkg.Structures, stru)

		for _, structField := range structType.Fields.List {
			if len(structField.Names) == 0 {
				continue
			}

			field := &Field{Name: structField.Names[0].Name}
			if err := readFieldType(structField.Type, field); err != nil {
				return err
			}
			stru.Fields = append(stru.Fields, field)
		}
	}

	return nil
}

func readFieldType(fieldType ast.Expr, field *Field) error {
	switch tp := fieldType.(type) {
	case *ast.Ident:
		field.KindName = tp.Name
	case *ast.ArrayType:
		field.Slice = true
		return readFieldType(tp.Elt, field)
	case *ast.StarExpr:
		field.Pointer = true
		return readFieldType(tp.X, field)
	case *ast.SelectorExpr:
		if x, ok := tp.X.(*ast.Ident); ok {
			field.KindPkg = x.Name
		}
		field.KindName = tp.Sel.Name
	default:
		return fmt.Errorf("field type unknow: %T", fieldType)
	}

	return nil
}

func getStructType(o *ast.Object) *ast.StructType {
	typeSpec, ok := o.Decl.(*ast.TypeSpec)
	if !ok {
		return nil
	}

	structType := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil
	}
	return structType
}