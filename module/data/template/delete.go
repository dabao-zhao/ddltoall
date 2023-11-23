package template

const (
	// DeleteMethod defines a delete template
	DeleteMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	res := rp.db.WithContext(ctx).Table(rp.TableName()).Delete({{.lowerStartCamelPrimaryKey}})
	return res.Error
}
`
)
