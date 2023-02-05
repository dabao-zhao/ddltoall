package gen

import (
	"github.com/dabao-zhao/ddltoall/module/service/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genTypes(table *parser.Table) (string, error) {
	text, err := filex.LoadTemplate(typeTemplateFile, template.TypeTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
