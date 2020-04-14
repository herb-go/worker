package tools

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"strings"
)

type Context struct {
	Ignored    map[string]bool
	Writer     io.Writer
	Overseers  map[string]string
	WorkerSuff string
	Root       string
	Prefix     string
	Checked    map[string]bool
}

func (c *Context) Printf(format string, v ...interface{}) {
	if c.Writer != nil {
		fmt.Fprintf(c.Writer, format, v...)
	}
}

func (c *Context) MustCheckFolder(path string) {
	fset := token.NewFileSet()
	f, err := parser.ParseDir(fset, path, nil, 0)
	if err != nil {
		panic(err)
	}
	var conf = types.Config{Importer: importer.ForCompiler(fset, "source", nil)}
	for _, ppkg := range f {
		files := make([]*ast.File, 0, len(ppkg.Files))
		for _, v := range ppkg.Files {
			files = append(files, v)
		}
		pkg, err := conf.Check(path, fset, files, nil)
		if err != nil {
			panic(err)
		}
		for _, v := range pkg.Imports() {
			c.CheckPackage(v)
		}
	}
}
func (c *Context) CheckPackage(pkg *types.Package) {
	path := pkg.Path()
	if c.Checked[path] {
		return
	}
	c.Checked[path] = true
	if !strings.HasPrefix(path, c.Prefix) {
		return
	}
	fmt.Println("package", pkg.Name())
	for _, v := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(v)
		if !obj.Exported() {
			continue
		}
		t := obj.Type().String()
		c.Printf("Worker \"%s\" (%s) found.\n", v, t)
	}

	for _, v := range pkg.Imports() {
		c.CheckPackage(v)
	}
}
func (c *Context) MustLoadOverseers(path string) {
	fset := token.NewFileSet()
	f, err := parser.ParseDir(fset, path, nil, 0)
	if err != nil {
		panic(err)
	}
	var conf = types.Config{Importer: importer.ForCompiler(fset, "source", nil)}
	for _, ppkg := range f {
		files := make([]*ast.File, 0, len(ppkg.Files))
		for _, v := range ppkg.Files {
			files = append(files, v)
		}
		pkg, err := conf.Check(path, fset, files, nil)
		if err != nil {
			panic(err)
		}
		for _, v := range pkg.Scope().Names() {
			if strings.HasSuffix(v, c.WorkerSuff) && v != c.WorkerSuff {
				obj := pkg.Scope().Lookup(v)
				if !obj.Exported() {
					continue
				}
				t := obj.Type().String()
				if !c.Ignored[t] {
					c.Printf("Worker \"%s\" (%s) found.\n", v, t)
					c.Overseers[t] = v
				}
			}
		}
	}
}

func NewContext() *Context {
	return &Context{
		Ignored:    map[string]bool{},
		Overseers:  map[string]string{},
		WorkerSuff: "Worker",
		Checked:    map[string]bool{},
	}
}
