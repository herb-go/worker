package tools

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/doc"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Worker struct {
	Name         string
	Type         string
	Introduction string
}
type Package struct {
	Name       string
	ID         string
	ImportPath string
	Path       string
	Workers    []*Worker
}

func (p *Package) AddWorker(w *Worker) {
	p.Workers = append(p.Workers, w)
}
func (p *Package) IsEmpty() bool {
	return len(p.Workers) == 0
}
func NewPackage() *Package {
	return &Package{
		Workers: []*Worker{},
	}
}

type Overseer struct {
	Name string
	Type types.Type
}
type Context struct {
	Ignored       map[string]bool
	Writer        io.Writer
	Overseers     map[string]*Overseer
	OverseersPath string
	WorkerSuff    string
	Root          string
	Checked       map[string]bool
	Result        []*Package
	Filename      string
	FileMode      os.FileMode
	Marker        string
	Assignable    bool
	GomodFolder   string
	fset          *token.FileSet
	Importer      types.Importer
}

func (c *Context) Printf(format string, v ...interface{}) {
	if c.Writer != nil {
		fmt.Fprintf(c.Writer, format, v...)
	}
}

func (c *Context) MustCheckFolder(path string) {
	defer func() {
		build.Default.Dir = ""
	}()
	build.Default.Dir = c.GomodFolder
	c.fset = token.NewFileSet()
	f, err := parser.ParseDir(c.fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var conf = types.Config{Importer: c.Importer}
	for _, ppkg := range f {
		files := []*ast.File{}
		for _, v := range ppkg.Files {
			files = append(files, v)
		}
		pkg, err := conf.Check(path, c.fset, files, nil)
		if err != nil {
			panic(err)
		}
		for _, v := range pkg.Imports() {
			c.CheckPackage(v)
		}
	}

}
func (c *Context) IsWorker(o types.Object) bool {
	t := o.Type()

	if c.Overseers[t.String()] != nil {
		return true
	}
	if c.Assignable {
		for _, v := range c.Overseers {
			if types.IsInterface(v.Type) {
				if types.AssignableTo(t, v.Type) {
					return true
				}
			}
		}
	}
	return false
}
func (c *Context) CheckPackage(pkg *types.Package) {
	if c.Root[len(c.Root)-1:] != "/" {
		c.Root = c.Root + "/"
	}
	path := pkg.Path()
	if c.Checked[path] {
		return
	}

	c.Checked[path] = true
	p, err := build.Default.Import(path, c.Root, build.FindOnly)

	if err != nil {
		panic(err)
	}
	if !strings.HasPrefix(p.Dir+"/", c.Root) {
		return
	}

	f, err := parser.ParseDir(c.fset, p.Dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	astpkg := f[pkg.Name()]
	if astpkg == nil {
		return
	}
	d := doc.New(astpkg, path, 0)

	intros := map[string]string{}
	if c.Marker != "" {
		for _, v := range d.Notes[c.Marker] {
			body := strconv.Quote(strings.TrimSpace(v.Body))
			if body != "\"\"" {
				intros[v.UID] = body
			}
		}
	}

	tp := NewPackage()
	tp.Name = pkg.Name()
	tp.ImportPath = path
	tp.Path = p.Dir
	tp.ID = strings.TrimPrefix(p.Dir, c.Root)
	if path != c.OverseersPath {
		for _, v := range pkg.Scope().Names() {
			obj := pkg.Scope().Lookup(v)
			if !obj.Exported() {
				continue
			}
			t := obj.Type().String()
			if !c.IsWorker(obj) {
				continue
			}
			w := &Worker{
				Name:         v,
				Type:         t,
				Introduction: intros[v],
			}
			tp.AddWorker(w)

			c.Printf("Worker \"%s\" (%s) in %s found.\n", v, t, tp.ID)
		}
	}
	if !tp.IsEmpty() {
		c.Result = append(c.Result, tp)
	}
	for _, v := range pkg.Imports() {
		c.CheckPackage(v)
	}
}
func (c *Context) MustLoadOverseers(path string) {
	defer func() {
		build.Default.Dir = ""
	}()
	build.Default.Dir = c.GomodFolder
	p, err := build.Default.Import(path, c.Root, build.FindOnly)
	if err != nil {
		panic(err)
	}
	c.OverseersPath = p.ImportPath
	c.fset = token.NewFileSet()
	f, err := parser.ParseDir(c.fset, p.Dir, nil, 0)
	if err != nil {
		panic(err)
	}
	c.Importer = importer.ForCompiler(c.fset, "source", nil)
	var conf = types.Config{Importer: c.Importer}
	for _, ppkg := range f {
		files := make([]*ast.File, 0, len(ppkg.Files))
		for _, v := range ppkg.Files {
			files = append(files, v)
		}
		pkg, err := conf.Check(p.Dir, c.fset, files, nil)
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
					c.Printf("Overseer \"%s\" (%s) found.\n", v, t)
					c.Overseers[t] = &Overseer{
						Name: v,
						Type: obj.Type(),
					}
				}
			}
		}
	}
}

func (c *Context) MustRender() map[string][]byte {
	result := map[string][]byte{}
	for _, v := range c.Result {
		buf := bytes.NewBuffer(nil)
		err := Template.Execute(buf, v)
		if err != nil {
			panic(err)
		}
		result[filepath.Join(v.Path, c.Filename)] = buf.Bytes()
	}
	return result
}

func (c *Context) MustRenderAndWrite() {
	result := c.MustRender()
	for k, v := range result {
		err := ioutil.WriteFile(k, v, c.FileMode)
		if err != nil {
			panic(err)
		}
		c.Printf("File %s generated. \n", k)
	}
}
func NewContext() *Context {
	return &Context{
		Ignored:     map[string]bool{},
		Overseers:   map[string]*Overseer{},
		WorkerSuff:  "Worker",
		GomodFolder: "",
		Checked:     map[string]bool{},
		Result:      []*Package{},
		Filename:    "workers.autogenerated.go",
		FileMode:    0660,
		Marker:      "WORKER",
		Assignable:  true,
	}
}
