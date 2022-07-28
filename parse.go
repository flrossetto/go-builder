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

func readAllFile(dir string, structures []string) (*Package, error) {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	pkg := Package{}

	var files []*ast.File

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
		files = append(files, ff)
	}

	for _, ff := range files {
		if err := readFuncs(fset, ff, &pkg); err != nil {
			return nil, err
		}
	}

	for _, ff := range files {
		if err := readStructures(fset, ff, structures, &pkg); err != nil {
			return nil, err
		}
	}

	return &pkg, nil
}

func readFuncs(fset *token.FileSet, file *ast.File, pkg *Package) error {
	for _, d := range file.Decls {
		funcDecl, ok := d.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || funcDecl.Name == nil {
			continue
		}

		if len(funcDecl.Recv.List) == 0 {
			continue
		}

		tp, ok := funcDecl.Recv.List[0].Type.(*ast.StarExpr)
		if !ok {
			continue
		}

		x, ok := tp.X.(*ast.Ident)
		if !ok {
			continue
		}

		pkg.Funcs = append(pkg.Funcs, &StructFunc{
			Name:   funcDecl.Name.Name,
			Struct: x.Name,
		})
	}
	return nil
}

func readStructures(fset *token.FileSet, file *ast.File, structures []string, pkg *Package) error {
	for _, o := range file.Scope.Objects {
		structType := getStructType(o)
		if structType == nil {
			continue
		}

		stru := &Struct{Name: o.Name}
		if len(structures) > 0 && !stringInSlice(stru.Name, structures) {
			continue
		}

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

		if len(stru.Fields) > 0 {
			pkg.Structures = append(pkg.Structures, stru)
		}
	}

	return nil
}

func readFieldType(expr ast.Expr, field *Field) error {
	switch tp := expr.(type) {
	case *ast.Ident:
		field.KindName = tp.Name
		field.PrimitiveType = primitiveType(tp)
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
	case *ast.InterfaceType:
		field.Interface = true
	case *ast.MapType:
		var key Field
		if err := readFieldType(tp.Key, &key); err != nil {
			return err
		}

		var value Field
		if err := readFieldType(tp.Value, &value); err != nil {
			return err
		}

		field.KindName = fmt.Sprintf("map[%s]%s", fieldType(key), fieldType(value))
	default:
		return fmt.Errorf("field type unknow: %T", fieldType)
	}

	return nil
}

func primitiveType(f *ast.Ident) string {
	if f == nil || f.Obj == nil || f.Obj.Decl == nil {
		return ""
	}

	typeSpec, ok := f.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return ""
	}

	ident, ok := typeSpec.Type.(*ast.Ident)
	if !ok {
		return ""
	}

	return ident.Name
}

func getStructType(o *ast.Object) *ast.StructType {
	if o == nil || o.Decl == nil {
		return nil
	}

	typeSpec, ok := o.Decl.(*ast.TypeSpec)
	if !ok || typeSpec.Type == nil {
		return nil
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if ok {
		return structType
	}

	ident, ok := typeSpec.Type.(*ast.Ident)
	if !ok || ident.Obj == nil {
		return nil
	}

	return getStructType(ident.Obj)
}

func stringInSlice(value string, list []string) bool {
	for _, item := range list {
		if value == item {
			return true
		}
	}
	return false
}
