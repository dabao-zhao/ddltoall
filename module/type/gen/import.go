package gen

import (
	"github.com/dabao-zhao/ddltoall/module/type/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
)

func genImports(table *parser.Table, timeImport bool) (string, error) {
	text, err := filex.LoadTemplate(importTemplateFile, template.ImportTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("import").Parse(text).Execute(map[string]interface{}{
		"time": timeImport,
		"data": table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
