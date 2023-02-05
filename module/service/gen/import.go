package gen

import (
	"github.com/dabao-zhao/ddltoall/module/service/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/packagex"
)

func genImports(table *parser.Table, parentPkg string) (string, error) {
	text, err := filex.LoadTemplate(importTemplateFile, template.ImportTpl)
	if err != nil {
		return "", err
	}

	rootPackage := packagex.GetRootPackage()
	importPackage := rootPackage + "/" + parentPkg

	buffer, err := output.With("import").Parse(text).Execute(map[string]interface{}{
		"data":         table,
		"bizPackage":   importPackage + "/biz",
		"typesPackage": importPackage + "/types",
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
