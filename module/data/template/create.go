package template

const (
	CreateMethod = `
func (rp *{{.upperStartCamelObject}}Repo) Create(ctx context.Context, data *biz.{{.upperStartCamelObject}}) (*biz.{{.upperStartCamelObject}}, error) {
	res := rp.db.WithContext(ctx).Create(&data)
	return data, res.Error
}
`
)
