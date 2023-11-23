package template

const (
	CreateMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) Create(ctx context.Context, data *biz.{{.upperStartCamelObject}}) (*biz.{{.upperStartCamelObject}}, error) {
	res := rp.db.WithContext(ctx).Table(rp.TableName()).Create(&data)
	return data, res.Error
}
`
)
