package gopygen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"text/template"
)

const classDefinitionStr = `
var {{.Ident}}Def = py.Class{
	Name: "{{.Ident}}",
	Pointer: (*{{.Ident}})(nil),
}
`

const objectAllocFunctionStr = `
// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func (obj {{.Ident}}) Alloc() (*{{.Ident}}, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := {{.Ident}}Def.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*{{.Ident}})
	// Copy fields
{{range .Fields.Fields}}
  {{if not .Anonymous}}
	alloc.{{.Name}} = obj.{{.Name}}
  {{end}}
{{end}}
	return alloc, nil
}
`

const classRegisterStr = `
// Registers this type with a python module
func Register{{.Ident}}(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = {{.Ident}}Def.Create(); err != nil {
		return err
	}

	if err = module.AddObject("{{.Ident}}", class); err != nil {
		return err
	}

	return nil
}
`

var classDefinitionTmpl = template.Must(template.New("class_definition").Parse(classDefinitionStr))
var objectAllocFunctionTmpl = template.Must(template.New("object_alloc").Parse(objectAllocFunctionStr))
var classRegisterTmpl = template.Must(template.New("class_register").Parse(classRegisterStr))

type TypeDeclData struct {
	Ident  Ident
	Fields FieldList
}

type TypeDecl struct {
	*TypeDeclData
	fileset *token.FileSet
}

func NewTypeDecl(fileset *token.FileSet) TypeDecl {
	return TypeDecl{
		fileset: fileset,
		TypeDeclData: &TypeDeclData{
			Fields: NewFieldList(fileset),
		},
	}
}

func (d TypeDecl) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.TypeSpec:
		d.Ident = NewIdent(d.fileset)
		ast.Walk(d.Ident, node.Name)
	case *ast.StructType:
		return d.Fields
	}
	return d
}

func (d *TypeDeclData) AllocateFunction() string {
	var buffer bytes.Buffer
	err := objectAllocFunctionTmpl.Execute(&buffer, d)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) RegisterFunction() string {
	var buffer bytes.Buffer
	err := classRegisterTmpl.Execute(&buffer, d)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) ClassDeclaration() string {
	var buffer bytes.Buffer
	err := classDefinitionTmpl.Execute(&buffer, d)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%v", d.Ident.String())
	return buffer.String()
}
