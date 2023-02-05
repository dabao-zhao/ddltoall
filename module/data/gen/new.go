package gen

import (
	"github.com/dabao-zhao/ddltoall/module/data/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
)

func genNew(table *parser.Table) (string, error) {
	text, err := filex.LoadTemplate(dataNewTemplateFile, template.NewTpl)
	if err != nil {
		return "", err
	}

	camel := table.Name.ToCamel()
	buffer, err := output.With("new").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
