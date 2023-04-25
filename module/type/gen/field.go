package gen

import (
	"strings"

	"github.com/dabao-zhao/ddltoall/module/type/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genFields(table *parser.Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genFieldsOnlyJsonTag(table *parser.Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genFieldOnlyJsonTag(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genField(table *parser.Table, field *parser.Field) (string, error) {
	tag, err := genTag(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	tagOnlyJson, err := genTagOnlyJson(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	text, err := filex.LoadTemplate(fieldTemplateFile, template.FieldTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":        stringx.SafeString(field.Name.ToCamel()),
			"type":        field.DataType,
			"tag":         tag,
			"tagOnlyJson": tagOnlyJson,
			"hasComment":  field.Comment != "",
			"comment":     field.Comment,
			"data":        table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func genFieldOnlyJsonTag(table *parser.Table, field *parser.Field) (string, error) {
	tag, err := genTagOnlyJson(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	text, err := filex.LoadTemplate(fieldTemplateFile, template.FieldTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":       stringx.SafeString(field.Name.ToCamel()),
			"type":       field.DataType,
			"tag":        tag,
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
			"data":       table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
