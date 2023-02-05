package gen

import (
	"github.com/dabao-zhao/ddltoall/module/type/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
)

func genTag(table *parser.Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := filex.LoadTemplate(tagTemplateFile, template.TagTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("tag").Parse(text).Execute(map[string]interface{}{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
