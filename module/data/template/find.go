package template

const (
	FindOneMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*biz.{{.upperStartCamelObject}}, error) {
	var data *biz.{{.upperStartCamelObject}}
	res := rp.db.WithContext(ctx).Table(rp.TableName()).Where("{{.lowerStartCamelPrimaryKey}} = ?", {{.lowerStartCamelPrimaryKey}}).First(&data)
	return data, res.Error
}
`
	FindAllMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) FindAll(ctx context.Context, where []where_builder.Expr) ([]*biz.{{.upperStartCamelObject}}, error) {
	var data []*biz.{{.upperStartCamelObject}}
	query, args := where_builder.ToWhere(where)
	res := rp.db.WithContext(ctx).Table(rp.TableName()).Where(query, args...).Find(&data)
	return data, res.Error
}
`
	PaginateMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) Paginate(ctx context.Context, where []where_builder.Expr, limit, offset int64, orderBy string) ([]*biz.{{.upperStartCamelObject}}, error) {
	var data []*biz.{{.upperStartCamelObject}}
	query, args := where_builder.ToWhere(where)
	res := rp.db.WithContext(ctx).Table(rp.TableName()).Where(query, args...).Limit(int(limit)).Offset(int(offset)).Order(orderBy).Find(&data)
	return data, res.Error
}
`
	CountMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) Count(ctx context.Context, where []where_builder.Expr) (int64, error) {
	var data int64
	query, args := where_builder.ToWhere(where)
	res := rp.db.WithContext(ctx).Table(rp.TableName()).Where(query, args...).Count(&data)
	return data, res.Error
}
`
)
